package arbol

import "main/ambito"

type Asignacion_variable struct {
	Id        string
	Expresion BaseNodo
}

func (a Asignacion_variable) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	encontrado := ambito_padre.BuscarID(a.Id)
	if encontrado != nil {
		if encontrado.Tipo == "constante" && encontrado.Tipo != "variable" {
			panic("No se puede modificar una constante")
		}
		resultado := a.Expresion.Ejecutar(ambito_padre)
		switch rr := resultado.(type) {
		case int:
			if encontrado.Primitivo == "Int" {
				if encontrado.Referencia {
					encontrado.Puntero_valor.Valor = rr
					return nil
				}
				encontrado.Valor = rr
				return nil
			}
			if encontrado.Primitivo == "Float" {
				if encontrado.Referencia {
					encontrado.Puntero_valor.Valor = float64(rr)
					return nil
				}
				encontrado.Valor = float64(rr)
				return nil
			}
		case float64:
			if encontrado.Primitivo == "Float" {
				if encontrado.Referencia {
					encontrado.Puntero_valor.Valor = rr
					return nil
				}
				encontrado.Valor = rr
				return nil
			}
		case string:
			if encontrado.Primitivo == "String" {
				if encontrado.Referencia {
					encontrado.Puntero_valor.Valor = rr
					return nil
				}
				encontrado.Valor = rr
				return nil
			}
		case bool:
			if encontrado.Primitivo == "Bool" {
				if encontrado.Referencia {
					encontrado.Puntero_valor.Valor = rr
					return nil
				}
				encontrado.Valor = rr
				return nil
			}
		case rune:
			if encontrado.Primitivo == "Character" {
				if encontrado.Referencia {
					encontrado.Puntero_valor.Valor = rr
					return nil
				}
				encontrado.Valor = rr
				return nil
			}
		case ambito.Objeto_struct:
			if encontrado.Primitivo == rr.Ambito_struct.NombreAmbito {
				if encontrado.Referencia {
					encontrado.Puntero_valor.Valor = rr
					return nil
				}
				encontrado.Valor = rr
				return nil
			}
		default:
			panic("Tipo no permitido " + a.Id)
		}
		panic("Los tipos no coinciden " + a.Id)
	}
	panic("Error el ID no existe " + a.Id)
}
