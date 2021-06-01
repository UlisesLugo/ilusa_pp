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
var globalCurrentScope map[string]*tables.VarRow
var globalVarTable *tables.VarTable
var globalOperatorsDict *semantic.HierarchyDict
var globalSemanticCube *semantic.SemanticCube

var quadsCounter int
var paramOrder int

func init() {
	globalFuncTable = tables.NewFuncTable() // Function Directory
	vmemory = memory.NewVirtualMemory()     // Virtual Memory
	constantsMap = vmemory.ConstantMap()
	globalOperatorsDict = semantic.NewHierarchyDict() // operators hierarchy table
	globalStackOperands = make(stacks.Stack, 0)
	globalStackOperators = make(stacks.Stack, 0)
	globalStackJumps = make(stacks.Stack, 0)
	globalCurrQuads = make([]quadruples.Cuadruplo, 0) // TODO change main to memory address
	globalStackTypes = make(stacks.Stack, 0)
	globalSemanticCube = semantic.NewSemanticCube()
	quadsCounter = 0
	paramOrder = 0

	fmt.Println("Defining globals")
}

/*
	NewProgram
	@param id Attrib
	reads the program name id
	returns progam name as a literal
*/
func NewProgram(id, func_est, main_est Attrib) (*Program, error) {
	fmt.Println("In NEW PROGRAM", globalStackOperators, globalStackOperands, globalFuncTable, constantsMap)
	// cast id Attrib to token literal string
	nombre := string(id.(*token.Token).Lit)
	// cast id Attrib to token
	new_id, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("Program " + nombre + "is not valid")
	}

	fmt.Println("\tmain stmts", main_est)
	curr_quads := make([]quadruples.Cuadruplo, 0)

	func_quads, ok := func_est.([]quadruples.Cuadruplo)
	if ok {
		curr_quads = append(curr_quads, func_quads...)
		quadsCounter += len(func_quads)
	}
	globalCurrQuads = append(globalCurrQuads, func_quads...)

	// Prepend main quad
	main_quad := quadruples.Cuadruplo{"GOTO", "-1", "-1", fmt.Sprint(quadsCounter + 1)}
	globalCurrQuads = append([]quadruples.Cuadruplo{main_quad}, globalCurrQuads...)

	est_quads, ok := main_est.([]quadruples.Cuadruplo)
	if ok {
		curr_quads = append(curr_quads, est_quads...)
	}
	globalCurrQuads = append(globalCurrQuads, est_quads...)

	end_quad := quadruples.Cuadruplo{"END", "-1", "-1", "-1"}
	globalCurrQuads = append(globalCurrQuads, end_quad)
	quadruples.ParseQuadruples(globalCurrQuads)
	quadruples.ParseFunctionAdresses(globalCurrQuads, globalFuncTable)
	return &Program{nombre, globalCurrQuads, new_id, constantsMap, nil}, nil
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
func NewFunction(id, attrib_map, var_map, est, est_list, rest_func Attrib) ([]quadruples.Cuadruplo, error) {
	tok, ok := id.(*token.Token)
	curr_quads := make([]quadruples.Cuadruplo, 0)
	if !ok {
		return nil, errors.New("problem reading function")
	}
	// cast id Attrib to string token literal
	idName := string(tok.Lit)
	row := new(tables.FuncRow)
	row.SetId(idName)
	function_map := make(map[string]*tables.VarRow)
	new_var_table := &tables.VarTable{}

	if attrib_map != nil {
		new_attrib_map := attrib_map.(map[string]*tables.VarRow)
		function_map = new_attrib_map
	}

	if var_map != nil {
		new_var_map, _ := var_map.(map[string]*tables.VarRow)
		for _, var_row := range new_var_map {
			function_map[var_row.Id()] = var_row
		}
	}
	new_var_table.SetParent(globalVarTable) // Sets global table as parent
	new_var_table.SetTable(function_map)
	row.SetLocalVars(new_var_table)
	globalFuncTable.AddRow(row)
	// TODO Add type checking and check to repeated func
	
	start_func := quadruples.Cuadruplo{"START_FUNC","-1","-1",idName}
	curr_quads = append(curr_quads, start_func)

	// Add inner statements
	new_est, _ := est.([]quadruples.Cuadruplo)
	curr_quads = append(curr_quads, new_est...)

	if est_list != nil {
		new_est_list, _ := est_list.([]quadruples.Cuadruplo)
		curr_quads = append(curr_quads, new_est_list...)
	}

	endfunc_quad := quadruples.Cuadruplo{"ENDPROC", "-1", "-1", "-1"}
	curr_quads = append(curr_quads, endfunc_quad)

	if rest_func != nil {
		new_func_quads, _ := rest_func.([]quadruples.Cuadruplo)
		curr_quads = append(curr_quads, new_func_quads...)
	}

	fmt.Println("Function:", row.Id(), curr_quads)

	// // Reset local memory and currFunc string
	// globalCurrentScope = nil
	// vmemory.ResetLocalMemory()

	return curr_quads, nil
}

