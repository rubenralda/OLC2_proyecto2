package arbol

import (
	"main/ambito"
	"main/valor"
)

type Incremento_variable struct {
	Id        string
	Expresion BaseNodo
}

func (i Incremento_variable) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
