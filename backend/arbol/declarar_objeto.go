package arbol

import (
	"main/ambito"
	"main/valor"
)

type Declarar_objeto struct {
	Id    string
	Dupla []Dupla_atributos
}

func (d Declarar_objeto) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Dupla_atributos struct {
	Id_externo string
	Referencia bool // no se usa pero esta para convertir de Lista_argumentos
	Expresion  BaseNodo
}
