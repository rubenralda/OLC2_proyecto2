package arbol

import (
	"main/ambito"
	"main/valor"
)

type Decremento_vector struct {
	Id        string
	Expresion BaseNodo
	Indice    BaseNodo
}

func (d Decremento_vector) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
