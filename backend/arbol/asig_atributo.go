package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Asignar_atributos struct {
	ID_inicial      Atributo_ide
	Lista_atributos []Atributo_ide
	Expresion       BaseNodo
}

func (a Asignar_atributos) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	variable, size := ambito_padre.BuscarVariable(a.ID_inicial.ID)
	if variable == nil {
		ambito_padre.Agregar_error("La variable no existe " + a.ID_inicial.ID)
		return valor.Value{}
	}
	if !variable.Is_instancia {
		ambito_padre.Agregar_error("Error no tiene atributos " + variable.Id)
		return valor.Value{}
	}
	//vive en el stack
	estruct, _ := ambito_padre.Buscar_struct(variable.Tipo_struct) //busco el struct para recuperar que atributo es
	posicion_struct := generador.Mi_generador.NewTemp()
	tmp_posicion_entorno := generador.Mi_generador.NewTemp()
	puntero_heap := generador.Mi_generador.NewTemp()
	valor_heap := generador.Mi_generador.NewTemp()
	if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
		generador.Mi_generador.AddAssign(tmp_posicion_entorno, "0")
	} else {
		generador.Mi_generador.AddExpression(tmp_posicion_entorno, "P", strconv.Itoa(size), "-")
	}
	generador.Mi_generador.AddExpression(posicion_struct, tmp_posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
	generador.Mi_generador.AddGetStack(puntero_heap, "(int)"+posicion_struct)
	var atributo_modificar *ambito.Variables                 //para hacer las validaciones despues del for
	for indice, atributo_buscar := range a.Lista_atributos { //viene uno si o si
		atributo, posicion := estruct.Buscar_atributo(atributo_buscar.ID)
		if atributo == nil {
			ambito_padre.Agregar_error("El atributo no existe al struct " + estruct.Nombre)
			return valor.Value{}
		}
		atributo_modificar = atributo
		//validad antes si es un atributo vector para calcular el valor en la posicion puesta
		generador.Mi_generador.AddExpression(puntero_heap, puntero_heap, strconv.Itoa(posicion), "+")

		if atributo.Is_instancia {
			//modifco el puntero con el del nuevo struct
			estruct, _ = ambito_padre.Buscar_struct(atributo.Tipo_struct)
			if indice < len(a.Lista_atributos)-1 { //si no es el ultimo
				generador.Mi_generador.AddGetHeap(valor_heap, "(int)"+puntero_heap)
				generador.Mi_generador.AddAssign(puntero_heap, valor_heap)
			}
		} else { //es un int, float, string, char, bool
			if indice < len(a.Lista_atributos)-1 { //si no es el ultimo error
				ambito_padre.Agregar_error("La variable " + atributo.Id + " no tiene atributos")
				return valor.Value{}
			}
		}
	}
	resultado := a.Expresion.Ejecutar(ambito_padre)
	if resultado.Type == valor.NULL {
		if atributo_modificar.Tipo_struct != resultado.Tipo_struct {
			ambito_padre.Agregar_error("La expresion no es del mismo tipo del atributo " + atributo_modificar.Id)
			return valor.Value{}
		}
		generador.Mi_generador.AddSetHeap("(int)"+puntero_heap, resultado.Value)
	} else if resultado.Type == atributo_modificar.Tipo {
		//el temporal puntero_heap se pudo haber borrado si es una llamada recursiva (validar)
		generador.Mi_generador.AddSetHeap("(int)"+puntero_heap, resultado.Value)
	} else {
		ambito_padre.Agregar_error("La expresion no es del mismo tipo del atributo " + atributo_modificar.Id)
	}
	return valor.Value{}
}
