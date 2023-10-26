package arbol

import (
	"main/ambito"
	"main/valor"
)

type Decremento_atributo struct {
	ID_inicial      Atributo_ide
	Lista_atributos []Atributo_ide
	Expresion       BaseNodo
}

func (d Decremento_atributo) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
