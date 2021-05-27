package vm

import (
	"fmt"
	"strconv"

	"github.com/uliseslugo/ilusa_pp/memory"
)

// Memory Segment for Run-time
type MemoryBlock struct {
	baseAddr memory.Address
	id       string // id for memory block (Example: Global)
	memlist  map[memory.Address]interface{}
}

func NewMemoryBlock(context_id string, context_start int) *MemoryBlock {
	return &MemoryBlock{
		memory.Address(context_start),
		context_id,
		make(map[memory.Address]interface{}),
	}
}

// Memory represents the virtual memory for the virtual machine
type Memory struct {
	mem_global   *MemoryBlock
	mem_local    *MemoryBlock
	mem_constant *MemoryBlock
	mem_scope    *MemoryBlock
}

func NewMemory() *Memory {
	return &Memory{
		NewMemoryBlock("GlobalContext", memory.GlobalContext),
		NewMemoryBlock("LocalContext", memory.LocalContext),
		NewMemoryBlock("ConstantsContext", memory.ConstantsContext),
		NewMemoryBlock("PointersContext", memory.PointersContext),
	}
}

/**
	GetValue
	gets value stored in given address of runtime memory
	@param addr address of stored value
	return value
**/
func (m *Memory) GetValue(addr string) interface{} {
	int_addr, err := strconv.Atoi(addr)

	// could not get int address
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if int_addr < 0 {
		fmt.Println("Invalid address.")
		return nil
	}

	mem_addr := memory.Address(int_addr)

	switch a := mem_addr; {
	case a >= 16000 && a < 20000:
		val := m.mem_constant.memlist[a]
		return val
	}
	return nil
}

/**
	SetValue
	sets value to a specific address for run-time
	@param value to put at address -> Hay que pasarle el int(address - offset) (Ejemplo: si la dir es 5003 -> pasarle 3)
	@param addr address for Run-time memory
**/
// func (mem_segment MemoryBlock) SetValue(value interface{}, addr memory.Address) error {
// 	if addr < 0 {
// 		return errors.New("Address out of scope.")
// 	}

// 	switch {
// 	case addr >= 0 && addr <= 999:
// 		// Insert integer
// 		value_int, ok := value.(int64)
// 		if ok {
// 			mem_segment.integers[addr] = int(value_int)
// 			return nil
// 		}
// 		return errors.New("Couldn't add integer to memory segment.")

// 	case addr >= 1000 && addr <= 1999:
// 		// Insert float
// 		value_float, ok := value.(float64)
// 		if ok {
// 			mem_segment.floats[addr] = value_float
// 			return nil
// 		}
// 		return errors.New("Couldn't add float to memory segment.")
// 	case addr >= 2000 && addr <= 2999:
// 		// Insert char
// 		value_char, ok := value.(rune)
// 		if ok {
// 			mem_segment.chars[addr] = value_char
// 			return nil
// 		}
// 		return errors.New("Couldn't add char to memory segment.")
// 	case addr >= 3000 && addr <= 3999:
// 		// Insert bool
// 		value_bool := value.(int) // 0 or 1
// 		if value_bool == 0 || value_bool == 1 {
// 			mem_segment.booleans[addr] = value_bool
// 			return nil
// 		}
// 		return errors.New("Couldn't add char to memory segment.")
// 	case addr >= 4000 && addr <= 4999:
// 		// Insert id
// 		value_id, ok := value.(string)
// 		if ok {
// 			mem_segment.ids[addr] = value_id
// 		}
// 		return errors.New("Couldn't add id to memory segment.")

// 	default:
// 		return errors.New("Address out of scope")
// 	}
// }
