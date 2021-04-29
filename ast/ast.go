package ast

import (
	"fmt"

	"github.com/uliseslugo/ilusa_pp/gocc/token"
	"github.com/uliseslugo/ilusa_pp/tables"
)

type Attrib interface{}

// Reads the program name id and returns it as a literal
func NewProgram(id Attrib) (string, error) {
	fmt.Println("In NewProgram Func")
	return string(id.(*token.Token).Lit), nil
}

func NewClass(id Attrib) (string, error) {
	fmt.Println("In NewClass Func")
	className := string(id.(*token.Token).Lit)
	fmt.Println(className)
	return className, nil
}

func NewVariable(id Attrib) (*tables.VarRow, error) {
	fmt.Println("In NewVariable Func")
	idName := string(id.(*token.Token).Lit)
	row := new(tables.VarRow)
	row.SetId(idName)
	fmt.Println("Variable:", row.Id())
	return row, nil
}
