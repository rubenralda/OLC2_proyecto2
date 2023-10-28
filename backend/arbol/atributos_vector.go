package arbol

import (
	"main/ambito"
	"main/valor"
)

type Vector_isempty struct {
	Id string
}

func (v Vector_isempty) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Vector_count struct {
	Id string
}

func (v Vector_count) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
