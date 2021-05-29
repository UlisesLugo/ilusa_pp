package ast

import (
	// go packages
	"errors"
	"fmt"
	"strconv"

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
var constantsMap map[string]int 

var globalStackOperators stacks.Stack
var globalStackOperands stacks.Stack
var globalStackTypes stacks.Stack
var globalStackJumps stacks.Stack
var globalCurrQuads []quadruples.Cuadruplo
var globalFuncTable *tables.FuncTable
var globalCurrentScope int
var globalVarTable *tables.VarTable
var globalOperatorsDict *semantic.HierarchyDict

var quadsCounter int

func init() {
	// globalSemanticCube := semantic.NewSemanticCube()
	globalFuncTable = tables.NewFuncTable()
	vmemory = memory.NewVirtualMemory()
	constantsMap = vmemory.ConstantMap()
	globalOperatorsDict = semantic.NewHierarchyDict() // operators hierarchy table
	globalStackOperands = make(stacks.Stack, 0)
	globalStackOperators = make(stacks.Stack, 0)
	globalStackJumps = make(stacks.Stack,0)
	globalCurrQuads = make([]quadruples.Cuadruplo, 0) // TODO change main to memory address
	globalStackTypes = make(stacks.Stack, 0)
	quadsCounter = 0

	fmt.Println("Defining globals")
}

/*
	NewProgram
	@param id Attrib
	reads the program name id
	returns progam name as a literal
*/
func NewProgram(id Attrib) (*Program, error) {
	fmt.Println("In NEW PROGRAM", globalStackOperators, globalStackOperands, globalFuncTable, constantsMap)
	// cast id Attrib to token literal string
	nombre := string(id.(*token.Token).Lit)
	// cast id Attrib to token
	new_id, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("Program " + nombre + "is not valid")
	}
	// Prepend main quad
	main_quad := quadruples.Cuadruplo{"GOTO", "-1", "-1", "main"}
	end_quad := quadruples.Cuadruplo{"END","-1","-1","-1"}
	globalCurrQuads = append([]quadruples.Cuadruplo{main_quad}, globalCurrQuads...)
	globalCurrQuads = append(globalCurrQuads, end_quad)
	quadsCounter+= 2
	return &Program{nombre, globalCurrQuads, new_id, constantsMap}, nil
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
func NewFunction(id, var_map Attrib) (*tables.FuncRow, error) {
	tok, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("problem reading function")
	}
	// cast id Attrib to string token literal
	idName := string(tok.Lit)
	row := new(tables.FuncRow)
	row.SetId(idName)
	if var_map != nil {
		new_var_map, _ := var_map.(map[string]*tables.VarRow)
		new_var_table := &tables.VarTable{}
		new_var_table.SetParent(globalVarTable) // Sets global table as parent
		new_var_table.SetTable(new_var_map)
		row.SetLocalVars(new_var_table)
	}
	globalFuncTable.AddRow(row)
	// TODO Add type checking and check to repeated func
	fmt.Println("Function:", row.Id())
	return row, nil
}

/* GlobalVarDec
*/
func GlobalVarDec(var_map Attrib) (int, error){
	if var_map == nil{
		return 0, nil
	}
	new_var_map, _ := var_map.(map[string]*tables.VarRow)
	globalVarTable = &tables.VarTable{}
	globalVarTable.SetTable(new_var_map)
	return 1,nil
}


