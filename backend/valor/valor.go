package valor

type Value struct {
	Value      string
	IsTemp     bool
	Type       TipoExpresion
	TrueLabel  []interface{}
	FalseLabel []interface{}
	OutLabel   []interface{}
	IntValue   int
}

type TipoExpresion int

const (
	INTEGER TipoExpresion = iota //0
	FLOAT                        //1
	STRING                       //2
	BOOLEAN                      //3
	CHAR                         //4
	NULL                         //5
	ARRAY                        //6
	STRUCT                       //7
)
