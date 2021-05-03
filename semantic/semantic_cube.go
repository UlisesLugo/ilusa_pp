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

// Create semantic cube matrix
// 11 operations
// 4 types for left operator
// 4 types for right operator
semantic_cube[11][3][3] = [][][]types.CoreType{
	{
		// Add
		{types.CoreType.Integer}, // Integer w Integer
		{types.CoreType.Float}, // Integer w Float
		{types.CoreType.Null}, // Integer w Char
	},
	{
		 // Sub
		{types.CoreType.Integer}, // Integer w Integer
		{types.CoreType.Float}, // Integer w Float
		{types.CoreType.Null}, // Integer w Char
	},
	{
		// Mult
		{types.CoreType.Integer}, // Integer w Integer
		{types.CoreType.Float}, // Integer w Float
		{types.CoreType.Null}, // Integer w Char
	},
	{
		// Div
		{types.CoreType.Integer}, // Integer w Integer
		{types.CoreType.Float}, // Integer w Float
		{types.CoreType.Null}, // Integer w Char
	},
	{
		// Mod
		{types.CoreType.Integer}, // Integer w Integer
		{types.CoreType.Float}, // Integer w Float
		{types.CoreType.Null}, // Integer w Char
	},
	{
		// And
		{types.CoreType.Integer}, // Integer w Integer
		{types.CoreType.Float}, // Integer w Float
		{types.CoreType.Null}, // Integer w Char
	},
	{
		// Or
		{types.CoreType.Integer}, // Integer w Integer
		{types.CoreType.Float}, // Integer w Float
		{types.CoreType.Null}, // Integer w Char
	},
	{
		// Not
		{types.CoreType.Integer}, // Integer w Integer
		{types.CoreType.Float}, // Integer w Float
		{types.CoreType.Null}, // Integer w Char
	},
	{
 		// LessT
		{types.CoreType.Integer}, // Integer w Integer
		{types.CoreType.Float}, // Integer w Float
		{types.CoreType.Null}, // Integer w Char
	},
	{
		// GreaterT
		{types.CoreType.Integer}, // Integer w Integer
		{types.CoreType.Float}, // Integer w Float
		{types.CoreType.Null}, // Integer w Char
	},
	{
		// Equal
		{types.CoreType.Integer}, // Integer w Integer
		{types.CoreType.Float}, // Integer w Float
		{types.CoreType.Null}, // Integer w Char
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
