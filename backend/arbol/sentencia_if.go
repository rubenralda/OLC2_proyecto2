package arbol

import (
	"main/ambito"
	"main/valor"
)

type Sentencia_if struct {
	Expresion       BaseNodo
	Sentencias      []BaseNodo
	Sentencias_else []BaseNodo
	Siguiente       BaseNodo
}

func (a Sentencia_if) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
