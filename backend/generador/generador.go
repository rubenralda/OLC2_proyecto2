package generador

import (
	"fmt"
)

type Generator struct {
	Temporal        int
	Label           int
	Code            []interface{}
	FinalCode       []interface{}
	Natives         []interface{}
	FuncCode        []interface{}
	TempList        []interface{}
	PrintStringFlag bool
	Is_int_print    bool //si ya fue agregada la funcion
	MainCode        bool
}

var Mi_generador *Generator

func (g Generator) GetCode() []interface{} {
	return g.Code
}

func (g Generator) GetFinalCode() []interface{} {
	return g.FinalCode
}

func (g Generator) GetTemps() []interface{} {
	return g.TempList
}

func (g *Generator) SetMainFlag(newVal bool) {
	g.MainCode = newVal
}

// Generar un nuevo temporal
func (g *Generator) NewTemp() string {
	temp := "t" + fmt.Sprintf("%v", g.Temporal)
	g.Temporal = g.Temporal + 1
	//Lo guardamos para declararlo
	g.TempList = append(g.TempList, temp)
	return temp
}

// Generador de Label
func (g *Generator) NewLabel() string {
	temp := g.Label
	g.Label = g.Label + 1
	return "L" + fmt.Sprintf("%v", temp)
}

// Añade Label al codigo
func (g *Generator) AddLabel(Label string) {
	if g.MainCode {
		g.Code = append(g.Code, Label+":\n")
	} else {
		g.FuncCode = append(g.FuncCode, Label+":\n")
	}
}

func (g *Generator) AddIf(left string, right string, operator string, Label string) {
	if g.MainCode {
		g.Code = append(g.Code, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
	}
}

func (g *Generator) AddGoto(Label string) {
	if g.MainCode {
		g.Code = append(g.Code, "goto "+Label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "goto "+Label+";\n")
	}
}

func (g *Generator) AddReturn() {
	if g.MainCode {
		g.Code = append(g.Code, "return"+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "return"+";\n")
	}
}

func (g *Generator) AddExpression(target string, left string, right string, operator string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = "+left+" "+operator+" "+right+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = "+left+" "+operator+" "+right+";\n")
	}
}

func (g *Generator) AddAssign(target, val string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = "+val+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = "+val+";\n")
	}
}

func (g *Generator) AddPrintf(typePrint string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "printf(\"%"+typePrint+"\", "+value+");\n")
	} else {
		g.FuncCode = append(g.FuncCode, "printf(\"%"+typePrint+"\", "+value+");\n")
	}
}

func (g *Generator) AddSetHeap(index string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "heap["+index+"] = "+value+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "heap["+index+"] = "+value+";\n")
	}
}

func (g *Generator) AddGetHeap(target string, index string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = heap["+index+"];\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = heap["+index+"];\n")
	}
}

func (g *Generator) AddSetStack(index string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "stack["+index+"] = "+value+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "stack["+index+"] = "+value+";\n")
	}
}

func (g *Generator) AddGetStack(target string, index string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = stack["+index+"];\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = stack["+index+"];\n")
	}
}

