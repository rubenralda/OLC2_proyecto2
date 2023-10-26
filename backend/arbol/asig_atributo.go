package arbol

import (
	"main/ambito"
	"main/valor"
)

type Asignar_atributos struct {
	ID_inicial      Atributo_ide
	Lista_atributos []Atributo_ide
	Expresion       BaseNodo
}

func (a Asignar_atributos) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
