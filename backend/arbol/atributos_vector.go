package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Vector_isempty struct {
	Id string
}

func (a Vector_isempty) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	variable, size := ambito_padre.BuscarVariable(a.Id)
	if variable == nil {
		panic("La variable no existe " + a.Id)
	}
	if variable.Tipo_dimension != valor.DIMENSION1 {
		panic("La variable no es un vector " + a.Id)
	}
	puntero_variable := generador.Mi_generador.NewTemp()
	puntero_heap := generador.Mi_generador.NewTemp()
	tmp_posicion_entorno := generador.Mi_generador.NewTemp()
	true_label := generador.Mi_generador.NewLabel()
	false_label := generador.Mi_generador.NewLabel()
	if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
		generador.Mi_generador.AddAssign(tmp_posicion_entorno, "0")
	} else {
		generador.Mi_generador.AddExpression(tmp_posicion_entorno, "P", strconv.Itoa(size), "-")
	}
	generador.Mi_generador.AddExpression(puntero_variable, tmp_posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
	generador.Mi_generador.AddGetStack(puntero_heap, "(int)"+puntero_variable)

	generador.Mi_generador.AddExpression(puntero_heap, puntero_heap, "1", "+") //puntero del siguiente elemento
	generador.Mi_generador.AddGetHeap(puntero_heap, "(int)"+puntero_heap)
	generador.Mi_generador.AddIf(puntero_heap, "-1", "==", true_label) //si es -1 entonces esta vacio
	generador.Mi_generador.AddGoto(false_label)
	resultado := valor.Value{Type: valor.BOOLEAN}
	resultado.TrueLabel = append(resultado.TrueLabel, true_label)
	resultado.FalseLabel = append(resultado.FalseLabel, false_label)
	return resultado
}

type Vector_count struct {
	Id string
}

func (a Vector_count) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	variable, size := ambito_padre.BuscarVariable(a.Id)
	if variable == nil {
		panic("La variable no existe " + a.Id)
	}
	if variable.Tipo_dimension != valor.DIMENSION1 {
		panic("La variable no es un vector " + a.Id)
	}
	puntero_variable := generador.Mi_generador.NewTemp()
	puntero_heap := generador.Mi_generador.NewTemp()
	tmp_posicion_entorno := generador.Mi_generador.NewTemp()
	label_inicio := generador.Mi_generador.NewLabel()
	label_salida := generador.Mi_generador.NewLabel()
	contador := generador.Mi_generador.NewTemp()
	if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
		generador.Mi_generador.AddAssign(tmp_posicion_entorno, "0")
	} else {
		generador.Mi_generador.AddExpression(tmp_posicion_entorno, "P", strconv.Itoa(size), "-")
	}
	generador.Mi_generador.AddExpression(puntero_variable, tmp_posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
	generador.Mi_generador.AddGetStack(puntero_heap, "(int)"+puntero_variable)

	generador.Mi_generador.AddAssign(contador, "0")
	generador.Mi_generador.AddLabel(label_inicio)
	generador.Mi_generador.AddExpression(puntero_heap, puntero_heap, "1", "+") //puntero del siguiente elemento
	generador.Mi_generador.AddGetHeap(puntero_heap, "(int)"+puntero_heap)
	generador.Mi_generador.AddIf(puntero_heap, "-1", "==", label_salida) //si es -1 no incremento el contador
	generador.Mi_generador.AddExpression(contador, contador, "1", "+")
	generador.Mi_generador.AddGoto(label_inicio)

	generador.Mi_generador.AddLabel(label_salida)
	return valor.Value{Value: contador, Type: valor.INTEGER}
}
