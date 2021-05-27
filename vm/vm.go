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

	// leftOperator := quad.Var1
	// rightOperator := quad.Var2

	//fmt.Println("lop"+leftOperator, "rop"+rightOperator)

	switch quad.Op {
	case "+":
		return nil
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

func (vm *VirtualMachine) LoadConstants() error {
	if vm.constants == nil {
		return errors.New("Constants map empty in VM.")
	}

	for key, val := range vm.constants {
		addr := memory.Address(val) - vm.mm.mem_constant.baseAddr

		// get value
		if addr >= 0 && addr < 1000 {
			int_val, err_cast := strconv.Atoi(key)
			vm.mm.mem_constant.integers[addr] = int_val
			if err_cast != nil {
				fmt.Println(err_cast)
			}
			return err_cast
		} else if addr >= 1000 && addr < 2000 {
			flt_val, err_cast := strconv.ParseFloat(key, 64)
			vm.mm.mem_constant.floats[addr] = flt_val
			if err_cast != nil {
				fmt.Println(err_cast)
			}
			return err_cast
		} else if addr >= 2000 && addr < 3000 {
			char_val := key[0]
			vm.mm.mem_constant.chars[addr] = rune(char_val)
			return nil
		} else if addr >= 3000 && addr < 4000 {
			bool_val, err_cast := strconv.Atoi(key)
			vm.mm.mem_constant.booleans[addr] = bool_val
			if err_cast != nil {
				fmt.Println(err_cast)
			}
			return err_cast
		} else {
			return errors.New("Error loading consant " + key + "in address " + fmt.Sprint(val))
		}
	}

	return nil
}

func (vm *VirtualMachine) RunMachine() {
	if len(vm.quads) <= 0 {
		fmt.Println("Quadruples list is empty.")
	}

	err_cte := vm.LoadConstants()

	if err_cte != nil {
		fmt.Println("Couldn't load constants to main memory.")
	}

	// Load Constants Map from VM to memory
	fmt.Println("Constants in MM")
	fmt.Println(vm.mm.mem_constant.integers)

	// TODO: Push main activation record

	// execute quad
	for i := range vm.quads {
		q := vm.quads[i]
		if q.Var1 == "-1" || q.Var2 == "-1" {
			fmt.Println("Running uniry quad", vm.ip)
			vm.RunUnaryQuad(q)
			vm.ip++
		} else {
			fmt.Println("Running binary quad", vm.ip)
			vm.RunBinaryQuad(q)
			vm.ip++
		}
	}
}
