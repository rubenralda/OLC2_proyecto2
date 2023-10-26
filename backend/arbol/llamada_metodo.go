package arbol

import (
	"main/ambito"
	"main/valor"
)

type Llamada_metodo struct {
	Id_metodo        string
	Id_objeto        string
	Lista_argumentos []Lista_argumentos
}

func (a Llamada_metodo) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
