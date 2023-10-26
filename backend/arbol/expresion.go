package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
)

type Expresion struct {
	Valor1    BaseNodo
	Valor2    BaseNodo
	Operacion string
}

// tabla de tipos para operaciones aritmeticas
var tipo_op_aritmetica = [6][6]valor.TipoExpresion{
	//INTEGER		FLOAT		 STRING		BOOLEAN		NULL			CHAR
	{valor.INTEGER, valor.FLOAT, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
	//FLOAT
	{valor.FLOAT, valor.FLOAT, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
	//STRING
	{valor.NULL, valor.NULL, valor.STRING, valor.NULL, valor.NULL, valor.NULL},
	//BOOLEAN
	{valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
	//CHAR
	{valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
	//NULL
	{valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
}

// tabla de tipos para operaciones relacionales
var tipo_op_relaciones = [6][6]valor.TipoExpresion{
	//INTEGER		FLOAT		STRING		BOOLEAN		NULL		CHAR
	{valor.INTEGER, valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
	//FLOAT
	{valor.NULL, valor.FLOAT, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
	//STRING
	{valor.NULL, valor.NULL, valor.STRING, valor.NULL, valor.NULL, valor.NULL},
	//BOOLEAN
	{valor.NULL, valor.NULL, valor.NULL, valor.BOOLEAN, valor.NULL, valor.NULL},
	//CHAR
	{valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.CHAR},
	//NULL
	{valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL, valor.NULL},
}

func (e Expresion) Ejecutar(ambito *ambito.Ambito) valor.Value {
	var result valor.Value
	newTemp := generador.Mi_generador.NewTemp()
	switch e.Operacion {
	case "+":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_aritmetica[resultado1.Type][resultado2.Type]
		if tipo_dominante == valor.INTEGER || tipo_dominante == valor.FLOAT {
			generador.Mi_generador.AddExpression(newTemp, resultado1.Value, resultado2.Value, "+")
			result = valor.Value{Value: newTemp, IsTemp: true, Type: tipo_dominante}
			result.IntValue = resultado1.IntValue + resultado2.IntValue
			return result
		} else if tipo_dominante == valor.STRING {
			// logica para manejar el string
			return result
		} else { // nil
			panic("Error tipo no valido")
		}
	case "negacion":
		resultado1 := e.Valor1.Ejecutar(ambito)
		if resultado1.Type == valor.INTEGER || resultado1.Type == valor.FLOAT {
			generador.Mi_generador.AddExpression(newTemp, "", resultado1.Value, "-")
			result = valor.Value{Value: newTemp, IsTemp: true, Type: resultado1.Type}
			result.IntValue = -resultado1.IntValue
			return result
		} else {
			panic("No se puede negar la operacion")
		}
	case "-":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_aritmetica[resultado1.Type][resultado2.Type]
		if tipo_dominante == valor.INTEGER || tipo_dominante == valor.FLOAT {
			generador.Mi_generador.AddExpression(newTemp, resultado1.Value, resultado2.Value, "-")
			result = valor.Value{Value: newTemp, IsTemp: true, Type: tipo_dominante}
			result.IntValue = resultado1.IntValue - resultado2.IntValue
			return result
		} else { // nil
			panic("Error tipo no valido")
		}
	case "*":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_aritmetica[resultado1.Type][resultado2.Type]
		if tipo_dominante == valor.INTEGER || tipo_dominante == valor.FLOAT {
			generador.Mi_generador.AddExpression(newTemp, resultado1.Value, resultado2.Value, "*")
			result = valor.Value{Value: newTemp, IsTemp: true, Type: tipo_dominante}
			result.IntValue = resultado1.IntValue * resultado2.IntValue
			return result
		} else { // nil
			panic("Error tipo no valido")
		}
	case "/":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_aritmetica[resultado1.Type][resultado2.Type]
		if tipo_dominante == valor.INTEGER || tipo_dominante == valor.FLOAT {
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
			result = valor.Value{Value: newTemp, IsTemp: true, Type: tipo_dominante}
			return result
		} else { // nil
			panic("Error tipo no valido")
		}
	case "%":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_aritmetica[resultado1.Type][resultado2.Type]
		if tipo_dominante == valor.INTEGER {
			generador.Mi_generador.AddExpression(newTemp, resultado1.Value, resultado2.Value, "%")
			result = valor.Value{Value: newTemp, IsTemp: true, Type: tipo_dominante}
			result.IntValue = resultado1.IntValue % resultado2.IntValue
			return result
		} else { // nil
			panic("Error tipo no valido")
		}
	case ">":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
		if tipo_dominante == valor.INTEGER || tipo_dominante == valor.FLOAT || tipo_dominante == valor.CHAR { // agregar char si funciona igual
			trueLabel := generador.Mi_generador.NewLabel()
			falseLabel := generador.Mi_generador.NewLabel()

			generador.Mi_generador.AddIf(resultado1.Value, resultado2.Value, ">", trueLabel)
			generador.Mi_generador.AddGoto(falseLabel)

			result = valor.Value{Value: "", IsTemp: false, Type: valor.BOOLEAN}
			result.TrueLabel = append(result.TrueLabel, trueLabel)
			result.FalseLabel = append(result.FalseLabel, falseLabel)
			return result
		} else if tipo_dominante == valor.STRING {
			//logica para manejar el string comparando cada posicion
			return result
		} else {
			panic("Error no se puede comparar >")
		}
	case "<":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
		if tipo_dominante == valor.INTEGER || tipo_dominante == valor.FLOAT || tipo_dominante == valor.CHAR { // agregar char si funciona igual
			trueLabel := generador.Mi_generador.NewLabel()
			falseLabel := generador.Mi_generador.NewLabel()

			generador.Mi_generador.AddIf(resultado1.Value, resultado2.Value, "<", trueLabel)
			generador.Mi_generador.AddGoto(falseLabel)

			result = valor.Value{Value: "", IsTemp: false, Type: valor.BOOLEAN}
			result.TrueLabel = append(result.TrueLabel, trueLabel)
			result.FalseLabel = append(result.FalseLabel, falseLabel)
			return result
		} else if tipo_dominante == valor.STRING {
			//logica para manejar el string comparando cada posicion
			return result
		} else {
			panic("Error no se puede comparar <")
		}
	case ">=":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
		if tipo_dominante == valor.INTEGER || tipo_dominante == valor.FLOAT || tipo_dominante == valor.CHAR { // agregar char si funciona igual
			trueLabel := generador.Mi_generador.NewLabel()
			falseLabel := generador.Mi_generador.NewLabel()

			generador.Mi_generador.AddIf(resultado1.Value, resultado2.Value, ">=", trueLabel)
			generador.Mi_generador.AddGoto(falseLabel)

			result = valor.Value{Value: "", IsTemp: false, Type: valor.BOOLEAN}
			result.TrueLabel = append(result.TrueLabel, trueLabel)
			result.FalseLabel = append(result.FalseLabel, falseLabel)
			return result
		} else if tipo_dominante == valor.STRING {
			//logica para manejar el string comparando cada posicion
			return result
		} else {
			panic("Error no se puede comparar >=")
		}
	case "<=":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
		if tipo_dominante == valor.INTEGER || tipo_dominante == valor.FLOAT || tipo_dominante == valor.CHAR { // agregar char si funciona igual
			trueLabel := generador.Mi_generador.NewLabel()
			falseLabel := generador.Mi_generador.NewLabel()

			generador.Mi_generador.AddIf(resultado1.Value, resultado2.Value, "<=", trueLabel)
			generador.Mi_generador.AddGoto(falseLabel)

			result = valor.Value{Value: "", IsTemp: false, Type: valor.BOOLEAN}
			result.TrueLabel = append(result.TrueLabel, trueLabel)
			result.FalseLabel = append(result.FalseLabel, falseLabel)
			return result
		} else if tipo_dominante == valor.STRING {
			//logica para manejar el string comparando cada posicion
			return result
		} else {
			panic("Error no se puede comparar <=")
		}
	case "==":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
		if tipo_dominante == valor.STRING {
			//logica para manejar el string comparando cada posicion
			return result
		} else if tipo_dominante != valor.NULL {
			trueLabel := generador.Mi_generador.NewLabel()
			falseLabel := generador.Mi_generador.NewLabel()

			generador.Mi_generador.AddIf(resultado1.Value, resultado2.Value, "==", trueLabel)
			generador.Mi_generador.AddGoto(falseLabel)

			result = valor.Value{Value: "", IsTemp: false, Type: valor.BOOLEAN}
			result.TrueLabel = append(result.TrueLabel, trueLabel)
			result.FalseLabel = append(result.FalseLabel, falseLabel)
			return result
		} else {
			panic("Error no se puede comparar ==")
		}
	case "!=":
		resultado1 := e.Valor1.Ejecutar(ambito)
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
		if tipo_dominante == valor.STRING {
			//logica para manejar el string comparando cada posicion
			return result
		} else if tipo_dominante != valor.NULL {
			trueLabel := generador.Mi_generador.NewLabel()
			falseLabel := generador.Mi_generador.NewLabel()

			generador.Mi_generador.AddIf(resultado1.Value, resultado2.Value, "!=", trueLabel)
			generador.Mi_generador.AddGoto(falseLabel)

			result = valor.Value{Value: "", IsTemp: false, Type: valor.BOOLEAN}
			result.TrueLabel = append(result.TrueLabel, trueLabel)
			result.FalseLabel = append(result.FalseLabel, falseLabel)
			return result
		} else {
			panic("Error no se puede comparar !=")
		}
	case "!":
		resultado1 := e.Valor1.Ejecutar(ambito)
		if resultado1.Type == valor.BOOLEAN {
			result = valor.Value{Value: "", IsTemp: false, Type: valor.BOOLEAN}
			result.TrueLabel = append(resultado1.FalseLabel, result.TrueLabel)
			result.FalseLabel = append(resultado1.TrueLabel, result.FalseLabel)
			return result
		} else {
			panic("No se puede hacer la operacion !")
		}
	case "&&":
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
		result = valor.Value{Value: "", IsTemp: false, Type: valor.BOOLEAN}
		result.TrueLabel = append(resultado2.TrueLabel, result.TrueLabel...)
		result.FalseLabel = append(resultado1.FalseLabel, result.FalseLabel...)
		result.FalseLabel = append(resultado2.FalseLabel, result.FalseLabel...)
		return result
	case "||":
		resultado1 := e.Valor1.Ejecutar(ambito)
		//add op1 labels
		for _, lvl := range resultado1.FalseLabel {
			generador.Mi_generador.AddLabel(lvl.(string))
		}
		resultado2 := e.Valor2.Ejecutar(ambito)
		tipo_dominante := tipo_op_relaciones[resultado1.Type][resultado2.Type]
		if tipo_dominante != valor.BOOLEAN {
			panic("Error los operadores no son boolean &&")
		}
		result = valor.Value{Value: "", IsTemp: false, Type: valor.BOOLEAN}
		result.TrueLabel = append(resultado1.TrueLabel, result.TrueLabel...)
		result.TrueLabel = append(resultado2.TrueLabel, result.TrueLabel...)
		result.FalseLabel = append(resultado2.FalseLabel, result.FalseLabel...)
		return result
	case "primitivo":
		return e.Valor1.Ejecutar(ambito)
	case "identificador":
		return e.Valor1.Ejecutar(ambito)
	case "vector":
		return e.Valor1.Ejecutar(ambito)
	case "atributos":
		return e.Valor1.Ejecutar(ambito)
	case "llamada":
		return e.Valor1.Ejecutar(ambito)
	case "objeto":
		return e.Valor1.Ejecutar(ambito)
	case "matriz":
		return e.Valor1.Ejecutar(ambito)
	default:
		return result
	}
}
