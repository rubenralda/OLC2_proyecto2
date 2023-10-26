package arbol

import (
	"main/ambito"
	"main/valor"
)

type Asignacion_vector struct {
	Id        string
	Expresion BaseNodo
	Indice    BaseNodo
}

func (a Asignacion_vector) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
