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
	nombre      string
	operaciones []Cuadruplo
	id          *token.Token
}

func (p *Program) String() string {
	return p.nombre
}

type Cuadruplo struct {
	operacion semantic.Operation
	var1      string
	var2      string
	res       string
}

// Reads the program name id and returns it as a literal
func NewProgram(id Attrib) (*Program, error) {
	nombre := string(id.(*token.Token).Lit)
	new_id, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("Program " + nombre + "is not valid")
	}
	fmt.Println("Program name", nombre)
	return &Program{nombre, nil, new_id}, nil
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

/*
	Super Expression Struct
	operadores: Stack
	operandos: Stack
	tipos: Stack
	saltos: Stack
*/
type S_exp struct {
	operadores stacks.Stack
	operandos  stacks.Stack
	tipos      stacks.Stack
	saltos     stacks.Stack
}

func (s_exp *S_exp) IsEmpty() bool {
	fmt.Println("S_exp is empty =", s_exp.operadores.Empty())
	return (s_exp.operadores == nil &&
		s_exp.operandos == nil &&
		s_exp.tipos == nil &&
		s_exp.saltos == nil)
}

/*
	Hyper Expression Struct

	super_exp1: left super expression
	logical_operator: operator
	super_exp2:
	tok
*/
type H_exp struct {
	super_exp1       S_exp
	Logical_operator semantic.Operation
	super_exp2       S_exp
	tok              *token.Token
}

func (h_exp *H_exp) isEmpty() bool {
	return (h_exp.super_exp1.IsEmpty() &&
		h_exp.super_exp2.IsEmpty())
}

// TODO () Add newExpression that checks stacks and appends cuadruplos
func NewHyperExpression(super_exp, hyper_exp Attrib) (*S_exp, error) {
	h_exp, _ := hyper_exp.(H_exp)
	fmt.Println("aqui")
	fmt.Println(h_exp.Logical_operator)
	// _, s_ok := super_exp.(S_exp)
	// if !h_ok {
	// 	return nil, errors.New("Problem in casting hyper_expression")
	// }
	// if !s_ok {
	// 	return nil, errors.New("Problem in casting super_expression")
	// }
	if !h_exp.isEmpty() {
		fmt.Println("Hyper exp not empty")
	}

	fmt.Println("New hyper_expression created")
	return &S_exp{nil, nil, nil, nil}, nil
}

func NewSuperExpression(log_op, super_exp Attrib) (*S_exp, error) {
	// logical_op, _ := log_op.(semantic.Operation)
	// _, s_ok := super_exp.(S_exp)
	// if !log_ok {
	// 	return nil, errors.New("Problem in casting logical operator")
	// }
	// if !s_ok {
	// 	return nil, errors.New("Problem in casting super_expression")
	// }
	s := make(stacks.Stack, 0)
	s = s.Push(2)
	fmt.Println("Stack added")
	return &S_exp{s, nil, nil, nil}, nil
}

// TODO? () Move structs into astx for cleanliness?
