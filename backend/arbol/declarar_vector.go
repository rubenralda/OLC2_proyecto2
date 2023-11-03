package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Declarar_vector struct {
	Tipo            string // String, Int, Bool, Float, char, (Nombre_struct)
	ID              string
	Lista_expresion []BaseNodo // 10, 20.5, "hola", true, objeto_strcut
	ID_otro_vector  string     //no viene para asignar otro vector
}

func (d Declarar_vector) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	existe, _ := ambito_padre.BuscarVariable(d.ID)
	if existe != nil {
		ambito_padre.Agregar_error("La variable ya existe " + d.ID)
		return valor.Value{}
	}
	variable := ambito.Variables{Id: d.ID, Tipo_dimension: valor.DIMENSION1, Posicion_relativa: ambito_padre.Size}
	variable.Tipo = Tipo_variable[d.Tipo]
	if variable.Tipo == valor.NULL {
		//es un array de objetos validar por d.Tipo
		estruct, _ := ambito_padre.Buscar_struct(d.Tipo)
		if estruct == nil {
			ambito_padre.Agregar_error("El struct no esta definido " + d.Tipo)
			return valor.Value{}
		}
		variable.Is_instancia = true
		variable.Tipo_struct = d.Tipo
	}
	//poner el nodo vacio
	posicion_valor := generador.Mi_generador.NewTemp()
	posicion_puntero := generador.Mi_generador.NewTemp()
	posicion_stack := generador.Mi_generador.NewTemp()
	generador.Mi_generador.AddComment("inicio declaracion")
	generador.Mi_generador.AddExpression(posicion_stack, "P", strconv.Itoa(ambito_padre.Size), "+")
	generador.Mi_generador.AddSetStack("(int)"+posicion_stack, "H")
	ambito_padre.Size++ //se reserva aunque no haya terminado de declarar porque puede que sobreescriban esa posicion
	generador.Mi_generador.AddAssign(posicion_valor, "H")
	generador.Mi_generador.AddExpression(posicion_puntero, posicion_valor, "1", "+")
	generador.Mi_generador.AddSetHeap("(int)"+posicion_puntero, "-1")
	generador.Mi_generador.AddExpression("H", "H", "2", "+") //reservo 2 espacios
	for _, item := range d.Lista_expresion {
		generador.Mi_generador.AddComment("ejecuta expresion")
		resultado := item.Ejecutar(ambito_padre)
		if variable.Is_instancia {
			if variable.Tipo_struct != resultado.Tipo_struct {
				ambito_padre.Agregar_error("El item no es del tipo del array " + variable.Id)
				return valor.Value{}
			}
		} else if variable.Tipo == valor.BOOLEAN {
			if resultado.Type != valor.BOOLEAN {
				ambito_padre.Agregar_error("El item no es del tipo del array " + variable.Id)
				return valor.Value{}
			}
			// imprimir etiquetas y en un temporal guardar 0 o 1 para darle setHeap ese valor
		} else if variable.Tipo != resultado.Type {
			ambito_padre.Agregar_error("El item no es del tipo del array " + variable.Id)
			return valor.Value{}
		}
		//falta el array de bool, que usa etiquetas true y false
		generador.Mi_generador.AddComment("metiendovalor")

		generador.Mi_generador.AddSetHeap("(int)"+posicion_valor, resultado.Value)
		generador.Mi_generador.AddExpression(posicion_puntero, posicion_valor, "1", "+")
		generador.Mi_generador.AddSetHeap("(int)"+posicion_puntero, "H")

		generador.Mi_generador.AddAssign(posicion_valor, "H")
		generador.Mi_generador.AddExpression(posicion_puntero, posicion_valor, "1", "+")
		generador.Mi_generador.AddSetHeap("(int)"+posicion_puntero, "-1")
		generador.Mi_generador.AddExpression("H", "H", "2", "+") //reservo 2 espacios
		generador.Mi_generador.AddComment("metiendovalortermino")
	}
	ambito_padre.AgregarVariable(variable)
	return valor.Value{}
}
