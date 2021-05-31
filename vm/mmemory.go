package vm

import (
	"errors"

	"github.com/uliseslugo/ilusa_pp/memory"
)

// Get value of address in memory block from Main memory
// Memory represents the virtual memory for the virtual machine
type Memory struct {
	mem_global   *MemoryBlock
	mem_local    *MemoryBlock
	mem_constant *MemoryBlock
	mem_pointers *MemoryBlock
	mem_scope    *MemoryBlock
}

func NewMemory() *Memory {
	return &Memory{
		NewMemoryBlock("GlobalContext", memory.GlobalContext),
		NewMemoryBlock("LocalContext", memory.LocalContext),
		NewMemoryBlock("ConstantsContext", memory.ConstantsContext),
		NewMemoryBlock("PointersContext", memory.PointersContext),
		NewMemoryBlock("Scope Context", memory.Scopestart),
	}
}

// Get Value from Main Memory
// First you need to know which context
func (mm *Memory) GetValue(addr memory.Address) (interface{}, error) {
	switch {
	case addr < memory.GlobalContext: // < 0
		return nil, errors.New("Address out of scope.")
	case addr < memory.LocalContext: // Referring to Global var 0 - 8000
		val, err := mm.mem_global.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return val, nil
	case addr < memory.ConstantsContext: // Referring to Local var 8 - 16
		val, err := mm.mem_local.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return val, nil
	case addr < memory.PointersContext: // Referring to Constant 16 - 20
		val, err := mm.mem_constant.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return val, nil
	case addr < memory.Scopestart: // Referring to Pointer Context
		val, err := mm.mem_pointers.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return val, nil
	}

	return nil, errors.New("Address out of scope")
}

/**
	SetValue
	@param addr memory address
	@param val value to be stored
	checks the context of the address and calls setValue of given context
**/
func (mm *Memory) SetValue(addr memory.Address, val interface{}) error {
	switch {
	case addr < memory.GlobalContext: // < 0
		return errors.New("Address out of scope.")
	case addr < memory.LocalContext: // Referring to Global var 0 - 8000
		err := mm.mem_global.SetValue(addr, val)
		if err != nil {
			return err
		}
		return nil
	case addr < memory.ConstantsContext: // Referring to Local var 8 - 16
		err := mm.mem_local.SetValue(addr, val)
		if err != nil {
			return err
		}
		return nil
	case addr < memory.PointersContext: // Referring to Constant 16 - 20
		err := mm.mem_constant.SetValue(addr, val)
		if err != nil {
			return err
		}
		return nil
	case addr < memory.Scopestart: // Referring to Pointers 20 - 30
		err := mm.mem_pointers.SetValue(addr, val)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("Couldn't set Value in Address out of scope")
}
