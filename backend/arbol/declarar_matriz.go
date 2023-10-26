package arbol

import (
	"main/ambito"
	"main/valor"
)

type Tipo_matriz struct {
	Dimension int
	Primitivo string // String, Int, Bool, Float, char
}

type Declarar_matriz struct {
	Tipo_matriz_ex Tipo_matriz // vacio si no viene
	Id             string
	Matriz         []interface{}
	/*
			[
		    [[37, 49, 61, 29, 44], [56, 60, 51, 68, 70], [47, 15, 39, 17, 74]],
		    [[69, 74, 52, 34, 36], [24, 44, 50, 18, 76], [74, 60, 32, 63, 78]],
		    [[78, 14, 23, 52, 33], [28, 79, 77, 55, 24], [23, 79, 47, 62, 44]],
		    [[73, 53, 11, 49, 52], [29, 16, 65, 34, 12], [72, 69, 30, 44, 37]]
			]
	*/
}

func (d Declarar_matriz) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
