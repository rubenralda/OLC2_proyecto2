package arbol

import "main/ambito"

type BaseNodo interface {
	Ejecutar(ambito *ambito.Ambito) interface{}
}
