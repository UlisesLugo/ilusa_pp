package virtual_machine

import (
	"errors"

	"github.com/uliseslugo/ilusa_pp/memory"
	"github.com/uliseslugo/ilusa_pp/quadruples"
	"github.com/uliseslugo/ilusa_pp/tables"
)

type VirtualMachine struct {
	funcTable tables.FuncTable
	quads     []quadruples.Cuadruplo
	memory    memory.VirtualMemory
	ip        int // instruction pointer
	paramp    int // param pointer
	// TODO: Add functions attributes (Act_Records)
}

// TODO: NewVirtualMachine
func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{
		tables.FuncTable{},
		make([]quadruples.Cuadruplo, 0),
		memory.VirtualMemory{},
		0,
		0,
	}
}

func (vm *VirtualMachine) RunMachine() error {
	if len(vm.quads) <= 0 {
		return errors.New("Quadruples list is empty.")
	}

	// TODO: Push main activation record
	// execute quad

	// TODO: switch for actions in quads
	return nil
}
