package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Tipo_matriz struct {
	Dimension int
	Primitivo string // String, Int, Bool, Float, char, ""
}

type Declarar_matriz struct {
	Tipo_matriz_ex Tipo_matriz // vacio si no viene
	Id             string
	Matriz         []interface{}
	/*
			[
		    [[37, 49, 61, 29, 44], [56, 60, 51, 68, 70], [47, 15, 39, 17, 74]],
		    [[69, 74, 52, 34, 36], [24, 44, 50, 18, 76], [74, 60, 32, 63, 78]],
		    [[78, 14, 23, 52, 33], [28, 79, 77, 55, 24], [23, 79, 47, 62, 44]],
		    [[73, 53, 11, 49, 52], [29, 16, 65, 34, 12], [72, 69, 30, 44, 37]]
			]
	*/
}

func (d Declarar_matriz) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	tipo := Tipo_variable[d.Tipo_matriz_ex.Primitivo]
	posicion_matriz := generador.Mi_generador.NewTemp()
	generador.Mi_generador.AddAssign(posicion_matriz, "H")
	tamano := 0
	variable := ambito.Variables{Id: d.Id, Posicion_relativa: ambito_padre.Size, Tipo_dimension: valor.DIMENSIONN, Tipo: tipo}
	if d.Tipo_matriz_ex.Dimension != 0 {
		tamano, variable.Len_dimension = comprobar_dimensiones(d.Matriz, ambito_padre, &tipo)
		if d.Tipo_matriz_ex.Dimension != tamano {
			ambito_padre.Agregar_error("La definicion de dimensiones no concuerda " + d.Id)
			return valor.Value{}
		}
	} else {
		tamano, variable.Len_dimension = comprobar_dimensiones(d.Matriz, ambito_padre, &tipo)
		d.Tipo_matriz_ex.Dimension = tamano
	}
	if d.Tipo_matriz_ex.Dimension < 2 {
		ambito_padre.Agregar_error("La definicion no es una matriz " + d.Id)
		return valor.Value{}
	}
	ambito_padre.AgregarVariable(variable)
	ambito_padre.Size++
	posicion_stack := generador.Mi_generador.NewTemp()
	generador.Mi_generador.AddExpression(posicion_stack, "P", strconv.Itoa(variable.Posicion_relativa), "+")
	generador.Mi_generador.AddSetStack("(int)"+posicion_stack, posicion_matriz)
	return valor.Value{}
}

/*
comprueba si el slice cada posicion tiene la misma profundidad, por ejemplo si en
una posicion hay otro slice pero en la posicion de la par no, no tiene la misma profundidad y
tira error
*/
func comprobar_dimensiones(vector []interface{}, ambito_padre *ambito.Ambito, primitivo *valor.TipoExpresion) (int, []int) {
	dimension := 0
	dim_hijo := 0
	len_dimension := []int{len(vector)}
	len_hijo := []int{}
	for i, item := range vector {
		switch rr := item.(type) {
		case []interface{}:
			dim_hijo, len_hijo = comprobar_dimensiones(rr, ambito_padre, primitivo)
			dim_hijo += 1
			if i == 0 { //si es el primer elemento determina la profundidad de sus hermanos
				dimension = dim_hijo
			} else if dimension != dim_hijo {
				ambito_padre.Agregar_error("No todas las dimensiones tienen la misma profundidad")
			}
		default:
			if i == 0 {
				dimension = 1
			} else if dimension != 1 {
				ambito_padre.Agregar_error("No todas las dimensiones tienen la misma profundidad")
			}
			resultado := rr.(BaseNodo).Ejecutar(ambito_padre)
			if *primitivo == valor.NULL {
				primitivo = &resultado.Type
			}
			if resultado.Type != *primitivo {
				ambito_padre.Agregar_error("No todas las dimensiones tienen la misma profundidad")
			}
			generador.Mi_generador.AddSetHeap("(int)H", resultado.Value)
			generador.Mi_generador.AddExpression("H", "H", "1", "+")
		}
	}
	len_dimension = append(len_dimension, len_hijo...)
	return dimension, len_dimension
}
