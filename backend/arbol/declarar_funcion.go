package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
)

type Declarar_funcion struct {
	Id               string
	Tipo_retorno     string
	Sentencias       []BaseNodo
	Lista_parametros []Lista_parametros

	Mutable bool //no se usa en el proyecto 2
}

func (d Declarar_funcion) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	ambito_local := &ambito.Ambito{NombreAmbito: "Funcion " + d.Id, Padre: ambito_padre}
	funcion := &ambito.Funciones{Nombre: d.Id, Tipo: Tipo_variable[d.Tipo_retorno]}
	ambito_padre.AgregarAmbito(ambito_local)
	ambito_padre.AgregarFuncion(funcion)
	ambito_padre.Modo_funcion = true

	label_return := generador.Mi_generador.NewLabel()
	ambito_padre.Label_salida = label_return
	ambito_padre.Tipo_funcion = Tipo_variable[d.Tipo_retorno]
	generador.Mi_generador.SetMainFlag(false)
	generador.Mi_generador.AddTittle(d.Id)
	ambito_local.Size++ //reservar el espacio de retorno

	//declarar parametros como variables
	for _, parametro := range d.Lista_parametros {
		variable := ambito.Variables{Id: parametro.Id_interno, Posicion_relativa: ambito_local.Size, Tipo_dimension: valor.DIMENSION0, Is_init: true}
		variable.Tipo = Tipo_variable[parametro.Primitivo]
		if variable.Tipo == valor.NULL {
			//buscar el struct y darle el tipo falta
			variable.Tipo_struct = parametro.Primitivo
			variable.Is_instancia = true
		}
		if !parametro.Referencia { //se declara como constante
			variable.Is_constante = true
		} else { //si es por referencia se activa
			variable.Is_referencia = true
		}
		// falta si es un vector o matriz
		ambito_local.AgregarVariable(variable)
		ambito_local.Size++
		//agregar nombres externo e interno
		funcion.Parametros = append(funcion.Parametros, ambito.Parametros{Id_interno: parametro.Id_interno, Id_externo: parametro.Id_externo, Tipo: variable.Tipo, Tipo_struct: variable.Tipo_struct})
	}
	for _, linea := range d.Sentencias {
		linea.Ejecutar(ambito_local)
	}
	ambito_padre.Modo_funcion = false
	generador.Mi_generador.AddLabel(label_return)
	generador.Mi_generador.AddEnd()
	generador.Mi_generador.SetMainFlag(true)
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
