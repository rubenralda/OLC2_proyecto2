package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Llamada_funcion struct {
	Id                  string
	Lista_argumentos    []Lista_argumentos
	Declarar_objeto_amb Declarar_objeto
}

func (l Llamada_funcion) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	funcion, _ := ambito_padre.Buscar_funcion(l.Id)
	if funcion == nil {
		panic("La funcion no existe " + l.Id)
	}
	if len(l.Lista_argumentos) != len(funcion.Parametros) {
		panic("Faltan parametros en la funcion " + l.Id)
	}
	generador.Mi_generador.AddComment("llamadaempieza")
	size := ambito_padre.Size
	posicion_argumento := generador.Mi_generador.NewTemp()
	generador.Mi_generador.AddExpression(posicion_argumento, "P", strconv.Itoa(size+1), "+")
	//verificar nombres internos, externos y guardar el valor en el stack
	for i, argumento := range l.Lista_argumentos {
		//comprobar los nombres de los parametros
		if funcion.Parametros[i].Id_externo == "" { //solo posee un ID
			if argumento.Id_externo != funcion.Parametros[i].Id_interno {
				panic("El nombre externo no coincide " + argumento.Id_externo)
			}
		} else {
			if funcion.Parametros[i].Id_externo == "_" {
				if argumento.Id_externo != "" {
					panic("La funcion no usa nombre externo" + l.Id)
				}
			} else if funcion.Parametros[i].Id_externo != argumento.Id_externo {
				panic("El nombre del parametro no coincide con la funcion " + funcion.Parametros[i].Id_externo)
			}
		}
		//verificar la expresion
		resultado := argumento.Expresion.Ejecutar(ambito_padre)
		if argumento.Referencia {
			if resultado.Referencia == "" {
				panic("La expresion se esperaba una referencia " + funcion.Parametros[i].Id_interno + " " + funcion.Nombre)
			}
			generador.Mi_generador.AddSetStack("(int)"+posicion_argumento, resultado.Referencia)
		} else {
			generador.Mi_generador.AddSetStack("(int)"+posicion_argumento, resultado.Value)
		}
		generador.Mi_generador.AddExpression(posicion_argumento, posicion_argumento, "1", "+")
	}
	tmp_return := generador.Mi_generador.NewTemp()
	generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(size), "+")
	generador.Mi_generador.AddCall(l.Id)
	generador.Mi_generador.AddGetStack(tmp_return, "(int)P")
	generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(size), "-")
	if funcion.Tipo == valor.NULL {
		generador.Mi_generador.AddComment("finLlamada")
		return valor.Value{}
	} else if funcion.Tipo == valor.BOOLEAN {
		true_label := generador.Mi_generador.NewLabel()
		false_label := generador.Mi_generador.NewLabel()

		generador.Mi_generador.AddIf(tmp_return, "1", "==", true_label)
		generador.Mi_generador.AddGoto(false_label)
		result := valor.Value{Type: valor.BOOLEAN}
		result.TrueLabel = append(result.TrueLabel, true_label)
		result.FalseLabel = append(result.FalseLabel, false_label)
		generador.Mi_generador.AddComment("finLlamada")
		return result
	} else {
		generador.Mi_generador.AddComment("finLlamada")
		return valor.Value{Value: tmp_return, Type: funcion.Tipo}
	}

}

type Lista_argumentos struct {
	Id_externo string
	Referencia bool // cuando se usa & verificar que sea id
	Expresion  BaseNodo
}
