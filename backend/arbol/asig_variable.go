package arbol

import (
	"main/ambito"
	"main/valor"
)

type Asignacion_variable struct {
	Id        string
	Expresion BaseNodo
}

func (a Asignacion_variable) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
