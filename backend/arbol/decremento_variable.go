package arbol

import (
	"main/ambito"
	"main/valor"
)

type Decremento_variable struct {
	Id        string
	Expresion BaseNodo
}

func (d Decremento_variable) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
