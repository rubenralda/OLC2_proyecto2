package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Expresion struct {
	Valor1    BaseNodo
	Valor2    BaseNodo
	Operacion string
}

// tabla de tipos para operaciones aritmeticas
var tipo_op_aritmetica = [6][6]valor.TipoExpresion{
	//NULL			FLOAT		 STRING		BOOLEAN		CHAR		INT
	{valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
	//FLOAT
	{valor.NULL, valor.FLOAT, valor.NULL, valor.NULL, valor.NULL, valor.FLOAT},
	//STRING
	{valor.NULL, valor.NULL, valor.STRING, valor.NULL, valor.NULL, valor.NULL},
	//BOOLEAN
	{valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
	//CHAR
	{valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
	//INT
	{valor.NULL, valor.FLOAT, valor.NULL, valor.NULL, valor.NULL, valor.INTEGER},
}

// tabla de tipos para operaciones relacionales
var tipo_op_relaciones = [6][6]valor.TipoExpresion{
	//NULL		FLOAT		STRING		BOOLEAN		CHAR		INT
	{valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
	//FLOAT
	{valor.NULL, valor.FLOAT, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
	//STRING
	{valor.NULL, valor.NULL, valor.STRING, valor.NULL, valor.NULL, valor.NULL},
	//BOOLEAN
	{valor.NULL, valor.NULL, valor.NULL, valor.BOOLEAN, valor.NULL, valor.NULL},
	//CHAR
	{valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.CHAR, valor.NULL},
	//INT
	{valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.INTEGER},
}

func (e Expresion) Ejecutar(ambito *ambito.Ambito) valor.Value {
	var result valor.Value
	if e.Operacion == "&&" {
		resultado1 := e.Valor1.Ejecutar(ambito)
		//add op1 labels
		for _, lvl := range resultado1.TrueLabel {
			generador.Mi_generador.AddLabel(lvl.(string))
		}
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
		if tipo_dominante != valor.BOOLEAN {
			panic("Error los operadores no son boolean &&")
		}
		result = valor.Value{Type: valor.BOOLEAN}
		result.TrueLabel = append(resultado2.TrueLabel, result.TrueLabel...)
		result.FalseLabel = append(resultado1.FalseLabel, result.FalseLabel...)
		result.FalseLabel = append(resultado2.FalseLabel, result.FalseLabel...)
		return result
	} else if e.Operacion == "||" {
		resultado1 := e.Valor1.Ejecutar(ambito)
		//add op1 labels
		for _, lvl := range resultado1.FalseLabel {
			generador.Mi_generador.AddLabel(lvl.(string))
		}
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
		if tipo_dominante != valor.BOOLEAN {
			panic("Error los operadores no son boolean ||")
		}
		result = valor.Value{Type: valor.BOOLEAN}
		result.TrueLabel = append(resultado1.TrueLabel, result.TrueLabel...)
		result.TrueLabel = append(resultado2.TrueLabel, result.TrueLabel...)
		result.FalseLabel = append(resultado2.FalseLabel, result.FalseLabel...)
		return result
	} else if e.Operacion == "!" {
		resultado1 := e.Valor1.Ejecutar(ambito)
		if resultado1.Type == valor.BOOLEAN {
			result = valor.Value{Type: valor.BOOLEAN}
			result.TrueLabel = append(result.TrueLabel, resultado1.FalseLabel...)
			result.FalseLabel = append(result.FalseLabel, resultado1.TrueLabel...)
			return result
		} else {
			panic("No se puede hacer la operacion (!)")
		}
	} else if e.Operacion == "negacion" {
		resultado1 := e.Valor1.Ejecutar(ambito)
		if resultado1.Type == valor.INTEGER || resultado1.Type == valor.FLOAT {
			newTemp := generador.Mi_generador.NewTemp()
			generador.Mi_generador.AddExpression(newTemp, "", resultado1.Value, "-")
			result = valor.Value{Value: newTemp, Type: resultado1.Type}
			return result
		} else {
			panic("No se puede negar la operacion")
		}
	} else { //todas las aritmeticas
		resultado1 := e.Valor1.Ejecutar(ambito)
		var resultado2 valor.Value
		if _, ok := e.Valor2.(Llamada_funcion); ok {
			tmp_posicion := generador.Mi_generador.NewTemp()
			generador.Mi_generador.AddExpression(tmp_posicion, "P", strconv.Itoa(ambito.Size), "+")
			generador.Mi_generador.AddSetStack(tmp_posicion, resultado1.Value)
			ambito.Size += 1
			resultado2 = e.Valor2.Ejecutar(ambito)
			ambito.Size -= 1
			tmp_posicion = generador.Mi_generador.NewTemp()
			generador.Mi_generador.AddExpression(tmp_posicion, "P", strconv.Itoa(ambito.Size), "+")
			generador.Mi_generador.AddGetStack(resultado1.Value, tmp_posicion)
		} else {
			resultado2 = e.Valor2.Ejecutar(ambito)
		}
		switch e.Operacion {
		case "+":
			newTemp := generador.Mi_generador.NewTemp()
			tipo_dominante := tipo_op_aritmetica[resultado1.Type][resultado2.Type]
			if tipo_dominante == valor.FLOAT {
				generador.Mi_generador.AddExpression(newTemp, resultado1.Value, resultado2.Value, "+")
				result = valor.Value{Value: newTemp, Type: tipo_dominante}
				return result
			} else if tipo_dominante == valor.INTEGER {
				generador.Mi_generador.AddExpression(newTemp, resultado1.Value, resultado2.Value, "+")
				generador.Mi_generador.AddAssign(newTemp, "(int)"+newTemp)
				result = valor.Value{Value: newTemp, Type: tipo_dominante}
				return result
			} else if tipo_dominante == valor.STRING {
				newTemp := generador.Mi_generador.NewTemp()
				//iguala a heap pointer
				generador.Mi_generador.AddAssign(newTemp, "H")
				tem_valor := generador.Mi_generador.NewTemp()
				true_label := generador.Mi_generador.NewLabel()
				false_label := generador.Mi_generador.NewLabel()

				// copiar string del resultado 1
				generador.Mi_generador.AddLabel(false_label)                                       //label true
				generador.Mi_generador.AddGetHeap(tem_valor, "(int)"+resultado1.Value)             //valor en el heep
				generador.Mi_generador.AddExpression(resultado1.Value, resultado1.Value, "1", "+") //incremento el valor del puntero

				generador.Mi_generador.AddIf(tem_valor, "-1", "==", true_label)
				generador.Mi_generador.AddSetHeap("(int)H", tem_valor)
				generador.Mi_generador.AddExpression("H", "H", "1", "+")
				generador.Mi_generador.AddGoto(false_label)
				generador.Mi_generador.AddLabel(true_label)

				//copiar string del resultado 2
				true_label2 := generador.Mi_generador.NewLabel()                                   //label true
				generador.Mi_generador.AddGetHeap(tem_valor, "(int)"+resultado2.Value)             //valor en el heep
				generador.Mi_generador.AddExpression(resultado2.Value, resultado2.Value, "1", "+") //incremento el valor del puntero

				generador.Mi_generador.AddIf(tem_valor, "-1", "==", true_label2)
				generador.Mi_generador.AddSetHeap("(int)H", tem_valor)
				generador.Mi_generador.AddExpression("H", "H", "1", "+")
				generador.Mi_generador.AddGoto(true_label) //caso que es falso
				generador.Mi_generador.AddLabel(true_label2)

				// delimitador del nuevo string
				generador.Mi_generador.AddSetHeap("(int)H", "-1")
				generador.Mi_generador.AddExpression("H", "H", "1", "+")
				generador.Mi_generador.AddBr()
				result = valor.Value{Value: newTemp, Type: tipo_dominante}
				return result
			} else { // nil
				panic("Error tipo no valido")
			}
		case "-":
			newTemp := generador.Mi_generador.NewTemp()
			tipo_dominante := tipo_op_aritmetica[resultado1.Type][resultado2.Type]
			if tipo_dominante == valor.FLOAT {
				generador.Mi_generador.AddExpression(newTemp, resultado1.Value, resultado2.Value, "-")
				result = valor.Value{Value: newTemp, Type: tipo_dominante}
				return result
			} else if tipo_dominante == valor.INTEGER {
				generador.Mi_generador.AddExpression(newTemp, resultado1.Value, resultado2.Value, "-")
				generador.Mi_generador.AddAssign(newTemp, "(int)"+newTemp)
				result = valor.Value{Value: newTemp, Type: tipo_dominante}
				return result
			} else { // nil
				panic("Error tipo no valido")
			}
		case "*":
			newTemp := generador.Mi_generador.NewTemp()
			tipo_dominante := tipo_op_aritmetica[resultado1.Type][resultado2.Type]
			if tipo_dominante == valor.FLOAT {
				generador.Mi_generador.AddExpression(newTemp, resultado1.Value, resultado2.Value, "*")
				result = valor.Value{Value: newTemp, Type: tipo_dominante}
				return result
			} else if tipo_dominante == valor.INTEGER {
				generador.Mi_generador.AddExpression(newTemp, resultado1.Value, resultado2.Value, "*")
				generador.Mi_generador.AddAssign(newTemp, "(int)"+newTemp)
				result = valor.Value{Value: newTemp, Type: tipo_dominante}
				return result
			} else { // nil
				panic("Error tipo no valido")
			}
		case "/":
			newTemp := generador.Mi_generador.NewTemp()
			tipo_dominante := tipo_op_aritmetica[resultado1.Type][resultado2.Type]
			if tipo_dominante == valor.FLOAT {
				lvl1 := generador.Mi_generador.NewLabel()
				lvl2 := generador.Mi_generador.NewLabel()
				generador.Mi_generador.AddIf(resultado2.Value, "0", "!=", lvl1)
				generador.Mi_generador.AddPrintf("c", "77")
				generador.Mi_generador.AddPrintf("c", "97")
				generador.Mi_generador.AddPrintf("c", "116")
				generador.Mi_generador.AddPrintf("c", "104")
				generador.Mi_generador.AddPrintf("c", "69")
				generador.Mi_generador.AddPrintf("c", "114")
				generador.Mi_generador.AddPrintf("c", "114")
				generador.Mi_generador.AddPrintf("c", "111")
				generador.Mi_generador.AddPrintf("c", "114")
				generador.Mi_generador.AddExpression(newTemp, "0", "", "")
				generador.Mi_generador.AddGoto(lvl2)
				generador.Mi_generador.AddLabel(lvl1)
				generador.Mi_generador.AddExpression(newTemp, resultado1.Value, resultado2.Value, "/")
				generador.Mi_generador.AddLabel(lvl2)
				result = valor.Value{Value: newTemp, Type: tipo_dominante}
				return result
			} else if tipo_dominante == valor.INTEGER {
				lvl1 := generador.Mi_generador.NewLabel()
				lvl2 := generador.Mi_generador.NewLabel()
				generador.Mi_generador.AddIf(resultado2.Value, "0", "!=", lvl1)
				generador.Mi_generador.AddPrintf("c", "77")
				generador.Mi_generador.AddPrintf("c", "97")
				generador.Mi_generador.AddPrintf("c", "116")
				generador.Mi_generador.AddPrintf("c", "104")
				generador.Mi_generador.AddPrintf("c", "69")
				generador.Mi_generador.AddPrintf("c", "114")
				generador.Mi_generador.AddPrintf("c", "114")
				generador.Mi_generador.AddPrintf("c", "111")
				generador.Mi_generador.AddPrintf("c", "114")
				generador.Mi_generador.AddExpression(newTemp, "0", "", "")
				generador.Mi_generador.AddGoto(lvl2)
				generador.Mi_generador.AddLabel(lvl1)
				generador.Mi_generador.AddExpression(newTemp, resultado1.Value, resultado2.Value, "/")
				generador.Mi_generador.AddAssign(newTemp, "(int)"+newTemp)
				generador.Mi_generador.AddLabel(lvl2)
				result = valor.Value{Value: newTemp, Type: tipo_dominante}
				return result
			} else { // nil
				panic("Error tipo no valido")
			}
		case "%":
			newTemp := generador.Mi_generador.NewTemp()
			tipo_dominante := tipo_op_aritmetica[resultado1.Type][resultado2.Type]
			if tipo_dominante == valor.INTEGER {
				generador.Mi_generador.AddExpression(newTemp, "(int)"+resultado1.Value, "(int)"+resultado2.Value, "%")
				generador.Mi_generador.AddAssign(newTemp, "(int)"+newTemp)
				result = valor.Value{Value: newTemp, Type: tipo_dominante}
				return result
			} else { // nil
				panic("Error tipo no valido")
			}
		case ">":
			tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
			if tipo_dominante == valor.INTEGER || tipo_dominante == valor.FLOAT || tipo_dominante == valor.CHAR {
				trueLabel := generador.Mi_generador.NewLabel()
				falseLabel := generador.Mi_generador.NewLabel()
				generador.Mi_generador.AddIf(resultado1.Value, resultado2.Value, ">", trueLabel)
				generador.Mi_generador.AddGoto(falseLabel)
				result = valor.Value{Type: valor.BOOLEAN}
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else if tipo_dominante == valor.STRING {
				return operacion_comparacion(">", resultado1, resultado2)
			} else {
				panic("Error no se puede comparar >")
			}
		case "<":
			tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
			if tipo_dominante == valor.INTEGER || tipo_dominante == valor.FLOAT || tipo_dominante == valor.CHAR {
				trueLabel := generador.Mi_generador.NewLabel()
				falseLabel := generador.Mi_generador.NewLabel()
				generador.Mi_generador.AddIf(resultado1.Value, resultado2.Value, "<", trueLabel)
				generador.Mi_generador.AddGoto(falseLabel)
				result = valor.Value{Type: valor.BOOLEAN}
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else if tipo_dominante == valor.STRING {
				return operacion_comparacion("<", resultado1, resultado2)
			} else {
				panic("Error no se puede comparar <")
			}
		case ">=":
			tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
			if tipo_dominante == valor.INTEGER || tipo_dominante == valor.FLOAT || tipo_dominante == valor.CHAR {
				trueLabel := generador.Mi_generador.NewLabel()
				falseLabel := generador.Mi_generador.NewLabel()
				generador.Mi_generador.AddIf(resultado1.Value, resultado2.Value, ">=", trueLabel)
				generador.Mi_generador.AddGoto(falseLabel)
				result = valor.Value{Type: valor.BOOLEAN}
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else if tipo_dominante == valor.STRING {
				return operacion_comparacion(">=", resultado1, resultado2)
			} else {
				panic("Error no se puede comparar >=")
			}
		case "<=":
			tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
			if tipo_dominante == valor.INTEGER || tipo_dominante == valor.FLOAT || tipo_dominante == valor.CHAR {
				trueLabel := generador.Mi_generador.NewLabel()
				falseLabel := generador.Mi_generador.NewLabel()
				generador.Mi_generador.AddIf(resultado1.Value, resultado2.Value, "<=", trueLabel)
				generador.Mi_generador.AddGoto(falseLabel)
				result = valor.Value{Type: valor.BOOLEAN}
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else if tipo_dominante == valor.STRING {
				return operacion_comparacion("<=", resultado1, resultado2)
			} else {
				panic("Error no se puede comparar <=")
			}
		case "==":
			tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
			if tipo_dominante == valor.STRING {
				return operacion_comparacion("==", resultado1, resultado2)
			} else if tipo_dominante != valor.NULL {
				trueLabel := generador.Mi_generador.NewLabel()
				falseLabel := generador.Mi_generador.NewLabel()
				generador.Mi_generador.AddIf(resultado1.Value, resultado2.Value, "==", trueLabel)
				generador.Mi_generador.AddGoto(falseLabel)
				result = valor.Value{Type: valor.BOOLEAN}
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				panic("Error no se puede comparar ==")
			}
		case "!=":
			tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
			if tipo_dominante == valor.STRING {
				return operacion_comparacion("!=", resultado1, resultado2)
			} else if tipo_dominante != valor.NULL {
				trueLabel := generador.Mi_generador.NewLabel()
				falseLabel := generador.Mi_generador.NewLabel()
				generador.Mi_generador.AddIf(resultado1.Value, resultado2.Value, "!=", trueLabel)
				generador.Mi_generador.AddGoto(falseLabel)
				result = valor.Value{Type: valor.BOOLEAN}
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				panic("Error no se puede comparar !=")
			}
		default:
			return result
		}
	}
}

// para string, resuelve que string es mayor
func operacion_comparacion(operacion string, resultado1 valor.Value, resultado2 valor.Value) valor.Value {
	temp_contador1 := generador.Mi_generador.NewTemp()
	temp_contador2 := generador.Mi_generador.NewTemp()
	temp_valor := generador.Mi_generador.NewTemp()

	true_label := generador.Mi_generador.NewLabel()
	false_label := generador.Mi_generador.NewLabel()
	generador.Mi_generador.AddLabel(false_label)
	generador.Mi_generador.AddGetHeap(temp_valor, "(int)"+resultado1.Value)
	generador.Mi_generador.AddExpression(resultado1.Value, resultado1.Value, "1", "+")
	generador.Mi_generador.AddIf(temp_valor, "-1", "==", true_label)
	generador.Mi_generador.AddExpression(temp_contador1, temp_contador1, "1", "+")
	generador.Mi_generador.AddGoto(false_label)

	generador.Mi_generador.AddLabel(true_label) //label false del siguiente if tambien
	true_label2 := generador.Mi_generador.NewLabel()
	generador.Mi_generador.AddGetHeap(temp_valor, "(int)"+resultado2.Value)
	generador.Mi_generador.AddExpression(resultado2.Value, resultado2.Value, "1", "+")
	generador.Mi_generador.AddIf(temp_valor, "-1", "==", true_label2)
	generador.Mi_generador.AddExpression(temp_contador2, temp_contador2, "1", "+")
	generador.Mi_generador.AddGoto(true_label)
	generador.Mi_generador.AddLabel(true_label2)

	true_label_comparacion := generador.Mi_generador.NewLabel()
	false_label_comparacion := generador.Mi_generador.NewLabel()
	generador.Mi_generador.AddIf(temp_contador1, temp_contador2, operacion, true_label_comparacion)
	generador.Mi_generador.AddGoto(false_label_comparacion)

	result := valor.Value{Type: valor.BOOLEAN}
	return result
}
