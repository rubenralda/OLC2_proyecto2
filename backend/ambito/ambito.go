package ambito

import "main/valor"

type Objeto_struct struct {
	Ambito_struct *Ambito
}

type Identificadores struct {
	Id                string              // tambien nombre interno de una funcion
	Primitivo         valor.TipoExpresion // String, Int, Bool, Float, char, (Nombre_struct)
	Posicion_relativa int                 // uso: p + posicion_relativa
	//Tipo              TipoVariable        //variable, constante, vector, funcion, struct, matriz y parametro
	Is_init bool // Si fue declarada con un valor o no

	Objeto       Objeto_struct // el objeto para el tipo struct
	Lista_vector []interface{} // array de datos para el tipo vector
	// parte de funciones
	Funcion       interface{}      // En realidad es arbol.Ejecutar_funcion pero recursividad indirecta de modulos
	Referencia    valor.Value      // true si es por referencia, tambien mutable o inmutable
	Puntero_valor *Identificadores // tiene un objeto identificadores que puede modificar
}

type Variables struct {
	Id                string              //nombre de la variable
	Posicion_relativa int                 //uso: p + posicion_relativa
	Tipo_dimension    valor.DIMENSION     //variable, vector, matriz
	Tipo              valor.TipoExpresion //INT, FLOAT, STRING, BOOL, CHAR, NULL
	Tipo_struct       string              //nombre tipo del struct (ID) solo si Is_init esta activada
	Is_instancia      bool                //si es un objeto, se usa el tipo struct
	Is_init           bool                //Si fue declarada con un valor o no
	Is_constante      bool                //Si fue declarada con let
}

type Funciones struct {
}

type Ambito struct {
	NombreAmbito string
	Padre        *Ambito
	AmbitosHijos []*Ambito
	Variables    []*Variables
	Funciones    []*Funciones
	Size         int

	//para llamadas
	Is_ambito_llamada bool // si esta activa desde aqui se hizo la llamada
	Tipo_funcion      valor.TipoExpresion

	//para ciclos
	Is_ciclo      bool //si esta activa es el ambito del ciclo
	BreakLabel    string
	ContinueLabel string
}

func (a *Ambito) AgregarVariable(variable Variables) {
	a.Variables = append(a.Variables, &variable)
}

func (a *Ambito) AgregarAmbito(ambito *Ambito) bool {
	a.AmbitosHijos = append(a.AmbitosHijos, ambito)
	return true
}

func (a *Ambito) BuscarVariable(id string) (*Variables, int) {
	anterior := a
	var elemento *Variables
	size := 0
	for {
		for _, value := range anterior.Variables {
			if value.Id == id {
				elemento = value
			}
		}
		if elemento != nil {
			break
		}
		anterior = anterior.Padre
		if anterior != nil {
			size += anterior.Size
		} else {
			break
		}
	}
	return elemento, size
}

//retorna el size al ambito donde se llamo y si se activo, 0 si es el actual
func (a *Ambito) Buscar_llamada_ambito() (int, bool, *Ambito) {
	anterior := a
	size := 0
	for {
		if anterior.Is_ambito_llamada {
			return size, true, anterior
		}
		anterior = anterior.Padre
		if anterior != nil {
			size += anterior.Size
		} else {
			break
		}
	}
	return size, false, nil
}

//retorna el size hasta el ambito del ultimo ciclo, etiquta break, etiqueta continue y si se encontro
func (a *Ambito) Activacion_ciclo() (int, string, string, bool) {
	anterior := a
	size := 0
	for {
		if anterior.Is_ciclo {
			size += anterior.Padre.Size
			return size, anterior.BreakLabel, anterior.ContinueLabel, true
		}
		anterior = anterior.Padre
		if anterior != nil {
			size += anterior.Size
		} else {
			break
		}
	}
	return size, anterior.BreakLabel, anterior.ContinueLabel, false
}