func NewFunctionCall(id, params Attrib) ([]quadruples.Cuadruplo, error) {
	tok, ok := id.(*token.Token)
	val := string(tok.Lit)
	curr_quads := make([]quadruples.Cuadruplo, 0)
	if !ok {
		return nil, errors.New("problem reading function")
	}
	if globalFuncTable == nil || globalFuncTable.Table() == nil {
		return nil, errors.New(fmt.Sprint("undefined function ", val))
	}
	func_row, ok := globalFuncTable.Table()[val]
	if !ok {
		return nil, errors.New(fmt.Sprint("undefined function ", val))
	}

	era_quad := quadruples.Cuadruplo{"ERA", "-1", "-1", func_row.Id()}
	curr_quads = append(curr_quads, era_quad)

	// TODO Add parameter verification
	sub_quad := quadruples.Cuadruplo{"GOSUB", "-1", "-1", func_row.Id()}
	curr_quads = append(curr_quads, sub_quad)
	return curr_quads, nil
}

func NewFunctionAttrib(tipo, id, rest Attrib) (map[string]*tables.VarRow, error) {
	tok, ok := id.(*token.Token)
	val := string(tok.Lit)
	// currFunc = val
	curr_map := make(map[string]*tables.VarRow)
	if !ok {
		return nil, errors.New("problem reading function attribute")
	}
	row := &tables.VarRow{}
	row.SetId(val)
	row.SetToken(tok)
	row.SetDim1(0)
	row.SetDim2(0)
	curr_type, _ := tipo.(types.CoreType)
	addr, _ := vmemory.NextLocalTemp(curr_type)
	row.SetDirV(addr)
	if rest == nil {
		curr_map[val] = row
		globalCurrentScope = curr_map
		fmt.Println("Reading attr", curr_map)
		return curr_map, nil
	}
	rest_map, _ := rest.(map[string]*tables.VarRow)
	if _, ok := rest_map[val]; ok {
		return nil, errors.New(fmt.Sprint("cannot declare multiple times variable", val))
	}
	rest_map[val] = row
	globalCurrentScope = rest_map
	return rest_map, nil
}

func NewStatements(est, est_list Attrib) ([]quadruples.Cuadruplo, error) {
	curr_quads := make([]quadruples.Cuadruplo, 0)
	new_quads, ok := est.([]quadruples.Cuadruplo)
	if ok {
		curr_quads = append(curr_quads, new_quads...)
	}
	new_est_list, _ := est_list.([]quadruples.Cuadruplo)
	curr_quads = append(curr_quads, new_est_list...)
	return curr_quads, nil
}

/* GlobalVarDec
 */
func GlobalVarDec(var_map Attrib) (int, error) {
	if var_map == nil {
		return 0, nil
	}
	new_var_map, _ := var_map.(map[string]*tables.VarRow)
	globalVarTable = &tables.VarTable{}
	globalVarTable.SetTable(new_var_map)
	return 1, nil
}

