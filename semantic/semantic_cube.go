// Semantic Cube
package semantic

import "github.com/uliseslugo/ilusa_pp/types"

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
)
//TODO Hierarchy dictionary (operation, num)

// Crete semantic cube map
// key - string of operation
// value - type of return of operation
type SemanticCube struct {
	operations map[string]types.CoreType
}

// NewSemanticCube
// creates a new semantic cube map
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
