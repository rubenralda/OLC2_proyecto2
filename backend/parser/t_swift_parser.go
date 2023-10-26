// Code generated from .\parser\T_swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // T_swift
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type T_swiftParser struct {
	*antlr.BaseParser
}

var T_swiftParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func t_swiftParserInit() {
	staticData := &T_swiftParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "'func'", "'('", "')'", "'->'", "','", "':'", "'inout'",
		"'{'", "'}'", "'struct'", "'let'", "'var'", "'='", "'mutating'", "'&'",
		"'.'", "'isEmpty'", "'count'", "'['", "']'", "'+='", "'-='", "'repeating'",
		"'for'", "'in'", "'...'", "'while'", "'if'", "'else'", "'guard'", "'continue'",
		"'break'", "'return'", "'switch'", "'case'", "'default'", "'?'", "'String'",
		"'Int'", "'Float'", "'Bool'", "'Character'", "'print'", "'append'",
		"'removeLast'", "'remove'", "'at'", "'!'", "'-'", "'/'", "'%'", "'*'",
		"'+'", "'<'", "'<='", "'>='", "'>'", "'=='", "'!='", "'&&'", "'||'",
		"'nil'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "Block_comment",
		"Line_comment", "Int", "Float", "Caracter", "String", "Bool", "Identificador",
	}
	staticData.RuleNames = []string{
		"inicio", "instrucciones_globales", "intruccion_global", "function_declaracion",
		"lista_parametros", "declaracion_parametro", "code_block", "struct_declaracion",
		"lista_atributos", "instrucciones", "instruccion", "declaracion", "loop_statement",
		"branch_statement", "control_transfer_statement", "llamadas_funciones",
		"llamada_normal", "lista_argumentos", "declaracion_argumento", "llamada_metodos",
		"atributos", "ide_atributo", "asignar_atributos", "matriz_declaracion",
		"tipo_matriz", "definicion_matriz", "lista_valores_matriz", "elementos_matriz",
		"elemento_matriz", "simple_matriz", "for_in_statement", "rango", "while_statement",
		"if_statement", "guard_statement", "switch_statement", "switch_case",
		"default_case", "break_statement", "continue_statement", "return_statement",
		"constant_declaracion", "variable_declaracion", "anotacion_tipo", "tipos",
		"array_declaracion", "definicion_vector", "lista_expresiones", "funcion_print",
		"funcion_append", "funcion_removeLast", "funcion_removeat", "funcion_int",
		"funcion_float", "funcion_string", "asignacion", "expresion", "l_duble",
		"primitivos",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 72, 741, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 123, 8, 1, 10, 1, 12, 1, 126,
		9, 1, 1, 2, 1, 2, 3, 2, 130, 8, 2, 1, 2, 1, 2, 3, 2, 134, 8, 2, 1, 2, 1,
		2, 3, 2, 138, 8, 2, 1, 2, 1, 2, 3, 2, 142, 8, 2, 1, 2, 1, 2, 3, 2, 146,
		8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 152, 8, 2, 3, 2, 154, 8, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 3, 3, 160, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 165, 8, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 172, 8, 4, 10, 4, 12, 4, 175, 9, 4, 1, 5,
		3, 5, 178, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 183, 8, 5, 1, 5, 1, 5, 3, 5, 187,
		8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 192, 8, 5, 1, 5, 3, 5, 195, 8, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 205, 8, 7, 10, 7, 12, 7, 208,
		9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 216, 8, 8, 1, 8, 1, 8,
		3, 8, 220, 8, 8, 1, 8, 3, 8, 223, 8, 8, 1, 8, 3, 8, 226, 8, 8, 1, 8, 3,
		8, 229, 8, 8, 1, 9, 5, 9, 232, 8, 9, 10, 9, 12, 9, 235, 9, 9, 1, 10, 1,
		10, 3, 10, 239, 8, 10, 1, 10, 1, 10, 3, 10, 243, 8, 10, 1, 10, 1, 10, 3,
		10, 247, 8, 10, 1, 10, 1, 10, 3, 10, 251, 8, 10, 1, 10, 1, 10, 3, 10, 255,
		8, 10, 1, 10, 1, 10, 3, 10, 259, 8, 10, 1, 10, 1, 10, 3, 10, 263, 8, 10,
		3, 10, 265, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 271, 8, 11, 1, 12,
		1, 12, 3, 12, 275, 8, 12, 1, 13, 1, 13, 1, 13, 3, 13, 280, 8, 13, 1, 14,
		1, 14, 1, 14, 3, 14, 285, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 3, 15, 296, 8, 15, 1, 16, 1, 16, 1, 16, 3, 16,
		301, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 5, 17, 308, 8, 17, 10, 17,
		12, 17, 311, 9, 17, 1, 18, 1, 18, 3, 18, 315, 8, 18, 1, 18, 3, 18, 318,
		8, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 327, 8,
		19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 4, 20, 340, 8, 20, 11, 20, 12, 20, 341, 3, 20, 344, 8, 20, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 352, 8, 21, 1, 22, 1, 22, 1,
		22, 4, 22, 357, 8, 22, 11, 22, 12, 22, 358, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 4, 22, 367, 8, 22, 11, 22, 12, 22, 368, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 4, 22, 377, 8, 22, 11, 22, 12, 22, 378, 1, 22,
		1, 22, 1, 22, 3, 22, 384, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 390,
		8, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 3, 24, 403, 8, 24, 1, 25, 1, 25, 3, 25, 407, 8, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 5, 27, 416, 8, 27, 10, 27, 12, 27,
		419, 9, 27, 1, 28, 1, 28, 3, 28, 423, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 447, 8, 29,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 454, 8, 30, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 3, 33, 482, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 496, 8, 35, 10,
		35, 12, 35, 499, 9, 35, 1, 35, 3, 35, 502, 8, 35, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 511, 8, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 3, 37, 517, 8, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 3, 40,
		525, 8, 40, 1, 41, 1, 41, 1, 41, 3, 41, 530, 8, 41, 1, 41, 1, 41, 1, 41,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 543, 8,
		42, 1, 42, 1, 42, 3, 42, 547, 8, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 571, 8, 45, 1, 46,
		1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 583,
		8, 46, 1, 47, 1, 47, 1, 47, 5, 47, 588, 8, 47, 10, 47, 12, 47, 591, 9,
		47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1,
		51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52,
		1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1,
		54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55,
		1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1,
		55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 3, 55,
		665, 8, 55, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1,
		56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 4, 56, 683, 8, 56,
		11, 56, 12, 56, 684, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1,
		56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 3, 56, 702, 8, 56,
		1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1,
		56, 1, 56, 5, 56, 716, 8, 56, 10, 56, 12, 56, 719, 9, 56, 1, 57, 1, 57,
		1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 5, 57, 728, 8, 57, 10, 57, 12, 57, 731,
		9, 57, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 3, 58, 739, 8, 58, 1,
		58, 0, 1, 112, 59, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100,
		102, 104, 106, 108, 110, 112, 114, 116, 0, 7, 1, 0, 12, 13, 1, 0, 32, 34,
		2, 0, 39, 43, 72, 72, 1, 0, 51, 53, 2, 0, 50, 50, 54, 54, 1, 0, 55, 60,
		1, 0, 61, 62, 799, 0, 118, 1, 0, 0, 0, 2, 124, 1, 0, 0, 0, 4, 153, 1, 0,
		0, 0, 6, 155, 1, 0, 0, 0, 8, 168, 1, 0, 0, 0, 10, 194, 1, 0, 0, 0, 12,
		196, 1, 0, 0, 0, 14, 200, 1, 0, 0, 0, 16, 228, 1, 0, 0, 0, 18, 233, 1,
		0, 0, 0, 20, 264, 1, 0, 0, 0, 22, 270, 1, 0, 0, 0, 24, 274, 1, 0, 0, 0,
		26, 279, 1, 0, 0, 0, 28, 284, 1, 0, 0, 0, 30, 295, 1, 0, 0, 0, 32, 297,
		1, 0, 0, 0, 34, 304, 1, 0, 0, 0, 36, 314, 1, 0, 0, 0, 38, 321, 1, 0, 0,
		0, 40, 343, 1, 0, 0, 0, 42, 351, 1, 0, 0, 0, 44, 383, 1, 0, 0, 0, 46, 385,
		1, 0, 0, 0, 48, 402, 1, 0, 0, 0, 50, 406, 1, 0, 0, 0, 52, 408, 1, 0, 0,
		0, 54, 412, 1, 0, 0, 0, 56, 422, 1, 0, 0, 0, 58, 446, 1, 0, 0, 0, 60, 448,
		1, 0, 0, 0, 62, 457, 1, 0, 0, 0, 64, 461, 1, 0, 0, 0, 66, 481, 1, 0, 0,
		0, 68, 483, 1, 0, 0, 0, 70, 491, 1, 0, 0, 0, 72, 505, 1, 0, 0, 0, 74, 512,
		1, 0, 0, 0, 76, 518, 1, 0, 0, 0, 78, 520, 1, 0, 0, 0, 80, 522, 1, 0, 0,
		0, 82, 526, 1, 0, 0, 0, 84, 546, 1, 0, 0, 0, 86, 548, 1, 0, 0, 0, 88, 551,
		1, 0, 0, 0, 90, 570, 1, 0, 0, 0, 92, 582, 1, 0, 0, 0, 94, 584, 1, 0, 0,
		0, 96, 592, 1, 0, 0, 0, 98, 597, 1, 0, 0, 0, 100, 604, 1, 0, 0, 0, 102,
		610, 1, 0, 0, 0, 104, 619, 1, 0, 0, 0, 106, 624, 1, 0, 0, 0, 108, 629,
		1, 0, 0, 0, 110, 664, 1, 0, 0, 0, 112, 701, 1, 0, 0, 0, 114, 720, 1, 0,
		0, 0, 116, 738, 1, 0, 0, 0, 118, 119, 3, 2, 1, 0, 119, 120, 5, 0, 0, 1,
		120, 1, 1, 0, 0, 0, 121, 123, 3, 4, 2, 0, 122, 121, 1, 0, 0, 0, 123, 126,
		1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 3, 1, 0, 0,
		0, 126, 124, 1, 0, 0, 0, 127, 129, 3, 22, 11, 0, 128, 130, 5, 1, 0, 0,
		129, 128, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 154, 1, 0, 0, 0, 131,
		133, 3, 24, 12, 0, 132, 134, 5, 1, 0, 0, 133, 132, 1, 0, 0, 0, 133, 134,
		1, 0, 0, 0, 134, 154, 1, 0, 0, 0, 135, 137, 3, 26, 13, 0, 136, 138, 5,
		1, 0, 0, 137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 154, 1, 0, 0,
		0, 139, 141, 3, 110, 55, 0, 140, 142, 5, 1, 0, 0, 141, 140, 1, 0, 0, 0,
		141, 142, 1, 0, 0, 0, 142, 154, 1, 0, 0, 0, 143, 145, 3, 30, 15, 0, 144,
		146, 5, 1, 0, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 154,
		1, 0, 0, 0, 147, 154, 3, 6, 3, 0, 148, 154, 3, 14, 7, 0, 149, 151, 3, 44,
		22, 0, 150, 152, 5, 1, 0, 0, 151, 150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0,
		152, 154, 1, 0, 0, 0, 153, 127, 1, 0, 0, 0, 153, 131, 1, 0, 0, 0, 153,
		135, 1, 0, 0, 0, 153, 139, 1, 0, 0, 0, 153, 143, 1, 0, 0, 0, 153, 147,
		1, 0, 0, 0, 153, 148, 1, 0, 0, 0, 153, 149, 1, 0, 0, 0, 154, 5, 1, 0, 0,
		0, 155, 156, 5, 2, 0, 0, 156, 157, 5, 72, 0, 0, 157, 159, 5, 3, 0, 0, 158,
		160, 3, 8, 4, 0, 159, 158, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161,
		1, 0, 0, 0, 161, 164, 5, 4, 0, 0, 162, 163, 5, 5, 0, 0, 163, 165, 3, 88,
		44, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0,
		166, 167, 3, 12, 6, 0, 167, 7, 1, 0, 0, 0, 168, 173, 3, 10, 5, 0, 169,
		170, 5, 6, 0, 0, 170, 172, 3, 10, 5, 0, 171, 169, 1, 0, 0, 0, 172, 175,
		1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 9, 1, 0, 0,
		0, 175, 173, 1, 0, 0, 0, 176, 178, 5, 72, 0, 0, 177, 176, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 5, 72, 0, 0, 180, 182,
		5, 7, 0, 0, 181, 183, 5, 8, 0, 0, 182, 181, 1, 0, 0, 0, 182, 183, 1, 0,
		0, 0, 183, 184, 1, 0, 0, 0, 184, 195, 3, 88, 44, 0, 185, 187, 5, 72, 0,
		0, 186, 185, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188,
		189, 5, 72, 0, 0, 189, 191, 5, 7, 0, 0, 190, 192, 5, 8, 0, 0, 191, 190,
		1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 195, 3, 48,
		24, 0, 194, 177, 1, 0, 0, 0, 194, 186, 1, 0, 0, 0, 195, 11, 1, 0, 0, 0,
		196, 197, 5, 9, 0, 0, 197, 198, 3, 18, 9, 0, 198, 199, 5, 10, 0, 0, 199,
		13, 1, 0, 0, 0, 200, 201, 5, 11, 0, 0, 201, 202, 5, 72, 0, 0, 202, 206,
		5, 9, 0, 0, 203, 205, 3, 16, 8, 0, 204, 203, 1, 0, 0, 0, 205, 208, 1, 0,
		0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 209, 1, 0, 0, 0,
		208, 206, 1, 0, 0, 0, 209, 210, 5, 10, 0, 0, 210, 15, 1, 0, 0, 0, 211,
		212, 7, 0, 0, 0, 212, 215, 5, 72, 0, 0, 213, 214, 5, 7, 0, 0, 214, 216,
		3, 88, 44, 0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 219, 1,
		0, 0, 0, 217, 218, 5, 14, 0, 0, 218, 220, 3, 112, 56, 0, 219, 217, 1, 0,
		0, 0, 219, 220, 1, 0, 0, 0, 220, 222, 1, 0, 0, 0, 221, 223, 5, 1, 0, 0,
		222, 221, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 229, 1, 0, 0, 0, 224,
		226, 5, 15, 0, 0, 225, 224, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 227,
		1, 0, 0, 0, 227, 229, 3, 6, 3, 0, 228, 211, 1, 0, 0, 0, 228, 225, 1, 0,
		0, 0, 229, 17, 1, 0, 0, 0, 230, 232, 3, 20, 10, 0, 231, 230, 1, 0, 0, 0,
		232, 235, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234,
		19, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 236, 238, 3, 22, 11, 0, 237, 239,
		5, 1, 0, 0, 238, 237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 265, 1, 0,
		0, 0, 240, 242, 3, 24, 12, 0, 241, 243, 5, 1, 0, 0, 242, 241, 1, 0, 0,
		0, 242, 243, 1, 0, 0, 0, 243, 265, 1, 0, 0, 0, 244, 246, 3, 26, 13, 0,
		245, 247, 5, 1, 0, 0, 246, 245, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247,
		265, 1, 0, 0, 0, 248, 250, 3, 28, 14, 0, 249, 251, 5, 1, 0, 0, 250, 249,
		1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 265, 1, 0, 0, 0, 252, 254, 3, 110,
		55, 0, 253, 255, 5, 1, 0, 0, 254, 253, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0,
		255, 265, 1, 0, 0, 0, 256, 258, 3, 30, 15, 0, 257, 259, 5, 1, 0, 0, 258,
		257, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 265, 1, 0, 0, 0, 260, 262,
		3, 44, 22, 0, 261, 263, 5, 1, 0, 0, 262, 261, 1, 0, 0, 0, 262, 263, 1,
		0, 0, 0, 263, 265, 1, 0, 0, 0, 264, 236, 1, 0, 0, 0, 264, 240, 1, 0, 0,
		0, 264, 244, 1, 0, 0, 0, 264, 248, 1, 0, 0, 0, 264, 252, 1, 0, 0, 0, 264,
		256, 1, 0, 0, 0, 264, 260, 1, 0, 0, 0, 265, 21, 1, 0, 0, 0, 266, 271, 3,
		82, 41, 0, 267, 271, 3, 84, 42, 0, 268, 271, 3, 90, 45, 0, 269, 271, 3,
		46, 23, 0, 270, 266, 1, 0, 0, 0, 270, 267, 1, 0, 0, 0, 270, 268, 1, 0,
		0, 0, 270, 269, 1, 0, 0, 0, 271, 23, 1, 0, 0, 0, 272, 275, 3, 60, 30, 0,
		273, 275, 3, 64, 32, 0, 274, 272, 1, 0, 0, 0, 274, 273, 1, 0, 0, 0, 275,
		25, 1, 0, 0, 0, 276, 280, 3, 66, 33, 0, 277, 280, 3, 68, 34, 0, 278, 280,
		3, 70, 35, 0, 279, 276, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279, 278, 1,
		0, 0, 0, 280, 27, 1, 0, 0, 0, 281, 285, 3, 76, 38, 0, 282, 285, 3, 78,
		39, 0, 283, 285, 3, 80, 40, 0, 284, 281, 1, 0, 0, 0, 284, 282, 1, 0, 0,
		0, 284, 283, 1, 0, 0, 0, 285, 29, 1, 0, 0, 0, 286, 296, 3, 96, 48, 0, 287,
		296, 3, 98, 49, 0, 288, 296, 3, 100, 50, 0, 289, 296, 3, 102, 51, 0, 290,
		296, 3, 104, 52, 0, 291, 296, 3, 106, 53, 0, 292, 296, 3, 108, 54, 0, 293,
		296, 3, 32, 16, 0, 294, 296, 3, 38, 19, 0, 295, 286, 1, 0, 0, 0, 295, 287,
		1, 0, 0, 0, 295, 288, 1, 0, 0, 0, 295, 289, 1, 0, 0, 0, 295, 290, 1, 0,
		0, 0, 295, 291, 1, 0, 0, 0, 295, 292, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0,
		295, 294, 1, 0, 0, 0, 296, 31, 1, 0, 0, 0, 297, 298, 5, 72, 0, 0, 298,
		300, 5, 3, 0, 0, 299, 301, 3, 34, 17, 0, 300, 299, 1, 0, 0, 0, 300, 301,
		1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 303, 5, 4, 0, 0, 303, 33, 1, 0,
		0, 0, 304, 309, 3, 36, 18, 0, 305, 306, 5, 6, 0, 0, 306, 308, 3, 36, 18,
		0, 307, 305, 1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309,
		310, 1, 0, 0, 0, 310, 35, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 312, 313, 5,
		72, 0, 0, 313, 315, 5, 7, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0,
		0, 315, 317, 1, 0, 0, 0, 316, 318, 5, 16, 0, 0, 317, 316, 1, 0, 0, 0, 317,
		318, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 320, 3, 112, 56, 0, 320, 37,
		1, 0, 0, 0, 321, 322, 5, 72, 0, 0, 322, 323, 5, 17, 0, 0, 323, 324, 5,
		72, 0, 0, 324, 326, 5, 3, 0, 0, 325, 327, 3, 34, 17, 0, 326, 325, 1, 0,
		0, 0, 326, 327, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 329, 5, 4, 0, 0,
		329, 39, 1, 0, 0, 0, 330, 331, 5, 72, 0, 0, 331, 332, 5, 17, 0, 0, 332,
		344, 5, 18, 0, 0, 333, 334, 5, 72, 0, 0, 334, 335, 5, 17, 0, 0, 335, 344,
		5, 19, 0, 0, 336, 339, 3, 42, 21, 0, 337, 338, 5, 17, 0, 0, 338, 340, 3,
		42, 21, 0, 339, 337, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 339, 1, 0,
		0, 0, 341, 342, 1, 0, 0, 0, 342, 344, 1, 0, 0, 0, 343, 330, 1, 0, 0, 0,
		343, 333, 1, 0, 0, 0, 343, 336, 1, 0, 0, 0, 344, 41, 1, 0, 0, 0, 345, 352,
		5, 72, 0, 0, 346, 347, 5, 72, 0, 0, 347, 348, 5, 20, 0, 0, 348, 349, 3,
		112, 56, 0, 349, 350, 5, 21, 0, 0, 350, 352, 1, 0, 0, 0, 351, 345, 1, 0,
		0, 0, 351, 346, 1, 0, 0, 0, 352, 43, 1, 0, 0, 0, 353, 356, 3, 42, 21, 0,
		354, 355, 5, 17, 0, 0, 355, 357, 3, 42, 21, 0, 356, 354, 1, 0, 0, 0, 357,
		358, 1, 0, 0, 0, 358, 356, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 360,
		1, 0, 0, 0, 360, 361, 5, 14, 0, 0, 361, 362, 3, 112, 56, 0, 362, 384, 1,
		0, 0, 0, 363, 366, 3, 42, 21, 0, 364, 365, 5, 17, 0, 0, 365, 367, 3, 42,
		21, 0, 366, 364, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0,
		368, 369, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 371, 5, 22, 0, 0, 371,
		372, 3, 112, 56, 0, 372, 384, 1, 0, 0, 0, 373, 376, 3, 42, 21, 0, 374,
		375, 5, 17, 0, 0, 375, 377, 3, 42, 21, 0, 376, 374, 1, 0, 0, 0, 377, 378,
		1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 380, 1, 0,
		0, 0, 380, 381, 5, 23, 0, 0, 381, 382, 3, 112, 56, 0, 382, 384, 1, 0, 0,
		0, 383, 353, 1, 0, 0, 0, 383, 363, 1, 0, 0, 0, 383, 373, 1, 0, 0, 0, 384,
		45, 1, 0, 0, 0, 385, 386, 5, 13, 0, 0, 386, 389, 5, 72, 0, 0, 387, 388,
		5, 7, 0, 0, 388, 390, 3, 48, 24, 0, 389, 387, 1, 0, 0, 0, 389, 390, 1,
		0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 392, 5, 14, 0, 0, 392, 393, 3, 50,
		25, 0, 393, 47, 1, 0, 0, 0, 394, 395, 5, 20, 0, 0, 395, 396, 3, 48, 24,
		0, 396, 397, 5, 21, 0, 0, 397, 403, 1, 0, 0, 0, 398, 399, 5, 20, 0, 0,
		399, 400, 3, 88, 44, 0, 400, 401, 5, 21, 0, 0, 401, 403, 1, 0, 0, 0, 402,
		394, 1, 0, 0, 0, 402, 398, 1, 0, 0, 0, 403, 49, 1, 0, 0, 0, 404, 407, 3,
		52, 26, 0, 405, 407, 3, 58, 29, 0, 406, 404, 1, 0, 0, 0, 406, 405, 1, 0,
		0, 0, 407, 51, 1, 0, 0, 0, 408, 409, 5, 20, 0, 0, 409, 410, 3, 54, 27,
		0, 410, 411, 5, 21, 0, 0, 411, 53, 1, 0, 0, 0, 412, 417, 3, 56, 28, 0,
		413, 414, 5, 6, 0, 0, 414, 416, 3, 56, 28, 0, 415, 413, 1, 0, 0, 0, 416,
		419, 1, 0, 0, 0, 417, 415, 1, 0, 0, 0, 417, 418, 1, 0, 0, 0, 418, 55, 1,
		0, 0, 0, 419, 417, 1, 0, 0, 0, 420, 423, 3, 52, 26, 0, 421, 423, 3, 112,
		56, 0, 422, 420, 1, 0, 0, 0, 422, 421, 1, 0, 0, 0, 423, 57, 1, 0, 0, 0,
		424, 425, 3, 48, 24, 0, 425, 426, 5, 3, 0, 0, 426, 427, 5, 24, 0, 0, 427,
		428, 5, 7, 0, 0, 428, 429, 3, 58, 29, 0, 429, 430, 5, 6, 0, 0, 430, 431,
		5, 19, 0, 0, 431, 432, 5, 7, 0, 0, 432, 433, 5, 67, 0, 0, 433, 434, 5,
		4, 0, 0, 434, 447, 1, 0, 0, 0, 435, 436, 3, 48, 24, 0, 436, 437, 5, 3,
		0, 0, 437, 438, 5, 24, 0, 0, 438, 439, 5, 7, 0, 0, 439, 440, 3, 112, 56,
		0, 440, 441, 5, 6, 0, 0, 441, 442, 5, 19, 0, 0, 442, 443, 5, 7, 0, 0, 443,
		444, 5, 67, 0, 0, 444, 445, 5, 4, 0, 0, 445, 447, 1, 0, 0, 0, 446, 424,
		1, 0, 0, 0, 446, 435, 1, 0, 0, 0, 447, 59, 1, 0, 0, 0, 448, 449, 5, 25,
		0, 0, 449, 450, 5, 72, 0, 0, 450, 453, 5, 26, 0, 0, 451, 454, 3, 112, 56,
		0, 452, 454, 3, 62, 31, 0, 453, 451, 1, 0, 0, 0, 453, 452, 1, 0, 0, 0,
		454, 455, 1, 0, 0, 0, 455, 456, 3, 12, 6, 0, 456, 61, 1, 0, 0, 0, 457,
		458, 3, 112, 56, 0, 458, 459, 5, 27, 0, 0, 459, 460, 3, 112, 56, 0, 460,
		63, 1, 0, 0, 0, 461, 462, 5, 28, 0, 0, 462, 463, 3, 112, 56, 0, 463, 464,
		3, 12, 6, 0, 464, 65, 1, 0, 0, 0, 465, 466, 5, 29, 0, 0, 466, 467, 3, 112,
		56, 0, 467, 468, 3, 12, 6, 0, 468, 482, 1, 0, 0, 0, 469, 470, 5, 29, 0,
		0, 470, 471, 3, 112, 56, 0, 471, 472, 3, 12, 6, 0, 472, 473, 5, 30, 0,
		0, 473, 474, 3, 12, 6, 0, 474, 482, 1, 0, 0, 0, 475, 476, 5, 29, 0, 0,
		476, 477, 3, 112, 56, 0, 477, 478, 3, 12, 6, 0, 478, 479, 5, 30, 0, 0,
		479, 480, 3, 66, 33, 0, 480, 482, 1, 0, 0, 0, 481, 465, 1, 0, 0, 0, 481,
		469, 1, 0, 0, 0, 481, 475, 1, 0, 0, 0, 482, 67, 1, 0, 0, 0, 483, 484, 5,
		31, 0, 0, 484, 485, 3, 112, 56, 0, 485, 486, 5, 30, 0, 0, 486, 487, 5,
		9, 0, 0, 487, 488, 3, 18, 9, 0, 488, 489, 7, 1, 0, 0, 489, 490, 5, 10,
		0, 0, 490, 69, 1, 0, 0, 0, 491, 492, 5, 35, 0, 0, 492, 493, 3, 112, 56,
		0, 493, 497, 5, 9, 0, 0, 494, 496, 3, 72, 36, 0, 495, 494, 1, 0, 0, 0,
		496, 499, 1, 0, 0, 0, 497, 495, 1, 0, 0, 0, 497, 498, 1, 0, 0, 0, 498,
		501, 1, 0, 0, 0, 499, 497, 1, 0, 0, 0, 500, 502, 3, 74, 37, 0, 501, 500,
		1, 0, 0, 0, 501, 502, 1, 0, 0, 0, 502, 503, 1, 0, 0, 0, 503, 504, 5, 10,
		0, 0, 504, 71, 1, 0, 0, 0, 505, 506, 5, 36, 0, 0, 506, 507, 3, 112, 56,
		0, 507, 508, 5, 7, 0, 0, 508, 510, 3, 18, 9, 0, 509, 511, 5, 33, 0, 0,
		510, 509, 1, 0, 0, 0, 510, 511, 1, 0, 0, 0, 511, 73, 1, 0, 0, 0, 512, 513,
		5, 37, 0, 0, 513, 514, 5, 7, 0, 0, 514, 516, 3, 18, 9, 0, 515, 517, 5,
		33, 0, 0, 516, 515, 1, 0, 0, 0, 516, 517, 1, 0, 0, 0, 517, 75, 1, 0, 0,
		0, 518, 519, 5, 33, 0, 0, 519, 77, 1, 0, 0, 0, 520, 521, 5, 32, 0, 0, 521,
		79, 1, 0, 0, 0, 522, 524, 5, 34, 0, 0, 523, 525, 3, 112, 56, 0, 524, 523,
		1, 0, 0, 0, 524, 525, 1, 0, 0, 0, 525, 81, 1, 0, 0, 0, 526, 527, 5, 12,
		0, 0, 527, 529, 5, 72, 0, 0, 528, 530, 3, 86, 43, 0, 529, 528, 1, 0, 0,
		0, 529, 530, 1, 0, 0, 0, 530, 531, 1, 0, 0, 0, 531, 532, 5, 14, 0, 0, 532,
		533, 3, 112, 56, 0, 533, 83, 1, 0, 0, 0, 534, 535, 5, 13, 0, 0, 535, 536,
		5, 72, 0, 0, 536, 537, 3, 86, 43, 0, 537, 538, 5, 38, 0, 0, 538, 547, 1,
		0, 0, 0, 539, 540, 5, 13, 0, 0, 540, 542, 5, 72, 0, 0, 541, 543, 3, 86,
		43, 0, 542, 541, 1, 0, 0, 0, 542, 543, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0,
		544, 545, 5, 14, 0, 0, 545, 547, 3, 112, 56, 0, 546, 534, 1, 0, 0, 0, 546,
		539, 1, 0, 0, 0, 547, 85, 1, 0, 0, 0, 548, 549, 5, 7, 0, 0, 549, 550, 3,
		88, 44, 0, 550, 87, 1, 0, 0, 0, 551, 552, 7, 2, 0, 0, 552, 89, 1, 0, 0,
		0, 553, 554, 5, 13, 0, 0, 554, 555, 5, 72, 0, 0, 555, 556, 5, 7, 0, 0,
		556, 557, 5, 20, 0, 0, 557, 558, 3, 88, 44, 0, 558, 559, 5, 21, 0, 0, 559,
		560, 3, 92, 46, 0, 560, 571, 1, 0, 0, 0, 561, 562, 5, 13, 0, 0, 562, 563,
		5, 72, 0, 0, 563, 564, 5, 14, 0, 0, 564, 565, 5, 20, 0, 0, 565, 566, 3,
		88, 44, 0, 566, 567, 5, 21, 0, 0, 567, 568, 5, 3, 0, 0, 568, 569, 5, 4,
		0, 0, 569, 571, 1, 0, 0, 0, 570, 553, 1, 0, 0, 0, 570, 561, 1, 0, 0, 0,
		571, 91, 1, 0, 0, 0, 572, 573, 5, 14, 0, 0, 573, 574, 5, 20, 0, 0, 574,
		575, 3, 94, 47, 0, 575, 576, 5, 21, 0, 0, 576, 583, 1, 0, 0, 0, 577, 578,
		5, 14, 0, 0, 578, 579, 5, 20, 0, 0, 579, 583, 5, 21, 0, 0, 580, 581, 5,
		14, 0, 0, 581, 583, 5, 72, 0, 0, 582, 572, 1, 0, 0, 0, 582, 577, 1, 0,
		0, 0, 582, 580, 1, 0, 0, 0, 583, 93, 1, 0, 0, 0, 584, 589, 3, 112, 56,
		0, 585, 586, 5, 6, 0, 0, 586, 588, 3, 112, 56, 0, 587, 585, 1, 0, 0, 0,
		588, 591, 1, 0, 0, 0, 589, 587, 1, 0, 0, 0, 589, 590, 1, 0, 0, 0, 590,
		95, 1, 0, 0, 0, 591, 589, 1, 0, 0, 0, 592, 593, 5, 44, 0, 0, 593, 594,
		5, 3, 0, 0, 594, 595, 3, 94, 47, 0, 595, 596, 5, 4, 0, 0, 596, 97, 1, 0,
		0, 0, 597, 598, 5, 72, 0, 0, 598, 599, 5, 17, 0, 0, 599, 600, 5, 45, 0,
		0, 600, 601, 5, 3, 0, 0, 601, 602, 3, 112, 56, 0, 602, 603, 5, 4, 0, 0,
		603, 99, 1, 0, 0, 0, 604, 605, 5, 72, 0, 0, 605, 606, 5, 17, 0, 0, 606,
		607, 5, 46, 0, 0, 607, 608, 5, 3, 0, 0, 608, 609, 5, 4, 0, 0, 609, 101,
		1, 0, 0, 0, 610, 611, 5, 72, 0, 0, 611, 612, 5, 17, 0, 0, 612, 613, 5,
		47, 0, 0, 613, 614, 5, 3, 0, 0, 614, 615, 5, 48, 0, 0, 615, 616, 5, 7,
		0, 0, 616, 617, 3, 112, 56, 0, 617, 618, 5, 4, 0, 0, 618, 103, 1, 0, 0,
		0, 619, 620, 5, 40, 0, 0, 620, 621, 5, 3, 0, 0, 621, 622, 3, 112, 56, 0,
		622, 623, 5, 4, 0, 0, 623, 105, 1, 0, 0, 0, 624, 625, 5, 41, 0, 0, 625,
		626, 5, 3, 0, 0, 626, 627, 3, 112, 56, 0, 627, 628, 5, 4, 0, 0, 628, 107,
		1, 0, 0, 0, 629, 630, 5, 39, 0, 0, 630, 631, 5, 3, 0, 0, 631, 632, 3, 112,
		56, 0, 632, 633, 5, 4, 0, 0, 633, 109, 1, 0, 0, 0, 634, 635, 5, 72, 0,
		0, 635, 636, 5, 14, 0, 0, 636, 665, 3, 112, 56, 0, 637, 638, 5, 72, 0,
		0, 638, 639, 5, 22, 0, 0, 639, 665, 3, 112, 56, 0, 640, 641, 5, 72, 0,
		0, 641, 642, 5, 23, 0, 0, 642, 665, 3, 112, 56, 0, 643, 644, 5, 72, 0,
		0, 644, 645, 5, 20, 0, 0, 645, 646, 3, 112, 56, 0, 646, 647, 5, 21, 0,
		0, 647, 648, 5, 14, 0, 0, 648, 649, 3, 112, 56, 0, 649, 665, 1, 0, 0, 0,
		650, 651, 5, 72, 0, 0, 651, 652, 5, 20, 0, 0, 652, 653, 3, 112, 56, 0,
		653, 654, 5, 21, 0, 0, 654, 655, 5, 22, 0, 0, 655, 656, 3, 112, 56, 0,
		656, 665, 1, 0, 0, 0, 657, 658, 5, 72, 0, 0, 658, 659, 5, 20, 0, 0, 659,
		660, 3, 112, 56, 0, 660, 661, 5, 21, 0, 0, 661, 662, 5, 23, 0, 0, 662,
		663, 3, 112, 56, 0, 663, 665, 1, 0, 0, 0, 664, 634, 1, 0, 0, 0, 664, 637,
		1, 0, 0, 0, 664, 640, 1, 0, 0, 0, 664, 643, 1, 0, 0, 0, 664, 650, 1, 0,
		0, 0, 664, 657, 1, 0, 0, 0, 665, 111, 1, 0, 0, 0, 666, 667, 6, 56, -1,
		0, 667, 702, 3, 116, 58, 0, 668, 702, 3, 40, 20, 0, 669, 670, 5, 72, 0,
		0, 670, 671, 5, 20, 0, 0, 671, 672, 3, 112, 56, 0, 672, 673, 5, 21, 0,
		0, 673, 702, 1, 0, 0, 0, 674, 675, 5, 72, 0, 0, 675, 676, 5, 20, 0, 0,
		676, 677, 3, 112, 56, 0, 677, 682, 5, 21, 0, 0, 678, 679, 5, 20, 0, 0,
		679, 680, 3, 112, 56, 0, 680, 681, 5, 21, 0, 0, 681, 683, 1, 0, 0, 0, 682,
		678, 1, 0, 0, 0, 683, 684, 1, 0, 0, 0, 684, 682, 1, 0, 0, 0, 684, 685,
		1, 0, 0, 0, 685, 702, 1, 0, 0, 0, 686, 702, 3, 30, 15, 0, 687, 688, 5,
		72, 0, 0, 688, 689, 5, 3, 0, 0, 689, 690, 3, 114, 57, 0, 690, 691, 5, 4,
		0, 0, 691, 702, 1, 0, 0, 0, 692, 702, 5, 72, 0, 0, 693, 694, 5, 3, 0, 0,
		694, 695, 3, 112, 56, 0, 695, 696, 5, 4, 0, 0, 696, 702, 1, 0, 0, 0, 697,
		698, 5, 49, 0, 0, 698, 702, 3, 112, 56, 6, 699, 700, 5, 50, 0, 0, 700,
		702, 3, 112, 56, 5, 701, 666, 1, 0, 0, 0, 701, 668, 1, 0, 0, 0, 701, 669,
		1, 0, 0, 0, 701, 674, 1, 0, 0, 0, 701, 686, 1, 0, 0, 0, 701, 687, 1, 0,
		0, 0, 701, 692, 1, 0, 0, 0, 701, 693, 1, 0, 0, 0, 701, 697, 1, 0, 0, 0,
		701, 699, 1, 0, 0, 0, 702, 717, 1, 0, 0, 0, 703, 704, 10, 4, 0, 0, 704,
		705, 7, 3, 0, 0, 705, 716, 3, 112, 56, 5, 706, 707, 10, 3, 0, 0, 707, 708,
		7, 4, 0, 0, 708, 716, 3, 112, 56, 4, 709, 710, 10, 2, 0, 0, 710, 711, 7,
		5, 0, 0, 711, 716, 3, 112, 56, 3, 712, 713, 10, 1, 0, 0, 713, 714, 7, 6,
		0, 0, 714, 716, 3, 112, 56, 2, 715, 703, 1, 0, 0, 0, 715, 706, 1, 0, 0,
		0, 715, 709, 1, 0, 0, 0, 715, 712, 1, 0, 0, 0, 716, 719, 1, 0, 0, 0, 717,
		715, 1, 0, 0, 0, 717, 718, 1, 0, 0, 0, 718, 113, 1, 0, 0, 0, 719, 717,
		1, 0, 0, 0, 720, 721, 5, 72, 0, 0, 721, 722, 5, 7, 0, 0, 722, 729, 3, 112,
		56, 0, 723, 724, 5, 6, 0, 0, 724, 725, 5, 72, 0, 0, 725, 726, 5, 7, 0,
		0, 726, 728, 3, 112, 56, 0, 727, 723, 1, 0, 0, 0, 728, 731, 1, 0, 0, 0,
		729, 727, 1, 0, 0, 0, 729, 730, 1, 0, 0, 0, 730, 115, 1, 0, 0, 0, 731,
		729, 1, 0, 0, 0, 732, 739, 5, 69, 0, 0, 733, 739, 5, 67, 0, 0, 734, 739,
		5, 68, 0, 0, 735, 739, 5, 70, 0, 0, 736, 739, 5, 71, 0, 0, 737, 739, 5,
		63, 0, 0, 738, 732, 1, 0, 0, 0, 738, 733, 1, 0, 0, 0, 738, 734, 1, 0, 0,
		0, 738, 735, 1, 0, 0, 0, 738, 736, 1, 0, 0, 0, 738, 737, 1, 0, 0, 0, 739,
		117, 1, 0, 0, 0, 74, 124, 129, 133, 137, 141, 145, 151, 153, 159, 164,
		173, 177, 182, 186, 191, 194, 206, 215, 219, 222, 225, 228, 233, 238, 242,
		246, 250, 254, 258, 262, 264, 270, 274, 279, 284, 295, 300, 309, 314, 317,
		326, 341, 343, 351, 358, 368, 378, 383, 389, 402, 406, 417, 422, 446, 453,
		481, 497, 501, 510, 516, 524, 529, 542, 546, 570, 582, 589, 664, 684, 701,
		715, 717, 729, 738,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// T_swiftParserInit initializes any static state used to implement T_swiftParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewT_swiftParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func T_swiftParserInit() {
	staticData := &T_swiftParserStaticData
	staticData.once.Do(t_swiftParserInit)
}

