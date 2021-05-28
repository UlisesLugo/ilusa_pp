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

type VirtualMachine struct {
	funcTable tables.FuncTable
	quads     []quadruples.Cuadruplo
	ip        int // instruction pointer
	paramp    int // param pointer
	constants map[string]int
	mm        *Memory
	// TODO: Add functions attributes (Act_Records)
	// TODO: Pasarle los contextos
}

// TODO: NewVirtualMachine
func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{
		tables.FuncTable{},
		make([]quadruples.Cuadruplo, 0),
		0,
		0,
		make(map[string]int), // constants map
		NewMemory(),          // main memory
	}
}

func (vm *VirtualMachine) RunBinaryQuad(q Attrib) error {
	quad, ok := q.(quadruples.Cuadruplo)

	if !ok {
		return errors.New("Couldn't cast to Cuadruplo.")
	}

	// cast string values from quads to address
	int_var1, err_v1 := strconv.Atoi(quad.Var1)
	int_var2, err_v2 := strconv.Atoi(quad.Var2)
	int_var3, err_v3 := strconv.Atoi(quad.Res)

	if err_v1 != nil || err_v2 != nil || err_v3 != nil {
		return errors.New("Couldn't cast quad addresses from string to memory Address type")
	}

	addr_1 := memory.Address(int_var1)
	addr_2 := memory.Address(int_var2)
	addr_res := memory.Address(int_var3)

	switch quad.Op {
	case "+":
		fmt.Println("Add ", addr_1, addr_2)
		op_err := vm.Add(addr_1, addr_2, addr_res)
		if op_err != nil {
			return op_err
		}
		vm.ip++
	case "-":
		return nil
	case "*":
		return nil
	case "/":
		return nil
	case "&&":
		return nil
	case "||":
		return nil
	case "!":
		return nil
	case "<":
		return nil
	case ">":
		return nil
	case "==":
		return nil
	case "!=":
		return nil
	case "=":
		return nil
		// GOSUB, ERA, PARAM, ENDFUNC, RETURN,
	}
	return nil
}

func (vm *VirtualMachine) RunUnaryQuad(q Attrib) error {
	quad, ok := q.(quadruples.Cuadruplo)

	if !ok {
		return errors.New("Couldn't cast to Cuadruplo.")
	}
	switch quad.Op {
	case "PRINT":
		return nil
	case "GOTO":
		return nil
	case "GOTOF":
		return nil
	case "GOTOV":
		return nil
	case "READ":
		return nil
	}
	return nil
}

func (vm *VirtualMachine) RunQuad(q Attrib) error {
	quad, ok := q.(quadruples.Cuadruplo)

	if !ok {
		return errors.New("Error running quad " + fmt.Sprint(vm.ip))
	}

	if quad.Var1 == "-1" || quad.Var2 == "-1" {
		fmt.Println("Running uniry quad", vm.ip)
		vm.RunUnaryQuad(q)
	} else {
		fmt.Println("Running binary quad", vm.ip)
		vm.RunBinaryQuad(q)
	}
	return nil
}

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
}
