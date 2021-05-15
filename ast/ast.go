package ast

import (
	// go packages
	"errors"
	"fmt"

	// internal packages
	"github.com/uliseslugo/ilusa_pp/gocc/token"
	"github.com/uliseslugo/ilusa_pp/memory"
	"github.com/uliseslugo/ilusa_pp/quadruples"
	"github.com/uliseslugo/ilusa_pp/semantic"
	"github.com/uliseslugo/ilusa_pp/tables"
	"github.com/uliseslugo/ilusa_pp/types"
)

var globalIntCount int
var globalFloatCount int
var globalIdCount int
var globalCurrQuads []quadruples.Cuadruplo

func init() {
	// globalSemanticCube := semantic.NewSemanticCube()
	// globalStackOperators
	// globalStackOperands
	// globalStackTypes
	// globalStackJumps
	globalIntCount, globalFloatCount, globalIdCount = 0, 0, 0
	globalCurrQuads = make([]quadruples.Cuadruplo, 0)
	globalCurrQuads := append(globalCurrQuads, quadruples.Cuadruplo{semantic.Operation("GOTO"), "", "", "main"}) // TODO change main to memory address
	fmt.Println("Added main quad:", globalCurrQuads[0])
	fmt.Println("In init")
}

/*
	NewProgram
	@param id Attrib
	reads the program name id
	returns progam name as a literal
*/
func NewProgram(id Attrib) (*Program, error) {
	fmt.Println("In NEW PROGRAM")
	// cast id Attrib to token literal string
	nombre := string(id.(*token.Token).Lit)
	// cast id Attrib to token
	new_id, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("Program " + nombre + "is not valid")
	}
	return &Program{nombre, globalCurrQuads, new_id}, nil
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
	new_term, _ := exp1.(*Exp) // term non-terminal
	new_const, ok := exp1.(*Constant)
	if ok {
		fmt.Println("New const id", new_const.Value())
		fmt.Println("\t with address", new_const.Address())
		new_op_exp, ok := exp2.(*Op_exp)
		if ok {
			return &Exp{nil, new_op_exp, new_const}, nil

		}
		return &Exp{nil, nil, new_const}, nil

	}
	return &Exp{new_term, nil, nil}, nil
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
	new_exp, _ := exp.(*Exp)
	new_op := semantic.Operation(tok.Lit)
	fmt.Println("Operator: ", string(tok.Lit))
	return &Op_exp{new_op, new_exp}, nil
}

/*
	NewIdConst
	@param id Attrib
*/
func NewIdConst(id Attrib) (*Constant, error) {
	val, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("Problem in id constants")
	}
	// calculate current address occuppied in context
	current_address := globalIdCount + memory.IdOffset
	globalIdCount++ // assign next available address
	return &Constant{string(val.Lit), val, types.Char, memory.Address(current_address)}, nil
}

/*
	NewIntConst
	@param value Attrib

*/
func NewIntConst(value Attrib) (*Constant, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errors.New("Problem in integer constants")
	}
	// calculate current address occuppied in context
	current_address := globalIdCount + memory.IntOffset
	globalIntCount++ // assign next available address
	return &Constant{string(val.Lit), val, types.Integer, memory.Address(current_address)}, nil
}

/*
	NewFloatConst
	@param value Attrib

*/
func NewFloatConst(value Attrib) (*Constant, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errors.New("Problem in float constants")
	}
	// calculate current address occuppied in context
	current_address := globalIdCount + memory.FloatOffset
	globalFloatCount++ // assign next available address
	return &Constant{string(val.Lit), val, types.Float, memory.Address(current_address)}, nil
}