// NewT_swiftParser produces a new parser instance for the optional input antlr.TokenStream.
func NewT_swiftParser(input antlr.TokenStream) *T_swiftParser {
	T_swiftParserInit()
	this := new(T_swiftParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &T_swiftParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "T_swift.g4"

	return this
}

// T_swiftParser tokens.
const (
	T_swiftParserEOF           = antlr.TokenEOF
	T_swiftParserT__0          = 1
	T_swiftParserT__1          = 2
	T_swiftParserT__2          = 3
	T_swiftParserT__3          = 4
	T_swiftParserT__4          = 5
	T_swiftParserT__5          = 6
	T_swiftParserT__6          = 7
	T_swiftParserT__7          = 8
	T_swiftParserT__8          = 9
	T_swiftParserT__9          = 10
	T_swiftParserT__10         = 11
	T_swiftParserT__11         = 12
	T_swiftParserT__12         = 13
	T_swiftParserT__13         = 14
	T_swiftParserT__14         = 15
	T_swiftParserT__15         = 16
	T_swiftParserT__16         = 17
	T_swiftParserT__17         = 18
	T_swiftParserT__18         = 19
	T_swiftParserT__19         = 20
	T_swiftParserT__20         = 21
	T_swiftParserT__21         = 22
	T_swiftParserT__22         = 23
	T_swiftParserT__23         = 24
	T_swiftParserT__24         = 25
	T_swiftParserT__25         = 26
	T_swiftParserT__26         = 27
	T_swiftParserT__27         = 28
	T_swiftParserT__28         = 29
	T_swiftParserT__29         = 30
	T_swiftParserT__30         = 31
	T_swiftParserT__31         = 32
	T_swiftParserT__32         = 33
	T_swiftParserT__33         = 34
	T_swiftParserT__34         = 35
	T_swiftParserT__35         = 36
	T_swiftParserT__36         = 37
	T_swiftParserT__37         = 38
	T_swiftParserT__38         = 39
	T_swiftParserT__39         = 40
	T_swiftParserT__40         = 41
	T_swiftParserT__41         = 42
	T_swiftParserT__42         = 43
	T_swiftParserT__43         = 44
	T_swiftParserT__44         = 45
	T_swiftParserT__45         = 46
	T_swiftParserT__46         = 47
	T_swiftParserT__47         = 48
	T_swiftParserT__48         = 49
	T_swiftParserT__49         = 50
	T_swiftParserT__50         = 51
	T_swiftParserT__51         = 52
	T_swiftParserT__52         = 53
	T_swiftParserT__53         = 54
	T_swiftParserT__54         = 55
	T_swiftParserT__55         = 56
	T_swiftParserT__56         = 57
	T_swiftParserT__57         = 58
	T_swiftParserT__58         = 59
	T_swiftParserT__59         = 60
	T_swiftParserT__60         = 61
	T_swiftParserT__61         = 62
	T_swiftParserT__62         = 63
	T_swiftParserWS            = 64
	T_swiftParserBlock_comment = 65
	T_swiftParserLine_comment  = 66
	T_swiftParserInt           = 67
	T_swiftParserFloat         = 68
	T_swiftParserCaracter      = 69
	T_swiftParserString_       = 70
	T_swiftParserBool          = 71
	T_swiftParserIdentificador = 72
)

// T_swiftParser rules.
const (
	T_swiftParserRULE_inicio                     = 0
	T_swiftParserRULE_instrucciones_globales     = 1
	T_swiftParserRULE_intruccion_global          = 2
	T_swiftParserRULE_function_declaracion       = 3
	T_swiftParserRULE_lista_parametros           = 4
	T_swiftParserRULE_declaracion_parametro      = 5
	T_swiftParserRULE_code_block                 = 6
	T_swiftParserRULE_struct_declaracion         = 7
	T_swiftParserRULE_lista_atributos            = 8
	T_swiftParserRULE_instrucciones              = 9
	T_swiftParserRULE_instruccion                = 10
	T_swiftParserRULE_declaracion                = 11
	T_swiftParserRULE_loop_statement             = 12
	T_swiftParserRULE_branch_statement           = 13
	T_swiftParserRULE_control_transfer_statement = 14
	T_swiftParserRULE_llamadas_funciones         = 15
	T_swiftParserRULE_llamada_normal             = 16
	T_swiftParserRULE_lista_argumentos           = 17
	T_swiftParserRULE_declaracion_argumento      = 18
	T_swiftParserRULE_llamada_metodos            = 19
	T_swiftParserRULE_atributos                  = 20
	T_swiftParserRULE_ide_atributo               = 21
	T_swiftParserRULE_asignar_atributos          = 22
	T_swiftParserRULE_matriz_declaracion         = 23
	T_swiftParserRULE_tipo_matriz                = 24
	T_swiftParserRULE_definicion_matriz          = 25
	T_swiftParserRULE_lista_valores_matriz       = 26
	T_swiftParserRULE_elementos_matriz           = 27
	T_swiftParserRULE_elemento_matriz            = 28
	T_swiftParserRULE_simple_matriz              = 29
	T_swiftParserRULE_for_in_statement           = 30
	T_swiftParserRULE_rango                      = 31
	T_swiftParserRULE_while_statement            = 32
	T_swiftParserRULE_if_statement               = 33
	T_swiftParserRULE_guard_statement            = 34
	T_swiftParserRULE_switch_statement           = 35
	T_swiftParserRULE_switch_case                = 36
	T_swiftParserRULE_default_case               = 37
	T_swiftParserRULE_break_statement            = 38
	T_swiftParserRULE_continue_statement         = 39
	T_swiftParserRULE_return_statement           = 40
	T_swiftParserRULE_constant_declaracion       = 41
	T_swiftParserRULE_variable_declaracion       = 42
	T_swiftParserRULE_anotacion_tipo             = 43
	T_swiftParserRULE_tipos                      = 44
	T_swiftParserRULE_array_declaracion          = 45
	T_swiftParserRULE_definicion_vector          = 46
	T_swiftParserRULE_lista_expresiones          = 47
	T_swiftParserRULE_funcion_print              = 48
	T_swiftParserRULE_funcion_append             = 49
	T_swiftParserRULE_funcion_removeLast         = 50
	T_swiftParserRULE_funcion_removeat           = 51
	T_swiftParserRULE_funcion_int                = 52
	T_swiftParserRULE_funcion_float              = 53
	T_swiftParserRULE_funcion_string             = 54
	T_swiftParserRULE_asignacion                 = 55
	T_swiftParserRULE_expresion                  = 56
	T_swiftParserRULE_l_duble                    = 57
	T_swiftParserRULE_primitivos                 = 58
)

// IInicioContext is an interface to support dynamic dispatch.
type IInicioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instrucciones_globales() IInstrucciones_globalesContext
	EOF() antlr.TerminalNode

	// IsInicioContext differentiates from other interfaces.
	IsInicioContext()
}

type InicioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInicioContext() *InicioContext {
	var p = new(InicioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_inicio
	return p
}

func InitEmptyInicioContext(p *InicioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_inicio
}

func (*InicioContext) IsInicioContext() {}

func NewInicioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InicioContext {
	var p = new(InicioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_inicio

	return p
}

func (s *InicioContext) GetParser() antlr.Parser { return s.parser }

func (s *InicioContext) Instrucciones_globales() IInstrucciones_globalesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstrucciones_globalesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstrucciones_globalesContext)
}

func (s *InicioContext) EOF() antlr.TerminalNode {
	return s.GetToken(T_swiftParserEOF, 0)
}

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InicioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitInicio(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Inicio() (localctx IInicioContext) {
	localctx = NewInicioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, T_swiftParserRULE_inicio)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Instrucciones_globales()
	}
	{
		p.SetState(119)
		p.Match(T_swiftParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstrucciones_globalesContext is an interface to support dynamic dispatch.
type IInstrucciones_globalesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIntruccion_global() []IIntruccion_globalContext
	Intruccion_global(i int) IIntruccion_globalContext

	// IsInstrucciones_globalesContext differentiates from other interfaces.
	IsInstrucciones_globalesContext()
}

type Instrucciones_globalesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstrucciones_globalesContext() *Instrucciones_globalesContext {
	var p = new(Instrucciones_globalesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_instrucciones_globales
	return p
}

func InitEmptyInstrucciones_globalesContext(p *Instrucciones_globalesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_instrucciones_globales
}

func (*Instrucciones_globalesContext) IsInstrucciones_globalesContext() {}

func NewInstrucciones_globalesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instrucciones_globalesContext {
	var p = new(Instrucciones_globalesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_instrucciones_globales

	return p
}

func (s *Instrucciones_globalesContext) GetParser() antlr.Parser { return s.parser }

func (s *Instrucciones_globalesContext) AllIntruccion_global() []IIntruccion_globalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIntruccion_globalContext); ok {
			len++
		}
	}

	tst := make([]IIntruccion_globalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIntruccion_globalContext); ok {
			tst[i] = t.(IIntruccion_globalContext)
			i++
		}
	}

	return tst
}

func (s *Instrucciones_globalesContext) Intruccion_global(i int) IIntruccion_globalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntruccion_globalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntruccion_globalContext)
}

func (s *Instrucciones_globalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instrucciones_globalesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instrucciones_globalesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitInstrucciones_globales(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Instrucciones_globales() (localctx IInstrucciones_globalesContext) {
	localctx = NewInstrucciones_globalesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, T_swiftParserRULE_instrucciones_globales)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&21477822838788) != 0) || _la == T_swiftParserIdentificador {
		{
			p.SetState(121)
			p.Intruccion_global()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntruccion_globalContext is an interface to support dynamic dispatch.
type IIntruccion_globalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaracion() IDeclaracionContext
	Loop_statement() ILoop_statementContext
	Branch_statement() IBranch_statementContext
	Asignacion() IAsignacionContext
	Llamadas_funciones() ILlamadas_funcionesContext
	Function_declaracion() IFunction_declaracionContext
	Struct_declaracion() IStruct_declaracionContext
	Asignar_atributos() IAsignar_atributosContext

	// IsIntruccion_globalContext differentiates from other interfaces.
	IsIntruccion_globalContext()
}

type Intruccion_globalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntruccion_globalContext() *Intruccion_globalContext {
	var p = new(Intruccion_globalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_intruccion_global
	return p
}

func InitEmptyIntruccion_globalContext(p *Intruccion_globalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_intruccion_global
}

func (*Intruccion_globalContext) IsIntruccion_globalContext() {}

func NewIntruccion_globalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Intruccion_globalContext {
	var p = new(Intruccion_globalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_intruccion_global

	return p
}

func (s *Intruccion_globalContext) GetParser() antlr.Parser { return s.parser }

func (s *Intruccion_globalContext) Declaracion() IDeclaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *Intruccion_globalContext) Loop_statement() ILoop_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoop_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoop_statementContext)
}

func (s *Intruccion_globalContext) Branch_statement() IBranch_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBranch_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBranch_statementContext)
}

func (s *Intruccion_globalContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *Intruccion_globalContext) Llamadas_funciones() ILlamadas_funcionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadas_funcionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadas_funcionesContext)
}

func (s *Intruccion_globalContext) Function_declaracion() IFunction_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_declaracionContext)
}

func (s *Intruccion_globalContext) Struct_declaracion() IStruct_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_declaracionContext)
}

func (s *Intruccion_globalContext) Asignar_atributos() IAsignar_atributosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignar_atributosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignar_atributosContext)
}

func (s *Intruccion_globalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Intruccion_globalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Intruccion_globalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitIntruccion_global(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Intruccion_global() (localctx IIntruccion_globalContext) {
	localctx = NewIntruccion_globalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, T_swiftParserRULE_intruccion_global)
	var _la int

	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Declaracion()
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(128)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Loop_statement()
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(132)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(135)
			p.Branch_statement()
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(136)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(139)
			p.Asignacion()
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(140)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(143)
			p.Llamadas_funciones()
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(144)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(147)
			p.Function_declaracion()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(148)
			p.Struct_declaracion()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(149)
			p.Asignar_atributos()
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(150)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_declaracionContext is an interface to support dynamic dispatch.
type IFunction_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Code_block() ICode_blockContext
	Lista_parametros() ILista_parametrosContext
	Tipos() ITiposContext

	// IsFunction_declaracionContext differentiates from other interfaces.
	IsFunction_declaracionContext()
}

type Function_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_declaracionContext() *Function_declaracionContext {
	var p = new(Function_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_function_declaracion
	return p
}

func InitEmptyFunction_declaracionContext(p *Function_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_function_declaracion
}

func (*Function_declaracionContext) IsFunction_declaracionContext() {}

func NewFunction_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declaracionContext {
	var p = new(Function_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_function_declaracion

	return p
}

func (s *Function_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declaracionContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Function_declaracionContext) Code_block() ICode_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *Function_declaracionContext) Lista_parametros() ILista_parametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_parametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_parametrosContext)
}

