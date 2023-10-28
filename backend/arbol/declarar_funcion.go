package arbol

import (
	"main/ambito"
	"main/valor"
)

type Ejecutar_funcion struct {
	Id               string
	Tipo_retorno     string
	Sentencias       []BaseNodo
	Lista_parametros []Lista_parametros
	Mutable          bool
}

func (e Ejecutar_funcion) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Declarar_funcion struct {
	Id               string
	Tipo_retorno     string
	Sentencias       []BaseNodo
	Lista_parametros []Lista_parametros
	Mutable          bool
}

func (d Declarar_funcion) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Lista_parametros struct {
	Id_interno string
	Id_externo string
	Referencia bool
	Primitivo  string // String, Int, Bool, Float, char, (Nombre_struct)
	Vector     bool   //true si es vector
	Matriz     bool
}
