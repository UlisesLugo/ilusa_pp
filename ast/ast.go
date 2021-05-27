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

var vmemory *memory.VirtualMemory

var globalStackOperators stacks.Stack
var globalStackOperands stacks.Stack
var globalStackTypes stacks.Stack
var globalStackJumps stacks.Stack
var globalCurrQuads []quadruples.Cuadruplo
var globalFuncTable *tables.FuncTable
var globalCurrentScope int

var quadsCounter int

func init() {
	// globalSemanticCube := semantic.NewSemanticCube()
	vmemory = memory.NewVirtualMemory()
	globalStackOperands := make(stacks.Stack, 0)
	globalStackOperators := make(stacks.Stack, 0)
	globalCurrQuads = make([]quadruples.Cuadruplo, 0) // TODO change main to memory address
	globalStackTypes = make(stacks.Stack, 0)

	fmt.Println("Defining globals")
	fmt.Println("\tOperatorsStack:", globalStackOperators)
	fmt.Println("\tOperandsStack:", globalStackOperands)
	fmt.Println("\tQuad:", globalCurrQuads)
	fmt.Println("\tVirtual Memory:", vmemory)
}

/*
	NewProgram
	@param id Attrib
	reads the program name id
	returns progam name as a literal
*/
func NewProgram(id Attrib) (*Program, error) {
	fmt.Println("In NEW PROGRAM", globalStackOperators, globalStackOperands, globalFuncTable)
	// cast id Attrib to token literal string
	nombre := string(id.(*token.Token).Lit)
	// cast id Attrib to token
	new_id, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("Program " + nombre + "is not valid")
	}
	// Prepend main quad
	main_quad := quadruples.Cuadruplo{"GOTO", "-1", "-1", "main"}
	globalCurrQuads = append([]quadruples.Cuadruplo{main_quad}, globalCurrQuads...)
	quadsCounter++
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
	if globalFuncTable == nil {
		globalFuncTable = tables.NewFuncTable()
	}
	tok, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("problem reading function")
	}
	// cast id Attrib to string token literal
	idName := string(tok.Lit)
	row := new(tables.FuncRow)
	row.SetId(idName)
	globalFuncTable.AddRow(row)
	// TODO Add type checking and check to repeated func
	fmt.Println("Function:", row.Id())
	return row, nil
}

/*
	NewTypeVariables
	@param id Attrib
	@param dim1 Attrib
	@param dim2 Attrib
	returns variable entry
*/
func NewTypeVariables(typed_var, var_list Attrib) (map[string]*tables.VarRow, error) {
	new_typed_var, ok := typed_var.([]*tables.VarRow)
	fmt.Println("In new typed var")
	curr_map := make(map[string]*tables.VarRow)
	if !ok || len(new_typed_var) != 1 {
		return nil, errors.New("problem in casting typed variable")
	}
	curr_map[new_typed_var[0].Id()] = new_typed_var[0]
	if var_list != nil {
		new_var_list := var_list.([]*tables.VarRow)
		fmt.Println("\t Len of list",len(new_var_list))
	}
	return curr_map, nil
}

/*
	NewVariable
	@param id Attrib
	@param dim1 Attrib
	@param dim2 Attrib
	returns variable entry
*/
func NewVariable(curr_type, id, dim1, dim2, rows Attrib) ([]*tables.VarRow, error) {
	// cast id to token
	tok, tok_ok := id.(*token.Token)
	new_rows, _ := rows.([]*tables.VarRow)
	fmt.Println("In New variable", dim1, dim2)
	if !tok_ok {
		return nil, errors.New("Problem in casting id token")
	}
	new_dim1, _ := dim1.(int)
	new_dim2, _ := dim1.(int)
	// create variable row
	row := &tables.VarRow{} // TODO Constructor for VarRow
	if curr_type != nil {
		row.SetType(curr_type.(types.CoreType))
	}
	row.SetDim1(new_dim1)
	row.SetDim2(new_dim2)
	// set values to varibale row
	row.SetId(string(tok.Lit))
	row.SetToken(tok)
	fmt.Println("New var:", row.Id())
	return append([]*tables.VarRow{row} ,new_rows...), nil
}

