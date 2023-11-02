package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Atributo_ide struct {
	ID     string
	Indice BaseNodo
	Vector bool
}

type Atributo_general struct {
	ID_inicial      Atributo_ide
	Lista_atributos []Atributo_ide
}

func (a Atributo_general) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	variable, size := ambito_padre.BuscarVariable(a.ID_inicial.ID)
	if variable == nil {
		panic("La variable no existe " + a.ID_inicial.ID)
	}
	if a.ID_inicial.Vector { //el id es vector: id[i](.atributo)*
		if variable.Tipo_dimension != valor.DIMENSION1 {
			panic("El id no es un vector" + a.ID_inicial.ID)
		}
		if !variable.Is_instancia {
			panic("El vector no es una instancia")
		}
		//codigo para acceder a la posicion
	}
	if variable.Is_instancia { //vive en el stack
		var resultado valor.Value
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
		for indice, atributo_buscar := range a.Lista_atributos { //viene uno si o si
			atributo, posicion := estruct.Buscar_atributo(atributo_buscar.ID)
			if atributo == nil {
				panic("El atributo no existe al struct " + estruct.Nombre)
			}
			if atributo_buscar.Vector { //el id es vector: id.atributo[i](.atributo)*
				if atributo.Tipo_dimension != valor.DIMENSION1 {
					panic("El id no es un vector" + atributo_buscar.ID)
				}
				//codigo para acceder a la posicion
			}
			generador.Mi_generador.AddExpression(puntero_heap, puntero_heap, strconv.Itoa(posicion), "+")
			generador.Mi_generador.AddGetHeap(valor_heap, "(int)"+puntero_heap)
			if atributo.Is_instancia { //
				estruct, _ = ambito_padre.Buscar_struct(atributo.Tipo_struct)
				generador.Mi_generador.AddAssign(puntero_heap, valor_heap)
				resultado = valor.Value{Value: valor_heap, Is_intancia: true, Tipo_struct: atributo.Tipo_struct}
			} else if atributo.Tipo == valor.BOOLEAN {
				true_label := generador.Mi_generador.NewLabel()
				false_label := generador.Mi_generador.NewLabel()
				generador.Mi_generador.AddIf(valor_heap, "1", "==", true_label)
				generador.Mi_generador.AddGoto(false_label)
				resultado = valor.Value{Type: valor.BOOLEAN}
				resultado.TrueLabel = append(resultado.TrueLabel, true_label)
				resultado.FalseLabel = append(resultado.FalseLabel, false_label)
				if indice < len(a.Lista_atributos)-1 { //si no es el ultimo error
					panic("La variable " + atributo.Id + " no tiene atributos")
				}
			} else {
				resultado = valor.Value{Value: valor_heap, Type: atributo.Tipo}
			}
		}
		return resultado
	} else {
		panic("Error no tiene atributos " + variable.Id)
	}
}
