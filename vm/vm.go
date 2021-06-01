package vm

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/uliseslugo/ilusa_pp/memory"
	"github.com/uliseslugo/ilusa_pp/quadruples"
	"github.com/uliseslugo/ilusa_pp/tables"
)

type Attrib interface{}

// Our virtual machine structure
type VirtualMachine struct {
	funcTable tables.FuncTable       // function directory
	quads     []quadruples.Cuadruplo // code
	ip        int                    // instruction pointer
	paramp    int                    // param pointer
	constants map[string]int         // constants in virtual memory
	mm        *Memory                // memory of vm
	// TODO: Add functions attributes (Act_Records)
}

/**
	NewVirtualMachine
	return vm
	initializes VM struct
**/
func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{
		tables.FuncTable{},              // functable
		make([]quadruples.Cuadruplo, 0), // quads[]
		0,                               // ip
		0,                               // paramp
		make(map[string]int),            // constants map
		NewMemory(),                     // main memory of machine
	}
}

/**
	RunBinaryQuad
	@param q quad
	returns error
	executes operation for binary quadruple
**/
func (vm *VirtualMachine) RunBinaryQuad(q Attrib) error {
	quad, ok := q.(quadruples.Cuadruplo)

	if !ok {
		return errors.New("Couldn't cast to Cuadruplo.")
	}

	// cast string values from quads to integers
	int_var1, err_v1 := strconv.Atoi(quad.Var1)
	int_var2, err_v2 := strconv.Atoi(quad.Var2)
	int_var3, err_v3 := strconv.Atoi(quad.Res)

	if err_v1 != nil || err_v2 != nil || err_v3 != nil {
		return errors.New("Couldn't cast quad addresses from string to memory Address type")
	}

	// cast int values to memory addresses
	addr_1 := memory.Address(int_var1)
	addr_2 := memory.Address(int_var2)
	addr_res := memory.Address(int_var3)

	// Execute binary operation
	switch quad.Op {
	case "+":
		op_err := vm.Add(addr_1, addr_2, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
	case "-":
		op_err := vm.Sub(addr_1, addr_2, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "*":
		op_err := vm.Mult(addr_1, addr_2, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "/":
		op_err := vm.Div(addr_1, addr_2, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "&&":
		op_err := vm.And(addr_1, addr_2, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "||":
		op_err := vm.Or(addr_1, addr_2, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "<":
		op_err := vm.LessT(addr_1, addr_2, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case ">":
		op_err := vm.GreaterT(addr_1, addr_2, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "==":
		op_err := vm.EqualT(addr_1, addr_2, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "!=":
		op_err := vm.NotEqualT(addr_1, addr_2, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "=":
		op_err := vm.Assign(addr_1, addr_2, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
		// GOSUB, ERA, PARAM, ENDFUNC, RETURN,
	}
	return nil
}

/**
	RunUnaryQuad
	@param q quad
	returns error
	executes operation for unary quadruple

	// TODO-ISA: add operator parameter
**/
func (vm *VirtualMachine) RunUnaryQuad(q Attrib) error {
	quad, ok := q.(quadruples.Cuadruplo)

	if quad.Res == "main" {
		vm.ip++
		return nil
	}

	if !ok {
		return errors.New("Couldn't cast to Cuadruplo.")
	}

	// cast string values from quads to integers
	int_var1, err_v1 := strconv.Atoi(quad.Var1)
	int_res, err_res := strconv.Atoi(quad.Res)

	if err_v1 != nil || err_res != nil {
		return errors.New("Couldn't cast quad addresses from string to memory Address type")
	}

	// cast int values to memory addresses
	addr_1 := memory.Address(int_var1)
	addr_res := memory.Address(int_res)

	switch quad.Op {
	case "!":
		op_err := vm.Not(addr_1, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "WRITE":
		op_err := vm.Write(addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "GOTO":
		op_err := vm.Goto(int(addr_res))
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "GOTOF":
		op_err := vm.GotoF(addr_1, int(addr_res))
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "READ":
		// TODO
		return nil
	}
	return nil
}

/**
	RunQuad
	@param q quad
	returns error
	calls RunBinaryQuad or RunUnaryQuad according to q
**/
func (vm *VirtualMachine) RunQuad(q Attrib) error {
	quad, ok := q.(quadruples.Cuadruplo)

	if !ok {
		return errors.New("Error running quad " + fmt.Sprint(vm.ip))
	}

	if quad.Var1 == "-1" || quad.Var2 == "-1" {
		vm.RunUnaryQuad(q)
	} else {
		vm.RunBinaryQuad(q)
	}
	return nil
}

/**
	RunMachine
	calls LoadConstants and executes vm quads
**/
func (vm *VirtualMachine) RunMachine() {
	if len(vm.quads) <= 0 {
		fmt.Println("Quadruples list is empty.")
	}

	err_ctes := vm.LoadConstants()

	if err_ctes != nil {
		fmt.Println("Couldn't load constants to main memory.")
	}

	// TODO: Push main activation record

	// execute quad
	for i := range vm.quads {
		q := vm.quads[i]
		vm.RunQuad(q)
	}

	// Debugging quads
	fmt.Println("QUADS:")
	for i := range vm.quads {
		fmt.Println(vm.quads[i])
	}
}
