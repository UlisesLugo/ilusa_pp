package semantic

import "github.com/uliseslugo/ilusa_pp/types"

type Operation string // operation type int

// enumarate operations by signs
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
)

// Crete semantic cube matrix
// 11 opeations
// 4 tpes for left operator
// 4 types for right operator
type SemanticCube struct {
	operations map[string]types.CoreType
}

// NewSemanticCube creates a new semantic cube struct
func NewSemanticCube() *SemanticCube {
	return &SemanticCube{
		map[string]types.CoreType{
			//Arithmetical Operators
			"+00": types.Integer, // 00 int con int
			"+11": types.Float,   // 11 float con float
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
			"==00": types.Integer, // 00 int con int
			"==11": types.Integer, // 11 float con float
			"==22": types.Integer, // 22 char con char
			"!=00": types.Integer,
			"!=11": types.Integer, // 11 float con float
			"!=22": types.Integer, // 22 char con char
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
