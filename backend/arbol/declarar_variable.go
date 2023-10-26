package arbol

import (
	"main/ambito"
	"main/valor"
)

// nodo declarar variable

type Declarar_variable struct {
	Id        string
	Tipo      string   // String, Int, Bool, Float, char, (Nombre_struct)
	Expresion BaseNodo // 10, 20.5, "hola", true, objeto_strcut
}

func (d Declarar_variable) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	return valor.Value{}
}

func Copiar_objeto_struct(original ambito.Objeto_struct) ambito.Objeto_struct {
	copia_hijos := make([]*ambito.Identificadores, len(original.Ambito_struct.Locales))
	for i, val := range original.Ambito_struct.Locales {
		// Crear una nueva variable para el puntero y copiar el valor apuntado
		copia_hijos[i] = &ambito.Identificadores{}
		*copia_hijos[i] = *val
	}
	copia_ambito := ambito.Ambito{}
	copia_ambito = *original.Ambito_struct
	valor := ambito.Objeto_struct{Ambito_struct: &copia_ambito}
	valor.Ambito_struct.Locales = copia_hijos
	return valor
}
