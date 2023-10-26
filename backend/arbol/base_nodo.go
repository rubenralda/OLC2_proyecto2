package arbol

import (
	"main/ambito"
	"main/valor"
)

type BaseNodo interface {
	Ejecutar(ambito *ambito.Ambito) valor.Value
}
