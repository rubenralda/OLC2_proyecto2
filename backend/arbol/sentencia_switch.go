package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
)

type Sentencia_switch struct {
	Expresion    BaseNodo
	Lista_case   []Sentencia_case
	Default_case BaseNodo
}

func (s Sentencia_switch) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	ambito_local := &ambito.Ambito{NombreAmbito: "Sentencia switch", Padre: ambito_padre}
	ambito_padre.AgregarAmbito(ambito_local)
	resultado := s.Expresion.Ejecutar(ambito_padre)
	if resultado.Tipo_dimension != valor.DIMENSION0 || resultado.Is_intancia || resultado.Type == valor.BOOLEAN {
		ambito_local.Agregar_error("No se puede evaluar los casos del case")
		return valor.Value{}
	}
	// quiza guardar en el stack el resultado a evaluar en el case para perder el valor porque
	// probablemente sea un temporal
	label_prueba := generador.Mi_generador.NewLabel()
	label_salida := generador.Mi_generador.NewLabel()
	generador.Mi_generador.AddGoto(label_prueba)
	var labels_prueba []string
	for _, sentencia_case := range s.Lista_case {
		label_case := generador.Mi_generador.NewLabel()
		labels_prueba = append(labels_prueba, label_case)
		generador.Mi_generador.AddLabel(label_case)
		sentencia_case.Ejecutar(ambito_local) //agregar sentencias del case
		generador.Mi_generador.AddGoto(label_salida)
	}
	//agregar el caso default
	label_default := generador.Mi_generador.NewLabel()
	generador.Mi_generador.AddLabel(label_default)
	if s.Default_case != nil {
		s.Default_case.Ejecutar(ambito_local)
	}
	generador.Mi_generador.AddGoto(label_salida)

	generador.Mi_generador.AddLabel(label_prueba)
	for i, sentencia_case := range s.Lista_case {
		resultado_case := sentencia_case.Expresion.Ejecutar(ambito_local)
		if resultado_case.Type != resultado.Type {
			ambito_local.Agregar_error("El tipo de la expresion del case no coincide")
			return valor.Value{}
		}
		if resultado.Type == valor.STRING {
			// comparar dos string de nuevo xd
		} else {
			generador.Mi_generador.AddIf(resultado.Value, resultado_case.Value, "==", labels_prueba[i])
		}
	}
	generador.Mi_generador.AddGoto(label_default)
	generador.Mi_generador.AddLabel(label_salida)
	return valor.Value{}
}

type Sentencia_case struct {
	Expresion  BaseNodo
	Sentencias []BaseNodo
}

func (a Sentencia_case) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	for _, linea := range a.Sentencias {
		linea.Ejecutar(ambito_padre)
	}
	return valor.Value{}
}

type Default_case struct {
	Sentencias []BaseNodo
}

func (a Default_case) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	for _, linea := range a.Sentencias {
		linea.Ejecutar(ambito_padre)
	}
	return valor.Value{}
}
