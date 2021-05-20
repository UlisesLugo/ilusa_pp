// Semantic Cube
package semantic

import (
	// internal packages
	"github.com/uliseslugo/ilusa_pp/types"
)

type Operation string // operation type string

// Assign string token to operation label
const (
	Add      = "+"
	Sub      = "-"
	Mult     = "*"
	Div      = "/"
	And      = "&&"
	Or       = "||"
	Not      = "!"
	LessT    = "<"
	GreaterT = ">"
	Equal    = "=="
	NotEqual = "!="
	Assign   = "="
	GOTO     = "GOTO"
)

type OperationsDict struct {
	Operators map[int]string // Reverse map
}

// Assign string token to operation label
func NewOperationsDict() *OperationsDict {
	return &OperationsDict{
		map[int]string{
			0:  "!",
			1:  "=",
			2:  "!=",
			3:  "&&",
			4:  "||",
			5:  "==",
			6:  "<",
			7:  ">",
			8:  "+",
			9:  "-",
			10: "*",
			11: "/",
			12: "GOTO",
		},
	}
}

/*
	Operators Dictionary
	key: int value
	value: string operator
*/
type HierarchyDict struct {
	Op_hierarchy map[string]int
}

func NewHierarchyDict() *HierarchyDict {
	// Hierarchy levels:
	// 	indexes 0-2: 0
	// 	indexes 3-5: 1,
	// 	indexes 6-7: 2,
	// 	indexes 8-9: 3,
	// 	indexes 10-11: 4
	return &HierarchyDict{
		map[string]int{
			// super low hierarchy (right association)
			"!":  0,
			"=":  1,
			"!=": 2,
			// logic operators very low hierarchy
			"&&": 3,
			"||": 4,
			"==": 5,
			// relations operators low hierarchy
			"<": 6,
			">": 7,
			// add or sub operators medium hierarchy
			"+": 8,
			"-": 9,
			// mult ord div high hierarchy
			"*": 10,
			"/": 11,
		},
	}
}

/*
	Semantic Cube struct
	operations: map of string key to types value
*/
type SemanticCube struct {
	operations map[string]types.CoreType
}

/*
	NewSemanticCube
	Returns new structure of semantic cube
*/
func NewSemanticCube() *SemanticCube {
	// Keys notation:
	// 00 int con int
	// 11 float con float
	// 22 char con char

	return &SemanticCube{
		map[string]types.CoreType{
			//Arithmetical Operators
			"+00": types.Integer, // ints
			"+11": types.Float,   // floats
			"-00": types.Integer,
			"-11": types.Float,
			"*00": types.Integer,
			"*11": types.Float,
			"/00": types.Integer,
			"/11": types.Float,
			//Relational Operators
			"<00":  types.Integer,
			"<11":  types.Integer,
			">00":  types.Integer,
			">11":  types.Integer,
			"==00": types.Integer,
			"==11": types.Integer,
			"==22": types.Integer, // chars
			"!=00": types.Integer,
			"!=11": types.Integer,
			"!=22": types.Integer,
			//Logical Operators
			"&&00": types.Integer,
			"&&11": types.Integer,
			"&&33": types.Null,
			"||00": types.Integer,
			"||11": types.Integer,
			"||33": types.Integer,
			"!0":   types.Integer,
			// Assignation
			"=00": types.Integer,
			"=11": types.Float,
			"=22": types.Char,
			"=33": types.Null,
		},
	}
}
