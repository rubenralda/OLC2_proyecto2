package arbol

import (
	"main/ambito"
	"main/valor"
)

type Sentencia_switch struct {
	Expresion    BaseNodo
	Lista_case   []Sentencia_case
	Default_case BaseNodo
}

func (a Sentencia_switch) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Sentencia_case struct {
	Expresion  BaseNodo
	Sentencias []BaseNodo
}

func (a Sentencia_case) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Default_case struct {
	Sentencias []BaseNodo
}

func (a Default_case) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
