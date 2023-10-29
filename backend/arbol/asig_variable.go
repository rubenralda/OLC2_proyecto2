package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Asignacion_variable struct {
	Id        string
	Expresion BaseNodo
}

func (d Asignacion_variable) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
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
	posicion_variable := generador.Mi_generador.NewTemp()
	resultado := d.Expresion.Ejecutar(ambito_padre)
	if variable.Is_instancia {
		if variable.Tipo_struct == resultado.Tipo_struct {
			//logica para cambiar la posicion relativa a la nueva del objeto struct en el heep
		} else {
			panic("Tipo no permitivo")
		}
	} else if variable.Tipo == valor.FLOAT {
		if resultado.Type == valor.INTEGER || resultado.Type == valor.FLOAT {
			//validar antes si es por referencia
			posicion_entorno := generador.Mi_generador.NewTemp()
			generador.Mi_generador.AddExpression(posicion_entorno, "P", strconv.Itoa(size_total), "-")
			generador.Mi_generador.AddExpression(posicion_variable, posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
			generador.Mi_generador.AddSetStack("(int)"+posicion_variable, resultado.Value)
			variable.Is_init = true
		} else {
			panic("Error de tipos al asignar " + d.Id)
		}
	} else if variable.Tipo == valor.BOOLEAN {
		if resultado.Type != valor.BOOLEAN {
			panic("Error de tipos al asignar " + d.Id)
		}
		newLabel := generador.Mi_generador.NewLabel()
		posicion_entorno := generador.Mi_generador.NewTemp()
		//add labels
		for _, lvl := range resultado.TrueLabel {
			generador.Mi_generador.AddLabel(lvl.(string))
		}
		generador.Mi_generador.AddExpression(posicion_entorno, "P", strconv.Itoa(size_total), "-")
		generador.Mi_generador.AddExpression(posicion_variable, posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
		generador.Mi_generador.AddSetStack("(int)"+posicion_variable, "1")
		generador.Mi_generador.AddGoto(newLabel)
		//add labels
		for _, lvl := range resultado.FalseLabel {
			generador.Mi_generador.AddLabel(lvl.(string))
		}
		generador.Mi_generador.AddExpression(posicion_entorno, "P", strconv.Itoa(size_total), "-")
		generador.Mi_generador.AddExpression(posicion_variable, posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
		generador.Mi_generador.AddSetStack("(int)"+posicion_variable, "0")
		generador.Mi_generador.AddGoto(newLabel)
		generador.Mi_generador.AddLabel(newLabel)
		generador.Mi_generador.AddBr()
		variable.Is_init = true
	} else if variable.Tipo == resultado.Type { //tipo primitivos
		//validar antes si es por referencia
		posicion_variable := generador.Mi_generador.NewTemp()
		posicion_entorno := generador.Mi_generador.NewTemp()
		generador.Mi_generador.AddExpression(posicion_entorno, "P", strconv.Itoa(size_total), "-")
		generador.Mi_generador.AddExpression(posicion_variable, posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
		generador.Mi_generador.AddSetStack("(int)"+posicion_variable, resultado.Value)
		variable.Is_init = true
	} else {
		panic("Error de tipo al asignar " + d.Id)
	}
	return result
}