func (s *Function_declaracionContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *Function_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFunction_declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Function_declaracion() (localctx IFunction_declaracionContext) {
	localctx = NewFunction_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, T_swiftParserRULE_function_declaracion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(T_swiftParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserIdentificador {
		{
			p.SetState(158)
			p.Lista_parametros()
		}

	}
	{
		p.SetState(161)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__4 {
		{
			p.SetState(162)
			p.Match(T_swiftParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Tipos()
		}

	}
	{
		p.SetState(166)
		p.Code_block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_parametrosContext is an interface to support dynamic dispatch.
type ILista_parametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaracion_parametro() []IDeclaracion_parametroContext
	Declaracion_parametro(i int) IDeclaracion_parametroContext

	// IsLista_parametrosContext differentiates from other interfaces.
	IsLista_parametrosContext()
}

type Lista_parametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_parametrosContext() *Lista_parametrosContext {
	var p = new(Lista_parametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_parametros
	return p
}

func InitEmptyLista_parametrosContext(p *Lista_parametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_parametros
}

func (*Lista_parametrosContext) IsLista_parametrosContext() {}

func NewLista_parametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_parametrosContext {
	var p = new(Lista_parametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_lista_parametros

	return p
}

func (s *Lista_parametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_parametrosContext) AllDeclaracion_parametro() []IDeclaracion_parametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracion_parametroContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracion_parametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracion_parametroContext); ok {
			tst[i] = t.(IDeclaracion_parametroContext)
			i++
		}
	}

	return tst
}

func (s *Lista_parametrosContext) Declaracion_parametro(i int) IDeclaracion_parametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_parametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_parametroContext)
}

func (s *Lista_parametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_parametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_parametrosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLista_parametros(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Lista_parametros() (localctx ILista_parametrosContext) {
	localctx = NewLista_parametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, T_swiftParserRULE_lista_parametros)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Declaracion_parametro()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__5 {
		{
			p.SetState(169)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Declaracion_parametro()
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracion_parametroContext is an interface to support dynamic dispatch.
type IDeclaracion_parametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclaracion_parametroContext differentiates from other interfaces.
	IsDeclaracion_parametroContext()
}

type Declaracion_parametroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracion_parametroContext() *Declaracion_parametroContext {
	var p = new(Declaracion_parametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_declaracion_parametro
	return p
}

func InitEmptyDeclaracion_parametroContext(p *Declaracion_parametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_declaracion_parametro
}

func (*Declaracion_parametroContext) IsDeclaracion_parametroContext() {}

func NewDeclaracion_parametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_parametroContext {
	var p = new(Declaracion_parametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_declaracion_parametro

	return p
}

func (s *Declaracion_parametroContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_parametroContext) CopyAll(ctx *Declaracion_parametroContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Declaracion_parametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_parametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declaracion_parametro_simpleContext struct {
	Declaracion_parametroContext
	refencia antlr.Token
}

func NewDeclaracion_parametro_simpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_parametro_simpleContext {
	var p = new(Declaracion_parametro_simpleContext)

	InitEmptyDeclaracion_parametroContext(&p.Declaracion_parametroContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_parametroContext))

	return p
}

func (s *Declaracion_parametro_simpleContext) GetRefencia() antlr.Token { return s.refencia }

func (s *Declaracion_parametro_simpleContext) SetRefencia(v antlr.Token) { s.refencia = v }

func (s *Declaracion_parametro_simpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_parametro_simpleContext) AllIdentificador() []antlr.TerminalNode {
	return s.GetTokens(T_swiftParserIdentificador)
}

func (s *Declaracion_parametro_simpleContext) Identificador(i int) antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, i)
}

func (s *Declaracion_parametro_simpleContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *Declaracion_parametro_simpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclaracion_parametro_simple(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declaracion_parametro_vectorContext struct {
	Declaracion_parametroContext
	refencia antlr.Token
}

func NewDeclaracion_parametro_vectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_parametro_vectorContext {
	var p = new(Declaracion_parametro_vectorContext)

	InitEmptyDeclaracion_parametroContext(&p.Declaracion_parametroContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_parametroContext))

	return p
}

func (s *Declaracion_parametro_vectorContext) GetRefencia() antlr.Token { return s.refencia }

func (s *Declaracion_parametro_vectorContext) SetRefencia(v antlr.Token) { s.refencia = v }

func (s *Declaracion_parametro_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_parametro_vectorContext) AllIdentificador() []antlr.TerminalNode {
	return s.GetTokens(T_swiftParserIdentificador)
}

func (s *Declaracion_parametro_vectorContext) Identificador(i int) antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, i)
}

func (s *Declaracion_parametro_vectorContext) Tipo_matriz() ITipo_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_matrizContext)
}

func (s *Declaracion_parametro_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclaracion_parametro_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Declaracion_parametro() (localctx IDeclaracion_parametroContext) {
	localctx = NewDeclaracion_parametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, T_swiftParserRULE_declaracion_parametro)
	var _la int

	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclaracion_parametro_simpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(177)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(176)
				p.Match(T_swiftParserIdentificador)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(179)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Match(T_swiftParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__7 {
			{
				p.SetState(181)

				var _m = p.Match(T_swiftParserT__7)

				localctx.(*Declaracion_parametro_simpleContext).refencia = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(184)
			p.Tipos()
		}

	case 2:
		localctx = NewDeclaracion_parametro_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(186)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(185)
				p.Match(T_swiftParserIdentificador)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(188)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Match(T_swiftParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__7 {
			{
				p.SetState(190)

				var _m = p.Match(T_swiftParserT__7)

				localctx.(*Declaracion_parametro_vectorContext).refencia = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(193)
			p.Tipo_matriz()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICode_blockContext is an interface to support dynamic dispatch.
type ICode_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instrucciones() IInstruccionesContext

	// IsCode_blockContext differentiates from other interfaces.
	IsCode_blockContext()
}

type Code_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCode_blockContext() *Code_blockContext {
	var p = new(Code_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_code_block
	return p
}

func InitEmptyCode_blockContext(p *Code_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_code_block
}

func (*Code_blockContext) IsCode_blockContext() {}

func NewCode_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Code_blockContext {
	var p = new(Code_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_code_block

	return p
}

func (s *Code_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Code_blockContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Code_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Code_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Code_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitCode_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Code_block() (localctx ICode_blockContext) {
	localctx = NewCode_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, T_swiftParserRULE_code_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(T_swiftParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Instrucciones()
	}
	{
		p.SetState(198)
		p.Match(T_swiftParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_declaracionContext is an interface to support dynamic dispatch.
type IStruct_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	AllLista_atributos() []ILista_atributosContext
	Lista_atributos(i int) ILista_atributosContext

	// IsStruct_declaracionContext differentiates from other interfaces.
	IsStruct_declaracionContext()
}

type Struct_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_declaracionContext() *Struct_declaracionContext {
	var p = new(Struct_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_struct_declaracion
	return p
}

func InitEmptyStruct_declaracionContext(p *Struct_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_struct_declaracion
}

func (*Struct_declaracionContext) IsStruct_declaracionContext() {}

func NewStruct_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_declaracionContext {
	var p = new(Struct_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_struct_declaracion

	return p
}

func (s *Struct_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_declaracionContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Struct_declaracionContext) AllLista_atributos() []ILista_atributosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_atributosContext); ok {
			len++
		}
	}

	tst := make([]ILista_atributosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_atributosContext); ok {
			tst[i] = t.(ILista_atributosContext)
			i++
		}
	}

	return tst
}

func (s *Struct_declaracionContext) Lista_atributos(i int) ILista_atributosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_atributosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_atributosContext)
}

func (s *Struct_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_declaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitStruct_declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Struct_declaracion() (localctx IStruct_declaracionContext) {
	localctx = NewStruct_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, T_swiftParserRULE_struct_declaracion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(T_swiftParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Match(T_swiftParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&45060) != 0 {
		{
			p.SetState(203)
			p.Lista_atributos()
		}

		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(209)
		p.Match(T_swiftParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_atributosContext is an interface to support dynamic dispatch.
type ILista_atributosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLista_atributosContext differentiates from other interfaces.
	IsLista_atributosContext()
}

type Lista_atributosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_atributosContext() *Lista_atributosContext {
	var p = new(Lista_atributosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_atributos
	return p
}

func InitEmptyLista_atributosContext(p *Lista_atributosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_atributos
}

func (*Lista_atributosContext) IsLista_atributosContext() {}

func NewLista_atributosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_atributosContext {
	var p = new(Lista_atributosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_lista_atributos

	return p
}

func (s *Lista_atributosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_atributosContext) CopyAll(ctx *Lista_atributosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Lista_atributosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_atributosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declarar_atributoContext struct {
	Lista_atributosContext
	tipo antlr.Token
}

func NewDeclarar_atributoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declarar_atributoContext {
	var p = new(Declarar_atributoContext)

	InitEmptyLista_atributosContext(&p.Lista_atributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_atributosContext))

	return p
}

func (s *Declarar_atributoContext) GetTipo() antlr.Token { return s.tipo }

func (s *Declarar_atributoContext) SetTipo(v antlr.Token) { s.tipo = v }

func (s *Declarar_atributoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_atributoContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Declarar_atributoContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *Declarar_atributoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Declarar_atributoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclarar_atributo(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declarar_funcion_scContext struct {
	Lista_atributosContext
	m antlr.Token
}

func NewDeclarar_funcion_scContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declarar_funcion_scContext {
	var p = new(Declarar_funcion_scContext)

	InitEmptyLista_atributosContext(&p.Lista_atributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_atributosContext))

	return p
}

func (s *Declarar_funcion_scContext) GetM() antlr.Token { return s.m }

func (s *Declarar_funcion_scContext) SetM(v antlr.Token) { s.m = v }

func (s *Declarar_funcion_scContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_funcion_scContext) Function_declaracion() IFunction_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_declaracionContext)
}

func (s *Declarar_funcion_scContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclarar_funcion_sc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Lista_atributos() (localctx ILista_atributosContext) {
	localctx = NewLista_atributosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, T_swiftParserRULE_lista_atributos)
	var _la int

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__11, T_swiftParserT__12:
		localctx = NewDeclarar_atributoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Declarar_atributoContext).tipo = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == T_swiftParserT__11 || _la == T_swiftParserT__12) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Declarar_atributoContext).tipo = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(212)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__6 {
			{
				p.SetState(213)
				p.Match(T_swiftParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(214)
				p.Tipos()
			}

		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__13 {
			{
				p.SetState(217)
				p.Match(T_swiftParserT__13)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(218)
				p.expresion(0)
			}

		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(221)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case T_swiftParserT__1, T_swiftParserT__14:
		localctx = NewDeclarar_funcion_scContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__14 {
			{
				p.SetState(224)

				var _m = p.Match(T_swiftParserT__14)

				localctx.(*Declarar_funcion_scContext).m = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(227)
			p.Function_declaracion()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInstruccion() []IInstruccionContext
	Instruccion(i int) IInstruccionContext

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_instrucciones
	return p
}

func InitEmptyInstruccionesContext(p *InstruccionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_instrucciones
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionContext); ok {
			tst[i] = t.(IInstruccionContext)
			i++
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitInstrucciones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, T_swiftParserRULE_instrucciones)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(230)
				p.Instruccion()
			}

		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaracion() IDeclaracionContext
	Loop_statement() ILoop_statementContext
	Branch_statement() IBranch_statementContext
	Control_transfer_statement() IControl_transfer_statementContext
	Asignacion() IAsignacionContext
	Llamadas_funciones() ILlamadas_funcionesContext
	Asignar_atributos() IAsignar_atributosContext

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_instruccion
	return p
}

func InitEmptyInstruccionContext(p *InstruccionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_instruccion
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Declaracion() IDeclaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *InstruccionContext) Loop_statement() ILoop_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoop_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoop_statementContext)
}

func (s *InstruccionContext) Branch_statement() IBranch_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBranch_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBranch_statementContext)
}

func (s *InstruccionContext) Control_transfer_statement() IControl_transfer_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControl_transfer_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControl_transfer_statementContext)
}

func (s *InstruccionContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *InstruccionContext) Llamadas_funciones() ILlamadas_funcionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadas_funcionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadas_funcionesContext)
}

func (s *InstruccionContext) Asignar_atributos() IAsignar_atributosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignar_atributosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignar_atributosContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitInstruccion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, T_swiftParserRULE_instruccion)
	var _la int

	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(236)
			p.Declaracion()
		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(237)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(240)
			p.Loop_statement()
		}
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(241)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(244)
			p.Branch_statement()
		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(245)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(248)
			p.Control_transfer_statement()
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(249)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(252)
			p.Asignacion()
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(253)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(256)
			p.Llamadas_funciones()
		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(257)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(260)
			p.Asignar_atributos()
		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__0 {
			{
				p.SetState(261)
				p.Match(T_swiftParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Constant_declaracion() IConstant_declaracionContext
	Variable_declaracion() IVariable_declaracionContext
	Array_declaracion() IArray_declaracionContext
	Matriz_declaracion() IMatriz_declaracionContext

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_declaracion
	return p
}

func InitEmptyDeclaracionContext(p *DeclaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_declaracion
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) Constant_declaracion() IConstant_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstant_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstant_declaracionContext)
}

func (s *DeclaracionContext) Variable_declaracion() IVariable_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariable_declaracionContext)
}

func (s *DeclaracionContext) Array_declaracion() IArray_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_declaracionContext)
}

func (s *DeclaracionContext) Matriz_declaracion() IMatriz_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatriz_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatriz_declaracionContext)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Declaracion() (localctx IDeclaracionContext) {
	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, T_swiftParserRULE_declaracion)
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.Constant_declaracion()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.Variable_declaracion()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(268)
			p.Array_declaracion()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(269)
			p.Matriz_declaracion()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoop_statementContext is an interface to support dynamic dispatch.
type ILoop_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	For_in_statement() IFor_in_statementContext
	While_statement() IWhile_statementContext

	// IsLoop_statementContext differentiates from other interfaces.
	IsLoop_statementContext()
}

type Loop_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoop_statementContext() *Loop_statementContext {
	var p = new(Loop_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_loop_statement
	return p
}

func InitEmptyLoop_statementContext(p *Loop_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_loop_statement
}

func (*Loop_statementContext) IsLoop_statementContext() {}

func NewLoop_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Loop_statementContext {
	var p = new(Loop_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_loop_statement

	return p
}

func (s *Loop_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Loop_statementContext) For_in_statement() IFor_in_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_in_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_in_statementContext)
}

func (s *Loop_statementContext) While_statement() IWhile_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_statementContext)
}

func (s *Loop_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Loop_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Loop_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLoop_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Loop_statement() (localctx ILoop_statementContext) {
	localctx = NewLoop_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, T_swiftParserRULE_loop_statement)
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__24:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			p.For_in_statement()
		}

	case T_swiftParserT__27:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)
			p.While_statement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBranch_statementContext is an interface to support dynamic dispatch.
type IBranch_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBranch_statementContext differentiates from other interfaces.
	IsBranch_statementContext()
}

type Branch_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranch_statementContext() *Branch_statementContext {
	var p = new(Branch_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_branch_statement
	return p
}

func InitEmptyBranch_statementContext(p *Branch_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_branch_statement
}

func (*Branch_statementContext) IsBranch_statementContext() {}

func NewBranch_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Branch_statementContext {
	var p = new(Branch_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_branch_statement

	return p
}

func (s *Branch_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Branch_statementContext) CopyAll(ctx *Branch_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Branch_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Branch_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declarar_guardContext struct {
	Branch_statementContext
}

func NewDeclarar_guardContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declarar_guardContext {
	var p = new(Declarar_guardContext)

	InitEmptyBranch_statementContext(&p.Branch_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Branch_statementContext))

	return p
}

func (s *Declarar_guardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_guardContext) Guard_statement() IGuard_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuard_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuard_statementContext)
}

func (s *Declarar_guardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclarar_guard(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declarar_ifContext struct {
	Branch_statementContext
}

func NewDeclarar_ifContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declarar_ifContext {
	var p = new(Declarar_ifContext)

	InitEmptyBranch_statementContext(&p.Branch_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Branch_statementContext))

	return p
}

func (s *Declarar_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_ifContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *Declarar_ifContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclarar_if(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declarar_switchContext struct {
	Branch_statementContext
}

func NewDeclarar_switchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declarar_switchContext {
	var p = new(Declarar_switchContext)

	InitEmptyBranch_statementContext(&p.Branch_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Branch_statementContext))

	return p
}

func (s *Declarar_switchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarar_switchContext) Switch_statement() ISwitch_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_statementContext)
}

func (s *Declarar_switchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclarar_switch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Branch_statement() (localctx IBranch_statementContext) {
	localctx = NewBranch_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, T_swiftParserRULE_branch_statement)
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__28:
		localctx = NewDeclarar_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.If_statement()
		}

	case T_swiftParserT__30:
		localctx = NewDeclarar_guardContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.Guard_statement()
		}

	case T_swiftParserT__34:
		localctx = NewDeclarar_switchContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(278)
			p.Switch_statement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IControl_transfer_statementContext is an interface to support dynamic dispatch.
type IControl_transfer_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Break_statement() IBreak_statementContext
	Continue_statement() IContinue_statementContext
	Return_statement() IReturn_statementContext

	// IsControl_transfer_statementContext differentiates from other interfaces.
	IsControl_transfer_statementContext()
}

type Control_transfer_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControl_transfer_statementContext() *Control_transfer_statementContext {
	var p = new(Control_transfer_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_control_transfer_statement
	return p
}

func InitEmptyControl_transfer_statementContext(p *Control_transfer_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_control_transfer_statement
}

func (*Control_transfer_statementContext) IsControl_transfer_statementContext() {}

func NewControl_transfer_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_transfer_statementContext {
	var p = new(Control_transfer_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_control_transfer_statement

	return p
}

func (s *Control_transfer_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_transfer_statementContext) Break_statement() IBreak_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreak_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreak_statementContext)
}

func (s *Control_transfer_statementContext) Continue_statement() IContinue_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinue_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinue_statementContext)
}

func (s *Control_transfer_statementContext) Return_statement() IReturn_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_statementContext)
}

func (s *Control_transfer_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_transfer_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_transfer_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitControl_transfer_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Control_transfer_statement() (localctx IControl_transfer_statementContext) {
	localctx = NewControl_transfer_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, T_swiftParserRULE_control_transfer_statement)
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__32:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Break_statement()
		}

	case T_swiftParserT__31:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(282)
			p.Continue_statement()
		}

	case T_swiftParserT__33:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(283)
			p.Return_statement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadas_funcionesContext is an interface to support dynamic dispatch.
type ILlamadas_funcionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Funcion_print() IFuncion_printContext
	Funcion_append() IFuncion_appendContext
	Funcion_removeLast() IFuncion_removeLastContext
	Funcion_removeat() IFuncion_removeatContext
	Funcion_int() IFuncion_intContext
	Funcion_float() IFuncion_floatContext
	Funcion_string() IFuncion_stringContext
	Llamada_normal() ILlamada_normalContext
	Llamada_metodos() ILlamada_metodosContext

	// IsLlamadas_funcionesContext differentiates from other interfaces.
	IsLlamadas_funcionesContext()
}