/*
	NewTypeVariables
	@param var_map Attrib
	@param next_var_map Attrib
	returns variable map[string]VarRow
*/
func NewBlockVariables(var_map, next_var_map Attrib) (map[string]*tables.VarRow, error) {
	new_var_map, ok := var_map.(map[string]*tables.VarRow)
	if !ok {
		return nil, errors.New("Problem in casting var map in block variable")
	}
	if next_var_map == nil {
		return new_var_map, nil
	}
	new_next_var_map, _ := next_var_map.(map[string]*tables.VarRow)
	for _, val := range new_next_var_map {
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
				return nil, errors.New(fmt.Sprint("Id redeclaration:", row.Id()))
			}
			curr_map[row.Id()] = row
		}
	}
	if globalCurrentScope != nil {
		for _, var_row := range curr_map {
			if _, ok := globalCurrentScope[var_row.Id()]; ok {
				return nil, errors.New("cannot redlecare variable " + var_row.Id())
			}
			globalCurrentScope[var_row.Id()] = var_row
		}
	} else {
		globalCurrentScope = curr_map
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

	// choose between local or global context
	// choose between local or global context
	var current_address memory.Address
	var err_addr error

	if globalCurrentScope != nil {
		// choose local context
		current_address, err_addr = vmemory.NextLocalTemp(row.Type())
		if err_addr != nil {
			fmt.Println("Error in new local temp: ", err_addr)
		}
	} else {
		// choose global context
		current_address, err_addr = vmemory.NextGlobalTemp(row.Type())
		if err_addr != nil {
			fmt.Println("Error in new global temp: ", err_addr)
		}
	}

	row.SetDirV(current_address)
	return append([]*tables.VarRow{row}, new_rows...), nil
}

/*
	New If
	@param id Attrib
	@param dim1 Attrib
	@param dim2 Attrib
	returns variable entry
*/
func NewIf(exp, est, est_list, else_res Attrib) ([]quadruples.Cuadruplo, error) {
	new_exp, ok := exp.(*Exp)
	curr_quads := make([]quadruples.Cuadruplo, 0)

	start_go := quadruples.Cuadruplo{"START_GO", "-1", "-1", "-1"}
	curr_quads = append(curr_quads, start_go)

	new_quads, _ := est.([]quadruples.Cuadruplo)
	if !ok {
		return nil, errors.New("problem in casting h_exp @if")
	}
	curr_quads = append(curr_quads, new_exp.quads_...)
	// TODO Validate boolean expr in stack types
	// get operand 1
	curr_top1, ok := globalStackOperands.Top() // Get result
	if !ok {
		return nil, errors.New("Cannot make if without expr")
	}
	globalStackOperands, _ = globalStackOperands.Pop()
	// Determine gotof location
	location := quadsCounter + len(new_exp.quads_) + 3 // One for main, one for gotof and one for est
	if est_list != nil {
		est_list_quads := est.([]quadruples.Cuadruplo)
		location += len(est_list_quads)
	}
	if else_res != nil {
		location += 1 // Adding to skip goto instruction
	}
	curr_quad := quadruples.Cuadruplo{"GOTOF", curr_top1, "-1", fmt.Sprint(location)}
	curr_quads = append(curr_quads, curr_quad)
	// quadsCounter++

	curr_quads = append(curr_quads, new_quads...)

	// TODO (Add append for est list)
	if est_list != nil {
		est_list_quads := est_list.([]quadruples.Cuadruplo)
		curr_quads = append(curr_quads, est_list_quads...)
	}

	if else_res != nil {
		else_quads, _ := else_res.([]quadruples.Cuadruplo)
		goto_location := location + len(else_quads) - 1
		goto_quad := quadruples.Cuadruplo{"GOTO", "-1", "-1", fmt.Sprint(goto_location)}
		curr_quads = append(curr_quads, goto_quad)
		curr_quads = append(curr_quads, else_quads...)
	}

	end_go := quadruples.Cuadruplo{"END_GO", "-1", "-1", "-1"}
	curr_quads = append(curr_quads, end_go)

	return curr_quads, nil
}

/*
	NewElse
	@param id Attrib
	@param dim1 Attrib
	@param dim2 Attrib
	returns variable entry
*/
func NewElse(est, est_list Attrib) ([]quadruples.Cuadruplo, error) {
	curr_quads := make([]quadruples.Cuadruplo, 0)

	new_quads, ok_1 := est.([]quadruples.Cuadruplo)
	if ok_1 {
		for _, quad := range new_quads {
			curr_quads = append(curr_quads, quad)
		}
	}
	fmt.Println("In new else", curr_quads, globalStackJumps)
	// TODO Add range for est_list
	return curr_quads, nil
}

