package arbol

import "main/ambito"

type Incremento_variable struct {
	Id        string
	Expresion BaseNodo
}

func (i Incremento_variable) Ejecutar(ambito *ambito.Ambito) interface{} {
	encontrado := ambito.BuscarID(i.Id)
	if encontrado != nil {
		resultado := i.Expresion.Ejecutar(ambito)
		switch rr := resultado.(type) {
		case int:
			if encontrado.Primitivo == "Int" {
				if encontrado.Referencia {
					encontrado.Puntero_valor.Valor = encontrado.Puntero_valor.Valor.(int) + rr
					return nil
				}
				encontrado.Valor = encontrado.Valor.(int) + rr
				return nil
			}
			if encontrado.Primitivo == "Float" {
				if encontrado.Referencia {
					encontrado.Puntero_valor.Valor = encontrado.Puntero_valor.Valor.(float64) + float64(rr)
					return nil
				}
				encontrado.Valor = encontrado.Valor.(float64) + float64(rr)
				return nil
			}
		case float64:
			if encontrado.Primitivo == "Float" {
				if encontrado.Referencia {
					encontrado.Puntero_valor.Valor = encontrado.Puntero_valor.Valor.(float64) + rr
					return nil
				}
				encontrado.Valor = encontrado.Valor.(float64) + rr
				return nil
			}
		case string:
			if encontrado.Primitivo == "String" {
				if encontrado.Referencia {
					encontrado.Puntero_valor.Valor = encontrado.Puntero_valor.Valor.(string) + rr
					return nil
				}
				encontrado.Valor = encontrado.Valor.(string) + rr
				return nil
			}
		default:
			panic("Operacion invalida += " + i.Id)
		}
	}
	panic("Error no se puede asignar, ID no existe " + i.Id)
}
