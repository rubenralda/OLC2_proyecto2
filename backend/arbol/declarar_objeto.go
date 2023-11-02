package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Declarar_objeto struct {
	Id    string
	Dupla []Dupla_atributos
}

func (d Declarar_objeto) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	estruct, _ := ambito_padre.Buscar_struct(d.Id)
	if estruct == nil {
		panic("El struct no esta definido " + d.Id)
	}
	if len(d.Dupla) != len(estruct.Atributos) {
		panic("La cantidad de atributos no coincide " + d.Id)
	}
	tmp_inicio_struct := generador.Mi_generador.NewTemp()
	copia_inicio := generador.Mi_generador.NewTemp()
	generador.Mi_generador.AddAssign(tmp_inicio_struct, "H")                        //si hay una llamada recursiva este valor se pierde xd
	generador.Mi_generador.AddExpression("H", "H", strconv.Itoa(len(d.Dupla)), "+") //reserver el espacio en el heap
	for i, atributo := range d.Dupla {
		if atributo.Id_externo != estruct.Atributos[i].Id {
			panic("El nombre externo no coincide con el nombre de atributo")
		}
		resultado := atributo.Expresion.Ejecutar(ambito_padre)
		if resultado.Type == valor.NULL {
			if resultado.Tipo_struct != estruct.Atributos[i].Tipo_struct {
				panic("El tipo no coincide con el del atributo " + estruct.Atributos[i].Tipo_struct)
			}
			generador.Mi_generador.AddExpression(copia_inicio, tmp_inicio_struct, strconv.Itoa(i), "+")
			generador.Mi_generador.AddSetHeap("(int)"+copia_inicio, resultado.Value)
		} else if resultado.Type == estruct.Atributos[i].Tipo {
			if resultado.Type == valor.BOOLEAN {
				label_salida := generador.Mi_generador.NewLabel()
				for _, label := range resultado.TrueLabel {
					generador.Mi_generador.AddLabel(label.(string))
				}
				generador.Mi_generador.AddExpression(copia_inicio, tmp_inicio_struct, strconv.Itoa(i), "+")
				generador.Mi_generador.AddSetHeap("(int)"+copia_inicio, "1")
				generador.Mi_generador.AddGoto(label_salida)
				for _, label := range resultado.FalseLabel {
					generador.Mi_generador.AddLabel(label.(string))
				}
				generador.Mi_generador.AddExpression(copia_inicio, tmp_inicio_struct, strconv.Itoa(i), "+")
				generador.Mi_generador.AddSetHeap("(int)"+copia_inicio, "0")
				generador.Mi_generador.AddLabel(label_salida)
				generador.Mi_generador.AddExpression("H", "H", "1", "+")
			} else {
				generador.Mi_generador.AddExpression(copia_inicio, tmp_inicio_struct, strconv.Itoa(i), "+")
				generador.Mi_generador.AddSetHeap("(int)"+copia_inicio, resultado.Value)
				generador.Mi_generador.AddExpression("H", "H", "1", "+")
			}
		} else {
			panic("El tipo no coincide con el del atributo " + estruct.Atributos[i].Tipo_struct)
		}
	}
	return valor.Value{Value: tmp_inicio_struct, Is_intancia: true, Tipo_struct: d.Id}
}

type Dupla_atributos struct {
	Id_externo string
	Referencia bool // no se usa pero esta para convertir de Lista_argumentos
	Expresion  BaseNodo
}
