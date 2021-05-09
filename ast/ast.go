package ast

import (
	"errors"
	"fmt"

	"github.com/uliseslugo/ilusa_pp/gocc/token"
	"github.com/uliseslugo/ilusa_pp/semantic"
	"github.com/uliseslugo/ilusa_pp/stacks"
	"github.com/uliseslugo/ilusa_pp/tables"
)

type Attrib interface{}

type Program struct {
	nombre string
	operaciones []cuadruplo 
	id *token.Token
}

func (p *Program) String() string {
	return p.nombre
}

type cuadruplo struct {
	operacion semantic.Operation
	var1 string 
	var2 string
	res string
}

// Reads the program name id and returns it as a literal
func NewProgram(id Attrib) (*Program, error) {
	nombre := string(id.(*token.Token).Lit)
	new_id,ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("Program " + nombre + "is not valid")
	}
	fmt.Println("Program name", nombre)
	return &Program{nombre,nil,new_id}, nil
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

func NewFunction(id Attrib) (*tables.FuncRow, error) {
	fmt.Println("In NewFunction Func")
	idName := string(id.(*token.Token).Lit)
	row := new(tables.FuncRow)
	row.SetId(idName)
	fmt.Println("Function:", row.Id())
	return row, nil
}

type s_exp struct{
	operadores stacks.Stack
	operandos stacks.Stack
	tipos stacks.Stack
	saltos stacks.Stack
}

// TODO () Add newExpression that checks stacks and appends cuadruplos
// TODO? () Move structs into astx for cleanliness?