package arbol

import (
	"main/ambito"
	"main/generador"
	"main/valor"
	"strconv"
)

type Control_transfer struct {
	Sentencia_break    bool
	Sentencia_continue bool
	Sentencia_return   bool
	Retorno            BaseNodo
}

func (l Control_transfer) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	generador.Mi_generador.AddComment("bananero")
	if l.Sentencia_return {
		l.Retorno.Ejecutar(ambito_padre)
	} else {
		size_total, break_label, continue_label, existe := ambito_padre.Activacion_ciclo()
		if !existe {
			panic("Sentencia de control invalido")
		}
		if l.Sentencia_break {
			//posicion al ambito de ciclo
			generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(size_total), "-")
			generador.Mi_generador.AddGoto(break_label)
		} else if l.Sentencia_continue {
			generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(size_total), "-")
			generador.Mi_generador.AddGoto(continue_label)
		}
	}
	return valor.Value{}
}

type Sentencia_return struct {
	Expresion    BaseNodo
	Tipo_retorno bool //true si tiene tipo, false si no
}

func (s Sentencia_return) Ejecutar(ambito_padre *ambito.Ambito) valor.Value {
	size, activado, ambito := ambito_padre.Buscar_llamada_ambito()
	if activado {
		if !s.Tipo_retorno {
			if ambito.Tipo_funcion != valor.NULL {
				panic("No se esperaba un valor de retorno")
			}
			generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(size), "-")
			generador.Mi_generador.AddGoto(ambito.Label_salida)
		} else {
			resultado := s.Expresion.Ejecutar(ambito_padre)
			if resultado.Type != ambito.Tipo_funcion {
				panic("El tipo del retorno no coincide con el de la funcion")
			}
			generador.Mi_generador.AddExpression("P", "P", strconv.Itoa(size), "-")
			generador.Mi_generador.AddSetStack("(int)P", resultado.Value)
			generador.Mi_generador.AddGoto(ambito.Label_salida)
		}
	} else {
		panic("El retorno no esta en una funcion")
	}
	return valor.Value{}
}
