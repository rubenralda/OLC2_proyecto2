package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Funcion_append struct {
	Id        string
	Expresion BaseNodo
}

func (f Funcion_append) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	variable, size := ambito_padre.BuscarVariable(f.Id)
	if variable == nil {
		panic("La variable no existe " + f.Id)
	}
	if variable.Tipo_dimension != valor.DIMENSION1 {
		panic("La variable no es un vector")
	}
	item := f.Expresion.Ejecutar(ambito_padre)
	if variable.Is_instancia {
		if variable.Tipo_struct != item.Tipo_struct {
			panic("El argumento no es del mismo tipo de la variable " + f.Id)
		}
	} else if variable.Tipo == valor.BOOLEAN {
		if item.Type != valor.BOOLEAN {
			panic("El argumento no es del mismo tipo de la variable " + f.Id)
		}
	} else if variable.Tipo != item.Type {
		panic("El argumento no es del mismo tipo de la variable " + f.Id)
	}
	puntero_variable := generador.Mi_generador.NewTemp()
	puntero_heap := generador.Mi_generador.NewTemp()
	tmp_posicion_entorno := generador.Mi_generador.NewTemp()
	valor_heap := generador.Mi_generador.NewTemp()
	label_inicio := generador.Mi_generador.NewLabel()
	//label_error := generador.Mi_generador.NewLabel()
	label_salida := generador.Mi_generador.NewLabel()
	if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
		generador.Mi_generador.AddAssign(tmp_posicion_entorno, "0")
	} else {
		generador.Mi_generador.AddExpression(tmp_posicion_entorno, "P", strconv.Itoa(size), "-")
	}
	//generador.Mi_generador.AddComment("appendvector")
	generador.Mi_generador.AddExpression(puntero_variable, tmp_posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
	generador.Mi_generador.AddGetStack(puntero_heap, "(int)"+puntero_variable)

	generador.Mi_generador.AddLabel(label_inicio)
	generador.Mi_generador.AddExpression(puntero_heap, puntero_heap, "1", "+") //puntero del siguiente elemento
	generador.Mi_generador.AddAssign(valor_heap, puntero_heap)
	generador.Mi_generador.AddGetHeap(puntero_heap, "(int)"+puntero_heap)
	generador.Mi_generador.AddIf(puntero_heap, "-1", "==", label_salida) //si es -1 estoy en el lugar para escribir
	generador.Mi_generador.AddGoto(label_inicio)

	generador.Mi_generador.AddLabel(label_salida)
	generador.Mi_generador.AddSetHeap("(int)"+valor_heap, "H")
	generador.Mi_generador.AddExpression(valor_heap, valor_heap, "1", "-") //posicion del nuevo item
	generador.Mi_generador.AddSetHeap("(int)"+valor_heap, item.Value)

	generador.Mi_generador.AddExpression("H", "H", "1", "+") //reservo espacio valor
	generador.Mi_generador.AddSetHeap("(int)H", "-1")
	generador.Mi_generador.AddExpression("H", "H", "1", "+") //reservo espacio puntero
	return valor.Value{}
}

type Funcion_removeLast struct {
	Id string
}

func (f Funcion_removeLast) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	variable, size := ambito_padre.BuscarVariable(f.Id)
	if variable == nil {
		panic("La variable no existe " + f.Id)
	}
	if variable.Tipo_dimension != valor.DIMENSION1 {
		panic("La variable no es un vector")
	}
	puntero_variable := generador.Mi_generador.NewTemp()
	puntero_heap := generador.Mi_generador.NewTemp()
	tmp_posicion_entorno := generador.Mi_generador.NewTemp()
	label_inicio := generador.Mi_generador.NewLabel()
	tmp1 := generador.Mi_generador.NewTemp()
	tmp2 := generador.Mi_generador.NewTemp()
	tmp3 := generador.Mi_generador.NewTemp()
	label_salida := generador.Mi_generador.NewLabel()
	label_eliminar := generador.Mi_generador.NewLabel()
	if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
		generador.Mi_generador.AddAssign(tmp_posicion_entorno, "0")
	} else {
		generador.Mi_generador.AddExpression(tmp_posicion_entorno, "P", strconv.Itoa(size), "-")
	}
	generador.Mi_generador.AddExpression(puntero_variable, tmp_posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
	generador.Mi_generador.AddGetStack(puntero_heap, "(int)"+puntero_variable)

	generador.Mi_generador.AddLabel(label_inicio)
	generador.Mi_generador.AddExpression(tmp1, puntero_heap, "1", "+") //puntero del siguiente elemento
	generador.Mi_generador.AddGetHeap(tmp2, "(int)"+tmp1)
	generador.Mi_generador.AddIf(tmp2, "-1", "==", label_salida) //si es -1 estoy en el lugar para escribir
	generador.Mi_generador.AddExpression(tmp1, tmp2, "1", "+")
	generador.Mi_generador.AddGetHeap(tmp3, "(int)"+tmp1)
	generador.Mi_generador.AddIf(tmp3, "-1", "==", label_eliminar)
	generador.Mi_generador.AddAssign(puntero_heap, tmp2)
	generador.Mi_generador.AddGoto(label_inicio)

	generador.Mi_generador.AddLabel(label_eliminar)
	generador.Mi_generador.AddSetHeap("(int)"+puntero_heap, "0")
	generador.Mi_generador.AddExpression(puntero_heap, puntero_heap, "1", "+")
	generador.Mi_generador.AddSetHeap("(int)"+puntero_heap, "-1")

	generador.Mi_generador.AddLabel(label_salida)
	return valor.Value{}
}

