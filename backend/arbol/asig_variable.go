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
		ambito_padre.Agregar_error("La variable no existe " + d.Id)
		return valor.Value{}
	}
	if variable.Is_constante {
		ambito_padre.Agregar_error("No se puede modificar una constante")
		return valor.Value{}
	}
	if variable.Tipo_dimension != valor.DIMENSION0 {
		ambito_padre.Agregar_error("La operacion asignacion no esta permitida " + d.Id)
		return valor.Value{}
	}
	posicion_variable := generador.Mi_generador.NewTemp()
	resultado := d.Expresion.Ejecutar(ambito_padre)
	if variable.Is_instancia {
		if variable.Tipo_struct == resultado.Tipo_struct {
			//logica para cambiar la posicion relativa a la nueva del objeto struct en el heep
		} else {
			ambito_padre.Agregar_error("Tipo no permitivo")
			return valor.Value{}
		}
	} else if variable.Tipo == valor.FLOAT {
		if resultado.Type == valor.INTEGER || resultado.Type == valor.FLOAT {
			posicion_entorno := generador.Mi_generador.NewTemp()
			if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
				generador.Mi_generador.AddAssign(posicion_entorno, "0")
			} else {
				generador.Mi_generador.AddExpression(posicion_entorno, "P", strconv.Itoa(size_total), "-")
			}
			generador.Mi_generador.AddExpression(posicion_variable, posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
			if variable.Is_referencia {
				tmp_puntero := generador.Mi_generador.NewTemp()
				//obtener el puntero y modificar en esa posicion el valor
				generador.Mi_generador.AddGetStack(tmp_puntero, "(int)"+posicion_variable)
				generador.Mi_generador.AddSetStack("(int)"+tmp_puntero, resultado.Value)
			} else {
				generador.Mi_generador.AddSetStack("(int)"+posicion_variable, resultado.Value)
			}
			variable.Is_init = true
		} else {
			ambito_padre.Agregar_error("Error de tipos al asignar " + d.Id)
			return valor.Value{}
		}
	} else if variable.Tipo == valor.BOOLEAN {
		if resultado.Type != valor.BOOLEAN {
			ambito_padre.Agregar_error("Error de tipos al asignar " + d.Id)
			return valor.Value{}
		}
		newLabel := generador.Mi_generador.NewLabel()
		posicion_entorno := generador.Mi_generador.NewTemp()
		//add labels
		for _, lvl := range resultado.TrueLabel {
			generador.Mi_generador.AddLabel(lvl.(string))
		}
		if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
			generador.Mi_generador.AddAssign(posicion_entorno, "0")
		} else {
			generador.Mi_generador.AddExpression(posicion_entorno, "P", strconv.Itoa(size_total), "-")
		}
		generador.Mi_generador.AddExpression(posicion_variable, posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
		if variable.Is_referencia {
			tmp_puntero := generador.Mi_generador.NewTemp()
			//obtener el puntero y modificar en esa posicion el valor
			generador.Mi_generador.AddGetStack(tmp_puntero, "(int)"+posicion_variable)
			generador.Mi_generador.AddSetStack("(int)"+tmp_puntero, "1")
		} else {
			generador.Mi_generador.AddSetStack("(int)"+posicion_variable, "1")
		}
		generador.Mi_generador.AddGoto(newLabel)
		//add labels
		for _, lvl := range resultado.FalseLabel {
			generador.Mi_generador.AddLabel(lvl.(string))
		}
		if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
			generador.Mi_generador.AddAssign(posicion_entorno, "0")
		} else {
			generador.Mi_generador.AddExpression(posicion_entorno, "P", strconv.Itoa(size_total), "-")
		}
		generador.Mi_generador.AddExpression(posicion_variable, posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
		if variable.Is_referencia {
			tmp_puntero := generador.Mi_generador.NewTemp()
			//obtener el puntero y modificar en esa posicion el valor
			generador.Mi_generador.AddGetStack(tmp_puntero, "(int)"+posicion_variable)
			generador.Mi_generador.AddSetStack("(int)"+tmp_puntero, "0")
		} else {
			generador.Mi_generador.AddSetStack("(int)"+posicion_variable, "0")
		}
		generador.Mi_generador.AddGoto(newLabel)
		generador.Mi_generador.AddLabel(newLabel)
		generador.Mi_generador.AddBr()
		variable.Is_init = true
	} else if variable.Tipo == resultado.Type { //tipo primitivos
		//validar antes si es por referencia
		posicion_entorno := generador.Mi_generador.NewTemp()
		if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
			generador.Mi_generador.AddAssign(posicion_entorno, "0")
		} else {
			generador.Mi_generador.AddExpression(posicion_entorno, "P", strconv.Itoa(size_total), "-")
		}
		generador.Mi_generador.AddExpression(posicion_variable, posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
		if variable.Is_referencia {
			tmp_puntero := generador.Mi_generador.NewTemp()
			//obtener el puntero y modificar en esa posicion el valor
			generador.Mi_generador.AddGetStack(tmp_puntero, "(int)"+posicion_variable)
			generador.Mi_generador.AddSetStack("(int)"+tmp_puntero, resultado.Value)
		} else {
			generador.Mi_generador.AddSetStack("(int)"+posicion_variable, resultado.Value)
		}
		variable.Is_init = true
	} else {
		ambito_padre.Agregar_error("Error de tipos al asignar " + d.Id)
		return valor.Value{}
	}
	return result
}
