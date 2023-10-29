package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Decremento_variable struct {
	Id        string
	Expresion BaseNodo
}

func (d Decremento_variable) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	result := valor.Value{Type: valor.NULL}
	variable, size_total := ambito_padre.BuscarVariable(d.Id)
	if variable == nil {
		panic("La variable no existe " + d.Id)
	}
	if variable.Is_constante {
		panic("No se puede modificar una constante")
	}
	if variable.Tipo_dimension != valor.DIMENSION0 {
		panic("La operacion asignacion no esta permitida " + d.Id)
	}
	resultado := d.Expresion.Ejecutar(ambito_padre)
	if variable.Tipo == valor.FLOAT {
		if resultado.Type != valor.INTEGER && resultado.Type != valor.FLOAT {
			panic("Error de tipos al decrementar " + d.Id)
		}
	} else if variable.Tipo == valor.INTEGER {
		if resultado.Type != valor.INTEGER {
			panic("Error de tipos al decrementar " + d.Id)
		}
	} else {
		panic("Error de tipo al asignar " + d.Id)
	}
	posicion_variable := generador.Mi_generador.NewTemp()
	//validar antes si es por referencia
	posicion_entorno := generador.Mi_generador.NewTemp()
	tmp_valor := generador.Mi_generador.NewTemp()
	tmp_final := generador.Mi_generador.NewTemp()
	generador.Mi_generador.AddExpression(posicion_entorno, "P", strconv.Itoa(size_total), "-")
	generador.Mi_generador.AddExpression(posicion_variable, posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
	generador.Mi_generador.AddGetStack(tmp_valor, "(int)"+posicion_variable)

	generador.Mi_generador.AddExpression(tmp_final, tmp_valor, resultado.Value, "-")
	generador.Mi_generador.AddSetStack("(int)"+posicion_variable, tmp_final)
	variable.Is_init = true
	return result
}
