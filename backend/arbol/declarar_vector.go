package arbol

import (
	"main/ambito"
	"main/valor"
)

type Declarar_vector struct {
	Tipo            string // String, Int, Bool, Float, char, (Nombre_struct)
	ID              string
	Lista_expresion []BaseNodo // 10, 20.5, "hola", true, objeto_strcut
	ID_otro_vector  string
}

func (d Declarar_vector) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
