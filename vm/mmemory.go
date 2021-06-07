package vm

import (
	"errors"
	"fmt"

	"github.com/uliseslugo/ilusa_pp/memory"
)

// Get value of address in memory block from Main memory
// Memory represents the virtual memory for the virtual machine
type Memory struct {
	mem_global     *MemoryBlock
	mem_local      *MemoryBlock
	mem_constant   *MemoryBlock   // TODO: Create memory block for constants
	mem_pointers   *MemoryBlock   // TODO: Create memory block for pointers
	mem_scope      *MemoryBlock   // TODO: ?????
	callStack      []*MemoryBlock // ss
	prevFuncsStack []*MemoryBlock // vf
}

func NewMemory() *Memory {
	return &Memory{
		NewMemoryBlock("GlobalContext", memory.GlobalContext),
		NewMemoryBlock("LocalContext", memory.LocalContext),
		NewMemoryBlock("ConstantsContext", memory.ConstantsContext),
		NewMemoryBlock("PointersContext", memory.PointersContext),
		NewMemoryBlock("Scope Context", memory.Scopestart),
		make([]*MemoryBlock, 0),
		make([]*MemoryBlock, 0),
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
		if len(mm.prevFuncsStack) > 0 {
			top := mm.prevFuncsStack[len(mm.prevFuncsStack)-1]
			val, err := top.GetValue(addr)
			if err != nil {
				return nil, err
			}
			return val, nil
		} else {
			val, err := mm.mem_local.GetValue(addr)
			if err != nil {
				return nil, err
			}
			return val, nil
		}
	case addr < memory.PointersContext: // Referring to Constant 16 - 20
		val, err := mm.mem_constant.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return val, nil
	case addr < memory.Scopestart: // Referring to Pointer Context
		// first we need to get the address of pointer
		val_addr, err_addr := mm.mem_pointers.GetValue(addr)
		if err_addr != nil {
			fmt.Println("Error getting address of pointer.")
			return nil, err_addr
		}
		// go to pointer address
		val_int := val_addr.(int)
		val_dir := memory.Address(val_int)
		val, err := mm.GetValue(val_dir)
		if err != nil {
			fmt.Println("Error in indirect addressing of pointer.")
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
		if len(mm.prevFuncsStack) > 0 {
			top := mm.prevFuncsStack[len(mm.prevFuncsStack)-1]
			err := top.SetValue(addr, val)
			if err != nil {
				return err
			}
		} else { // no functions
			err := mm.mem_local.SetValue(addr, val)
			if err != nil {
				return err
			}
		}
		return nil
	case addr < memory.PointersContext: // Referring to Constant 16 - 20
		err := mm.mem_constant.SetValue(addr, val)
		if err != nil {
			return err
		}
		return nil
	case addr < memory.Scopestart: // Referring to Pointers 20 - 30
		// check if already stored address in pointer
		val_ptr, err_ptr := mm.mem_pointers.GetValue(addr)
		if err_ptr != nil {
			fmt.Println("Error getting value of pointer address.")
			return err_ptr
		}

		// if val_ptr is not 0, it is an address
		if val_ptr != 0 {
			val_int := val_ptr.(int)
			val_dir := memory.Address(val_int) // get address

			// set address to indirect address
			err_indirect := mm.SetValue(val_dir, val)
			if err_indirect != nil {
				fmt.Println("Error setting value in indirect addressing.")
				return err_indirect
			}
		} else {
			// set new address in indirect address
			err_pt := mm.mem_pointers.SetValue(addr, val)
			if err_pt != nil {
				fmt.Println("Error setting value in pointer address.")
				return err_pt
			}
		}
		return nil
	}

	return errors.New("Couldn't set Value in Address out of scope")
}