func (g *Generator) AddCall(target string) {
	if g.MainCode {
		g.Code = append(g.Code, target+"();\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+"();\n")
	}
}

func (g *Generator) AddBr() {
	if g.MainCode {
		g.Code = append(g.Code, "\n")
	} else {
		g.FuncCode = append(g.FuncCode, "\n")
	}
}

func (g *Generator) AddComment(target string) {
	if g.MainCode {
		g.Code = append(g.Code, "//"+target+"\n")
	} else {
		g.FuncCode = append(g.FuncCode, "//"+target+"\n")
	}
}

func (g *Generator) AddTittle(target string) {
	if g.MainCode {
		g.Code = append(g.Code, "void "+target+"() {\n")
	} else {
		g.FuncCode = append(g.FuncCode, "void "+target+"() {\n")
	}
}

func (g *Generator) AddEnd() {
	if g.MainCode {
		g.Code = append(g.Code, "\treturn;\n")
		g.Code = append(g.Code, "}\n\n")
	} else {
		g.FuncCode = append(g.FuncCode, "\treturn;\n")
		g.FuncCode = append(g.FuncCode, "}\n\n")
	}
}

// agregar headers
func (g *Generator) GenerateFinalCode() {
	//****************** add head
	g.FinalCode = append(g.FinalCode, "/*------HEADER------*/\n")
	g.FinalCode = append(g.FinalCode, "#include <stdio.h>\n")
	g.FinalCode = append(g.FinalCode, "#include <math.h>\n")
	g.FinalCode = append(g.FinalCode, "double heap[30101999];\n")
	g.FinalCode = append(g.FinalCode, "double stack[30101999];\n")
	g.FinalCode = append(g.FinalCode, "double P;\n")
	g.FinalCode = append(g.FinalCode, "double H;\n")
	g.FinalCode = append(g.FinalCode, "double ")
	//****************** add temporal declaration
	tempArr := g.GetTemps()
	if len(tempArr) > 0 {
		tmpDec := fmt.Sprintf("%v", tempArr[0])
		tempArr = tempArr[1:]
		for _, s := range tempArr {
			tmpDec += ", "
			tmpDec += fmt.Sprintf("%v", s)
		}
		tmpDec += ";\n\n"
		g.FinalCode = append(g.FinalCode, tmpDec)
	}
	//****************** add natives functions
	if len(g.Natives) > 0 {
		g.FinalCode = append(g.FinalCode, "/*------NATIVES------*/\n")
		g.FinalCode = append(g.FinalCode, g.Natives...)

	}
	//****************** add functions
	if len(g.FuncCode) > 0 {
		g.FinalCode = append(g.FinalCode, "/*------FUNCTIONS------*/\n")
		g.FinalCode = append(g.FinalCode, g.FuncCode...)

	}
	//****************** add main
	g.FinalCode = append(g.FinalCode, "/*------MAIN------*/\n")
	g.FinalCode = append(g.FinalCode, "int main() {\n")
	g.FinalCode = append(g.FinalCode, "\tP = 0; H = 0;\n\n")
	for _, s := range g.Code {
		g.FinalCode = append(g.FinalCode, "\t"+s.(string))
	}
	g.FinalCode = append(g.FinalCode, "\n\treturn 0;\n}\n")
}

func (g *Generator) GeneratePrintString() {
	if g.PrintStringFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		//se genera la funcion printstring
		g.Natives = append(g.Natives, "void dbrust_printString() {\n")
		g.Natives = append(g.Natives, "\t"+newTemp1+" = P + 1;\n")
		g.Natives = append(g.Natives, "\t"+newTemp2+" = stack[(int)"+newTemp1+"];\n")
		g.Natives = append(g.Natives, "\t"+newLvl2+":\n")
		g.Natives = append(g.Natives, "\t"+newTemp3+" = heap[(int)"+newTemp2+"];\n")
		g.Natives = append(g.Natives, "\tif("+newTemp3+" == -1) goto "+newLvl1+";\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", (char)"+newTemp3+");\n")
		g.Natives = append(g.Natives, "\t"+newTemp2+" = "+newTemp2+" + 1;\n")
		g.Natives = append(g.Natives, "\tgoto "+newLvl2+";\n")
		g.Natives = append(g.Natives, "\t"+newLvl1+":\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")
		g.PrintStringFlag = false
	}
}

