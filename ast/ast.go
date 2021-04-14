package ast

import (
	"github.com/uliseslugo/ilusa_pp/token"
)

type Attrib interface {}

// Reads the program name id and returns it as a literal
func NewProgram(id Attrib) (string, error){
	return string(id.(*token.Token).Lit), nil
}