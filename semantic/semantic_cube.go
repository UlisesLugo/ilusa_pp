package semantic

import (
	"github.com/uliseslugo/ilusa_pp/types"
)

type Operation int // operation type int

// enumarate opeartion starting in 0
const (
	Add = iota
	Sub
	Mult
	Div
	Mod
	And
	Or
	Not
	LessT
	GreaterT
	Equal
)

// type Operators struct {
// 	left_op  types.CoreType
// 	right_op types.CoreType
// }

// const p1 = Operators{left_op: int, right_op: int}

// Create semantic cube matrix
// 11 operations
// 4 types for left operator
// 4 types for right operator
var semantic_cube = [][][]types.CoreType{
	{
		// Add
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Sub
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Mult
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Div
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Mod
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// And
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Or
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Not
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// LessT
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// GreaterT
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Equal
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
}

/**
	Translate
	Translate operation label to operation token
	return -> string
**/
func (o Operation) Translate() string {
	switch o {
	case Add:
		return "+"
	case Sub:
		return "-"
	case Mult:
		return "*"
	case Div:
		return "/"
	case Mod:
		return "%"
	case And:
		return "&&"
	case Or:
		return "||"
	case Not:
		return "!"
	case LessT:
		return "<"
	case GreaterT:
		return ">"
	case Equal:
		return "=="
	}
	return "noop" // return string no operation
}

func GetOperationLabel(operationStr string) Operation {
	switch operationStr {
	case "+":
		return Add
	case "-":
		return Sub
	case "*":
		return Mult
	case "/":
		return Div
	case "%":
		return Mod
	case "&&":
		return And
	case "||":
		return Or
	case "!":
		return Not
	case "<":
		return LessT
	case ">":
		return GreaterT
	case "==":
		return Equal
	}
	return -1 // invalid operation
}

// TODO: Implement Semantic Cube
// github.com/Loptt/lambdish-compiler/blob/master/sem/semanticcube.go
// type SemanticCube struct {
// 	operations map[string]types.CoreType
// }