/*
	NewTypeVariables
	@param var_map Attrib
	@param next_var_map Attrib
	returns variable map[string]VarRow
*/
func NewBlockVariables(var_map, next_var_map Attrib)(map[string]*tables.VarRow,error){
	new_var_map, ok := var_map.(map[string]*tables.VarRow)
	if !ok {
		return nil, errors.New("Problem in casting var map in block variable")
	}
	if next_var_map == nil {
		return new_var_map, nil
	}
	new_next_var_map, _ := next_var_map.(map[string]*tables.VarRow)
	for _, val := range new_next_var_map {
		if _, ok := new_var_map[val.Id()]; ok {
			return nil, errors.New(fmt.Sprint("Id redeclaration:",val.Id()))
		}
		new_var_map[val.Id()] = val
	}
	return new_var_map, nil
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
	curr_map := make(map[string]*tables.VarRow)
	if !ok || len(new_typed_var) != 1 {
		return nil, errors.New("problem in casting typed variable")
	}
	curr_map[new_typed_var[0].Id()] = new_typed_var[0]
	if var_list != nil {
		new_var_list := var_list.([]*tables.VarRow)
		for _, row := range new_var_list {
			// Check if id is already declared (same type only)
			if _, ok := curr_map[row.Id()]; ok {
				return nil, errors.New(fmt.Sprint("Id redeclaration:",row.Id()))
			}
			curr_map[row.Id()] = row
		}
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
	current_address, err := vmemory.NextGlobal(types.Ids)
	if err != nil {
		return nil, err
	}
	row.SetDirV(current_address)
	return append([]*tables.VarRow{row} ,new_rows...), nil
}


/*
	New If
	@param id Attrib
	@param dim1 Attrib
	@param dim2 Attrib
	returns variable entry
*/
func NewIf(exp, est, est_list Attrib) (quadruples.Cuadruplo, error) {
	_, ok := exp.(*Exp)
	if !ok {
		return quadruples.Cuadruplo{}, errors.New("problem in casting h_exp @if")
	}
	// TODO Validate boolean expr in stack types
	// get operand 1
	curr_top1, ok := globalStackOperands.Top() // Get result
	if !ok {
		return quadruples.Cuadruplo{}, errors.New("Cannot make if without expr")
	}
	globalStackOperands, _ = globalStackOperands.Pop()
	curr_quad := quadruples.Cuadruplo{"GOTOF", curr_top1,"-1","-2"}
	globalStackJumps = globalStackJumps.Push(fmt.Sprint(quadsCounter))
	globalCurrQuads = append(globalCurrQuads, curr_quad)
	quadsCounter++

	new_quads, ok := est.([]quadruples.Cuadruplo)
	if ok {
		for _, quad := range new_quads {
			globalCurrQuads = append(globalCurrQuads, quad)
			quadsCounter++
		}
	}
	return curr_quad,nil
}

func FinishIf(decision Attrib) (int, error) {
	new_end, ok := globalStackJumps.Top()
	if !ok {
		return -1, errors.New("expected finish of if")
	}
	int_end, _ := strconv.Atoi(new_end)
	globalCurrQuads[int_end].Res = fmt.Sprint(quadsCounter+1)
	globalStackJumps, _ = globalStackJumps.Pop()
	return 1, nil
}

/*
	NewElse
	@param id Attrib
	@param dim1 Attrib
	@param dim2 Attrib
	returns variable entry
*/
func NewElse(est, est_list Attrib) (int, error) {
	fmt.Println("In new else", est)
	return 0,nil
}

/*
	NewAssignation
	@param id Attrib
	@param dim1 Attrib
	@param dim2 Attrib
	returns variable entry
*/
func NewAssignation(id, exp Attrib) (quadruples.Cuadruplo, error) {
	// cast id to token
	tok, tok_ok := id.(*token.Token)
	if !tok_ok {
		return quadruples.Cuadruplo{}, errors.New("Problem in casting id token")
	}

	
	var_row, ok := globalVarTable.Table()[string(tok.Lit)]
	if !ok {
		return quadruples.Cuadruplo{}, errors.New(fmt.Sprint("Variable",string(tok.Lit),"has not been declared"))
	}
	current_address := memory.Address(var_row.DirV())
	
	// get operand 1
	curr_top1, ok := globalStackOperands.Top()
	if !ok {
		return quadruples.Cuadruplo{}, errors.New("Cannot assign to bad expr")
	}
	// pop operand 1
	globalStackOperands, _ = globalStackOperands.Pop()
	fmt.Println("Assign:",var_row.Id(), var_row.DirV()) // TODO Check Types
	fmt.Println("\t Stacks status",globalStackOperands, globalStackOperators)

	// Add to scope &Constant{string(val.Lit), val, types.Char, memory.Address(current_address)}, nil
	// _, exp_ok := exp.(*Exp)
	// if exp_ok {
	// 	createUnaryQuadruple(semantic.Equal)
	// }
	// curr_quad := quadruples.Cuadruplo{top, curr_top2, "-1", curr_top1}
	// globalCurrQuads = append(globalCurrQuads, curr_quad)
	// quadsCounter++;
	return quadruples.Cuadruplo{semantic.Assign,fmt.Sprint(current_address),"-1", curr_top1}, nil // return row
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
	// operatorsKey := semantic.NewOperatorKey() // operators table with keys

	level_id := globalOperatorsDict.Op_hierarchy[string(new_op)] // get hierarchy level of operator level

	top, ok := globalStackOperators.Top()        // get top operator
	top_level := globalOperatorsDict.Op_hierarchy[top] // get hierarchy level of top operator
	quadsToAdd := make([]quadruples.Cuadruplo, 0)

	for ok && top_level <= level_id { // top level has higher hierarchy level
		fmt.Println("Curr Level", new_op, ", top_level",top)
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
		fmt.Println("Adding quad temp", current_address)
		curr_quad := quadruples.Cuadruplo{top, curr_top1, curr_top2, fmt.Sprint(current_address)}
		globalStackOperands = globalStackOperands.Push(fmt.Sprint(current_address))
		quadsToAdd = append([]quadruples.Cuadruplo{curr_quad}, quadsToAdd...)
		quadsCounter++;

		top, ok = globalStackOperators.Top()
		top_level = globalOperatorsDict.Op_hierarchy[top]
	}
	globalCurrQuads = append(globalCurrQuads, quadsToAdd...)
}

func createUnaryQuadruple(new_op semantic.Operation) {
	fmt.Println("\tCreating unary cuad") // operators hierarchy table
	// operatorsKey := semantic.NewOperatorKey() // operators table with keys

	level_id := globalOperatorsDict.Op_hierarchy[string(new_op)] // get hierarchy level of operator level

	top, ok := globalStackOperators.Top()        // get top operator
	top_level := globalOperatorsDict.Op_hierarchy[top] // get hierarchy level of top operator

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
		quadsCounter++;

		top, ok = globalStackOperators.Top()
		top_level = globalOperatorsDict.Op_hierarchy[top]
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
	if globalVarTable.Table() == nil {
		return nil, errors.New(fmt.Sprint("Variable",string(val.Lit),"has not been declared"))
	}
	addr, ok := globalVarTable.Table()[string(val.Lit)] // Checking varTable
	if !ok {
		return nil, errors.New(fmt.Sprint("Variable",string(val.Lit),"has not been declared"))
	}
	current_address := memory.Address(addr.DirV())
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
	var current_address memory.Address
	if addr, ok := constantsMap[string(val.Lit)]; ok {
		current_address = memory.Address(addr)
	} else {
		current_address, _ = vmemory.NextConst(types.Integer) // TODO Check Types (Validate type with semantic cube)
		constantsMap[string(val.Lit)] = int(current_address)
	}
	fmt.Println("id=", string(val.Lit), " addr=", current_address)
	globalStackOperands = globalStackOperands.Push(fmt.Sprint(current_address))
	curr_constant := &Constant{string(val.Lit), val, types.Integer, current_address}
	return curr_constant, nil
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
	var current_address memory.Address
	if addr, ok := constantsMap[string(val.Lit)]; ok {
		current_address = memory.Address(addr)
	} else {
		current_address, _ = vmemory.NextConst(types.Float) // TODO Check Types (Validate type with semantic cube)
		constantsMap[string(val.Lit)] = int(current_address)
	}
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
		quadsCounter++
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

func FinishOutput(idList Attrib) ([]quadruples.Cuadruplo, error) {
	id_list, ok := idList.([]*Exp)
	curr_quads := make([]quadruples.Cuadruplo, 0)

	if !ok {
		return nil, errors.New("problem casting constant in input")
	}
	for i, _ := range id_list {
		fmt.Println("In write id list#", i)
		output_str, ok := globalStackOperands.Top()
		if !ok {
			return nil, errors.New("stack is empty in writing")
		}
		globalStackOperands, _ = globalStackOperands.Pop()

		curr_quad := quadruples.Cuadruplo{"WRITE", "-1", "-1", output_str}
		curr_quads = append(curr_quads, curr_quad)
	}
	return curr_quads, nil
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
	quadsCounter++
	return new_exp, nil
}