/*
	NewAssignation
	@param id Attrib
	@param dim1 Attrib
	@param dim2 Attrib
	returns variable entry
*/
func NewAssignation(id, exp Attrib) ([]quadruples.Cuadruplo, error) {
	// cast id to token
	tok, tok_ok := id.(*token.Token)
	if !tok_ok {
		return nil, errors.New("Problem in casting id token")
	}
	val := string(tok.Lit)

	curr_quads := make([]quadruples.Cuadruplo, 0)
	new_exp, ok := exp.(*Exp)
	if !ok {
		return nil, errors.New("problem casting exp in assign")
	}
	curr_quads = append(curr_quads, new_exp.quads_...)

	var current_address memory.Address
	if globalCurrentScope != nil {
		var_row, ok := globalCurrentScope[val]
		if !ok {
			return nil, errors.New(fmt.Sprint("Variable", string(tok.Lit), "has not been declared"))
		}
		current_address = memory.Address(var_row.DirV()) // TODO Check Types
	}

	// get operand 1
	fmt.Println("In new assign", globalStackOperands)
	curr_top1, ok := globalStackOperands.Top()
	if !ok {
		return nil, errors.New("Cannot assign to bad expr")
	}
	// pop operand 1
	globalStackOperands, _ = globalStackOperands.Pop()
	quadToAdd := quadruples.Cuadruplo{semantic.Assign, curr_top1, "-1", fmt.Sprint(current_address)}
	curr_quads = append(curr_quads, quadToAdd)
	return curr_quads, nil // return row
}

func NewWhile(exp, est Attrib) ([]quadruples.Cuadruplo, error) {
	curr_quads := make([]quadruples.Cuadruplo, 0)

	start_go := quadruples.Cuadruplo{"START_GO", "-1", "-1", "-1"}
	curr_quads = append(curr_quads, start_go)

	new_exp, ok := exp.(*Exp)
	if !ok {
		return nil, errors.New("while must have a valid expression")
	}
	curr_quads = append(curr_quads, new_exp.quads_...)
	est_quads, _ := est.([]quadruples.Cuadruplo)

	fmt.Println("In new while", curr_quads)
	curr_top1, ok := globalStackOperands.Top() // Get result
	if !ok {
		return nil, errors.New("Cannot make while without expr")
	}
	globalStackOperands, _ = globalStackOperands.Pop()
	gotof_loc := fmt.Sprint(quadsCounter + len(new_exp.quads_) + len(est_quads) + 3)
	gotof_quad := quadruples.Cuadruplo{"GOTOF", curr_top1, "-1", gotof_loc}
	curr_quads = append(curr_quads, gotof_quad)

	curr_quads = append(curr_quads, est_quads...)

	goto_quad := quadruples.Cuadruplo{"GOTO", "-1", "-1", fmt.Sprint(quadsCounter)}
	curr_quads = append(curr_quads, goto_quad)

	end_go := quadruples.Cuadruplo{"END_GO", "-1", "-1", "-1"}
	curr_quads = append(curr_quads, end_go)

	fmt.Println("In new while", curr_quads)
	return curr_quads, nil
}

func LoopStatements(est, est_list Attrib) ([]quadruples.Cuadruplo, error) {
	curr_quads, ok := est.([]quadruples.Cuadruplo)
	if !ok {
		return nil, errors.New("loop must contain one statement")
	}
	if est_list != nil {
		new_quads := est_list.([]quadruples.Cuadruplo)
		curr_quads = append(curr_quads, new_quads...)
	}
	return curr_quads, nil
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
	curr_quads := make([]quadruples.Cuadruplo, 0)
	if exp1_ok {
		curr_quads = append(new_exp1.Quads(), curr_quads...)
		if new_exp1.const_ != nil {
			new_const = new_exp1.const_
		}
		if new_exp1.op_exp_ != nil {
			quads_to_add := createBinaryQuadruple(new_exp1.op_exp_.operation)
			curr_quads = append(curr_quads,quads_to_add...)
		}

	}
	new_exp2, exp2_ok := exp2.(*Op_exp)
	if exp2_ok {
		quads_to_add := createBinaryQuadruple(new_exp2.operation)
		curr_quads = append(quads_to_add, curr_quads...)
		if new_exp2.exp != nil {
			curr_quads = append(new_exp2.exp.quads_,curr_quads...)
		}
	}
	// fmt.Println("Adding quad in exp",new_exp1, new_exp2, curr_quads)
	return &Exp{new_exp1, new_exp2, new_const, curr_quads}, nil
}

