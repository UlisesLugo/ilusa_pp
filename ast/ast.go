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
	"github.com/uliseslugo/ilusa_pp/stacks"
	"github.com/uliseslugo/ilusa_pp/tables"
	"github.com/uliseslugo/ilusa_pp/types"
)

var globalIntCount int
var globalFloatCount int
var globalIdCount int
var globalStackOperators stacks.Stack
var globalStackOperands stacks.Stack
var globalStackTypes stacks.Stack
var globalStackJumps stacks.Stack
var globalCurrQuads []quadruples.Cuadruplo
var globalOperatorsDict semantic.OperatorsDict

func init() {
	// globalSemanticCube := semantic.NewSemanticCube()
	// globalOperatorsDict := semantic.NewOperatorsDict()
	// globalStackOperands := make(stacks.Stack, 0)
	// globalStackOperantors := make(stacks.Stack, 0)
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
	new_exp, expr_ok := exp1.(*Exp) // term non-terminal
	new_op_exp, _ := exp1.(*Op_exp)
	if expr_ok {
		if new_exp.const1_ != nil {
			fmt.Println("Reading const in new exp:", new_exp.const1_.Value())
		}
		fmt.Println("Reading exp:", new_exp)
		// fmt.Println("Full Expression encountered")
		// 	new_const1 := new_term1.const1_
		// 	// new_op := new_term.operator
		// 	fmt.Println("New const 1 id", new_const1.Value())
		// 	fmt.Println("\t with address", new_const1.Address())
		// 	// fmt.Println("Operator: ", string(new_op))
		// 	new_term2, expr_ok2 := exp2.(*Exp)
		// 	if expr_ok2 {
		// 		new_const2 := new_term2.const1_
		// 		fmt.Println("New const 2 id", new_const2.Value())
		// 		fmt.Println("\t with address", new_const1.Address())
		// 		return &Exp{nil, nil, new_const1, semantic.Operation(""), new_const2}, nil
		// 	}

	}

	// a + b < c - d
	new_const1, ok := exp1.(*Constant)
	if ok {
		// fmt.Println("Reading constant 1:", new_const1.Value())
		return &Exp{nil, nil, new_const1, new_op_exp}, nil
	}
	new_const2, ok := exp2.(*Constant)
	if ok {
		fmt.Println("Reading constant 2:", new_const2.Value())
		return &Exp{nil, nil, nil, new_op_exp}, nil
	}

	return &Exp{new_exp, nil, nil, nil}, nil
}

/*
	NewOpExpression
	@param op Attrib
	@param exp Attrib

*/
func NewOperation(op, exp Attrib) (*Op_exp, error) {
	tok, t_ok := op.(*token.Token)
	if !t_ok {
		return &Op_exp{semantic.Operation(""), nil}, errors.New("problem in casting operator")
	}
	new_op := semantic.Operation(tok.Lit)
	fmt.Println("Operator: ", string(tok.Lit))
	new_exp, ok := exp.(*Exp)
	if ok {
		new_const := new_exp.exp1.const1_
		fmt.Println("\tConstant in op:", new_exp.exp1.const1_.value)
		return &Op_exp{new_op, new_const}, nil
	}
	// curr_hierarchy := globalOperatorsDict[string(tok.Lit)]
	// globalStackOperators := globalStackOperators.Push()
	return &Op_exp{new_op, nil}, nil
}

/*
	NewIdConst
	@param id Attrib
*/
func NewIdConst(id Attrib) (*Constant, error) {
	val, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("problem in id constants")
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
		return nil, errors.New("problem in integer constants")
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
		return nil, errors.New("problem in float constants")
	}
	// calculate current address occuppied in context
	current_address := globalIdCount + memory.FloatOffset
	globalFloatCount++ // assign next available address
	return &Constant{string(val.Lit), val, types.Float, memory.Address(current_address)}, nil
}
