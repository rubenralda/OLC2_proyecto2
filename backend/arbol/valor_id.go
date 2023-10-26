package arbol

import (
	"main/ambito"
	"main/valor"
)

type Id_expresion struct {
	Id string
}

func (a Id_expresion) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Id_vector struct {
	Id     string
	Indice BaseNodo
}

func (a Id_vector) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Id_matriz struct {
	Id      string
	Indices []BaseNodo
}

func (a Id_matriz) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
