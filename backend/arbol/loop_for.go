package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Loop_for_in struct {
	Expresion  BaseNodo
	Id         string
	Inicio     BaseNodo
	Final      BaseNodo
	Sentencias []BaseNodo
}

func (s Loop_for_in) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	ambito_local := &ambito.Ambito{NombreAmbito: "sentencia for", Padre: ambito_padre}
	ambito_padre.AgregarAmbito(ambito_local)

	if s.Expresion == nil {
		inicio := s.Inicio.Ejecutar(ambito_padre)
		final := s.Final.Ejecutar(ambito_padre)
		if inicio.Type != valor.INTEGER && final.Type != valor.INTEGER {
			panic("El rango no es entero")
		}
		label_final := generador.Mi_generador.NewLabel()
		label_error := generador.Mi_generador.NewLabel()
		generador.Mi_generador.AddIf(inicio.Value, final.Value, ">", label_error) //si es mayor el inicio error
		tmp_contador := generador.Mi_generador.NewTemp()
		// muevo el puntero al nuevo ambito
		generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(ambito_padre.Size), "+")
		generador.Mi_generador.AddExpression(tmp_contador, inicio.Value, "1", "-") //porque empiezo sumando
		//agregar la variable
		ambito_local.AgregarVariable(ambito.Variables{Id: s.Id, Posicion_relativa: ambito_local.Size, Tipo: valor.INTEGER, Is_init: true, Is_constante: true})
		generador.Mi_generador.AddSetStack("(int)P", tmp_contador)
		ambito_local.Size++

		label_inicio := generador.Mi_generador.NewLabel()
		generador.Mi_generador.AddLabel(label_inicio)

		generador.Mi_generador.AddGetStack(tmp_contador, "(int)P")
		generador.Mi_generador.AddExpression(tmp_contador, tmp_contador, "1", "+")
		generador.Mi_generador.AddSetStack("(int)P", tmp_contador)

		generador.Mi_generador.AddIf(tmp_contador, final.Value, ">", label_final)
		//ejecutar las sentencias
		for _, linea := range s.Sentencias {
			linea.Ejecutar(ambito_local)
		}
		generador.Mi_generador.AddGoto(label_inicio)
		generador.Mi_generador.AddLabel(label_final)
		generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(ambito_padre.Size), "-")
		generador.Mi_generador.AddLabel(label_error)
	} else {
		resultado := s.Expresion.Ejecutar(ambito_padre)
		if resultado.Type == valor.STRING {
			generador.Mi_generador.AddComment("chucho")
			label_final := generador.Mi_generador.NewLabel()
			tmp_puntero := generador.Mi_generador.NewTemp()
			tmp_valor := generador.Mi_generador.NewTemp()
			tmp_posicion := generador.Mi_generador.NewTemp()
			// muevo el puntero al nuevo ambito
			generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(ambito_padre.Size), "+")
			generador.Mi_generador.AddExpression(tmp_puntero, resultado.Value, "1", "-") //porque empiezo sumando
			//agregar temporal puntero al stack
			generador.Mi_generador.AddSetStack("(int)P", tmp_puntero)
			ambito_local.Size++
			//agregar la variable
			variable := ambito.Variables{Id: s.Id, Posicion_relativa: ambito_local.Size, Tipo: valor.CHAR, Is_init: true, Is_constante: true}
			ambito_local.AgregarVariable(variable)
			generador.Mi_generador.AddExpression(tmp_posicion, "P", strconv.Itoa(variable.Posicion_relativa), "+")
			generador.Mi_generador.AddGetHeap(tmp_valor, "(int)"+resultado.Value)
			generador.Mi_generador.AddSetStack("(int)"+tmp_posicion, tmp_valor) //valor del char
			ambito_local.Size++

			label_inicio := generador.Mi_generador.NewLabel()
			generador.Mi_generador.AddLabel(label_inicio)

			generador.Mi_generador.AddGetStack(tmp_puntero, "(int)P") //recupero
			generador.Mi_generador.AddExpression(tmp_puntero, tmp_puntero, "1", "+")
			generador.Mi_generador.AddSetStack("(int)P", tmp_puntero) //posicionado

			//obtener valor char del heap
			generador.Mi_generador.AddExpression(tmp_posicion, "P", strconv.Itoa(variable.Posicion_relativa), "+")
			generador.Mi_generador.AddGetHeap(tmp_valor, "(int)"+tmp_puntero)
			generador.Mi_generador.AddSetStack("(int)"+tmp_posicion, tmp_valor) //valor del char

			generador.Mi_generador.AddIf(tmp_valor, "-1", "==", label_final)

			//ejecutar las sentencias
			for _, linea := range s.Sentencias {
				linea.Ejecutar(ambito_local)
			}
			generador.Mi_generador.AddGoto(label_inicio)
			generador.Mi_generador.AddLabel(label_final)
			generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(ambito_padre.Size), "-")
		} else if resultado.Tipo_dimension == valor.DIMENSION1 {
			//logica para recorrer un vector
		} else {
			panic("Se esperaba un string o vector en la expresion")
		}
	}
	return valor.Value{}
}
