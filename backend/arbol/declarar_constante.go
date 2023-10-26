package arbol

import (
	"main/ambito"
	"main/valor"
)

// declarar constante

type Declarar_constante struct {
	Id        string
	Tipo      string   // String, Int, Bool, Float, char, (Nombre_struct)
	Expresion BaseNodo // 10, 20.5, "hola", true, objeto_strcut
}

func (d Declarar_constante) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