func ResetLocalMemory() (int, error) {
	fmt.Println("Resets Local Memory for new function.")
	vmemory.ResetLocalMemory()
	globalCurrentScope = nil
	return 0, nil
}

/*
	NewOpExpression
	@param op Attrib
	@param exp Attrib

*/
func NewOperation(op, exp Attrib) (*Op_exp, error) {
	tok, t_ok := op.(*token.Token)
	if !t_ok {
		return nil, errors.New("problem in casting operator")
	}
	new_op := semantic.Operation(tok.Lit)
	globalStackOperators = globalStackOperators.Push(string(new_op))
	new_const, _ := exp.(*Constant)
	new_exp, _ := exp.(*Exp)
	return &Op_exp{new_op, new_const, new_exp}, nil
}

func createBinaryQuadruple(new_op semantic.Operation) []quadruples.Cuadruplo {
	// operatorsKey := semantic.NewOperatorKey() // operators table with keys

	level_id := globalOperatorsDict.Op_hierarchy[string(new_op)] // get hierarchy level of operator level

	top, ok := globalStackOperators.Top()              // get top operator
	top_level := globalOperatorsDict.Op_hierarchy[top] // get hierarchy level of top operator
	quadsToAdd := make([]quadruples.Cuadruplo, 0)

	for ok && top_level <= level_id { // top level has higher hierarchy level
		fmt.Println("Curr Level", new_op, ", top_level", top)
		// pop top operator
		globalStackOperators, _ = globalStackOperators.Pop()
		// get operand 2 & type 2
		curr_top2, _ := globalStackOperands.Top()
		curr_type2, _ := globalStackTypes.Top()
		// pop operand 2 & type 2
		globalStackOperands, _ = globalStackOperands.Pop()
		globalStackTypes, _ = globalStackTypes.Pop()

		// get operand 1 & type 1
		curr_top1, _ := globalStackOperands.Top()
		curr_type1, _ := globalStackTypes.Top()
		// pop operand 1 & type 1
		globalStackOperands, _ = globalStackOperands.Pop()
		globalStackTypes, _ = globalStackTypes.Pop()

		// casting from string to int
		type_1, _ := strconv.Atoi(curr_type1)
		type_2, _ := strconv.Atoi(curr_type2)

		// Check Types (Validate type with semantic cube)
		cube_type, err_cube := globalSemanticCube.GetReturnType(semantic.Operation(top), types.CoreType(type_1), types.CoreType(type_2))
		if err_cube != nil {
			// TODO: Return error
			fmt.Println("Error in Semantic Cube:", err_cube)
		}

		// generate quad
		// choose between local or global context
		var current_address memory.Address
		var err_addr error

		if globalCurrentScope != nil {
			// choose local context
			current_address, err_addr = vmemory.NextLocalTemp(cube_type)
			if err_addr != nil {
				fmt.Println("Error in new local temp: ", err_addr)
			}
		} else {
			// choose global context
			fmt.Println("No global current scope")
			current_address, err_addr = vmemory.NextGlobalTemp(cube_type)
			if err_addr != nil {
				fmt.Println("Error in new global temp: ", err_addr)
			}
		}

		curr_quad := quadruples.Cuadruplo{top, curr_top1, curr_top2, fmt.Sprint(current_address)}
		globalStackOperands = globalStackOperands.Push(fmt.Sprint(current_address))
		globalStackTypes = globalStackTypes.Push(fmt.Sprint(cube_type))
		quadsToAdd = append([]quadruples.Cuadruplo{curr_quad}, quadsToAdd...)
		fmt.Println("Adding quad temp", current_address, quadsToAdd)

		top, ok = globalStackOperators.Top()
		top_level = globalOperatorsDict.Op_hierarchy[top]
	}
	return quadsToAdd
	// globalCurrQuads = append(globalCurrQuads, quadsToAdd...)
}