type Llamadas_funcionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamadas_funcionesContext() *Llamadas_funcionesContext {
	var p = new(Llamadas_funcionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_llamadas_funciones
	return p
}

func InitEmptyLlamadas_funcionesContext(p *Llamadas_funcionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_llamadas_funciones
}

func (*Llamadas_funcionesContext) IsLlamadas_funcionesContext() {}

func NewLlamadas_funcionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Llamadas_funcionesContext {
	var p = new(Llamadas_funcionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_llamadas_funciones

	return p
}

func (s *Llamadas_funcionesContext) GetParser() antlr.Parser { return s.parser }

func (s *Llamadas_funcionesContext) Funcion_print() IFuncion_printContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_printContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_printContext)
}

func (s *Llamadas_funcionesContext) Funcion_append() IFuncion_appendContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_appendContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_appendContext)
}

func (s *Llamadas_funcionesContext) Funcion_removeLast() IFuncion_removeLastContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_removeLastContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_removeLastContext)
}

func (s *Llamadas_funcionesContext) Funcion_removeat() IFuncion_removeatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_removeatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_removeatContext)
}

func (s *Llamadas_funcionesContext) Funcion_int() IFuncion_intContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_intContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_intContext)
}

func (s *Llamadas_funcionesContext) Funcion_float() IFuncion_floatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_floatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_floatContext)
}

func (s *Llamadas_funcionesContext) Funcion_string() IFuncion_stringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncion_stringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncion_stringContext)
}

func (s *Llamadas_funcionesContext) Llamada_normal() ILlamada_normalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamada_normalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamada_normalContext)
}

func (s *Llamadas_funcionesContext) Llamada_metodos() ILlamada_metodosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamada_metodosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamada_metodosContext)
}

func (s *Llamadas_funcionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llamadas_funcionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Llamadas_funcionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLlamadas_funciones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Llamadas_funciones() (localctx ILlamadas_funcionesContext) {
	localctx = NewLlamadas_funcionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, T_swiftParserRULE_llamadas_funciones)
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(286)
			p.Funcion_print()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.Funcion_append()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(288)
			p.Funcion_removeLast()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(289)
			p.Funcion_removeat()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(290)
			p.Funcion_int()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(291)
			p.Funcion_float()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(292)
			p.Funcion_string()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(293)
			p.Llamada_normal()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(294)
			p.Llamada_metodos()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamada_normalContext is an interface to support dynamic dispatch.
type ILlamada_normalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Lista_argumentos() ILista_argumentosContext

	// IsLlamada_normalContext differentiates from other interfaces.
	IsLlamada_normalContext()
}

type Llamada_normalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamada_normalContext() *Llamada_normalContext {
	var p = new(Llamada_normalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_llamada_normal
	return p
}

func InitEmptyLlamada_normalContext(p *Llamada_normalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_llamada_normal
}

func (*Llamada_normalContext) IsLlamada_normalContext() {}

func NewLlamada_normalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Llamada_normalContext {
	var p = new(Llamada_normalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_llamada_normal

	return p
}

func (s *Llamada_normalContext) GetParser() antlr.Parser { return s.parser }

func (s *Llamada_normalContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Llamada_normalContext) Lista_argumentos() ILista_argumentosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_argumentosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_argumentosContext)
}

func (s *Llamada_normalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llamada_normalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Llamada_normalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLlamada_normal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Llamada_normal() (localctx ILlamada_normalContext) {
	localctx = NewLlamada_normalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, T_swiftParserRULE_llamada_normal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9221661746517704696) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&63) != 0) {
		{
			p.SetState(299)
			p.Lista_argumentos()
		}

	}
	{
		p.SetState(302)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_argumentosContext is an interface to support dynamic dispatch.
type ILista_argumentosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaracion_argumento() []IDeclaracion_argumentoContext
	Declaracion_argumento(i int) IDeclaracion_argumentoContext

	// IsLista_argumentosContext differentiates from other interfaces.
	IsLista_argumentosContext()
}

type Lista_argumentosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_argumentosContext() *Lista_argumentosContext {
	var p = new(Lista_argumentosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_argumentos
	return p
}

func InitEmptyLista_argumentosContext(p *Lista_argumentosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_argumentos
}

func (*Lista_argumentosContext) IsLista_argumentosContext() {}

func NewLista_argumentosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_argumentosContext {
	var p = new(Lista_argumentosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_lista_argumentos

	return p
}

func (s *Lista_argumentosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_argumentosContext) AllDeclaracion_argumento() []IDeclaracion_argumentoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclaracion_argumentoContext); ok {
			len++
		}
	}

	tst := make([]IDeclaracion_argumentoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclaracion_argumentoContext); ok {
			tst[i] = t.(IDeclaracion_argumentoContext)
			i++
		}
	}

	return tst
}

func (s *Lista_argumentosContext) Declaracion_argumento(i int) IDeclaracion_argumentoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_argumentoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_argumentoContext)
}

func (s *Lista_argumentosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_argumentosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_argumentosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLista_argumentos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Lista_argumentos() (localctx ILista_argumentosContext) {
	localctx = NewLista_argumentosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, T_swiftParserRULE_lista_argumentos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Declaracion_argumento()
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__5 {
		{
			p.SetState(305)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Declaracion_argumento()
		}

		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracion_argumentoContext is an interface to support dynamic dispatch.
type IDeclaracion_argumentoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetR returns the r token.
	GetR() antlr.Token

	// SetR sets the r token.
	SetR(antlr.Token)

	// Getter signatures
	Expresion() IExpresionContext
	Identificador() antlr.TerminalNode

	// IsDeclaracion_argumentoContext differentiates from other interfaces.
	IsDeclaracion_argumentoContext()
}

type Declaracion_argumentoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	r      antlr.Token
}

func NewEmptyDeclaracion_argumentoContext() *Declaracion_argumentoContext {
	var p = new(Declaracion_argumentoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_declaracion_argumento
	return p
}

func InitEmptyDeclaracion_argumentoContext(p *Declaracion_argumentoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_declaracion_argumento
}

func (*Declaracion_argumentoContext) IsDeclaracion_argumentoContext() {}

func NewDeclaracion_argumentoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_argumentoContext {
	var p = new(Declaracion_argumentoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_declaracion_argumento

	return p
}

func (s *Declaracion_argumentoContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_argumentoContext) GetR() antlr.Token { return s.r }

func (s *Declaracion_argumentoContext) SetR(v antlr.Token) { s.r = v }

func (s *Declaracion_argumentoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Declaracion_argumentoContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Declaracion_argumentoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_argumentoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracion_argumentoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDeclaracion_argumento(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Declaracion_argumento() (localctx IDeclaracion_argumentoContext) {
	localctx = NewDeclaracion_argumentoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, T_swiftParserRULE_declaracion_argumento)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(314)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(312)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.Match(T_swiftParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__15 {
		{
			p.SetState(316)

			var _m = p.Match(T_swiftParserT__15)

			localctx.(*Declaracion_argumentoContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(319)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamada_metodosContext is an interface to support dynamic dispatch.
type ILlamada_metodosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentificador() []antlr.TerminalNode
	Identificador(i int) antlr.TerminalNode
	Lista_argumentos() ILista_argumentosContext

	// IsLlamada_metodosContext differentiates from other interfaces.
	IsLlamada_metodosContext()
}

type Llamada_metodosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamada_metodosContext() *Llamada_metodosContext {
	var p = new(Llamada_metodosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_llamada_metodos
	return p
}

func InitEmptyLlamada_metodosContext(p *Llamada_metodosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_llamada_metodos
}

func (*Llamada_metodosContext) IsLlamada_metodosContext() {}

func NewLlamada_metodosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Llamada_metodosContext {
	var p = new(Llamada_metodosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_llamada_metodos

	return p
}

func (s *Llamada_metodosContext) GetParser() antlr.Parser { return s.parser }

func (s *Llamada_metodosContext) AllIdentificador() []antlr.TerminalNode {
	return s.GetTokens(T_swiftParserIdentificador)
}

func (s *Llamada_metodosContext) Identificador(i int) antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, i)
}

func (s *Llamada_metodosContext) Lista_argumentos() ILista_argumentosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_argumentosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_argumentosContext)
}

func (s *Llamada_metodosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llamada_metodosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Llamada_metodosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLlamada_metodos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Llamada_metodos() (localctx ILlamada_metodosContext) {
	localctx = NewLlamada_metodosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, T_swiftParserRULE_llamada_metodos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Match(T_swiftParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9221661746517704696) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&63) != 0) {
		{
			p.SetState(325)
			p.Lista_argumentos()
		}

	}
	{
		p.SetState(328)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtributosContext is an interface to support dynamic dispatch.
type IAtributosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtributosContext differentiates from other interfaces.
	IsAtributosContext()
}

type AtributosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtributosContext() *AtributosContext {
	var p = new(AtributosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_atributos
	return p
}

func InitEmptyAtributosContext(p *AtributosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_atributos
}

func (*AtributosContext) IsAtributosContext() {}

func NewAtributosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtributosContext {
	var p = new(AtributosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_atributos

	return p
}

func (s *AtributosContext) GetParser() antlr.Parser { return s.parser }

func (s *AtributosContext) CopyAll(ctx *AtributosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AtributosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtributosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Atributos_vector_emptyContext struct {
	AtributosContext
}

func NewAtributos_vector_emptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Atributos_vector_emptyContext {
	var p = new(Atributos_vector_emptyContext)

	InitEmptyAtributosContext(&p.AtributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtributosContext))

	return p
}

func (s *Atributos_vector_emptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atributos_vector_emptyContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Atributos_vector_emptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAtributos_vector_empty(s)

	default:
		return t.VisitChildren(s)
	}
}

type Atributos_generalesContext struct {
	AtributosContext
}

func NewAtributos_generalesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Atributos_generalesContext {
	var p = new(Atributos_generalesContext)

	InitEmptyAtributosContext(&p.AtributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtributosContext))

	return p
}

func (s *Atributos_generalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atributos_generalesContext) AllIde_atributo() []IIde_atributoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIde_atributoContext); ok {
			len++
		}
	}

	tst := make([]IIde_atributoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIde_atributoContext); ok {
			tst[i] = t.(IIde_atributoContext)
			i++
		}
	}

	return tst
}

func (s *Atributos_generalesContext) Ide_atributo(i int) IIde_atributoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIde_atributoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIde_atributoContext)
}

func (s *Atributos_generalesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAtributos_generales(s)

	default:
		return t.VisitChildren(s)
	}
}

type Atributos_vector_countContext struct {
	AtributosContext
}

func NewAtributos_vector_countContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Atributos_vector_countContext {
	var p = new(Atributos_vector_countContext)

	InitEmptyAtributosContext(&p.AtributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtributosContext))

	return p
}

func (s *Atributos_vector_countContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atributos_vector_countContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Atributos_vector_countContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAtributos_vector_count(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Atributos() (localctx IAtributosContext) {
	localctx = NewAtributosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, T_swiftParserRULE_atributos)
	var _alt int

	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAtributos_vector_emptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(T_swiftParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(T_swiftParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewAtributos_vector_countContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(333)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Match(T_swiftParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(T_swiftParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewAtributos_generalesContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(336)
			p.Ide_atributo()
		}
		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(337)
					p.Match(T_swiftParserT__16)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(338)
					p.Ide_atributo()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(341)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIde_atributoContext is an interface to support dynamic dispatch.
type IIde_atributoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIde_atributoContext differentiates from other interfaces.
	IsIde_atributoContext()
}

type Ide_atributoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIde_atributoContext() *Ide_atributoContext {
	var p = new(Ide_atributoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_ide_atributo
	return p
}

func InitEmptyIde_atributoContext(p *Ide_atributoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_ide_atributo
}

func (*Ide_atributoContext) IsIde_atributoContext() {}

func NewIde_atributoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ide_atributoContext {
	var p = new(Ide_atributoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_ide_atributo

	return p
}

func (s *Ide_atributoContext) GetParser() antlr.Parser { return s.parser }

func (s *Ide_atributoContext) CopyAll(ctx *Ide_atributoContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Ide_atributoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ide_atributoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Ide_atributo_simpleContext struct {
	Ide_atributoContext
}

func NewIde_atributo_simpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Ide_atributo_simpleContext {
	var p = new(Ide_atributo_simpleContext)

	InitEmptyIde_atributoContext(&p.Ide_atributoContext)
	p.parser = parser
	p.CopyAll(ctx.(*Ide_atributoContext))

	return p
}

func (s *Ide_atributo_simpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ide_atributo_simpleContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Ide_atributo_simpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitIde_atributo_simple(s)

	default:
		return t.VisitChildren(s)
	}
}

type Ide_atributo_vectorContext struct {
	Ide_atributoContext
}

func NewIde_atributo_vectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Ide_atributo_vectorContext {
	var p = new(Ide_atributo_vectorContext)

	InitEmptyIde_atributoContext(&p.Ide_atributoContext)
	p.parser = parser
	p.CopyAll(ctx.(*Ide_atributoContext))

	return p
}

func (s *Ide_atributo_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ide_atributo_vectorContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Ide_atributo_vectorContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Ide_atributo_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitIde_atributo_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Ide_atributo() (localctx IIde_atributoContext) {
	localctx = NewIde_atributoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, T_swiftParserRULE_ide_atributo)
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIde_atributo_simpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewIde_atributo_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)
			p.expresion(0)
		}
		{
			p.SetState(349)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignar_atributosContext is an interface to support dynamic dispatch.
type IAsignar_atributosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsignar_atributosContext differentiates from other interfaces.
	IsAsignar_atributosContext()
}

type Asignar_atributosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignar_atributosContext() *Asignar_atributosContext {
	var p = new(Asignar_atributosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_asignar_atributos
	return p
}

func InitEmptyAsignar_atributosContext(p *Asignar_atributosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_asignar_atributos
}

func (*Asignar_atributosContext) IsAsignar_atributosContext() {}

func NewAsignar_atributosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asignar_atributosContext {
	var p = new(Asignar_atributosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_asignar_atributos

	return p
}

func (s *Asignar_atributosContext) GetParser() antlr.Parser { return s.parser }

func (s *Asignar_atributosContext) CopyAll(ctx *Asignar_atributosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Asignar_atributosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignar_atributosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Incre_atributos_normalContext struct {
	Asignar_atributosContext
}

func NewIncre_atributos_normalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Incre_atributos_normalContext {
	var p = new(Incre_atributos_normalContext)

	InitEmptyAsignar_atributosContext(&p.Asignar_atributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asignar_atributosContext))

	return p
}

func (s *Incre_atributos_normalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Incre_atributos_normalContext) AllIde_atributo() []IIde_atributoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIde_atributoContext); ok {
			len++
		}
	}

	tst := make([]IIde_atributoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIde_atributoContext); ok {
			tst[i] = t.(IIde_atributoContext)
			i++
		}
	}

	return tst
}

func (s *Incre_atributos_normalContext) Ide_atributo(i int) IIde_atributoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIde_atributoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIde_atributoContext)
}

func (s *Incre_atributos_normalContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Incre_atributos_normalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitIncre_atributos_normal(s)

	default:
		return t.VisitChildren(s)
	}
}

type Decre_atributos_normalContext struct {
	Asignar_atributosContext
}

func NewDecre_atributos_normalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Decre_atributos_normalContext {
	var p = new(Decre_atributos_normalContext)

	InitEmptyAsignar_atributosContext(&p.Asignar_atributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asignar_atributosContext))

	return p
}

func (s *Decre_atributos_normalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decre_atributos_normalContext) AllIde_atributo() []IIde_atributoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIde_atributoContext); ok {
			len++
		}
	}

	tst := make([]IIde_atributoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIde_atributoContext); ok {
			tst[i] = t.(IIde_atributoContext)
			i++
		}
	}

	return tst
}

func (s *Decre_atributos_normalContext) Ide_atributo(i int) IIde_atributoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIde_atributoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIde_atributoContext)
}

func (s *Decre_atributos_normalContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Decre_atributos_normalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDecre_atributos_normal(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignar_atributos_normalContext struct {
	Asignar_atributosContext
}

func NewAsignar_atributos_normalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignar_atributos_normalContext {
	var p = new(Asignar_atributos_normalContext)

	InitEmptyAsignar_atributosContext(&p.Asignar_atributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asignar_atributosContext))

	return p
}

func (s *Asignar_atributos_normalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignar_atributos_normalContext) AllIde_atributo() []IIde_atributoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIde_atributoContext); ok {
			len++
		}
	}

	tst := make([]IIde_atributoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIde_atributoContext); ok {
			tst[i] = t.(IIde_atributoContext)
			i++
		}
	}

	return tst
}

func (s *Asignar_atributos_normalContext) Ide_atributo(i int) IIde_atributoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIde_atributoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIde_atributoContext)
}

func (s *Asignar_atributos_normalContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignar_atributos_normalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignar_atributos_normal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Asignar_atributos() (localctx IAsignar_atributosContext) {
	localctx = NewAsignar_atributosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, T_swiftParserRULE_asignar_atributos)
	var _la int

	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignar_atributos_normalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)
			p.Ide_atributo()
		}
		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == T_swiftParserT__16 {
			{
				p.SetState(354)
				p.Match(T_swiftParserT__16)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(355)
				p.Ide_atributo()
			}

			p.SetState(358)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(360)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.expresion(0)
		}

	case 2:
		localctx = NewIncre_atributos_normalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(363)
			p.Ide_atributo()
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == T_swiftParserT__16 {
			{
				p.SetState(364)
				p.Match(T_swiftParserT__16)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(365)
				p.Ide_atributo()
			}

			p.SetState(368)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(370)
			p.Match(T_swiftParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.expresion(0)
		}

	case 3:
		localctx = NewDecre_atributos_normalContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.Ide_atributo()
		}
		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == T_swiftParserT__16 {
			{
				p.SetState(374)
				p.Match(T_swiftParserT__16)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(375)
				p.Ide_atributo()
			}

			p.SetState(378)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(380)
			p.Match(T_swiftParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatriz_declaracionContext is an interface to support dynamic dispatch.
type IMatriz_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Definicion_matriz() IDefinicion_matrizContext
	Tipo_matriz() ITipo_matrizContext

	// IsMatriz_declaracionContext differentiates from other interfaces.
	IsMatriz_declaracionContext()
}

type Matriz_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatriz_declaracionContext() *Matriz_declaracionContext {
	var p = new(Matriz_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_matriz_declaracion
	return p
}

func InitEmptyMatriz_declaracionContext(p *Matriz_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_matriz_declaracion
}

func (*Matriz_declaracionContext) IsMatriz_declaracionContext() {}

func NewMatriz_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matriz_declaracionContext {
	var p = new(Matriz_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_matriz_declaracion

	return p
}

func (s *Matriz_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Matriz_declaracionContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Matriz_declaracionContext) Definicion_matriz() IDefinicion_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinicion_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinicion_matrizContext)
}

func (s *Matriz_declaracionContext) Tipo_matriz() ITipo_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_matrizContext)
}