/*
	NewAssignation
	@param id Attrib
	@param dim1 Attrib
	@param dim2 Attrib
	returns variable entry
*/
func NewAssignation(id, exp Attrib) (int, error) {
	// cast id to token
	tok, tok_ok := id.(*token.Token)
	if !tok_ok {
		return -1, errors.New("Problem in casting id token")
	}

	globalStackOperators = globalStackOperators.Push(semantic.Assign)
	fmt.Println("Id: ", string(tok.Lit))
	current_address, addr_ok := vmemory.NextGlobal(types.Integer) // TODO Check Types

	if addr_ok != nil {
		return -1, addr_ok
	}

	globalStackOperands = globalStackOperands.Push(fmt.Sprintf("%v", current_address))
	// Add to scope &Constant{string(val.Lit), val, types.Char, memory.Address(current_address)}, nil
	_, exp_ok := exp.(*Exp)
	if exp_ok {
		createUnaryQuadruple(semantic.Equal)
	}
	return 1, nil // return row
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
	new_exp1, exp1_ok := exp1.(*Exp) // term non-terminal
	new_const, _ := exp1.(*Constant)
	if exp1_ok {
		if new_exp1.const_ != nil {
			new_const = new_exp1.const_
		}
		if new_exp1.op_exp_ != nil {
			createBinaryQuadruple(new_exp1.op_exp_.operation)
		}

	}
	new_exp2, exp2_ok := exp2.(*Op_exp)
	if exp2_ok {
		createBinaryQuadruple(new_exp2.operation)
	}
	return &Exp{new_exp1, new_exp2, new_const}, nil
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
	globalStackOperators = globalStackOperators.Push(string(new_op))
	new_const, _ := exp.(*Constant)
	return &Op_exp{new_op, new_const}, nil
}

func createBinaryQuadruple(new_op semantic.Operation) {
	operatorsDict := semantic.NewHierarchyDict() // operators hierarchy table
	// operatorsKey := semantic.NewOperatorKey() // operators table with keys

	level_id := operatorsDict.Op_hierarchy[string(new_op)] // get hierarchy level of operator level

	top, ok := globalStackOperators.Top()        // get top operator
	top_level := operatorsDict.Op_hierarchy[top] // get hierarchy level of top operator

	for ok && top_level <= level_id { // top level has higher hierarchy level
		// pop top operator
		globalStackOperators, _ = globalStackOperators.Pop()
		// get operand 2
		curr_top2, _ := globalStackOperands.Top()
		// pop operand 2
		globalStackOperands, _ = globalStackOperands.Pop()
		// get operand 1
		curr_top1, _ := globalStackOperands.Top()
		// pop operand 1
		globalStackOperands, _ = globalStackOperands.Pop()

		// generate quad
		current_address, _ := vmemory.NextGlobalTemp(types.Integer) // TODO Check Types (Validate type with semantic cube)
		curr_quad := quadruples.Cuadruplo{top, curr_top1, curr_top2, fmt.Sprint(current_address)}
		globalStackOperands = globalStackOperands.Push(fmt.Sprint(current_address))
		globalCurrQuads = append(globalCurrQuads, curr_quad)

		top, ok = globalStackOperators.Top()
		top_level = operatorsDict.Op_hierarchy[top]
	}
}

