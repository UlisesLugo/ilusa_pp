package vm

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/uliseslugo/ilusa_pp/memory"
	"github.com/uliseslugo/ilusa_pp/quadruples"
	"github.com/uliseslugo/ilusa_pp/stacks"
	"github.com/uliseslugo/ilusa_pp/tables"
)

type Attrib interface{}

// Our virtual machine structure
type VirtualMachine struct {
	funcTable []tables.FuncRow       // function directory
	quads     []quadruples.Cuadruplo // code
	ip        int                    // instruction pointer
	paramp    int                    // param pointer
	constants map[string]int         // constants in virtual memory
	mm        *Memory                // memory of vm
	jumps     stacks.Stack           // jump stack for execution
	// TODO: Add functions attributes (Act_Records)
}

/**
	NewVirtualMachine
	return vm
	initializes VM struct
**/
func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{
		make([]tables.FuncRow, 0),       // functable
		make([]quadruples.Cuadruplo, 0), // quads[]
		0,                               // ip
		0,                               // paramp
		make(map[string]int),            // constants map
		NewMemory(),                     // main memory of machine
		make(stacks.Stack, 0),           // jumps stack
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
		return nil
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
func (vm *VirtualMachine) RunUnaryQuad(q Attrib, f *os.File) error {
	quad, ok := q.(quadruples.Cuadruplo)

	if !ok {
		return errors.New("Couldn't cast to Cuadruplo.")
	}

	// cast string values from quads to integers
	int_var1, err_v1 := strconv.Atoi(quad.Var1)

	if err_v1 != nil {
		return errors.New("Couldn't cast quad addresses from string to memory Address type")
	}

	// cast int values to memory addresses
	addr_1 := memory.Address(int_var1)
	// addr_res := memory.Address(int_res)

	switch quad.Op {
	case "=":
		int_res, err_res := strconv.Atoi(quad.Res)
		addr_res := memory.Address(int_res)
		if err_res != nil {
			return errors.New("Couldn't cast q.Op to int")
		}
		op_err := vm.Assign(addr_1, addr_res)
		if op_err != nil {
			fmt.Println("error in assign")
			return op_err
		}
		vm.ip++
		return nil
	case "!":
		int_res, err_res := strconv.Atoi(quad.Res)
		addr_res := memory.Address(int_res)
		if err_res != nil {
			return errors.New("Couldn't cast q.Op to int")
		}
		op_err := vm.Not(addr_1, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "WRITE":
		int_res, err_res := strconv.Atoi(quad.Res)
		addr_res := memory.Address(int_res)
		if err_res != nil {
			_, err_out := f.WriteString(quad.Res)
			if err_out != nil {
				fmt.Println("Error in output file.")
				return err_out
			}
			vm.ip++
			return nil
		}
		op_err := vm.Write(addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "GOTO":
		int_res, err_res := strconv.Atoi(quad.Res)
		addr_res := memory.Address(int_res)
		if err_res != nil {
			return errors.New("Couldn't cast q.Op to int")
		}
		op_err := vm.Goto(int(addr_res))
		if op_err != nil {
			return op_err
		}
		return nil
	case "GOTOF":
		int_res, err_res := strconv.Atoi(quad.Res)
		addr_res := memory.Address(int_res)
		if err_res != nil {
			return errors.New("Couldn't cast q.Op to int")
		}
		op_err := vm.GotoF(addr_1, int(addr_res))
		if op_err != nil {
			return op_err
		}
		return nil
		// Functions
	case "PARAM":
		int_res, err_res := strconv.Atoi(quad.Res)
		addr_res := memory.Address(int_res)
		if err_res != nil {
			return errors.New("Couldn't cast q.Res to int")
		}
		op_err := vm.Param(addr_1, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "ENDPROC":
		op_err := vm.EndFunc()
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "ERA":
		// quad.Res is func name
		op_err := vm.Era(quad.Res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "GOSUB":
		op_err := vm.Gosub(quad.Res)
		if op_err != nil {
			return op_err
		}
		return nil
	case "RETURN":
		int_res, err_res := strconv.Atoi(quad.Res)
		addr_res := memory.Address(int_res)
		if err_res != nil {
			return errors.New("Couldn't cast q.Op to int")
		}
		op_err := vm.Return(addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
		return nil
	case "START_GO":
		vm.ip++
		return nil
	case "END_GO":
		vm.ip++
		return nil
		// Arrays
		// VER
	case "END":
		vm.ip++ // last quad
		return nil
	}
	return nil
}

/**
	RunNextQuad
	@param q quad
	returns error
	calls RunBinaryQuad or RunUnaryQuad according to q
**/
func (vm *VirtualMachine) RunNextQuad(file *os.File) error {
	quad := vm.quads[vm.ip]

	if quad.Var1 == "-1" || quad.Var2 == "-1" {
		err_u := vm.RunUnaryQuad(quad, file)
		if err_u != nil {
			return err_u
		}
	} else {
		err_b := vm.RunBinaryQuad(quad)
		if err_b != nil {
			return err_b
		}
	}
	return nil
}

/**
	RunMachine
	calls LoadConstants and executes vm quads
**/
func (vm *VirtualMachine) RunMachine() {
	// Create execution file
	file, err := os.Create("./exec_result.txt")

	defer file.Close()

	_, err_file := file.WriteString("--------Welcome to ILUSA Virtual Machine-------\n")

	if err != nil || err_file != nil {
		fmt.Println("Error creating execution file.")
	}

	if len(vm.quads) <= 0 {
		fmt.Println("Quadruples list is empty.")
	}

	err_ctes := vm.LoadConstants()

	if err_ctes != nil {
		fmt.Println("Couldn't load constants to main memory.")
	}

	// TODO: Push main activation record

	// execute quad
	for vm.ip < len(vm.quads)-1 {
		err := vm.RunNextQuad(file)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
