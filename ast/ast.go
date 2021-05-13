package ast

import (
	// go packages
	"errors"
	"fmt"

	// internal packages
	"github.com/uliseslugo/ilusa_pp/gocc/token"
	"github.com/uliseslugo/ilusa_pp/semantic"
	"github.com/uliseslugo/ilusa_pp/tables"
)

/*
	NewProgram
	@param id Attrib
	reads the program name id
	returns progam name as a literal
*/
func NewProgram(id Attrib) (*Program, error) {
	// cast id Attrib to token literal string
	nombre := string(id.(*token.Token).Lit)
	// cast id Attrib to token
	new_id, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("Program " + nombre + "is not valid")
	}
	return &Program{nombre, nil, new_id}, nil
}

/*
	NewClass
	@param id Attrib
	reads the class name id
	returns class name as a literal
*/
func NewClass(id Attrib) (string, error) {
	fmt.Println("In NewClass Func")
	// cast id Attrib to token literal string
	className := string(id.(*token.Token).Lit)
	fmt.Println("New class: ", className)
	return className, nil
}

/*
	NewFunction
	@param id Attrib
	reads the function name id and function entry from table
	returns function row in funciton directory
*/
func NewFunction(id Attrib) (*tables.FuncRow, error) {
	fmt.Println("In NewFunction Func")
	// cast id Attrib to string token literal
	idName := string(id.(*token.Token).Lit)
	// create new function row
	row := new(tables.FuncRow)
	// set id to funciton
	row.SetId(idName)
	fmt.Println("Function:", row.Id())
	return row, nil
}

/*
	NewVariable
	@param id Attrib
	@param dim1 Attrib
	@param dim2 Attrib
	returns variable entry
*/
func NewVariable(id, dim1, dim2 Attrib) (*tables.VarRow, error) {
	// cast id to token
	tok, tok_ok := id.(*token.Token)
	if !tok_ok {
		return nil, errors.New("Problem in casting id token")
	}
	// create variable row
	row := &tables.VarRow{} // TODO Constructor for VarRow
	// set values to varibale row
	row.SetId(string(tok.Lit))
	row.SetToken(tok)
	fmt.Println("New var:", row.Id())
	return row, nil // return row
}

/*
	NewExpression
	handles logic of creation of new expressions.
	@param exp1 Attrib
	@param exp2 Attrib
	reads expression
	returns Exp struct
*/
func NewExpression(exp1, exp2 Attrib) (*Exp, error) {
	fmt.Println("New expression created")
	new_term, _ := exp1.(Exp) // term non-terminal
	new_exp, _ := exp2.(Exp)  // exp non-terminal
	return &Exp{&new_term, &new_exp}, nil
}

/*
	NewOpExpression
	@param op Attrib
	@param exp Attrib

*/
func NewOpExpression(op, exp Attrib) (*Op_exp, error) {
	tok, t_ok := op.(*token.Token)
	if !t_ok {
		return nil, errors.New("Problem in casting operator")
	}
	new_op := semantic.Operation(tok.Lit)
	fmt.Println("Operator: ", string(tok.Lit))
	return &Op_exp{new_op, nil}, nil
}

/*
	NewIdConst
	@param id Attrib
	TODO: Missing Constant struct
*/
func NewIdConst(id Attrib) (string, error) {
	tok, ok := id.(*token.Token)
	if !ok {
		return "", errors.New("Problem in id constants")
	}
	fmt.Println("New id exp", string(tok.Lit))
	return string(tok.Lit), nil
}

/*
	NewIntConst
	@param value Attrib

*/
func NewIntConst(value Attrib) (int, error) {
	num, ok := value.(*token.Token).Int32Value()
	if ok != nil {
		return -1, errors.New("Problem in integer constants")
	}
	return int(num), nil
}

/*
	NewFloatConst
	@param value Attrib

*/
func NewFloatConst(value Attrib) (float64, error) {
	num, ok := value.(*token.Token).Float64Value()
	if ok != nil {
		return -1, errors.New("Problem in float constants")
	}
	return float64(num), nil
}
