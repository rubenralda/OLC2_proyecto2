package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Incremento_variable struct {
	Id        string
	Expresion BaseNodo
}

func (d Incremento_variable) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	result := valor.Value{Type: valor.NULL}
	variable, size_total := ambito_padre.BuscarVariable(d.Id)
	if variable == nil {
		panic("La variable no existe " + d.Id)
	}
	if variable.Is_constante {
		panic("No se puede modificar una constante")
	}
	if variable.Tipo_dimension != valor.DIMENSION0 {
		panic("La operacion asignacion no esta permitida " + d.Id)
	}
	resultado := d.Expresion.Ejecutar(ambito_padre)
	if variable.Tipo == valor.FLOAT {
		if resultado.Type != valor.INTEGER && resultado.Type != valor.FLOAT {
			panic("Error de tipos al decrementar " + d.Id)
		}
	} else if variable.Tipo == valor.INTEGER {
		if resultado.Type != valor.INTEGER {
			panic("Error de tipos al decrementar " + d.Id)
		}
	} else if variable.Tipo == valor.STRING {
		if resultado.Type != valor.STRING {
			panic("Error de tipos al incrementar" + d.Id)
		}
		tem_valor := generador.Mi_generador.NewTemp()
		true_label := generador.Mi_generador.NewLabel()
		false_label := generador.Mi_generador.NewLabel()
		tmp_sumado := generador.Mi_generador.NewTemp()
		posicion_variable := generador.Mi_generador.NewTemp()
		posicion_entorno := generador.Mi_generador.NewTemp()
		tmp_puntero1 := generador.Mi_generador.NewTemp() //puntero al string en el heap
		//iguala a heap pointer del nuevo string
		generador.Mi_generador.AddAssign(tmp_sumado, "H")
		//obtener el valor del puntero al heap, validar antes si es por referencia aqui
		if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
			generador.Mi_generador.AddAssign(posicion_entorno, "0")
		} else {
			generador.Mi_generador.AddExpression(posicion_entorno, "P", strconv.Itoa(size_total), "-")
		}
		generador.Mi_generador.AddExpression(posicion_variable, posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
		tmp_puntero := generador.Mi_generador.NewTemp() //puntero referencia
		if variable.Is_referencia {
			//obtener el puntero y modificar en esa posicion el valor
			generador.Mi_generador.AddGetStack(tmp_puntero, "(int)"+posicion_variable)
			generador.Mi_generador.AddGetStack(tmp_puntero1, "(int)"+tmp_puntero)
		} else {
			generador.Mi_generador.AddGetStack(tmp_puntero1, "(int)"+posicion_variable)
		}

		// copiar string del tmp_puntero1
		generador.Mi_generador.AddLabel(false_label)                               //label true
		generador.Mi_generador.AddGetHeap(tem_valor, "(int)"+tmp_puntero1)         //valor en el heep
		generador.Mi_generador.AddExpression(tmp_puntero1, tmp_puntero1, "1", "+") //incremento el valor del puntero

		generador.Mi_generador.AddIf(tem_valor, "-1", "==", true_label)
		generador.Mi_generador.AddSetHeap("(int)H", tem_valor)
		generador.Mi_generador.AddExpression("H", "H", "1", "+")
		generador.Mi_generador.AddGoto(false_label)
		generador.Mi_generador.AddLabel(true_label)

		//copiar string del resultado 2
		true_label2 := generador.Mi_generador.NewLabel()                                 //label true
		generador.Mi_generador.AddGetHeap(tem_valor, "(int)"+resultado.Value)            //valor en el heep, resultado es un temporal
		generador.Mi_generador.AddExpression(resultado.Value, resultado.Value, "1", "+") //incremento el valor del puntero

		generador.Mi_generador.AddIf(tem_valor, "-1", "==", true_label2)
		generador.Mi_generador.AddSetHeap("(int)H", tem_valor)
		generador.Mi_generador.AddExpression("H", "H", "1", "+")
		generador.Mi_generador.AddGoto(true_label) //caso que es falso
		generador.Mi_generador.AddLabel(true_label2)

		// delimitador del nuevo string
		generador.Mi_generador.AddSetHeap("(int)H", "-1")
		generador.Mi_generador.AddExpression("H", "H", "1", "+")
		//escribir el nuevo puntero del string en el stack
		if variable.Is_referencia {
			generador.Mi_generador.AddSetStack("(int)"+tmp_puntero, tmp_sumado)
		} else {
			generador.Mi_generador.AddSetStack("(int)"+posicion_variable, tmp_sumado)
		}
		generador.Mi_generador.AddBr()
		return result
	} else {
		panic("Error de tipo al asignar " + d.Id)
	}
	posicion_variable := generador.Mi_generador.NewTemp()
	//validar antes si es por referencia
	posicion_entorno := generador.Mi_generador.NewTemp()
	tmp_valor := generador.Mi_generador.NewTemp()
	tmp_final := generador.Mi_generador.NewTemp()
	if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
		generador.Mi_generador.AddAssign(posicion_entorno, "0")
	} else {
		generador.Mi_generador.AddExpression(posicion_entorno, "P", strconv.Itoa(size_total), "+")
	}
	generador.Mi_generador.AddExpression(posicion_variable, posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
	if variable.Is_referencia {
		tmp_puntero := generador.Mi_generador.NewTemp()
		//obtener el puntero y modificar en esa posicion el valor
		generador.Mi_generador.AddGetStack(tmp_puntero, "(int)"+posicion_variable)
		generador.Mi_generador.AddGetStack(tmp_valor, "(int)"+tmp_puntero)
		generador.Mi_generador.AddExpression(tmp_final, tmp_valor, resultado.Value, "+")
		generador.Mi_generador.AddSetStack("(int)"+tmp_puntero, tmp_final)
	} else {
		generador.Mi_generador.AddGetStack(tmp_valor, "(int)"+posicion_variable)
		generador.Mi_generador.AddExpression(tmp_final, tmp_valor, resultado.Value, "+")
		generador.Mi_generador.AddSetStack("(int)"+posicion_variable, tmp_final)
	}
	variable.Is_init = true
	return result
}
