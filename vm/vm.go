package vm

import (
	"errors"
	"fmt"

	"github.com/uliseslugo/ilusa_pp/quadruples"
	"github.com/uliseslugo/ilusa_pp/tables"
)

type Attrib interface{}

type VirtualMachine struct {
	funcTable tables.FuncTable
	quads     []quadruples.Cuadruplo
	// memory    memory.VirtualMemory
	ip     int // instruction pointer
	paramp int // param pointer
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
	}
}

func (vm *VirtualMachine) RunBinaryQuad(q Attrib) error {
	quad, ok := q.(quadruples.Cuadruplo)

	if !ok {
		return errors.New("Couldn't cast to Cuadruplo.")
	}

	leftOperator := quad.Var1
	rightOperator := quad.Var2

	print(leftOperator, rightOperator)

	// TODO: ask ulises operators are strings or numbers ??
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

func (vm *VirtualMachine) RunMachine() error {
	if len(vm.quads) <= 0 {
		fmt.Print("Quadruples list is empty.")
		//return errors.New("Quadruples list is empty.")
	}

	// TODO: Push main activation record

	// execute quad

	// TODO: switch for actions in quads
	return nil
}
