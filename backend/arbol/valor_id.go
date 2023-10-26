package arbol

import "main/ambito"

type Id_expresion struct {
	Id string
}

func (i Id_expresion) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	id_buscado := ambito_padre.BuscarID(i.Id)
	if id_buscado == nil {
		panic("Error el identificador no existe: " + i.Id)
	}
	if id_buscado.Tipo == "vector" || id_buscado.Tipo == "matriz" {
		return id_buscado.Lista_vector
	}
	if id_buscado.Tipo == "struct" {
		return id_buscado.Objeto
	}
	if id_buscado.Referencia {
		return id_buscado.Puntero_valor.Valor
	}
	return id_buscado.Valor
}

type Id_vector struct {
	Id     string
	Indice BaseNodo
}

func (i Id_vector) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	resul_indice := i.Indice.Ejecutar(ambito_padre)
	indice := 0
	switch rr := resul_indice.(type) {
	case int:
		indice = rr
	default:
		panic("El indice no es un entero")
	}
	id_buscado := ambito_padre.BuscarID(i.Id)
	if id_buscado == nil {
		panic("Error el vector no existe: " + i.Id)
	}
	if id_buscado.Tipo != "vector" {
		panic("El Id no corresponde a una variable o constante " + i.Id)
	}
	if indice >= 0 && indice < len(id_buscado.Lista_vector) {
		return id_buscado.Lista_vector[indice]
	} else if id_buscado.Referencia {
		if indice >= 0 && indice < len(id_buscado.Puntero_valor.Lista_vector) {
			return id_buscado.Puntero_valor.Lista_vector[indice]
		}
		//return nil // esto para errores
	}
	panic("El indice no existe")
}

type Id_matriz struct {
	Id      string
	Indices []BaseNodo
}

func (i Id_matriz) Ejecutar(ambito_padre *ambito.Ambito) interface{} {
	id_buscado := ambito_padre.BuscarID(i.Id)
	if id_buscado == nil {
		panic("Error la matriz no existe: " + i.Id)
	}
	if id_buscado.Tipo != "matriz" {
		panic("El Id no corresponde a una variable o constante " + i.Id)
	}
	valor := id_buscado.Lista_vector
	for j, indice := range i.Indices {
		resul_indice := indice.Ejecutar(ambito_padre)
		indice := 0
		switch rr := resul_indice.(type) {
		case int:
			indice = rr
		default:
			panic("El indice no es un entero")
		}
		if indice >= 0 && indice < len(valor) {
			switch r2 := valor[indice].(type) {
			case []interface{}:
				valor = r2
				continue
			default:
				if j != len(i.Indices)-1 {
					panic("Dimension extra " + id_buscado.Id)
				}
				return r2
			}
		}
		/*
			else if id_buscado.Referencia {
				if indice >= 0 && indice < len(id_buscado.Puntero_valor.Lista_vector) {
					return id_buscado.Puntero_valor.Lista_vector[indice]
				}
				//return nil // esto para errores
			}
		*/
	}
	return valor
}
