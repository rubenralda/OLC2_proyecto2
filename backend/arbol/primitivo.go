package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Primitivos struct {
	Tipo  valor.TipoExpresion
	Valor string
}

func (p Primitivos) Ejecutar(ambito *ambito.Ambito) valor.Value {
	var result valor.Value
	switch p.Tipo {
	case valor.INTEGER:
		result = valor.Value{Value: p.Valor, Type: p.Tipo}
	case valor.BOOLEAN:
		trueLabel := generador.Mi_generador.NewLabel()
		falseLabel := generador.Mi_generador.NewLabel()
		booleanValue, _ := strconv.ParseBool(p.Valor)
		if booleanValue {
			generador.Mi_generador.AddGoto(trueLabel)
		} else {
			generador.Mi_generador.AddGoto(falseLabel)
		}
		result = valor.Value{Type: p.Tipo}
		result.TrueLabel = append(result.TrueLabel, trueLabel)
		result.FalseLabel = append(result.FalseLabel, falseLabel)
	case valor.STRING:
		//nuevo temporal
		newTemp := generador.Mi_generador.NewTemp()
		//iguala a heap pointer
		generador.Mi_generador.AddAssign(newTemp, "H")
		//recorremos string en ascii
		myString := p.Valor[1 : len(p.Valor)-1]
		byteArray := []byte(myString)
		for _, asc := range byteArray {
			//se agrega ascii al heap
			generador.Mi_generador.AddSetHeap("(int)H", strconv.Itoa(int(asc)))
			//suma heap pointer
			generador.Mi_generador.AddExpression("H", "H", "1", "+")
		}
		//caracteres de escape
		generador.Mi_generador.AddSetHeap("(int)H", "-1")
		generador.Mi_generador.AddExpression("H", "H", "1", "+")
		generador.Mi_generador.AddBr()
		result = valor.Value{Value: newTemp, Type: p.Tipo}
	case valor.FLOAT:
		result = valor.Value{Value: p.Valor, Type: p.Tipo}
	case valor.CHAR:
		result = valor.Value{Value: p.Valor, Type: p.Tipo}
	case valor.NULL:
		result = valor.Value{Value: "", Type: p.Tipo}
	}
	return result
}
