package arbol

import (
	"main/ambito"
	"main/valor"
)

type Incremento_vector struct {
	Id        string
	Expresion BaseNodo
	Indice    BaseNodo
}

func (i Incremento_vector) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
