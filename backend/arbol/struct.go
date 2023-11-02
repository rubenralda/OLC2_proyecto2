package arbol

import (
	"main/ambito"
	"main/valor"
)

type Declarar_struct struct {
	Id        string
	Atributos []Declaracion_atributo
}

func (a Declarar_struct) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	var atributos []*ambito.Variables
	for _, atributo := range a.Atributos {
		atributos = append(atributos, atributo.Ejecutar(ambito_padre))
	}
	ambito_padre.Agregar_struct(&ambito.Estruct{Nombre: a.Id, Atributos: atributos})
	return valor.Value{}
}

type Declaracion_atributo struct {
	Id        string
	Primitivo string   // String, Int, Bool, Float, char, (Nombre_struct)
	Expresion BaseNodo // no forma parte del proyecto 2 (compilacion)
	Tipo      string   // constante o variable
}

func (a Declaracion_atributo) Ejecutar(ambito_padre *ambito.Ambito) *ambito.Variables {
	var variable ambito.Variables
	variable.Id = a.Id
	variable.Tipo = Tipo_variable[a.Primitivo]
	if variable.Tipo == valor.NULL {
		existe, _ := ambito_padre.Buscar_struct(a.Primitivo)
		if existe == nil {
			panic("El struct " + a.Primitivo + " no esta definido")
		}
		variable.Is_instancia = true
		variable.Tipo_struct = a.Primitivo
	}
	if a.Tipo == "constante" {
		variable.Is_constante = true //no se evalua en el proyecto
	}
	return &variable
}
