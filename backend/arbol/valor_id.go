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
	return valor.Value{}
}

type Id_matriz struct {
	Id      string
	Indices []BaseNodo
}

func (a Id_matriz) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