type Funcion_removeat struct {
	Id        string
	Expresion BaseNodo
}

func (f Funcion_removeat) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	variable, size := ambito_padre.BuscarVariable(f.Id)
	if variable == nil {
		panic("La variable no existe " + f.Id)
	}
	if variable.Tipo_dimension != valor.DIMENSION1 {
		panic("La variable no es un vector")
	}
	indice := f.Expresion.Ejecutar(ambito_padre)
	if indice.Type != valor.INTEGER {
		panic("El indice no es entero")
	}
	tmp_indice := generador.Mi_generador.NewTemp()
	generador.Mi_generador.AddAssign(tmp_indice, indice.Value)
	puntero_variable := generador.Mi_generador.NewTemp()
	puntero_heap := generador.Mi_generador.NewTemp()
	contador := generador.Mi_generador.NewTemp()
	tmp_posicion_entorno := generador.Mi_generador.NewTemp()
	label_inicio := generador.Mi_generador.NewLabel()
	tmp1 := generador.Mi_generador.NewTemp()
	tmp2 := generador.Mi_generador.NewTemp()
	tmp3 := generador.Mi_generador.NewTemp()
	label_salida := generador.Mi_generador.NewLabel()
	label_eliminar := generador.Mi_generador.NewLabel()
	label_error := generador.Mi_generador.NewLabel()
	if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
		generador.Mi_generador.AddAssign(tmp_posicion_entorno, "0")
	} else {
		generador.Mi_generador.AddExpression(tmp_posicion_entorno, "P", strconv.Itoa(size), "-")
	}
	generador.Mi_generador.AddComment("removelastvector")
	generador.Mi_generador.AddExpression(puntero_variable, tmp_posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
	generador.Mi_generador.AddGetStack(puntero_heap, "(int)"+puntero_variable)

	generador.Mi_generador.AddAssign(contador, "0")
	generador.Mi_generador.AddLabel(label_inicio)
	generador.Mi_generador.AddExpression(tmp1, puntero_heap, "1", "+") //puntero del siguiente elemento
	generador.Mi_generador.AddGetHeap(tmp2, "(int)"+tmp1)
	generador.Mi_generador.AddIf(tmp2, "-1", "==", label_error) //error porque no existe
	generador.Mi_generador.AddIf(contador, tmp_indice, "==", label_eliminar)
	generador.Mi_generador.AddAssign(puntero_heap, tmp2)
	generador.Mi_generador.AddExpression(contador, contador, "1", "+")
	generador.Mi_generador.AddGoto(label_inicio)

	generador.Mi_generador.AddLabel(label_eliminar)
	generador.Mi_generador.AddGetHeap(tmp3, "(int)"+tmp2)
	generador.Mi_generador.AddSetHeap("(int)"+puntero_heap, tmp3)
	generador.Mi_generador.AddExpression(tmp2, tmp2, "1", "+")
	generador.Mi_generador.AddExpression(puntero_heap, tmp2, "1", "+")
	generador.Mi_generador.AddSetHeap("(int)"+puntero_heap, tmp3)
	generador.Mi_generador.AddGoto(label_salida)

	generador.Mi_generador.AddLabel(label_error)
	// error index
	generador.Mi_generador.AddPrintf("c", "69")
	generador.Mi_generador.AddPrintf("c", "114")
	generador.Mi_generador.AddPrintf("c", "114")
	generador.Mi_generador.AddPrintf("c", "111")
	generador.Mi_generador.AddPrintf("c", "114")

	generador.Mi_generador.AddLabel(label_salida)
	return valor.Value{}
}