func createUnaryQuadruple(new_op semantic.Operation) {
	fmt.Println("\tCreating unary cuad")
	operatorsDict := semantic.NewHierarchyDict() // operators hierarchy table
	// operatorsKey := semantic.NewOperatorKey() // operators table with keys

	level_id := operatorsDict.Op_hierarchy[string(new_op)] // get hierarchy level of operator level

	top, ok := globalStackOperators.Top()        // get top operator
	top_level := operatorsDict.Op_hierarchy[top] // get hierarchy level of top operator

	fmt.Println("\tTop", globalStackOperands)
	for ok && top_level <= level_id { // top level has higher hierarchy level
		// pop top operator
		globalStackOperators, _ = globalStackOperators.Pop()
		// get operand 1
		curr_top1, _ := globalStackOperands.Top()
		// pop operand 1
		globalStackOperands, _ = globalStackOperands.Pop()
		// get operand 2
		curr_top2, _ := globalStackOperands.Top()
		// pop operand 2
		globalStackOperands, _ = globalStackOperands.Pop()

		// TODO (Add type validation)

		// generate quad
		curr_quad := quadruples.Cuadruplo{top, curr_top2, "-1", curr_top1}
		globalCurrQuads = append(globalCurrQuads, curr_quad)

		top, ok = globalStackOperators.Top()
		top_level = operatorsDict.Op_hierarchy[top]
	}
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
	current_address, _ := vmemory.NextGlobalTemp(types.Ids) // TODO Check Types (Validate type with semantic cube)
	// calculate current address occuppied in context
	globalStackOperands = globalStackOperands.Push(fmt.Sprint(current_address))
	return &Constant{string(val.Lit), val, types.Char, current_address}, nil
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
	current_address, _ := vmemory.NextConst(types.Integer) // TODO Check Types (Validate type with semantic cube)
	fmt.Println("id=", string(val.Lit), " addr=", current_address)
	globalStackOperands = globalStackOperands.Push(fmt.Sprint(current_address))
	return &Constant{string(val.Lit), val, types.Integer, current_address}, nil
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
	current_address, _ := vmemory.NextGlobalTemp(types.Float) // TODO Check Types (Validate type with semantic cube)
	fmt.Println("id=", val.Lit, " addr=", current_address)
	globalStackOperands = globalStackOperands.Push(fmt.Sprint(current_address))
	return &Constant{string(val.Lit), val, types.Float, current_address}, nil
}

func FinishInput(idList Attrib) (int, error) {
	id_list, ok := idList.([]*Constant)
	if !ok {
		return -1, errors.New("problem casting constant in input")
	}
	for _, id := range id_list {
		curr_quad := quadruples.Cuadruplo{"READ", "-1", "-1", fmt.Sprint(id.Address())}
		globalCurrQuads = append(globalCurrQuads, curr_quad)
	}
	return 1, nil
}

func NewInput(id, idList Attrib) ([]*Constant, error) {
	new_id, ok := id.(*Constant)
	id_list, _ := idList.([]*Constant)
	if !ok {
		return nil, errors.New("problem casting constant in input")
	}
	return append([]*Constant{new_id}, id_list...), nil // Prepend (Add first)
}

/*
	GetIdDimConst
	@param id Attrib
*/
func GetIdDimConst(id, dim1, dim2 Attrib) (*Constant, error) {
	fmt.Println("In IdDim Const")
	val, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("problem in id dim constants")
	}
	// TODO (Access id address from vartable scope instead of curr address)
	// TODO (Check dimensions)
	// calculate current address occuppied in context
	current_address, _ := vmemory.NextGlobalTemp(types.Ids) // TODO Check Types (Validate type with semantic cube)
	return &Constant{string(val.Lit), val, types.Ids, current_address}, nil
}

func FinishOutput(idList Attrib) (int, error) {
	id_list, ok := idList.([]*Exp)

	if !ok {
		return -1, errors.New("problem casting constant in input")
	}
	for i, _ := range id_list {
		fmt.Println("In write id list#", i)
		output_str, ok := globalStackOperands.Top()
		if !ok {
			return -1, errors.New("stack is empty in writing")
		}
		globalStackOperands, _ = globalStackOperands.Pop()

		curr_quad := quadruples.Cuadruplo{"WRITE", "-1", "-1", output_str}
		globalCurrQuads = append(globalCurrQuads, curr_quad)
	}
	return 1, nil
}

func NewOutput(id, idList Attrib) ([]*Exp, error) {
	new_id, ok := id.(*Exp)
	id_list, _ := idList.([]*Exp)
	if !ok {
		return nil, errors.New("problem casting constant in output")
	}
	return append([]*Exp{new_id}, id_list...), nil // Prepend (Add first)
}

func Return(exp Attrib) (*Exp, error) {
	new_exp, ok := exp.(*Exp)
	if !ok {
		return nil, errors.New("problem casting exp in return")
	}
	curr_top, ok := globalStackOperands.Top()
	if !ok {
		return nil, errors.New("stack is empty in return")
	}
	globalStackOperands, _ = globalStackOperands.Pop()
	curr_quad := quadruples.Cuadruplo{"RETURN", "-1", "-1", curr_top}
	globalCurrQuads = append(globalCurrQuads, curr_quad)
	return new_exp, nil
}
