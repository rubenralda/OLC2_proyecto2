package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Sentencia_guard struct {
	Expresion  BaseNodo
	Sentencias []BaseNodo
}

func (s Sentencia_guard) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	var result valor.Value
	ambito_local := &ambito.Ambito{NombreAmbito: "Sentencia guard", Padre: ambito_padre}
	ambito_padre.AgregarAmbito(ambito_local)
	resultado := s.Expresion.Ejecutar(ambito_padre)
	if resultado.Type != valor.BOOLEAN {
		ambito_local.Agregar_error("Error la expresion no es un bool que espera el guard")
		return valor.Value{}
	}
	label_salida := generador.Mi_generador.NewLabel()
	//false labels
	for _, lvl := range resultado.FalseLabel {
		generador.Mi_generador.AddLabel(lvl.(string))
	}
	// muevo el puntero al nuevo ambito
	generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(ambito_padre.Size), "+")
	//sentencias
	for _, linea := range s.Sentencias {
		linea.Ejecutar(ambito_padre)
	}
	//add out labels
	generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(ambito_padre.Size), "-")
	generador.Mi_generador.AddGoto(label_salida)
	//add true labels
	for _, lvl := range resultado.TrueLabel {
		generador.Mi_generador.AddLabel(lvl.(string))
	}
	//agregando etiquetas de salida
	generador.Mi_generador.AddLabel(label_salida)
	return result
}
