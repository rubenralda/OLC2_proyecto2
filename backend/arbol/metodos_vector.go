package arbol

import (
	"main/ambito"
	"main/valor"
)

type Funcion_append struct {
	Id        string
	Expresion BaseNodo
}

func (f Funcion_append) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Funcion_removeLast struct {
	Id string
}

func (f Funcion_removeLast) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Funcion_removeat struct {
	Id        string
	Expresion BaseNodo
}

func (f Funcion_removeat) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
