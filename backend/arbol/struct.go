package arbol

import (
	"main/ambito"
	"main/valor"
)

type Declarar_struct struct {
	Id        string
	Atributos []BaseNodo
}

func (a Declarar_struct) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Declaracion_atributo struct {
	Id        string
	Primitivo string   // String, Int, Bool, Float, char, (Nombre_struct)
	Expresion BaseNodo // 10, 20.5, "hola", true, objeto_strcut
	Tipo      string   // constante o variable
}

func (a Declaracion_atributo) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
