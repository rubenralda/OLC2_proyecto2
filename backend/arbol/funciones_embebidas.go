package arbol

import (
	"fmt"
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

var Salid_programa string

type Funcion_print struct {
	Lista_expresion []BaseNodo
}

func (i Funcion_print) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	var result valor.Value
	for _, expresion := range i.Lista_expresion {
		result = expresion.Ejecutar(ambito_padre)
		if result.Type == valor.INTEGER || result.Type == valor.FLOAT || result.Type == valor.CHAR {
			generador.Mi_generador.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))
			generador.Mi_generador.AddPrintf("c", "32")
			generador.Mi_generador.AddBr()
		} else if result.Type == valor.BOOLEAN {
			if result.IsTemp {
				//cuando es variable, si pasa?
				panic("no se si pasa en funcion print, no implementado")
			}
			newLabel := generador.Mi_generador.NewLabel()
			//add labels
			for _, lvl := range result.TrueLabel {
				generador.Mi_generador.AddLabel(lvl.(string))
			}
			generador.Mi_generador.AddPrintf("c", "(char)116")
			generador.Mi_generador.AddPrintf("c", "(char)114")
			generador.Mi_generador.AddPrintf("c", "(char)117")
			generador.Mi_generador.AddPrintf("c", "(char)101")
			generador.Mi_generador.AddGoto(newLabel)
			//add labels
			for _, lvl := range result.FalseLabel {
				generador.Mi_generador.AddLabel(lvl.(string))
			}
			generador.Mi_generador.AddPrintf("c", "(char)102")
			generador.Mi_generador.AddPrintf("c", "(char)97")
			generador.Mi_generador.AddPrintf("c", "(char)108")
			generador.Mi_generador.AddPrintf("c", "(char)115")
			generador.Mi_generador.AddPrintf("c", "(char)101")
			generador.Mi_generador.AddLabel(newLabel)
			generador.Mi_generador.AddPrintf("c", "32")
			generador.Mi_generador.AddBr()
		} else if result.Type == valor.STRING {
			//llamar a generar printstring
			generador.Mi_generador.GeneratePrintString()
			//agregar codigo en el main
			newTemp1 := generador.Mi_generador.NewTemp()
			newTemp2 := generador.Mi_generador.NewTemp()
			size := strconv.Itoa(ambito_padre.Size)
			generador.Mi_generador.AddExpression(newTemp1, "P", size, "+")     //nuevo temporal en pos vacia
			generador.Mi_generador.AddExpression(newTemp1, newTemp1, "1", "+") //se deja espacio de retorno
			generador.Mi_generador.AddSetStack("(int)"+newTemp1, result.Value) //se coloca string en parametro que se manda
			generador.Mi_generador.AddExpression("P", "P", size, "+")          // cambio de entorno
			generador.Mi_generador.AddCall("dbrust_printString")               //Llamada
			generador.Mi_generador.AddGetStack(newTemp2, "(int)P")             //obtencion retorno
			generador.Mi_generador.AddExpression("P", "P", size, "-")          //regreso del entorno
			generador.Mi_generador.AddPrintf("c", "32")                        //salto de linea
			generador.Mi_generador.AddBr()
		} else if result.Type == valor.NULL {
			generador.Mi_generador.AddPrintf("c", "(char)110")
			generador.Mi_generador.AddPrintf("c", "(char)105")
			generador.Mi_generador.AddPrintf("c", "(char)108")
			generador.Mi_generador.AddPrintf("c", "32") //espacio
			generador.Mi_generador.AddBr()
		}
	}
	generador.Mi_generador.AddPrintf("c", "10")
	return valor.Value{}
}

type Funcion_int struct {
	Expresion BaseNodo
}

func (f Funcion_int) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Funcion_float struct {
	Expresion BaseNodo
}

func (f Funcion_float) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

type Funcion_string struct {
	Expresion BaseNodo
}

func (f Funcion_string) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}