func (s *Matriz_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matriz_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matriz_declaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitMatriz_declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Matriz_declaracion() (localctx IMatriz_declaracionContext) {
	localctx = NewMatriz_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, T_swiftParserRULE_matriz_declaracion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(T_swiftParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__6 {
		{
			p.SetState(387)
			p.Match(T_swiftParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.Tipo_matriz()
		}

	}
	{
		p.SetState(391)
		p.Match(T_swiftParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.Definicion_matriz()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipo_matrizContext is an interface to support dynamic dispatch.
type ITipo_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipo_matriz() ITipo_matrizContext
	Tipos() ITiposContext

	// IsTipo_matrizContext differentiates from other interfaces.
	IsTipo_matrizContext()
}

type Tipo_matrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipo_matrizContext() *Tipo_matrizContext {
	var p = new(Tipo_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_tipo_matriz
	return p
}

func InitEmptyTipo_matrizContext(p *Tipo_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_tipo_matriz
}

func (*Tipo_matrizContext) IsTipo_matrizContext() {}

func NewTipo_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_matrizContext {
	var p = new(Tipo_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_tipo_matriz

	return p
}

func (s *Tipo_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_matrizContext) Tipo_matriz() ITipo_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_matrizContext)
}

func (s *Tipo_matrizContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *Tipo_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_matrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitTipo_matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Tipo_matriz() (localctx ITipo_matrizContext) {
	localctx = NewTipo_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, T_swiftParserRULE_tipo_matriz)
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(394)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
			p.Tipo_matriz()
		}
		{
			p.SetState(396)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(398)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.Tipos()
		}
		{
			p.SetState(400)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefinicion_matrizContext is an interface to support dynamic dispatch.
type IDefinicion_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lista_valores_matriz() ILista_valores_matrizContext
	Simple_matriz() ISimple_matrizContext

	// IsDefinicion_matrizContext differentiates from other interfaces.
	IsDefinicion_matrizContext()
}

type Definicion_matrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinicion_matrizContext() *Definicion_matrizContext {
	var p = new(Definicion_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_definicion_matriz
	return p
}

func InitEmptyDefinicion_matrizContext(p *Definicion_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_definicion_matriz
}

func (*Definicion_matrizContext) IsDefinicion_matrizContext() {}

func NewDefinicion_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Definicion_matrizContext {
	var p = new(Definicion_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_definicion_matriz

	return p
}

func (s *Definicion_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Definicion_matrizContext) Lista_valores_matriz() ILista_valores_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_valores_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_valores_matrizContext)
}

func (s *Definicion_matrizContext) Simple_matriz() ISimple_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_matrizContext)
}

func (s *Definicion_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Definicion_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Definicion_matrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDefinicion_matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Definicion_matriz() (localctx IDefinicion_matrizContext) {
	localctx = NewDefinicion_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, T_swiftParserRULE_definicion_matriz)
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(404)
			p.Lista_valores_matriz()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(405)
			p.Simple_matriz()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_valores_matrizContext is an interface to support dynamic dispatch.
type ILista_valores_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Elementos_matriz() IElementos_matrizContext

	// IsLista_valores_matrizContext differentiates from other interfaces.
	IsLista_valores_matrizContext()
}

type Lista_valores_matrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_valores_matrizContext() *Lista_valores_matrizContext {
	var p = new(Lista_valores_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_valores_matriz
	return p
}

func InitEmptyLista_valores_matrizContext(p *Lista_valores_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_valores_matriz
}

func (*Lista_valores_matrizContext) IsLista_valores_matrizContext() {}

func NewLista_valores_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_valores_matrizContext {
	var p = new(Lista_valores_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_lista_valores_matriz

	return p
}

func (s *Lista_valores_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_valores_matrizContext) Elementos_matriz() IElementos_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementos_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementos_matrizContext)
}

func (s *Lista_valores_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_valores_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_valores_matrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLista_valores_matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Lista_valores_matriz() (localctx ILista_valores_matrizContext) {
	localctx = NewLista_valores_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, T_swiftParserRULE_lista_valores_matriz)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(408)
		p.Match(T_swiftParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(409)
		p.Elementos_matriz()
	}
	{
		p.SetState(410)
		p.Match(T_swiftParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElementos_matrizContext is an interface to support dynamic dispatch.
type IElementos_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllElemento_matriz() []IElemento_matrizContext
	Elemento_matriz(i int) IElemento_matrizContext

	// IsElementos_matrizContext differentiates from other interfaces.
	IsElementos_matrizContext()
}

type Elementos_matrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementos_matrizContext() *Elementos_matrizContext {
	var p = new(Elementos_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_elementos_matriz
	return p
}

func InitEmptyElementos_matrizContext(p *Elementos_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_elementos_matriz
}

func (*Elementos_matrizContext) IsElementos_matrizContext() {}

func NewElementos_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Elementos_matrizContext {
	var p = new(Elementos_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_elementos_matriz

	return p
}

func (s *Elementos_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Elementos_matrizContext) AllElemento_matriz() []IElemento_matrizContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElemento_matrizContext); ok {
			len++
		}
	}

	tst := make([]IElemento_matrizContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElemento_matrizContext); ok {
			tst[i] = t.(IElemento_matrizContext)
			i++
		}
	}

	return tst
}

func (s *Elementos_matrizContext) Elemento_matriz(i int) IElemento_matrizContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElemento_matrizContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElemento_matrizContext)
}

func (s *Elementos_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Elementos_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Elementos_matrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitElementos_matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Elementos_matriz() (localctx IElementos_matrizContext) {
	localctx = NewElementos_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, T_swiftParserRULE_elementos_matriz)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(412)
		p.Elemento_matriz()
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__5 {
		{
			p.SetState(413)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.Elemento_matriz()
		}

		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElemento_matrizContext is an interface to support dynamic dispatch.
type IElemento_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lista_valores_matriz() ILista_valores_matrizContext
	Expresion() IExpresionContext

	// IsElemento_matrizContext differentiates from other interfaces.
	IsElemento_matrizContext()
}

type Elemento_matrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElemento_matrizContext() *Elemento_matrizContext {
	var p = new(Elemento_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_elemento_matriz
	return p
}

func InitEmptyElemento_matrizContext(p *Elemento_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_elemento_matriz
}

func (*Elemento_matrizContext) IsElemento_matrizContext() {}

func NewElemento_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Elemento_matrizContext {
	var p = new(Elemento_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_elemento_matriz

	return p
}

func (s *Elemento_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Elemento_matrizContext) Lista_valores_matriz() ILista_valores_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_valores_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_valores_matrizContext)
}

func (s *Elemento_matrizContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Elemento_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Elemento_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Elemento_matrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitElemento_matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Elemento_matriz() (localctx IElemento_matrizContext) {
	localctx = NewElemento_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, T_swiftParserRULE_elemento_matriz)
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(420)
			p.Lista_valores_matriz()
		}

	case T_swiftParserT__2, T_swiftParserT__38, T_swiftParserT__39, T_swiftParserT__40, T_swiftParserT__43, T_swiftParserT__48, T_swiftParserT__49, T_swiftParserT__62, T_swiftParserInt, T_swiftParserFloat, T_swiftParserCaracter, T_swiftParserString_, T_swiftParserBool, T_swiftParserIdentificador:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(421)
			p.expresion(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimple_matrizContext is an interface to support dynamic dispatch.
type ISimple_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipo_matriz() ITipo_matrizContext
	Simple_matriz() ISimple_matrizContext
	Int() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsSimple_matrizContext differentiates from other interfaces.
	IsSimple_matrizContext()
}

type Simple_matrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_matrizContext() *Simple_matrizContext {
	var p = new(Simple_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_simple_matriz
	return p
}

func InitEmptySimple_matrizContext(p *Simple_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_simple_matriz
}

func (*Simple_matrizContext) IsSimple_matrizContext() {}

func NewSimple_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_matrizContext {
	var p = new(Simple_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_simple_matriz

	return p
}

func (s *Simple_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_matrizContext) Tipo_matriz() ITipo_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_matrizContext)
}

func (s *Simple_matrizContext) Simple_matriz() ISimple_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_matrizContext)
}

func (s *Simple_matrizContext) Int() antlr.TerminalNode {
	return s.GetToken(T_swiftParserInt, 0)
}

