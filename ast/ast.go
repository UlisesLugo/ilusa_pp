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

func NewVariable(id, dim1, dim2 Attrib) (*tables.VarRow, error) {
	fmt.Println("New variable beginning");
	tok, tok_ok := id.(*token.Token)
	curr_id := string(tok.Lit)
	if !tok_ok {
		return nil, errors.New("Problem in casting id token")
	}
	row := &tables.VarRow{}
	row.SetId(curr_id)
	row.SetToken(tok)
	fmt.Println("New var:", curr_id)
	return row, nil
}

func NewIdConst(id Attrib) (string, error) {
	fmt.Println("In New Id const")
	tok, ok := id.(*token.Token)
	if !ok {
		return "", errors.New("Problem in id constants")
	}
	return string(tok.Lit), nil
}

func NewIntConst(value Attrib) (int, error){
	fmt.Println("In New Int Const")
	num, ok := value.(*token.Token).Int32Value()
	if ok != nil{
		return -1, errors.New("Problem in integer constants")
	}
	return int(num), nil
}

func NewFloatConst(value Attrib) (float64, error){
	fmt.Println("In New Float Const")
	num, ok := value.(*token.Token).Float64Value()
	if ok != nil{
		return -1, errors.New("Problem in float constants")
	}
	return float64(num), nil
}
// TODO? () Move structs into astx for cleanliness?
