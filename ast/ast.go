package ast

import (
	"fmt"

	"github.com/uliseslugo/ilusa_pp/token"
)

type Attrib interface {}

// Reads the program name id and returns it as a literal
func NewProgram(id Attrib) (string, error){
	fmt.Println("In NewProgram Func");
	return string(id.(*token.Token).Lit), nil
}

func NewClass(id Attrib) (string, error){
	fmt.Println("In NewClass Func");
	className := string(id.(*token.Token).Lit)
	fmt.Println(className)
	return className, nil
}