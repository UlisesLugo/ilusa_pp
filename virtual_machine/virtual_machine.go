package virtual_machine

import (
	"github.com/uliseslugo/ilusa_pp/memory"
	"github.com/uliseslugo/ilusa_pp/quadruples"
	"github.com/uliseslugo/ilusa_pp/tables"
)

type VirtualMachine struct {
	funcTable tables.FuncTable
	quads     []quadruples.Cuadruplo
	memory    memory.VirtualMemory
	ip        int // instruction pointer
	// TODO: Add functions attributes (Act_Records)
}

// TODO: setValue
// SetValue takes a value and an address and finds the corresponding position to that address and
// tries to save the given value
