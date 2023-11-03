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
		ambito_padre.Agregar_error("La variable no existe " + d.Id)
		return valor.Value{}
	}
	if variable.Is_constante {
		ambito_padre.Agregar_error("No se puede modificar una constante " + d.Id)
		return valor.Value{}
	}
	if variable.Tipo_dimension != valor.DIMENSION0 {
		ambito_padre.Agregar_error("La operacion asignacion no esta permitida " + d.Id)
		return valor.Value{}
	}
	resultado := d.Expresion.Ejecutar(ambito_padre)
	if variable.Tipo == valor.FLOAT {
		if resultado.Type != valor.INTEGER && resultado.Type != valor.FLOAT {
			ambito_padre.Agregar_error("Error de tipos al decrementar " + d.Id)
			return valor.Value{}
		}
	} else if variable.Tipo == valor.INTEGER {
		if resultado.Type != valor.INTEGER {
			ambito_padre.Agregar_error("Error de tipos al decrementar " + d.Id)
			return valor.Value{}
		}
	} else {
		ambito_padre.Agregar_error("Error de tipo al asignar " + d.Id)
		return valor.Value{}
	}
	posicion_variable := generador.Mi_generador.NewTemp()
	//validar antes si es por referencia
	posicion_entorno := generador.Mi_generador.NewTemp()
	tmp_valor := generador.Mi_generador.NewTemp()
	tmp_final := generador.Mi_generador.NewTemp()
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
		generador.Mi_generador.AddGetStack(tmp_valor, "(int)"+tmp_puntero)
		generador.Mi_generador.AddExpression(tmp_final, tmp_valor, resultado.Value, "-")
		generador.Mi_generador.AddSetStack("(int)"+tmp_puntero, tmp_final)
	} else {
		generador.Mi_generador.AddGetStack(tmp_valor, "(int)"+posicion_variable)
		generador.Mi_generador.AddExpression(tmp_final, tmp_valor, resultado.Value, "-")
		generador.Mi_generador.AddSetStack("(int)"+posicion_variable, tmp_final)
	}
	variable.Is_init = true
	return result
}
