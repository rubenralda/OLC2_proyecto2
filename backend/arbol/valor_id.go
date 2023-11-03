package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Id_expresion struct {
	Id string
}

func (a Id_expresion) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	//generador.Mi_generador.AddComment("Llamando una variable")
	var result valor.Value
	id_buscado, size_total := ambito_padre.BuscarVariable(a.Id)
	if id_buscado == nil {
		panic("La variable no existe " + a.Id)
	}
	if !id_buscado.Is_init {
		panic("La variable no se ha inicializado")
	}
	tmp_posicion := generador.Mi_generador.NewTemp()
	temp2 := generador.Mi_generador.NewTemp()
	tmp_posicion_entorno := generador.Mi_generador.NewTemp()
	posicion_referencia := ""
	if ambito.Ambito_global.Is_global_variable(id_buscado.Id) && !generador.Mi_generador.MainCode {
		generador.Mi_generador.AddAssign(tmp_posicion_entorno, "0")
	} else {
		generador.Mi_generador.AddExpression(tmp_posicion_entorno, "P", strconv.Itoa(size_total), "-")
	}
	generador.Mi_generador.AddExpression(tmp_posicion, tmp_posicion_entorno, strconv.Itoa(id_buscado.Posicion_relativa), "+")
	posicion_referencia = tmp_posicion
	if id_buscado.Is_referencia {
		tmp_puntero := generador.Mi_generador.NewTemp()
		//obtener el puntero y modificar en esa posicion el valor
		generador.Mi_generador.AddGetStack(tmp_puntero, "(int)"+tmp_posicion)
		posicion_referencia = tmp_puntero
		generador.Mi_generador.AddGetStack(temp2, "(int)"+tmp_puntero)
	} else {
		generador.Mi_generador.AddGetStack(temp2, "(int)"+tmp_posicion)
	}
	if id_buscado.Tipo_dimension == valor.DIMENSION1 || id_buscado.Tipo_dimension == valor.DIMENSIONN {
		// logica para retornar un vector o una matriz
		panic("no implementado")
	} else if id_buscado.Tipo == valor.BOOLEAN {
		true_label := generador.Mi_generador.NewLabel()
		false_label := generador.Mi_generador.NewLabel()

		generador.Mi_generador.AddIf(temp2, "1", "==", true_label)
		generador.Mi_generador.AddGoto(false_label)
		result = valor.Value{Type: valor.BOOLEAN, Referencia: posicion_referencia}
		result.TrueLabel = append(result.TrueLabel, true_label)
		result.FalseLabel = append(result.FalseLabel, false_label)
	} else {
		if id_buscado.Is_instancia {
			result = valor.Value{Value: temp2, Is_intancia: true, Tipo_struct: id_buscado.Tipo_struct, Referencia: posicion_referencia}
		} else {
			result = valor.Value{Value: temp2, Type: id_buscado.Tipo, Referencia: posicion_referencia}
		}
	}
	return result
}

type Id_vector struct {
	Id     string
	Indice BaseNodo
}

func (a Id_vector) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	variable, size := ambito_padre.BuscarVariable(a.Id)
	if variable == nil {
		panic("La variable no existe " + a.Id)
	}
	if variable.Tipo_dimension != valor.DIMENSION1 {
		panic("La variable no es un vector " + a.Id)
	}
	indice := a.Indice.Ejecutar(ambito_padre)
	if indice.Type != valor.INTEGER {
		panic("El indice no es entero")
	}
	tmp_indice := generador.Mi_generador.NewTemp()
	generador.Mi_generador.AddAssign(tmp_indice, indice.Value)
	puntero_variable := generador.Mi_generador.NewTemp()
	puntero_heap := generador.Mi_generador.NewTemp()
	tmp_posicion_entorno := generador.Mi_generador.NewTemp()
	valor_heap := generador.Mi_generador.NewTemp()
	label_inicio := generador.Mi_generador.NewLabel()
	label_error := generador.Mi_generador.NewLabel()
	label_salida := generador.Mi_generador.NewLabel()
	if ambito.Ambito_global.Is_global_variable(variable.Id) && !generador.Mi_generador.MainCode {
		generador.Mi_generador.AddAssign(tmp_posicion_entorno, "0")
	} else {
		generador.Mi_generador.AddExpression(tmp_posicion_entorno, "P", strconv.Itoa(size), "-")
	}
	generador.Mi_generador.AddExpression(puntero_variable, tmp_posicion_entorno, strconv.Itoa(variable.Posicion_relativa), "+")
	generador.Mi_generador.AddGetStack(puntero_heap, "(int)"+puntero_variable)

	generador.Mi_generador.AddLabel(label_inicio)
	generador.Mi_generador.AddGetHeap(valor_heap, "(int)"+puntero_heap)
	generador.Mi_generador.AddExpression(puntero_heap, puntero_heap, "1", "+") //puntero del siguiente elemento
	generador.Mi_generador.AddGetHeap(puntero_heap, "(int)"+puntero_heap)
	generador.Mi_generador.AddIf(puntero_heap, "-1", "==", label_error) //si es -1 se paso del index
	generador.Mi_generador.AddIf(tmp_indice, "0", "<=", label_salida)   //termino y el valor_heap es el resultado
	generador.Mi_generador.AddExpression(tmp_indice, tmp_indice, "1", "-")
	generador.Mi_generador.AddGoto(label_inicio)

	generador.Mi_generador.AddLabel(label_error)
	// error index
	generador.Mi_generador.AddPrintf("c", "69")
	generador.Mi_generador.AddPrintf("c", "114")
	generador.Mi_generador.AddPrintf("c", "114")
	generador.Mi_generador.AddPrintf("c", "111")
	generador.Mi_generador.AddPrintf("c", "114")
	generador.Mi_generador.AddAssign(valor_heap, "9999999827968.00") //es nil
	//generador.Mi_generador.AddGoto(label_salida)

	generador.Mi_generador.AddLabel(label_salida)
	if variable.Tipo == valor.BOOLEAN {
		true_label := generador.Mi_generador.NewLabel()
		false_label := generador.Mi_generador.NewLabel()

		generador.Mi_generador.AddIf(valor_heap, "1", "==", true_label)
		generador.Mi_generador.AddGoto(false_label)
		result := valor.Value{Type: valor.BOOLEAN}
		result.TrueLabel = append(result.TrueLabel, true_label)
		result.FalseLabel = append(result.FalseLabel, false_label)
		return result
	}
	return valor.Value{Value: valor_heap, Type: variable.Tipo, Is_intancia: variable.Is_instancia, Tipo_struct: variable.Tipo_struct}
}

type Id_matriz struct {
	Id      string
	Indices []BaseNodo
}

func (a Id_matriz) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
