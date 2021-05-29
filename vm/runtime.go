package vm

import (
	"errors"

	"github.com/uliseslugo/ilusa_pp/memory"
)

// Memory Segment for Run-time Memory
type MemoryBlock struct {
	baseAddr memory.Address
	id       string       // id for memory block (Example: Global)
	integers [100]int     // integeres memory
	floats   [100]float64 // floats memory
	chars    [100]rune    // chars memory
	bools    [100]int     // bools memory
	ids      [100]string  // ids memory
}

func NewMemoryBlock(context_id string, context_start int) *MemoryBlock {
	return &MemoryBlock{
		memory.Address(context_start),
		context_id,
		[100]int{},
		[100]float64{},
		[100]rune{},
		[100]int{},
		[100]string{},
		// TODO: pass size of arrays
	}
}

// Get value of address in memory block from Main memory
func (mb *MemoryBlock) GetValue(addr memory.Address) (interface{}, error) {
	idx := addr - mb.baseAddr

	switch {
	case idx < memory.IntOffset:
		return nil, errors.New("Invalid negative address.")

	case idx < memory.FloatOffset: // integer
		typeAddr := int(idx - memory.IntOffset)
		if len(mb.integers)-1 < typeAddr {
			return nil, errors.New("Integer address out of scope.")
		}
		return mb.integers[typeAddr], nil
	case idx < memory.CharOffset: // float
		typeAddr := int(idx - memory.FloatOffset)
		if len(mb.floats)-1 < typeAddr {
			return nil, errors.New("Float address out of scope.")
		}
		return mb.floats[typeAddr], nil
	case idx < memory.BoolOffset: // char
		typeAddr := int(idx - memory.CharOffset)
		if len(mb.chars)-1 < typeAddr {
			return nil, errors.New("Character address out of scope.")
		}
		return mb.chars[typeAddr], nil
		// TODO: case idx < memory.IdOffset: // id
		// 	return nil, nil
	}
	return nil, errors.New("Address out of scope")
}

// TODO: Get value of address from main memory

/**
	SetValue
	sets value to a specific address for run-time
	@param value to put at address -> Hay que pasarle el int(address - offset) (Ejemplo: si la dir es 5003 -> pasarle 3)
	@param addr address for Run-time memory
**/
func (mb *MemoryBlock) SetValue(addr memory.Address, val interface{}) error {
	idx := addr - mb.baseAddr

	switch {
	case idx < memory.IntOffset:
		return errors.New("Invalid negative address.")

	case idx < memory.FloatOffset: // integer
		typeAddr := int(idx - memory.IntOffset)
		mb.integers[typeAddr] = val.(int) // set
		return nil
	case idx < memory.CharOffset: // float
		typeAddr := int(idx - memory.FloatOffset)
		mb.floats[typeAddr] = val.(float64) // set
		return nil
	case idx < memory.BoolOffset: // char
		typeAddr := int(idx - memory.CharOffset)
		mb.chars[typeAddr] = val.(rune) // set
		return nil
		// TODO-ISA: case idx < memory.IdOffset: // id
		// 	return nil, nil
	}
	return errors.New("Address out of scope")
}
