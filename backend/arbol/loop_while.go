package arbol

import (
	"main/ambito"
	"main/valor"
)

type Loop_while struct {
	Expresion  BaseNodo
	Sentencias []BaseNodo
}

func (a Loop_while) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
