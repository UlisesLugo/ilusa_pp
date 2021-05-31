package vm

import (
	"errors"
	"fmt"

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

	temp_ints   [100]int
	temp_floats [100]float64
	temp_chars  [100]rune
	temp_bools  [100]int
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
		[100]int{},
		[100]float64{},
		[100]rune{},
		[100]int{},
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
	case idx < memory.TempIntOffset:
		typeAddr := int(idx - memory.IdOffset)
		if len(mb.ids)-1 < typeAddr {
			return nil, errors.New("Ids address out of scope.")
		}
		return mb.ids[typeAddr], nil
	case idx < memory.TempFloatOffset: // int temp
		typeAddr := int(idx - memory.TempIntOffset)
		if len(mb.temp_ints)-1 < typeAddr {
			return nil, errors.New("Temp Integers address out of scope.")
		}
		return mb.integers[typeAddr], nil
	case idx < memory.TempCharOffset: // float temp
		typeAddr := int(idx - memory.TempFloatOffset)
		if len(mb.temp_floats)-1 < typeAddr {
			return nil, errors.New("Temp Floats address out of scope.")
		}
		return mb.ids[typeAddr], nil
	case idx < memory.TempBoolOffset: // char temp
		typeAddr := int(idx - memory.TempCharOffset)
		if len(mb.temp_chars)-1 < typeAddr {
			return nil, errors.New("Temp Char address out of scope.")
		}
		return mb.ids[typeAddr], nil
	case idx < 9000: // bool temp
		typeAddr := int(idx - memory.TempBoolOffset)
		if len(mb.temp_bools)-1 < typeAddr {
			return nil, errors.New("Temp Booleans address out of scope.")
		}
		return mb.ids[typeAddr], nil
	}
	return nil, errors.New("Address out of scope")
}

/**
	SetValue
	sets value to a specific address for run-time
	@param value to put at address -> Hay que pasarle el int(address - offset) (Ejemplo: si la dir es 5003 -> pasarle 3)
	@param addr address for Run-time memory
**/
func (mb *MemoryBlock) SetValue(addr memory.Address, val interface{}) error {
	idx := addr - mb.baseAddr
	fmt.Println("Index to set value", idx)
	fmt.Println("VALUE", val)

	switch {
	case idx < 0:
		return errors.New("Invalid negative address.")
	case idx >= 0 && idx < 1000: // integer
		typeAddr := int(idx - memory.IntOffset)
		mb.integers[typeAddr] = val.(int) // set
		return nil
	case idx >= 1000 && idx < 2000: // float
		typeAddr := int(idx - memory.FloatOffset)
		fmt.Println("REFERRING TO", typeAddr)
		mb.floats[typeAddr] = val.(float64) // set
		return nil
	case idx >= 2000 && idx < 3000: // char
		typeAddr := int(idx - memory.CharOffset)
		mb.chars[typeAddr] = val.(rune) // set
		return nil
	case idx >= 3000 && idx < 4000:
		typeAddr := int(idx - memory.BoolOffset)
		mb.bools[typeAddr] = val.(int)
		return nil
	case idx >= 4000 && idx < 5000:
		typeAddr := int(idx - memory.IdOffset)
		mb.ids[typeAddr] = val.(string)
		return nil
	case idx >= 5000 && idx < 6000: // int temp
		fmt.Println("In Integer temporal set")
		typeAddr := int(idx - memory.TempIntOffset)
		mb.temp_ints[typeAddr] = val.(int)
		return nil
	case idx >= 6000 && idx < 7000: // float temp
		typeAddr := int(idx - memory.TempFloatOffset)
		mb.temp_floats[typeAddr] = float64(val.(float64))
		return nil
	case idx >= 7000 && idx < 8000: // char temp
		typeAddr := int(idx - memory.CharOffset)
		mb.temp_chars[typeAddr] = val.(rune)
		return nil
	case idx >= 8000 && idx < 9000: // bool temp
		typeAddr := int(idx - memory.BoolOffset)
		mb.temp_bools[typeAddr] = val.(int)
		return nil
	}
	return errors.New("Address out of scope")
}
