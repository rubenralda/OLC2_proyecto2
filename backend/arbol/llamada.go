package arbol

import (
	"main/ambito"
	"main/valor"
)

type Llamada_funcion struct {
	Id                  string
	Lista_argumentos    []Lista_argumentos
	Declarar_objeto_amb Declarar_objeto
}

func (a Llamada_funcion) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Lista_argumentos struct {
	Id_externo string
	Referencia bool // cuando se usa & verificar que sea id
	Expresion  BaseNodo
}
