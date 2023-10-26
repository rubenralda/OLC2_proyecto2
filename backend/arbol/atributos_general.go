package arbol

import (
	"main/ambito"
	"main/valor"
)

type Atributo_ide struct {
	ID     string
	Indice BaseNodo
	Vector bool
}

type Atributo_general struct {
	ID_inicial      Atributo_ide
	Lista_atributos []Atributo_ide
}

func (a Atributo_general) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
