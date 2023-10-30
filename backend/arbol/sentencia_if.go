package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Sentencia_if struct {
	Expresion       BaseNodo
	Sentencias      []BaseNodo
	Sentencias_else []BaseNodo
	Siguiente       BaseNodo //siguiente if
}

func (s Sentencia_if) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	var result valor.Value
	ambito_local := &ambito.Ambito{NombreAmbito: "sentencia if", Padre: ambito_padre}
	ambito_padre.AgregarAmbito(ambito_local)
	resultado := s.Expresion.Ejecutar(ambito_padre)
	if resultado.Type != valor.BOOLEAN {
		panic("Error la expresion no es un bool que espera el if")
	}
	label_salida := generador.Mi_generador.NewLabel()
	//true labels
	for _, lvl := range resultado.TrueLabel {
		generador.Mi_generador.AddLabel(lvl.(string))
	}
	// muevo el puntero al nuevo ambito
	generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(ambito_padre.Size), "+")
	//sentencias verdaderas
	for _, linea := range s.Sentencias {
		linea.Ejecutar(ambito_padre)
	}
	//add out labels
	generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(ambito_padre.Size), "-")
	generador.Mi_generador.AddGoto(label_salida)
	//add false labels
	for _, lvl := range resultado.FalseLabel {
		generador.Mi_generador.AddLabel(lvl.(string))
	}
	if s.Siguiente != nil { // ejecutamos el siguiente if
		s.Siguiente.Ejecutar(ambito_padre) //ambito padre porque los if estan al mismo nivel
	} else { //sentencias falsas
		// muevo el puntero al nuevo ambito
		generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(ambito_padre.Size), "+")
		for _, linea := range s.Sentencias_else {
			linea.Ejecutar(ambito_padre)
		}
		generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(ambito_padre.Size), "-")
	}
	//add out label del false
	generador.Mi_generador.AddGoto(label_salida)
	//agregando etiquetas de salida
	generador.Mi_generador.AddLabel(label_salida)
	return result
}
