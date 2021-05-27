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

func (vm *VirtualMachine) GetConstValue(addr int) interface{} {
	if vm.constants == nil {
		return errors.New("Constants map empty in VM.")
	}

	a := memory.Address(addr)
	return vm.mm.mem_constant.memlist[a].(int)

}

func (vm *VirtualMachine) GetContext(addr string) string {
	switch a := addr; {
	case a >= "0" && a < "8000":
		return "Global"
	case a >= "8000" && a < "16000":
		return "Local"
	case a >= "16000" && a < "20000":
		return "Constants"
	case a >= "20000" && a < "30000":
		return "Pointers"
	}

	return "Null"
}

func (vm *VirtualMachine) RunBinaryQuad(q Attrib) error {
	quad, ok := q.(quadruples.Cuadruplo)

	if !ok {
		return errors.New("Couldn't cast to Cuadruplo.")
	}

	lop, _ := strconv.Atoi(quad.Var1)
	leftVal := vm.GetConstValue(lop).(int)

	rop, _ := strconv.Atoi(quad.Var2)
	rightVal := vm.GetConstValue(rop).(int)

	fmt.Println(leftVal)
	fmt.Println(rightVal)
	// fmt.Println("lop"+leftOperator, "rop"+rightOperator)

	switch quad.Op {
	case "+":
		fmt.Println(leftVal + rightVal)
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
	fmt.Println(vm.mm.mem_constant.memlist)

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
