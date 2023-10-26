package arbol

import (
	"fmt"
	"main/ambito"
	"main/generador"
	"main/valor"
)

var Salid_programa string

type Funcion_print struct {
	Lista_expresion []BaseNodo
}

func (i Funcion_print) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	var result valor.Value
	for _, expresion := range i.Lista_expresion {
		result = expresion.Ejecutar(ambito_padre)
		if result.Type == valor.INTEGER || result.Type == valor.FLOAT || result.Type == valor.ARRAY {
			generador.Mi_generador.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))
			generador.Mi_generador.AddPrintf("c", "10")
			generador.Mi_generador.AddBr()
		} else if result.Type == valor.BOOLEAN {
			if result.IsTemp {
				//cuando es variable
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
			generador.Mi_generador.AddPrintf("c", "10")
			generador.Mi_generador.AddBr()
		} else if result.Type == valor.STRING {
			// falta modificar esa logica de cuando es string
		}
	}
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
