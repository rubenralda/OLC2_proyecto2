package arbol

import (
	"main/ambito"
	"main/valor"
)

type Incremento_atributo struct {
	ID_inicial      Atributo_ide
	Lista_atributos []Atributo_ide
	Expresion       BaseNodo
}

func (i Incremento_atributo) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