func (g *Generator) Generar_funcion_int() {
	if g.Is_int_print {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		valor_char := g.NewTemp()
		label_salida := g.NewLabel()
		label_inicio := g.NewLabel()
		label_punto := g.NewLabel()
		label_fin_cadena := g.NewLabel()
		label_preparativos := g.NewLabel()
		puntero_heap := g.NewTemp()
		numero := g.NewTemp()
		multiplicador := g.NewTemp()
		//se genera la funcion
		g.Natives = append(g.Natives, "void funcion_int_float_rubencin() {\n")
		g.Natives = append(g.Natives, "\t"+newTemp1+" = P + 1;\n")
		g.Natives = append(g.Natives, "\t"+newTemp2+" = stack[(int)"+newTemp1+"];\n") //valor float
		g.Natives = append(g.Natives, "\t"+newTemp2+" = (int)"+newTemp2+";\n")
		g.Natives = append(g.Natives, "\tstack[(int)P] = "+newTemp2+";\n") //valor de retorno
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")
		//funcion int para string
		g.Natives = append(g.Natives, "void funcion_int_string_rubencin() {\n")
		g.Natives = append(g.Natives, "\t"+newTemp1+" = P + 1;\n")
		g.Natives = append(g.Natives, "\t"+puntero_heap+" = stack[(int)"+newTemp1+"];\n") //puntero heap string
		g.Natives = append(g.Natives, "\t"+label_fin_cadena+":\n")                        //inicio ciclo
		g.Natives = append(g.Natives, "\t"+valor_char+" = heap[(int)"+puntero_heap+"];\n")
		g.Natives = append(g.Natives, "\tif("+valor_char+" == -1) goto "+label_preparativos+";\n")
		g.Natives = append(g.Natives, "\tif("+valor_char+" == 46) goto "+label_preparativos+";\n")
		g.Natives = append(g.Natives, "\t"+puntero_heap+" = "+puntero_heap+" + 1;\n")
		g.Natives = append(g.Natives, "\tgoto "+label_fin_cadena+";\n")
		g.Natives = append(g.Natives, "\t"+label_preparativos+":\n")

		g.Natives = append(g.Natives, "\t"+puntero_heap+" = "+puntero_heap+" - 1;\n")
		g.Natives = append(g.Natives, "\t"+multiplicador+" = 1;\n")
		g.Natives = append(g.Natives, "\t"+numero+" = 0;\n")

		g.Natives = append(g.Natives, "\t"+label_inicio+":\n") //inicio ciclo
		g.Natives = append(g.Natives, "\t"+valor_char+" = heap[(int)"+puntero_heap+"];\n")
		g.Natives = append(g.Natives, "\tif("+valor_char+" == -1) goto "+label_salida+";\n")
		g.Natives = append(g.Natives, "\t"+puntero_heap+" = "+puntero_heap+" - 1;\n")
		g.Natives = append(g.Natives, "\tif("+valor_char+" < 48) goto "+label_punto+";\n")
		g.Natives = append(g.Natives, "\tif("+valor_char+" > 57) goto "+label_punto+";\n")

		g.Natives = append(g.Natives, "\t"+valor_char+" = "+valor_char+" - 48;\n")
		g.Natives = append(g.Natives, "\t"+valor_char+" = "+valor_char+" * "+multiplicador+";\n")
		g.Natives = append(g.Natives, "\t"+multiplicador+" = "+multiplicador+" * 10;\n")
		g.Natives = append(g.Natives, "\t"+numero+" = "+numero+" + "+valor_char+";\n")
		g.Natives = append(g.Natives, "\tgoto "+label_inicio+";\n")

		g.Natives = append(g.Natives, "\t"+label_punto+":\n")
		g.Natives = append(g.Natives, "\tif("+valor_char+" == 46) goto "+label_salida+";\n")
		//si no es punto error y retorno cero
		g.Natives = append(g.Natives, "\tprintf(\"%c\", (char) 69);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", (char) 82);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", (char) 82);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", (char) 79);\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", (char) 82);\n")
		g.Natives = append(g.Natives, "\t"+numero+" = 0;\n")

		g.Natives = append(g.Natives, "\t"+label_salida+":\n")
		g.Natives = append(g.Natives, "\tstack[(int)P] = "+numero+";\n") //valor de retorno
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")
		g.Is_int_print = false
	}
}
