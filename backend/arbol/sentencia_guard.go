package arbol

import (
	"main/ambito"
	"main/valor"
)

type Sentencia_guard struct {
	Expresion  BaseNodo
	Sentencias []BaseNodo
}

func (a Sentencia_guard) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
