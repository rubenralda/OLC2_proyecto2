package valor

type Value struct {
	Value          string        //temp1, 5, 6, "hola", posicion_heap_struct
	Tipo_dimension DIMENSION     //variable, vector, matriz
	Type           TipoExpresion //INT, FLOAT, STRING, BOOL, CHAR, NULL
	Tipo_struct    string        //nombre tipo del struct (ID)
	Is_intancia    bool          //false
	// expresion bool
	TrueLabel  []interface{}
	FalseLabel []interface{}

	//verificar y borrar algunos
	//			|
	//			V
	Referencia bool
}

type TipoExpresion int

const (
	NULL    TipoExpresion = iota //0
	FLOAT                        //1
	STRING                       //2
	BOOLEAN                      //3
	CHAR                         //4
	INTEGER                      //5
)

type DIMENSION int

const (
	DIMENSION0 DIMENSION = iota //
	DIMENSION1
	DIMENSIONN
)