func createUnaryQuadruple(new_op semantic.Operation) {
	fmt.Println("\tCreating unary cuad") // operators hierarchy table
	// operatorsKey := semantic.NewOperatorKey() // operators table with keys

	level_id := globalOperatorsDict.Op_hierarchy[string(new_op)] // get hierarchy level of operator level

	top, ok := globalStackOperators.Top()              // get top operator
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
		quadsCounter++

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
	str_val := string(val.Lit)
	if !ok {
		return nil, errors.New("problem in id constants")
	}
	if globalCurrentScope != nil {
		addr, ok := globalCurrentScope[str_val]
		if ok {
			current_address := memory.Address(addr.DirV())
			globalStackOperands = globalStackOperands.Push(fmt.Sprint(current_address))
			return &Constant{str_val, val, types.Char, current_address}, nil
		}
	}

	if globalVarTable == nil || globalVarTable.Table() == nil {
		return nil, errors.New(fmt.Sprint("Variable ", str_val, " has not been declared"))
	}
	addr, ok := globalVarTable.Table()[str_val] // Checking varTable
	if !ok {
		return nil, errors.New(fmt.Sprint("Variable", str_val, "has not been declared"))
	}
	current_address := memory.Address(addr.DirV())
	// calculate current address occuppied in context
	globalStackOperands = globalStackOperands.Push(fmt.Sprint(current_address))
	return &Constant{str_val, val, types.Char, current_address}, nil
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
	str_val := string(val.Lit)
	cte_addr, err_addr := vmemory.InsertConstant(str_val, types.Integer)
	if err_addr != nil {
		return nil, err_addr
	}

	fmt.Println("id=", str_val, " addr=", cte_addr)
	globalStackOperands = globalStackOperands.Push(fmt.Sprint(cte_addr))
	globalStackTypes = globalStackTypes.Push("0")
	curr_constant := &Constant{str_val, val, types.Integer, cte_addr}
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
	str_val := string(val.Lit)
	cte_addr, err_addr := vmemory.InsertConstant(str_val, types.Float)
	if err_addr != nil {
		return nil, err_addr
	}

	fmt.Println("id=", str_val, " addr=", cte_addr)
	globalStackOperands = globalStackOperands.Push(fmt.Sprint(cte_addr))
	globalStackTypes = globalStackTypes.Push("1")
	return &Constant{str_val, val, types.Float, cte_addr}, nil
}

/*
	NewCharConst
	@param value Attrib
*/
func NewCharConst(value Attrib) (*Constant, error) {
	val, ok := value.(*token.Token)
	if !ok {
		return nil, errors.New("problem in char constants")
	}
	// calculate current address occuppied in context
	str_val := string(val.Lit)
	cte_addr, err_addr := vmemory.InsertConstant(str_val, types.Char)
	if err_addr != nil {
		return nil, err_addr
	}

	fmt.Println("id=", str_val, " addr=", cte_addr)
	globalStackOperands = globalStackOperands.Push(fmt.Sprint(cte_addr))
	globalStackTypes = globalStackTypes.Push("2")
	return &Constant{string(val.Lit), val, types.Char, cte_addr}, nil
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
	temp := make([]string,0)
	for i := range id_list {
		output_str, ok := globalStackOperands.Top()
		if !ok {
			return nil, errors.New(fmt.Sprint("stack is empty in writing",i))
		}
		globalStackOperands, _ = globalStackOperands.Pop()
		temp = append([]string{output_str}, temp...)

	}
	for i, curr_temp := range temp {
		curr_quads = append(curr_quads, id_list[i].Quads()...)
		curr_quad := quadruples.Cuadruplo{"WRITE", "-1", "-1", curr_temp}
		curr_quads = append(curr_quads, curr_quad)

	}
	fmt.Println("Output quads", curr_quads)
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

func Return(exp Attrib) ([]quadruples.Cuadruplo, error) {
	curr_top, ok := globalStackOperands.Top()
	if !ok {
		return nil, errors.New("stack is empty in return")
	}
	globalStackOperands, _ = globalStackOperands.Pop()
	curr_quad := quadruples.Cuadruplo{"RETURN", "-1", "-1", curr_top}
	return []quadruples.Cuadruplo{curr_quad}, nil
}
