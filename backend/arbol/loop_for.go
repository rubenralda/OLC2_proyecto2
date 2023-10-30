package arbol

import (
	"main/ambito"
	"main/valor"
)

type Loop_for_in struct {
	Expresion  BaseNodo
	Id         string
	Inicio     BaseNodo
	Final      BaseNodo
	Sentencias []BaseNodo
}

func (s Loop_for_in) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
