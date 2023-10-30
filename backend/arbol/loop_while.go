package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Loop_while struct {
	Expresion  BaseNodo
	Sentencias []BaseNodo
}

func (s Loop_while) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	generador.Mi_generador.AddComment("while quiero ver el 0 &&&")
	ambito_local := &ambito.Ambito{NombreAmbito: "sentencia while", Padre: ambito_padre}
	ambito_padre.AgregarAmbito(ambito_local)
	//crear la etiqueta de inicio
	label_inicio := generador.Mi_generador.NewLabel()
	generador.Mi_generador.AddLabel(label_inicio)
	//agregar la etiqueta de inicio para sentencias de transferencia
	generador.Mi_generador.ContinueLabel = label_inicio

	resultado := s.Expresion.Ejecutar(ambito_padre)
	//agregar la primera etiqueta de false para break
	generador.Mi_generador.BreakLabel = resultado.FalseLabel[0].(string)
	if resultado.Type != valor.BOOLEAN {
		panic("Error la expresion no es un bool que espera el while")
	}
	//true labels
	for _, lvl := range resultado.TrueLabel {
		generador.Mi_generador.AddLabel(lvl.(string))
	}
	// muevo el puntero al nuevo ambito
	generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(ambito_padre.Size), "+")
	//sentencias verdaderas
	for _, linea := range s.Sentencias {
		linea.Ejecutar(ambito_local)
	}
	//add out labels
	generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(ambito_padre.Size), "-")
	generador.Mi_generador.AddGoto(label_inicio)
	//add false labels
	for _, lvl := range resultado.FalseLabel {
		generador.Mi_generador.AddLabel(lvl.(string))
	}
	generador.Mi_generador.ContinueLabel = ""
	generador.Mi_generador.BreakLabel = ""
	return valor.Value{}
}
