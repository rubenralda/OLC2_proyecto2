package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

var Tipo_variable = map[string]valor.TipoExpresion{
	"Int": valor.INTEGER, "Float": valor.FLOAT, "String": valor.STRING, "Character": valor.CHAR, "Bool": valor.BOOLEAN}

type Declarar_variable struct {
	Id        string
	Tipo      string   // String, Int, Bool, Float, char
	Expresion BaseNodo // 10, 20.5, "hola", true, objeto_strcut
}

func (d Declarar_variable) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	result := valor.Value{Type: valor.NULL}
	if ambito_padre.BuscarVariable(d.Id) != nil {
		panic("Error la variable ya existe: " + d.Id)
	}
	variable := ambito.Variables{Id: d.Id, Posicion_relativa: ambito_padre.Size, Tipo_dimension: valor.DIMENSION0}
	if d.Expresion == nil {
		if d.Tipo != "" {
			variable.Tipo = Tipo_variable[d.Tipo]
			if variable.Tipo == valor.NULL {
				variable.Tipo_struct = d.Tipo
				variable.Is_instancia = true
			}
			variable.Is_init = false
			ambito_padre.AgregarVariable(variable)
			ambito_padre.Size++
			return result
		} else {
			panic("Error falta tipo al declarar variable")
		}
	}
	resultado := d.Expresion.Ejecutar(ambito_padre)
	if resultado.Type == valor.BOOLEAN {
		if d.Tipo != "Bool" && d.Tipo != "" {
			panic("Error tipos no coinciden")
		}
		//si no es temp (boolean)
		newLabel := generador.Mi_generador.NewLabel()
		//add labels
		for _, lvl := range resultado.TrueLabel {
			generador.Mi_generador.AddLabel(lvl.(string))
		}
		generador.Mi_generador.AddSetStack(strconv.Itoa(variable.Posicion_relativa), "1")
		generador.Mi_generador.AddGoto(newLabel)
		//add labels
		for _, lvl := range resultado.FalseLabel {
			generador.Mi_generador.AddLabel(lvl.(string))
		}
		generador.Mi_generador.AddSetStack(strconv.Itoa(variable.Posicion_relativa), "0")
		generador.Mi_generador.AddGoto(newLabel)
		generador.Mi_generador.AddLabel(newLabel)
		generador.Mi_generador.AddBr()
		variable.Tipo = resultado.Type
	} else if resultado.Type == valor.NULL {
		panic("Error")
	} else {
		variable.Tipo = Tipo_variable[d.Tipo]
		if d.Tipo == "" { //no viene con tipo
			if resultado.Is_intancia {
				variable.Is_instancia = true
				variable.Tipo_struct = resultado.Tipo_struct
			} else {
				variable.Is_instancia = false
				variable.Tipo = resultado.Type
			}
		} else if variable.Tipo == valor.NULL { //el tipo de una clase
			if resultado.Tipo_struct == d.Tipo {
				variable.Is_instancia = true
				variable.Tipo_struct = result.Tipo_struct
			} else {
				panic("Error tipos no coinciden")
			}
		} else if variable.Tipo != resultado.Type {
			panic("Error el tipo de la expresion y la variable no coinciden")
		}
		generador.Mi_generador.AddSetStack(strconv.Itoa(variable.Posicion_relativa), resultado.Value)
		generador.Mi_generador.AddBr()
	}
	variable.Is_init = true
	ambito_padre.AgregarVariable(variable)
	ambito_padre.Size++
	return result
}