func (s *Simple_matrizContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Simple_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_matrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitSimple_matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Simple_matriz() (localctx ISimple_matrizContext) {
	localctx = NewSimple_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, T_swiftParserRULE_simple_matriz)
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(424)
			p.Tipo_matriz()
		}
		{
			p.SetState(425)
			p.Match(T_swiftParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
			p.Match(T_swiftParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.Match(T_swiftParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.Simple_matriz()
		}
		{
			p.SetState(429)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.Match(T_swiftParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.Match(T_swiftParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)
			p.Match(T_swiftParserInt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.Match(T_swiftParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(435)
			p.Tipo_matriz()
		}
		{
			p.SetState(436)
			p.Match(T_swiftParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)
			p.Match(T_swiftParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)
			p.Match(T_swiftParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(439)
			p.expresion(0)
		}
		{
			p.SetState(440)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(441)
			p.Match(T_swiftParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(442)
			p.Match(T_swiftParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(443)
			p.Match(T_swiftParserInt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
			p.Match(T_swiftParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_in_statementContext is an interface to support dynamic dispatch.
type IFor_in_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Code_block() ICode_blockContext
	Expresion() IExpresionContext
	Rango() IRangoContext

	// IsFor_in_statementContext differentiates from other interfaces.
	IsFor_in_statementContext()
}

type For_in_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_in_statementContext() *For_in_statementContext {
	var p = new(For_in_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_for_in_statement
	return p
}

func InitEmptyFor_in_statementContext(p *For_in_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_for_in_statement
}

func (*For_in_statementContext) IsFor_in_statementContext() {}

func NewFor_in_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_in_statementContext {
	var p = new(For_in_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_for_in_statement

	return p
}

func (s *For_in_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *For_in_statementContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *For_in_statementContext) Code_block() ICode_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *For_in_statementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *For_in_statementContext) Rango() IRangoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangoContext)
}

func (s *For_in_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_in_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_in_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFor_in_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) For_in_statement() (localctx IFor_in_statementContext) {
	localctx = NewFor_in_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, T_swiftParserRULE_for_in_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.Match(T_swiftParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(449)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(450)
		p.Match(T_swiftParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(451)
			p.expresion(0)
		}

	case 2:
		{
			p.SetState(452)
			p.Rango()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(455)
		p.Code_block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangoContext is an interface to support dynamic dispatch.
type IRangoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext

	// IsRangoContext differentiates from other interfaces.
	IsRangoContext()
}

type RangoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangoContext() *RangoContext {
	var p = new(RangoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_rango
	return p
}

func InitEmptyRangoContext(p *RangoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_rango
}

func (*RangoContext) IsRangoContext() {}

func NewRangoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangoContext {
	var p = new(RangoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_rango

	return p
}

func (s *RangoContext) GetParser() antlr.Parser { return s.parser }

func (s *RangoContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *RangoContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *RangoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitRango(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Rango() (localctx IRangoContext) {
	localctx = NewRangoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, T_swiftParserRULE_rango)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(457)
		p.expresion(0)
	}
	{
		p.SetState(458)
		p.Match(T_swiftParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(459)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_statementContext is an interface to support dynamic dispatch.
type IWhile_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext
	Code_block() ICode_blockContext

	// IsWhile_statementContext differentiates from other interfaces.
	IsWhile_statementContext()
}

type While_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_statementContext() *While_statementContext {
	var p = new(While_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_while_statement
	return p
}

func InitEmptyWhile_statementContext(p *While_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_while_statement
}

func (*While_statementContext) IsWhile_statementContext() {}

func NewWhile_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_statementContext {
	var p = new(While_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_while_statement

	return p
}

func (s *While_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *While_statementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *While_statementContext) Code_block() ICode_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *While_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitWhile_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) While_statement() (localctx IWhile_statementContext) {
	localctx = NewWhile_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, T_swiftParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(461)
		p.Match(T_swiftParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(462)
		p.expresion(0)
	}
	{
		p.SetState(463)
		p.Code_block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) CopyAll(ctx *If_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Else_finalContext struct {
	If_statementContext
}

func NewElse_finalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Else_finalContext {
	var p = new(Else_finalContext)

	InitEmptyIf_statementContext(&p.If_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_statementContext))

	return p
}

func (s *Else_finalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_finalContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Else_finalContext) AllCode_block() []ICode_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICode_blockContext); ok {
			len++
		}
	}

	tst := make([]ICode_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICode_blockContext); ok {
			tst[i] = t.(ICode_blockContext)
			i++
		}
	}

	return tst
}

func (s *Else_finalContext) Code_block(i int) ICode_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *Else_finalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitElse_final(s)

	default:
		return t.VisitChildren(s)
	}
}

type If_simpleContext struct {
	If_statementContext
}

func NewIf_simpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *If_simpleContext {
	var p = new(If_simpleContext)

	InitEmptyIf_statementContext(&p.If_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_statementContext))

	return p
}

func (s *If_simpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_simpleContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *If_simpleContext) Code_block() ICode_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *If_simpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitIf_simple(s)

	default:
		return t.VisitChildren(s)
	}
}

type Siguiente_ifContext struct {
	If_statementContext
}

func NewSiguiente_ifContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Siguiente_ifContext {
	var p = new(Siguiente_ifContext)

	InitEmptyIf_statementContext(&p.If_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_statementContext))

	return p
}

func (s *Siguiente_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Siguiente_ifContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Siguiente_ifContext) Code_block() ICode_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *Siguiente_ifContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *Siguiente_ifContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitSiguiente_if(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, T_swiftParserRULE_if_statement)
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIf_simpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(465)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(466)
			p.expresion(0)
		}
		{
			p.SetState(467)
			p.Code_block()
		}

	case 2:
		localctx = NewElse_finalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(469)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
			p.expresion(0)
		}
		{
			p.SetState(471)
			p.Code_block()
		}
		{
			p.SetState(472)
			p.Match(T_swiftParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(473)
			p.Code_block()
		}

	case 3:
		localctx = NewSiguiente_ifContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(475)
			p.Match(T_swiftParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)
			p.expresion(0)
		}
		{
			p.SetState(477)
			p.Code_block()
		}
		{
			p.SetState(478)
			p.Match(T_swiftParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)
			p.If_statement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuard_statementContext is an interface to support dynamic dispatch.
type IGuard_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext
	Instrucciones() IInstruccionesContext

	// IsGuard_statementContext differentiates from other interfaces.
	IsGuard_statementContext()
}

type Guard_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuard_statementContext() *Guard_statementContext {
	var p = new(Guard_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_guard_statement
	return p
}

func InitEmptyGuard_statementContext(p *Guard_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_guard_statement
}

func (*Guard_statementContext) IsGuard_statementContext() {}

func NewGuard_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Guard_statementContext {
	var p = new(Guard_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_guard_statement

	return p
}

func (s *Guard_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Guard_statementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Guard_statementContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Guard_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Guard_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Guard_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitGuard_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Guard_statement() (localctx IGuard_statementContext) {
	localctx = NewGuard_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, T_swiftParserRULE_guard_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Match(T_swiftParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(484)
		p.expresion(0)
	}
	{
		p.SetState(485)
		p.Match(T_swiftParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(486)
		p.Match(T_swiftParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(487)
		p.Instrucciones()
	}
	{
		p.SetState(488)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30064771072) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(489)
		p.Match(T_swiftParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_statementContext is an interface to support dynamic dispatch.
type ISwitch_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext
	AllSwitch_case() []ISwitch_caseContext
	Switch_case(i int) ISwitch_caseContext
	Default_case() IDefault_caseContext

	// IsSwitch_statementContext differentiates from other interfaces.
	IsSwitch_statementContext()
}

type Switch_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_statementContext() *Switch_statementContext {
	var p = new(Switch_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_switch_statement
	return p
}

func InitEmptySwitch_statementContext(p *Switch_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_switch_statement
}

func (*Switch_statementContext) IsSwitch_statementContext() {}

func NewSwitch_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_statementContext {
	var p = new(Switch_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_switch_statement

	return p
}

func (s *Switch_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_statementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Switch_statementContext) AllSwitch_case() []ISwitch_caseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			len++
		}
	}

	tst := make([]ISwitch_caseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitch_caseContext); ok {
			tst[i] = t.(ISwitch_caseContext)
			i++
		}
	}

	return tst
}

func (s *Switch_statementContext) Switch_case(i int) ISwitch_caseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_caseContext)
}

func (s *Switch_statementContext) Default_case() IDefault_caseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefault_caseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefault_caseContext)
}

func (s *Switch_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitSwitch_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Switch_statement() (localctx ISwitch_statementContext) {
	localctx = NewSwitch_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, T_swiftParserRULE_switch_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
		p.Match(T_swiftParserT__34)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(492)
		p.expresion(0)
	}
	{
		p.SetState(493)
		p.Match(T_swiftParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__35 {
		{
			p.SetState(494)
			p.Switch_case()
		}

		p.SetState(499)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__36 {
		{
			p.SetState(500)
			p.Default_case()
		}

	}
	{
		p.SetState(503)
		p.Match(T_swiftParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_caseContext is an interface to support dynamic dispatch.
type ISwitch_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext
	Instrucciones() IInstruccionesContext

	// IsSwitch_caseContext differentiates from other interfaces.
	IsSwitch_caseContext()
}

type Switch_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_caseContext() *Switch_caseContext {
	var p = new(Switch_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_switch_case
	return p
}

func InitEmptySwitch_caseContext(p *Switch_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_switch_case
}

func (*Switch_caseContext) IsSwitch_caseContext() {}

func NewSwitch_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_caseContext {
	var p = new(Switch_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_switch_case

	return p
}

func (s *Switch_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_caseContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Switch_caseContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Switch_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_caseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitSwitch_case(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Switch_case() (localctx ISwitch_caseContext) {
	localctx = NewSwitch_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, T_swiftParserRULE_switch_case)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)
		p.Match(T_swiftParserT__35)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(506)
		p.expresion(0)
	}
	{
		p.SetState(507)
		p.Match(T_swiftParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(508)
		p.Instrucciones()
	}
	p.SetState(510)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__32 {
		{
			p.SetState(509)
			p.Match(T_swiftParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefault_caseContext is an interface to support dynamic dispatch.
type IDefault_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instrucciones() IInstruccionesContext

	// IsDefault_caseContext differentiates from other interfaces.
	IsDefault_caseContext()
}

type Default_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_caseContext() *Default_caseContext {
	var p = new(Default_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_default_case
	return p
}

func InitEmptyDefault_caseContext(p *Default_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_default_case
}

func (*Default_caseContext) IsDefault_caseContext() {}

func NewDefault_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_caseContext {
	var p = new(Default_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_default_case

	return p
}

func (s *Default_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_caseContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Default_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Default_caseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDefault_case(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Default_case() (localctx IDefault_caseContext) {
	localctx = NewDefault_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, T_swiftParserRULE_default_case)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		p.Match(T_swiftParserT__36)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(513)
		p.Match(T_swiftParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(514)
		p.Instrucciones()
	}
	p.SetState(516)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__32 {
		{
			p.SetState(515)
			p.Match(T_swiftParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreak_statementContext is an interface to support dynamic dispatch.
type IBreak_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBreak_statementContext differentiates from other interfaces.
	IsBreak_statementContext()
}

type Break_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_statementContext() *Break_statementContext {
	var p = new(Break_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_break_statement
	return p
}

func InitEmptyBreak_statementContext(p *Break_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_break_statement
}

func (*Break_statementContext) IsBreak_statementContext() {}

func NewBreak_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_statementContext {
	var p = new(Break_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_break_statement

	return p
}

func (s *Break_statementContext) GetParser() antlr.Parser { return s.parser }
func (s *Break_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitBreak_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Break_statement() (localctx IBreak_statementContext) {
	localctx = NewBreak_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, T_swiftParserRULE_break_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(518)
		p.Match(T_swiftParserT__32)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinue_statementContext is an interface to support dynamic dispatch.
type IContinue_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsContinue_statementContext differentiates from other interfaces.
	IsContinue_statementContext()
}

type Continue_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_statementContext() *Continue_statementContext {
	var p = new(Continue_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_continue_statement
	return p
}

func InitEmptyContinue_statementContext(p *Continue_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_continue_statement
}

func (*Continue_statementContext) IsContinue_statementContext() {}

func NewContinue_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_statementContext {
	var p = new(Continue_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_continue_statement

	return p
}

func (s *Continue_statementContext) GetParser() antlr.Parser { return s.parser }
func (s *Continue_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitContinue_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Continue_statement() (localctx IContinue_statementContext) {
	localctx = NewContinue_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, T_swiftParserRULE_continue_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(520)
		p.Match(T_swiftParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturn_statementContext is an interface to support dynamic dispatch.
type IReturn_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext

	// IsReturn_statementContext differentiates from other interfaces.
	IsReturn_statementContext()
}

type Return_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_statementContext() *Return_statementContext {
	var p = new(Return_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_return_statement
	return p
}

func InitEmptyReturn_statementContext(p *Return_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_return_statement
}

func (*Return_statementContext) IsReturn_statementContext() {}

func NewReturn_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_statementContext {
	var p = new(Return_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_return_statement

	return p
}

func (s *Return_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_statementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Return_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitReturn_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Return_statement() (localctx IReturn_statementContext) {
	localctx = NewReturn_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, T_swiftParserRULE_return_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(522)
		p.Match(T_swiftParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(524)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(523)
			p.expresion(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstant_declaracionContext is an interface to support dynamic dispatch.
type IConstant_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Expresion() IExpresionContext
	Anotacion_tipo() IAnotacion_tipoContext

	// IsConstant_declaracionContext differentiates from other interfaces.
	IsConstant_declaracionContext()
}

type Constant_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstant_declaracionContext() *Constant_declaracionContext {
	var p = new(Constant_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_constant_declaracion
	return p
}

func InitEmptyConstant_declaracionContext(p *Constant_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_constant_declaracion
}

func (*Constant_declaracionContext) IsConstant_declaracionContext() {}

func NewConstant_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Constant_declaracionContext {
	var p = new(Constant_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_constant_declaracion

	return p
}

func (s *Constant_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Constant_declaracionContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Constant_declaracionContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Constant_declaracionContext) Anotacion_tipo() IAnotacion_tipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnotacion_tipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnotacion_tipoContext)
}

func (s *Constant_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constant_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Constant_declaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitConstant_declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Constant_declaracion() (localctx IConstant_declaracionContext) {
	localctx = NewConstant_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, T_swiftParserRULE_constant_declaracion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(526)
		p.Match(T_swiftParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(527)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == T_swiftParserT__6 {
		{
			p.SetState(528)
			p.Anotacion_tipo()
		}

	}
	{
		p.SetState(531)
		p.Match(T_swiftParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(532)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariable_declaracionContext is an interface to support dynamic dispatch.
type IVariable_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Anotacion_tipo() IAnotacion_tipoContext
	Expresion() IExpresionContext

	// IsVariable_declaracionContext differentiates from other interfaces.
	IsVariable_declaracionContext()
}

type Variable_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariable_declaracionContext() *Variable_declaracionContext {
	var p = new(Variable_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_variable_declaracion
	return p
}

func InitEmptyVariable_declaracionContext(p *Variable_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_variable_declaracion
}

func (*Variable_declaracionContext) IsVariable_declaracionContext() {}

func NewVariable_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_declaracionContext {
	var p = new(Variable_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_variable_declaracion

	return p
}

func (s *Variable_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_declaracionContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Variable_declaracionContext) Anotacion_tipo() IAnotacion_tipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnotacion_tipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnotacion_tipoContext)
}

func (s *Variable_declaracionContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Variable_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Variable_declaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitVariable_declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Variable_declaracion() (localctx IVariable_declaracionContext) {
	localctx = NewVariable_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, T_swiftParserRULE_variable_declaracion)
	var _la int

	p.SetState(546)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(534)
			p.Match(T_swiftParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(535)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(536)
			p.Anotacion_tipo()
		}
		{
			p.SetState(537)
			p.Match(T_swiftParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(539)
			p.Match(T_swiftParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(540)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == T_swiftParserT__6 {
			{
				p.SetState(541)
				p.Anotacion_tipo()
			}

		}
		{
			p.SetState(544)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(545)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnotacion_tipoContext is an interface to support dynamic dispatch.
type IAnotacion_tipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipos() ITiposContext

	// IsAnotacion_tipoContext differentiates from other interfaces.
	IsAnotacion_tipoContext()
}

type Anotacion_tipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnotacion_tipoContext() *Anotacion_tipoContext {
	var p = new(Anotacion_tipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_anotacion_tipo
	return p
}

func InitEmptyAnotacion_tipoContext(p *Anotacion_tipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_anotacion_tipo
}

func (*Anotacion_tipoContext) IsAnotacion_tipoContext() {}

func NewAnotacion_tipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Anotacion_tipoContext {
	var p = new(Anotacion_tipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_anotacion_tipo

	return p
}

func (s *Anotacion_tipoContext) GetParser() antlr.Parser { return s.parser }

func (s *Anotacion_tipoContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *Anotacion_tipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Anotacion_tipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Anotacion_tipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAnotacion_tipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Anotacion_tipo() (localctx IAnotacion_tipoContext) {
	localctx = NewAnotacion_tipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, T_swiftParserRULE_anotacion_tipo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(548)
		p.Match(T_swiftParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(549)
		p.Tipos()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITiposContext is an interface to support dynamic dispatch.
type ITiposContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode

	// IsTiposContext differentiates from other interfaces.
	IsTiposContext()
}

type TiposContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTiposContext() *TiposContext {
	var p = new(TiposContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_tipos
	return p
}

func InitEmptyTiposContext(p *TiposContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_tipos
}

func (*TiposContext) IsTiposContext() {}

func NewTiposContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TiposContext {
	var p = new(TiposContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_tipos

	return p
}

func (s *TiposContext) GetParser() antlr.Parser { return s.parser }

func (s *TiposContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *TiposContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TiposContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TiposContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitTipos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Tipos() (localctx ITiposContext) {
	localctx = NewTiposContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, T_swiftParserRULE_tipos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(551)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-39)) & ^0x3f) == 0 && ((int64(1)<<(_la-39))&8589934623) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_declaracionContext is an interface to support dynamic dispatch.
type IArray_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArray_declaracionContext differentiates from other interfaces.
	IsArray_declaracionContext()
}

type Array_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_declaracionContext() *Array_declaracionContext {
	var p = new(Array_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_array_declaracion
	return p
}

func InitEmptyArray_declaracionContext(p *Array_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_array_declaracion
}

func (*Array_declaracionContext) IsArray_declaracionContext() {}

func NewArray_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_declaracionContext {
	var p = new(Array_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_array_declaracion

	return p
}

func (s *Array_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_declaracionContext) CopyAll(ctx *Array_declaracionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Array_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Array_comunContext struct {
	Array_declaracionContext
}

func NewArray_comunContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Array_comunContext {
	var p = new(Array_comunContext)

	InitEmptyArray_declaracionContext(&p.Array_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Array_declaracionContext))

	return p
}

func (s *Array_comunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_comunContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Array_comunContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *Array_comunContext) Definicion_vector() IDefinicion_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinicion_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinicion_vectorContext)
}

func (s *Array_comunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitArray_comun(s)

	default:
		return t.VisitChildren(s)
	}
}

type Array_vacioContext struct {
	Array_declaracionContext
}

func NewArray_vacioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Array_vacioContext {
	var p = new(Array_vacioContext)

	InitEmptyArray_declaracionContext(&p.Array_declaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Array_declaracionContext))

	return p
}

func (s *Array_vacioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_vacioContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Array_vacioContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *Array_vacioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitArray_vacio(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Array_declaracion() (localctx IArray_declaracionContext) {
	localctx = NewArray_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, T_swiftParserRULE_array_declaracion)
	p.SetState(570)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArray_comunContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(553)
			p.Match(T_swiftParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(554)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(555)
			p.Match(T_swiftParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(556)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(557)
			p.Tipos()
		}
		{
			p.SetState(558)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(559)
			p.Definicion_vector()
		}

	case 2:
		localctx = NewArray_vacioContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(561)
			p.Match(T_swiftParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(562)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(563)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(564)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(565)
			p.Tipos()
		}
		{
			p.SetState(566)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(567)
			p.Match(T_swiftParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(568)
			p.Match(T_swiftParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefinicion_vectorContext is an interface to support dynamic dispatch.
type IDefinicion_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lista_expresiones() ILista_expresionesContext
	Identificador() antlr.TerminalNode

	// IsDefinicion_vectorContext differentiates from other interfaces.
	IsDefinicion_vectorContext()
}

type Definicion_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinicion_vectorContext() *Definicion_vectorContext {
	var p = new(Definicion_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_definicion_vector
	return p
}

func InitEmptyDefinicion_vectorContext(p *Definicion_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_definicion_vector
}

func (*Definicion_vectorContext) IsDefinicion_vectorContext() {}

func NewDefinicion_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Definicion_vectorContext {
	var p = new(Definicion_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_definicion_vector

	return p
}

func (s *Definicion_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Definicion_vectorContext) Lista_expresiones() ILista_expresionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_expresionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_expresionesContext)
}

func (s *Definicion_vectorContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Definicion_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Definicion_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Definicion_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitDefinicion_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Definicion_vector() (localctx IDefinicion_vectorContext) {
	localctx = NewDefinicion_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, T_swiftParserRULE_definicion_vector)
	p.SetState(582)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(572)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(573)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(574)
			p.Lista_expresiones()
		}
		{
			p.SetState(575)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(577)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(578)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(579)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(580)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(581)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_expresionesContext is an interface to support dynamic dispatch.
type ILista_expresionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext

	// IsLista_expresionesContext differentiates from other interfaces.
	IsLista_expresionesContext()
}

type Lista_expresionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_expresionesContext() *Lista_expresionesContext {
	var p = new(Lista_expresionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_expresiones
	return p
}

func InitEmptyLista_expresionesContext(p *Lista_expresionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_lista_expresiones
}

func (*Lista_expresionesContext) IsLista_expresionesContext() {}

func NewLista_expresionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_expresionesContext {
	var p = new(Lista_expresionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_lista_expresiones

	return p
}

func (s *Lista_expresionesContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_expresionesContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Lista_expresionesContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Lista_expresionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_expresionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_expresionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitLista_expresiones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Lista_expresiones() (localctx ILista_expresionesContext) {
	localctx = NewLista_expresionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, T_swiftParserRULE_lista_expresiones)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(584)
		p.expresion(0)
	}
	p.SetState(589)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__5 {
		{
			p.SetState(585)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(586)
			p.expresion(0)
		}

		p.SetState(591)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_printContext is an interface to support dynamic dispatch.
type IFuncion_printContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lista_expresiones() ILista_expresionesContext

	// IsFuncion_printContext differentiates from other interfaces.
	IsFuncion_printContext()
}

type Funcion_printContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_printContext() *Funcion_printContext {
	var p = new(Funcion_printContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_print
	return p
}

func InitEmptyFuncion_printContext(p *Funcion_printContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_print
}

func (*Funcion_printContext) IsFuncion_printContext() {}

func NewFuncion_printContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_printContext {
	var p = new(Funcion_printContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_print

	return p
}

func (s *Funcion_printContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_printContext) Lista_expresiones() ILista_expresionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_expresionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_expresionesContext)
}

func (s *Funcion_printContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_printContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_printContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_print(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_print() (localctx IFuncion_printContext) {
	localctx = NewFuncion_printContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, T_swiftParserRULE_funcion_print)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(592)
		p.Match(T_swiftParserT__43)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(593)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(594)
		p.Lista_expresiones()
	}
	{
		p.SetState(595)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_appendContext is an interface to support dynamic dispatch.
type IFuncion_appendContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsFuncion_appendContext differentiates from other interfaces.
	IsFuncion_appendContext()
}

type Funcion_appendContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_appendContext() *Funcion_appendContext {
	var p = new(Funcion_appendContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_append
	return p
}

func InitEmptyFuncion_appendContext(p *Funcion_appendContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_append
}

func (*Funcion_appendContext) IsFuncion_appendContext() {}

func NewFuncion_appendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_appendContext {
	var p = new(Funcion_appendContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_append

	return p
}

func (s *Funcion_appendContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_appendContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Funcion_appendContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Funcion_appendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_appendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_appendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_append(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_append() (localctx IFuncion_appendContext) {
	localctx = NewFuncion_appendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, T_swiftParserRULE_funcion_append)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(597)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(598)
		p.Match(T_swiftParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(599)
		p.Match(T_swiftParserT__44)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(600)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(601)
		p.expresion(0)
	}
	{
		p.SetState(602)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_removeLastContext is an interface to support dynamic dispatch.
type IFuncion_removeLastContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode

	// IsFuncion_removeLastContext differentiates from other interfaces.
	IsFuncion_removeLastContext()
}

type Funcion_removeLastContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_removeLastContext() *Funcion_removeLastContext {
	var p = new(Funcion_removeLastContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_removeLast
	return p
}

func InitEmptyFuncion_removeLastContext(p *Funcion_removeLastContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_removeLast
}

func (*Funcion_removeLastContext) IsFuncion_removeLastContext() {}

func NewFuncion_removeLastContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_removeLastContext {
	var p = new(Funcion_removeLastContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_removeLast

	return p
}

func (s *Funcion_removeLastContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_removeLastContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Funcion_removeLastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_removeLastContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_removeLastContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_removeLast(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_removeLast() (localctx IFuncion_removeLastContext) {
	localctx = NewFuncion_removeLastContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, T_swiftParserRULE_funcion_removeLast)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(604)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(605)
		p.Match(T_swiftParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(606)
		p.Match(T_swiftParserT__45)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(607)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(608)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_removeatContext is an interface to support dynamic dispatch.
type IFuncion_removeatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identificador() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsFuncion_removeatContext differentiates from other interfaces.
	IsFuncion_removeatContext()
}

type Funcion_removeatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_removeatContext() *Funcion_removeatContext {
	var p = new(Funcion_removeatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_removeat
	return p
}

func InitEmptyFuncion_removeatContext(p *Funcion_removeatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_removeat
}

func (*Funcion_removeatContext) IsFuncion_removeatContext() {}

func NewFuncion_removeatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_removeatContext {
	var p = new(Funcion_removeatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_removeat

	return p
}

func (s *Funcion_removeatContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_removeatContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Funcion_removeatContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Funcion_removeatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_removeatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_removeatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_removeat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_removeat() (localctx IFuncion_removeatContext) {
	localctx = NewFuncion_removeatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, T_swiftParserRULE_funcion_removeat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(610)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(611)
		p.Match(T_swiftParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(612)
		p.Match(T_swiftParserT__46)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(613)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(614)
		p.Match(T_swiftParserT__47)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(615)
		p.Match(T_swiftParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(616)
		p.expresion(0)
	}
	{
		p.SetState(617)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_intContext is an interface to support dynamic dispatch.
type IFuncion_intContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext

	// IsFuncion_intContext differentiates from other interfaces.
	IsFuncion_intContext()
}

type Funcion_intContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_intContext() *Funcion_intContext {
	var p = new(Funcion_intContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_int
	return p
}

func InitEmptyFuncion_intContext(p *Funcion_intContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_int
}

func (*Funcion_intContext) IsFuncion_intContext() {}

func NewFuncion_intContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_intContext {
	var p = new(Funcion_intContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_int

	return p
}

func (s *Funcion_intContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_intContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Funcion_intContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_intContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_intContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_int(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_int() (localctx IFuncion_intContext) {
	localctx = NewFuncion_intContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, T_swiftParserRULE_funcion_int)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(619)
		p.Match(T_swiftParserT__39)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(620)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(621)
		p.expresion(0)
	}
	{
		p.SetState(622)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_floatContext is an interface to support dynamic dispatch.
type IFuncion_floatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext

	// IsFuncion_floatContext differentiates from other interfaces.
	IsFuncion_floatContext()
}

type Funcion_floatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_floatContext() *Funcion_floatContext {
	var p = new(Funcion_floatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_float
	return p
}

func InitEmptyFuncion_floatContext(p *Funcion_floatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_float
}

func (*Funcion_floatContext) IsFuncion_floatContext() {}

func NewFuncion_floatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_floatContext {
	var p = new(Funcion_floatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_float

	return p
}

func (s *Funcion_floatContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_floatContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Funcion_floatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_floatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_floatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_float(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_float() (localctx IFuncion_floatContext) {
	localctx = NewFuncion_floatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, T_swiftParserRULE_funcion_float)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(624)
		p.Match(T_swiftParserT__40)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(625)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(626)
		p.expresion(0)
	}
	{
		p.SetState(627)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncion_stringContext is an interface to support dynamic dispatch.
type IFuncion_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext

	// IsFuncion_stringContext differentiates from other interfaces.
	IsFuncion_stringContext()
}

type Funcion_stringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncion_stringContext() *Funcion_stringContext {
	var p = new(Funcion_stringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_string
	return p
}

func InitEmptyFuncion_stringContext(p *Funcion_stringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_funcion_string
}

func (*Funcion_stringContext) IsFuncion_stringContext() {}

func NewFuncion_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_stringContext {
	var p = new(Funcion_stringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_funcion_string

	return p
}

func (s *Funcion_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_stringContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Funcion_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_stringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitFuncion_string(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Funcion_string() (localctx IFuncion_stringContext) {
	localctx = NewFuncion_stringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, T_swiftParserRULE_funcion_string)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(629)
		p.Match(T_swiftParserT__38)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(630)
		p.Match(T_swiftParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(631)
		p.expresion(0)
	}
	{
		p.SetState(632)
		p.Match(T_swiftParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_asignacion
	return p
}

func InitEmptyAsignacionContext(p *AsignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_asignacion
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) CopyAll(ctx *AsignacionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Asignacion_normalContext struct {
	AsignacionContext
}

func NewAsignacion_normalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_normalContext {
	var p = new(Asignacion_normalContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_normalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_normalContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Asignacion_normalContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_normalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignacion_normal(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignacion_decrementoContext struct {
	AsignacionContext
}

func NewAsignacion_decrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_decrementoContext {
	var p = new(Asignacion_decrementoContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_decrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_decrementoContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Asignacion_decrementoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_decrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignacion_decremento(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignacion_vectorContext struct {
	AsignacionContext
}

func NewAsignacion_vectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_vectorContext {
	var p = new(Asignacion_vectorContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_vectorContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Asignacion_vectorContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Asignacion_vectorContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignacion_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignacion_incremento_vectorContext struct {
	AsignacionContext
}

func NewAsignacion_incremento_vectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_incremento_vectorContext {
	var p = new(Asignacion_incremento_vectorContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_incremento_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_incremento_vectorContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Asignacion_incremento_vectorContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Asignacion_incremento_vectorContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_incremento_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignacion_incremento_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignacion_incrementoContext struct {
	AsignacionContext
}

func NewAsignacion_incrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_incrementoContext {
	var p = new(Asignacion_incrementoContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_incrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_incrementoContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Asignacion_incrementoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_incrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignacion_incremento(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignacion_decremento_vectorContext struct {
	AsignacionContext
}

func NewAsignacion_decremento_vectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_decremento_vectorContext {
	var p = new(Asignacion_decremento_vectorContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *Asignacion_decremento_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_decremento_vectorContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Asignacion_decremento_vectorContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Asignacion_decremento_vectorContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Asignacion_decremento_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitAsignacion_decremento_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, T_swiftParserRULE_asignacion)
	p.SetState(664)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignacion_normalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(634)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(635)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(636)
			p.expresion(0)
		}

	case 2:
		localctx = NewAsignacion_incrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(637)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(638)
			p.Match(T_swiftParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(639)
			p.expresion(0)
		}

	case 3:
		localctx = NewAsignacion_decrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(640)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(641)
			p.Match(T_swiftParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(642)
			p.expresion(0)
		}

	case 4:
		localctx = NewAsignacion_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(643)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(644)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(645)
			p.expresion(0)
		}
		{
			p.SetState(646)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(647)
			p.Match(T_swiftParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(648)
			p.expresion(0)
		}

	case 5:
		localctx = NewAsignacion_incremento_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(650)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(651)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(652)
			p.expresion(0)
		}
		{
			p.SetState(653)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(654)
			p.Match(T_swiftParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(655)
			p.expresion(0)
		}

	case 6:
		localctx = NewAsignacion_decremento_vectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(657)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(658)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(659)
			p.expresion(0)
		}
		{
			p.SetState(660)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(661)
			p.Match(T_swiftParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(662)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) CopyAll(ctx *ExpresionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Expresion_llamadaContext struct {
	ExpresionContext
}

func NewExpresion_llamadaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_llamadaContext {
	var p = new(Expresion_llamadaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_llamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_llamadaContext) Llamadas_funciones() ILlamadas_funcionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadas_funcionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadas_funcionesContext)
}

func (s *Expresion_llamadaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_llamada(s)

	default:
		return t.VisitChildren(s)
	}
}

type Valor_primitivoContext struct {
	ExpresionContext
}

func NewValor_primitivoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Valor_primitivoContext {
	var p = new(Valor_primitivoContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Valor_primitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Valor_primitivoContext) Primitivos() IPrimitivosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivosContext)
}

func (s *Valor_primitivoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitValor_primitivo(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_atributosContext struct {
	ExpresionContext
}

func NewExpresion_atributosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_atributosContext {
	var p = new(Expresion_atributosContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_atributosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_atributosContext) Atributos() IAtributosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtributosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtributosContext)
}

func (s *Expresion_atributosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_atributos(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_relaContext struct {
	ExpresionContext
	op antlr.Token
}

func NewExpresion_relaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_relaContext {
	var p = new(Expresion_relaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_relaContext) GetOp() antlr.Token { return s.op }

func (s *Expresion_relaContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expresion_relaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_relaContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Expresion_relaContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_relaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_rela(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_aritContext struct {
	ExpresionContext
	op antlr.Token
}

func NewExpresion_aritContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_aritContext {
	var p = new(Expresion_aritContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_aritContext) GetOp() antlr.Token { return s.op }

func (s *Expresion_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expresion_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_aritContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Expresion_aritContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_aritContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_arit(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_compaContext struct {
	ExpresionContext
	op antlr.Token
}

func NewExpresion_compaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_compaContext {
	var p = new(Expresion_compaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_compaContext) GetOp() antlr.Token { return s.op }

func (s *Expresion_compaContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expresion_compaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_compaContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Expresion_compaContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_compaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_compa(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_struct_duplaContext struct {
	ExpresionContext
}

func NewExpresion_struct_duplaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_struct_duplaContext {
	var p = new(Expresion_struct_duplaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_struct_duplaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_struct_duplaContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Expresion_struct_duplaContext) L_duble() IL_dubleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_dubleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_dubleContext)
}

func (s *Expresion_struct_duplaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_struct_dupla(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_vectorContext struct {
	ExpresionContext
}

func NewExpresion_vectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_vectorContext {
	var p = new(Expresion_vectorContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_vectorContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Expresion_vectorContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_vectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_negaContext struct {
	ExpresionContext
	op antlr.Token
}

func NewExpresion_negaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_negaContext {
	var p = new(Expresion_negaContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_negaContext) GetOp() antlr.Token { return s.op }

func (s *Expresion_negaContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expresion_negaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_negaContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_negaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_nega(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_idContext struct {
	ExpresionContext
}

func NewExpresion_idContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_idContext {
	var p = new(Expresion_idContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_idContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Expresion_idContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_id(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_unarioContext struct {
	ExpresionContext
}

func NewExpresion_unarioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_unarioContext {
	var p = new(Expresion_unarioContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_unarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_unarioContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_unarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_unario(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_parenContext struct {
	ExpresionContext
}

func NewExpresion_parenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_parenContext {
	var p = new(Expresion_parenContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_parenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_parenContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_parenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_paren(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expresion_matrizContext struct {
	ExpresionContext
}

func NewExpresion_matrizContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expresion_matrizContext {
	var p = new(Expresion_matrizContext)

	InitEmptyExpresionContext(&p.ExpresionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpresionContext))

	return p
}

func (s *Expresion_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expresion_matrizContext) Identificador() antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, 0)
}

func (s *Expresion_matrizContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *Expresion_matrizContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Expresion_matrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitExpresion_matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *T_swiftParser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 112
	p.EnterRecursionRule(localctx, 112, T_swiftParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(701)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 69, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValor_primitivoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(667)
			p.Primitivos()
		}

	case 2:
		localctx = NewExpresion_atributosContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(668)
			p.Atributos()
		}

	case 3:
		localctx = NewExpresion_vectorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(669)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(670)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(671)
			p.expresion(0)
		}
		{
			p.SetState(672)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExpresion_matrizContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(674)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(675)
			p.Match(T_swiftParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(676)
			p.expresion(0)
		}
		{
			p.SetState(677)
			p.Match(T_swiftParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(682)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(678)
					p.Match(T_swiftParserT__19)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(679)
					p.expresion(0)
				}
				{
					p.SetState(680)
					p.Match(T_swiftParserT__20)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(684)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 68, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 5:
		localctx = NewExpresion_llamadaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(686)
			p.Llamadas_funciones()
		}

	case 6:
		localctx = NewExpresion_struct_duplaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(687)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(688)
			p.Match(T_swiftParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(689)
			p.L_duble()
		}
		{
			p.SetState(690)
			p.Match(T_swiftParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewExpresion_idContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(692)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewExpresion_parenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(693)
			p.Match(T_swiftParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(694)
			p.expresion(0)
		}
		{
			p.SetState(695)
			p.Match(T_swiftParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewExpresion_negaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(697)

			var _m = p.Match(T_swiftParserT__48)

			localctx.(*Expresion_negaContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(698)
			p.expresion(6)
		}

	case 10:
		localctx = NewExpresion_unarioContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(699)
			p.Match(T_swiftParserT__49)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(700)
			p.expresion(5)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(717)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 71, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(715)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 70, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresion_aritContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(703)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(704)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expresion_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15762598695796736) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expresion_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(705)
					p.expresion(5)
				}

			case 2:
				localctx = NewExpresion_aritContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(706)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(707)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expresion_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == T_swiftParserT__49 || _la == T_swiftParserT__53) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expresion_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(708)
					p.expresion(4)
				}

			case 3:
				localctx = NewExpresion_compaContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(709)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(710)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expresion_compaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2269814212194729984) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expresion_compaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(711)
					p.expresion(3)
				}

			case 4:
				localctx = NewExpresion_relaContext(p, NewExpresionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, T_swiftParserRULE_expresion)
				p.SetState(712)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(713)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expresion_relaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == T_swiftParserT__60 || _la == T_swiftParserT__61) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expresion_relaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(714)
					p.expresion(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(719)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 71, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IL_dubleContext is an interface to support dynamic dispatch.
type IL_dubleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentificador() []antlr.TerminalNode
	Identificador(i int) antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext

	// IsL_dubleContext differentiates from other interfaces.
	IsL_dubleContext()
}

type L_dubleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyL_dubleContext() *L_dubleContext {
	var p = new(L_dubleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_l_duble
	return p
}

func InitEmptyL_dubleContext(p *L_dubleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_l_duble
}

func (*L_dubleContext) IsL_dubleContext() {}

func NewL_dubleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_dubleContext {
	var p = new(L_dubleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_l_duble

	return p
}

func (s *L_dubleContext) GetParser() antlr.Parser { return s.parser }

func (s *L_dubleContext) AllIdentificador() []antlr.TerminalNode {
	return s.GetTokens(T_swiftParserIdentificador)
}

func (s *L_dubleContext) Identificador(i int) antlr.TerminalNode {
	return s.GetToken(T_swiftParserIdentificador, i)
}

func (s *L_dubleContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *L_dubleContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *L_dubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_dubleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *L_dubleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitL_duble(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) L_duble() (localctx IL_dubleContext) {
	localctx = NewL_dubleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, T_swiftParserRULE_l_duble)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(720)
		p.Match(T_swiftParserIdentificador)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(721)
		p.Match(T_swiftParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(722)
		p.expresion(0)
	}
	p.SetState(729)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == T_swiftParserT__5 {
		{
			p.SetState(723)
			p.Match(T_swiftParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(724)
			p.Match(T_swiftParserIdentificador)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(725)
			p.Match(T_swiftParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(726)
			p.expresion(0)
		}

		p.SetState(731)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitivosContext is an interface to support dynamic dispatch.
type IPrimitivosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimitivosContext differentiates from other interfaces.
	IsPrimitivosContext()
}

type PrimitivosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitivosContext() *PrimitivosContext {
	var p = new(PrimitivosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_primitivos
	return p
}

func InitEmptyPrimitivosContext(p *PrimitivosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = T_swiftParserRULE_primitivos
}

func (*PrimitivosContext) IsPrimitivosContext() {}

func NewPrimitivosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivosContext {
	var p = new(PrimitivosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = T_swiftParserRULE_primitivos

	return p
}

func (s *PrimitivosContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivosContext) CopyAll(ctx *PrimitivosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimitivosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Primitivo_boolContext struct {
	PrimitivosContext
}

func NewPrimitivo_boolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primitivo_boolContext {
	var p = new(Primitivo_boolContext)

	InitEmptyPrimitivosContext(&p.PrimitivosContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimitivosContext))

	return p
}

func (s *Primitivo_boolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_boolContext) Bool() antlr.TerminalNode {
	return s.GetToken(T_swiftParserBool, 0)
}

func (s *Primitivo_boolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitPrimitivo_bool(s)

	default:
		return t.VisitChildren(s)
	}
}

type Primitivo_stringContext struct {
	PrimitivosContext
}

func NewPrimitivo_stringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primitivo_stringContext {
	var p = new(Primitivo_stringContext)

	InitEmptyPrimitivosContext(&p.PrimitivosContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimitivosContext))

	return p
}

func (s *Primitivo_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_stringContext) String_() antlr.TerminalNode {
	return s.GetToken(T_swiftParserString_, 0)
}

func (s *Primitivo_stringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitPrimitivo_string(s)

	default:
		return t.VisitChildren(s)
	}
}

type Primitivo_floatContext struct {
	PrimitivosContext
}

func NewPrimitivo_floatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primitivo_floatContext {
	var p = new(Primitivo_floatContext)

	InitEmptyPrimitivosContext(&p.PrimitivosContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimitivosContext))

	return p
}

func (s *Primitivo_floatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_floatContext) Float() antlr.TerminalNode {
	return s.GetToken(T_swiftParserFloat, 0)
}

func (s *Primitivo_floatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitPrimitivo_float(s)

	default:
		return t.VisitChildren(s)
	}
}

type Primitivo_nuloContext struct {
	PrimitivosContext
}

func NewPrimitivo_nuloContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primitivo_nuloContext {
	var p = new(Primitivo_nuloContext)

	InitEmptyPrimitivosContext(&p.PrimitivosContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimitivosContext))

	return p
}

func (s *Primitivo_nuloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_nuloContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitPrimitivo_nulo(s)

	default:
		return t.VisitChildren(s)
	}
}

type Primitivo_caracterContext struct {
	PrimitivosContext
}

func NewPrimitivo_caracterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primitivo_caracterContext {
	var p = new(Primitivo_caracterContext)

	InitEmptyPrimitivosContext(&p.PrimitivosContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimitivosContext))

	return p
}

func (s *Primitivo_caracterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_caracterContext) Caracter() antlr.TerminalNode {
	return s.GetToken(T_swiftParserCaracter, 0)
}

func (s *Primitivo_caracterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitPrimitivo_caracter(s)

	default:
		return t.VisitChildren(s)
	}
}

type Primitivo_intContext struct {
	PrimitivosContext
}

func NewPrimitivo_intContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Primitivo_intContext {
	var p = new(Primitivo_intContext)

	InitEmptyPrimitivosContext(&p.PrimitivosContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimitivosContext))

	return p
}

func (s *Primitivo_intContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitivo_intContext) Int() antlr.TerminalNode {
	return s.GetToken(T_swiftParserInt, 0)
}

func (s *Primitivo_intContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case T_swiftVisitor:
		return t.VisitPrimitivo_int(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *T_swiftParser) Primitivos() (localctx IPrimitivosContext) {
	localctx = NewPrimitivosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, T_swiftParserRULE_primitivos)
	p.SetState(738)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case T_swiftParserCaracter:
		localctx = NewPrimitivo_caracterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(732)
			p.Match(T_swiftParserCaracter)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case T_swiftParserInt:
		localctx = NewPrimitivo_intContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(733)
			p.Match(T_swiftParserInt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case T_swiftParserFloat:
		localctx = NewPrimitivo_floatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(734)
			p.Match(T_swiftParserFloat)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case T_swiftParserString_:
		localctx = NewPrimitivo_stringContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(735)
			p.Match(T_swiftParserString_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case T_swiftParserBool:
		localctx = NewPrimitivo_boolContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(736)
			p.Match(T_swiftParserBool)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case T_swiftParserT__62:
		localctx = NewPrimitivo_nuloContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(737)
			p.Match(T_swiftParserT__62)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *T_swiftParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 56:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *T_swiftParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
