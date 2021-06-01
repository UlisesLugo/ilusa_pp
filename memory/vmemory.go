package memory

import (
	"errors"

	"github.com/uliseslugo/ilusa_pp/types"
)

// Virtual memory struct with address counters & constants map
// TODO: esto no se como lo vamos a usar
type VirtualMemory struct {
	// global segment
	global_int_count   int
	global_float_count int
	global_char_count  int
	global_bool_count  int
	global_ids_count   int

	// global temporal segment
	global_temp_int_count   int
	global_temp_float_count int
	global_temp_char_count  int
	global_temp_bool_count  int
	// global_temp_ids_count   int

	// local segment
	local_int_count   int
	local_float_count int
	local_char_count  int
	local_bool_count  int
	local_ids_count   int

	// local temporal segment
	local_temp_int_count   int
	local_temp_float_count int
	local_temp_char_count  int
	local_temp_bool_count  int
	// local_temp_ids_count   int

	// constants segment
	const_int_count   int
	const_float_count int
	const_char_count  int
	//const_bool_count  int
	const_ids_count int

	// pointers
	pointers_int_count   int
	pointers_float_count int
	pointers_char_count  int
	pointers_bool_count  int
	// pointers_ids_count   int

	constants_map map[string]int // address for constants map
}

/**
	ResetLocalMemory
	resets counters of local scope in Virtual memory
**/
func (vm *VirtualMemory) ResetLocalMemory() {
	vm.local_int_count = 0
	vm.local_char_count = 0
	vm.local_bool_count = 0
	vm.local_ids_count = 0

	vm.local_temp_int_count = 0
	vm.local_temp_char_count = 0
	vm.local_temp_bool_count = 0
	// vm.local_temp_ids_count = 0
}

/**
	NewVirtualMemory
	returns virtual memory struct
**/
func NewVirtualMemory() *VirtualMemory {
	return &VirtualMemory{
		0, // vm counters
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		make(map[string]int), // constants map
	}
}

// Functions for next available address

/**
	NextGlobal
	@param t type of global variable
	returns next available address for variable
**/
func (vm *VirtualMemory) NextGlobal(t types.CoreType) (Address, error) {
	switch t {
	case 0: // integer constant
		if vm.global_int_count >= segmentSize {
			return Address(-1), errors.New("Too many integers in global memory.")
		}
		addr := vm.global_int_count + IntOffset + GlobalContext
		// increment counter
		vm.global_int_count++
		return Address(addr), nil

	case 1: // float constant
		if vm.global_float_count >= segmentSize {
			return Address(-1), errors.New("Too many global floats.")
		}
		addr := vm.global_float_count + FloatOffset + GlobalContext
		// increment counter
		vm.global_float_count++
		return Address(addr), nil

	case 2: // chars
		if vm.global_char_count >= segmentSize {
			return Address(-1), errors.New("Too many chars in global memory.")
		}
		addr := vm.global_char_count + CharOffset + GlobalContext
		vm.global_char_count++
		return Address(addr), nil

	case 3: // booleans
		if vm.global_bool_count >= segmentSize {
			return Address(-1), errors.New("Too many booleans in global memory.")
		}
		addr := vm.global_bool_count + BoolOffset + GlobalContext
		vm.global_bool_count++
		return Address(addr), nil

	case 4: // ids
		if vm.global_ids_count >= segmentSize {
			return Address(-1), errors.New("Too many ids in global memory.")
		}
		addr := vm.global_ids_count + IdOffset + GlobalContext
		vm.global_ids_count++
		return Address(addr), nil
	}
	// type mismatch
	return Address(-1), errors.New("Type Error: Type not defined.")
}

/**
	NextGlobalTemp
	@param t type of global variable
	returns next available address for global temporal
**/
func (vm *VirtualMemory) NextGlobalTemp(t types.CoreType) (Address, error) {
	switch t {
	case 0: // integer constant
		if vm.global_temp_int_count >= segmentSize {
			return Address(-1), errors.New("Too many temporal integers in global memory.")
		}
		addr := vm.global_temp_int_count + TempIntOffset + GlobalContext
		// increment counter
		vm.global_temp_int_count++
		return Address(addr), nil

	case 1: // float constant
		if vm.global_temp_float_count >= segmentSize {
			return Address(-1), errors.New("Too many temporal float constants.")
		}
		addr := vm.global_temp_float_count + TempFloatOffset + GlobalContext
		// increment counter
		vm.global_temp_float_count++
		return Address(addr), nil

	case 2: // chars
		if vm.global_temp_char_count >= segmentSize {
			return Address(-1), errors.New("Too many temporal chars in local memory.")
		}
		addr := vm.global_temp_char_count + TempCharOffset + GlobalContext
		vm.global_temp_char_count++
		return Address(addr), nil

	case 3: // booleans
		if vm.global_temp_bool_count >= segmentSize {
			return Address(-1), errors.New("Too many temporal booleans in local memory.")
		}
		addr := vm.global_temp_bool_count + TempBoolOffset + GlobalContext
		vm.global_temp_bool_count++
		return Address(addr), nil

	}
	// type mismatch
	return Address(-1), errors.New("Type Error: Temporal type not defined.")
}

/**
	NextLocal
	@param t type of local variable
	returns next available address for variable
**/
func (vm *VirtualMemory) NextLocal(t types.CoreType) (Address, error) {
	switch t {
	case 0: // integer constant
		if vm.local_int_count >= segmentSize {
			return Address(-1), errors.New("Too many integers in local memory.")
		}
		addr := vm.local_int_count + IntOffset + LocalContext
		// increment counter
		vm.local_int_count++
		return Address(addr), nil

	case 1: // float constant
		if vm.local_float_count >= segmentSize {
			return Address(-1), errors.New("Too many float constants.")
		}
		addr := vm.local_float_count + FloatOffset + LocalContext
		// increment counter
		vm.local_float_count++
		return Address(addr), nil

	case 2: // chars
		if vm.local_char_count >= segmentSize {
			return Address(-1), errors.New("Too many chars in local memory.")
		}
		addr := vm.local_char_count + CharOffset + LocalContext
		vm.local_char_count++
		return Address(addr), nil

	case 3: // booleans
		if vm.local_bool_count >= segmentSize {
			return Address(-1), errors.New("Too many booleans in local memory.")
		}
		addr := vm.local_bool_count + BoolOffset + LocalContext
		vm.local_bool_count++
		return Address(addr), nil

	case 4: // ids
		if vm.local_ids_count >= segmentSize {
			return Address(-1), errors.New("Too many ids in local memory.")
		}
		addr := vm.local_ids_count + IdOffset + LocalContext
		vm.local_ids_count++
		return Address(addr), nil
	}
	// type mismatch
	return Address(-1), errors.New("Type Error: Constant type not defined.")
}

/**
	NextLocalTemp
	@param t type of local variable
	returns next available address for variable
**/
func (vm *VirtualMemory) NextLocalTemp(t types.CoreType) (Address, error) {
	switch t {
	case 0: // integer constant
		if vm.local_temp_int_count >= segmentSize {
			return Address(-1), errors.New("Too many temporal integers in local memory.")
		}
		addr := vm.local_temp_int_count + TempIntOffset + LocalContext
		// increment counter
		vm.local_temp_int_count++
		return Address(addr), nil

	case 1: // float constant
		if vm.local_temp_float_count >= segmentSize {
			return Address(-1), errors.New("Too many temporal float constants.")
		}
		addr := vm.local_temp_float_count + TempFloatOffset + LocalContext
		// increment counter
		vm.local_temp_float_count++
		return Address(addr), nil

	case 2: // chars
		if vm.local_temp_char_count >= segmentSize {
			return Address(-1), errors.New("Too many temporal chars in local memory.")
		}
		addr := vm.local_temp_char_count + TempCharOffset + LocalContext
		vm.local_temp_char_count++
		return Address(addr), nil

	case 3: // booleans
		if vm.local_temp_bool_count >= segmentSize {
			return Address(-1), errors.New("Too many temporal booleans in local memory.")
		}
		addr := vm.local_temp_bool_count + TempBoolOffset + LocalContext
		vm.local_bool_count++
		return Address(addr), nil

	}
	// type mismatch
	return Address(-1), errors.New("Type Error: Temporal type not defined.")
}

/**
	NextConst
	@param t type of constant to be added
	returns next available address of new constant in map
**/
func (vm *VirtualMemory) NextConst(t types.CoreType) (Address, error) {
	switch t {
	case 0: // integer constant
		if vm.const_int_count >= segmentSize {
			return Address(-1), errors.New("Too many integer constants.")
		}
		addr := vm.const_int_count + IntOffset + ConstantsContext
		vm.const_int_count++ // increment counter
		return Address(addr), nil

	case 1: // float constant
		if vm.const_float_count >= segmentSize {
			return Address(-1), errors.New("Too many float constants.")
		}
		addr := vm.const_float_count + FloatOffset + ConstantsContext
		vm.const_float_count++
		return Address(addr), nil

	case 2: // char constant
		if vm.const_char_count >= segmentSize {
			return Address(-1), errors.New("Too many char constants.")
		}
		addr := vm.const_char_count + CharOffset + ConstantsContext
		vm.const_char_count++
		return Address(addr), nil
	// TODO: case 3 - bools
	case 4:
		if vm.const_ids_count >= segmentSize {
			return Address(-1), errors.New("Too many ids constants.")
		}
		addr := vm.const_ids_count + IdOffset + ConstantsContext
		vm.const_ids_count++
		return Address(addr), nil
	}
	return Address(-1), errors.New("Type Error: Constant type not defined.")
}

/**
	NextPointer
	@param t type of constant to be added
	returns next available address of new constant in map
**/
func (vm *VirtualMemory) NextPointer(t types.CoreType) (Address, error) {
	switch t {
	case 0: // integer constant
		if vm.pointers_int_count >= segmentSize {
			return Address(-1), errors.New("Too many integer pointers.")
		}
		addr := vm.pointers_int_count + IntOffset + PointersContext
		vm.pointers_int_count++ // increment counter
		return Address(addr), nil

	case 1: // float constant
		if vm.pointers_float_count >= segmentSize {
			return Address(-1), errors.New("Too many float pointers.")
		}
		addr := vm.pointers_float_count + FloatOffset + PointersContext
		vm.pointers_float_count++ // increment counter
		return Address(addr), nil

	case 2: // char constant
		if vm.pointers_char_count >= segmentSize {
			return Address(-1), errors.New("Too many char pointers.")
		}
		addr := vm.pointers_char_count + CharOffset + PointersContext
		vm.pointers_char_count++ // increment counter
		return Address(addr), nil

	case 3: // bool constant
		if vm.pointers_bool_count >= segmentSize {
			return Address(-1), errors.New("Too many boolean pointers.")
		}
		addr := vm.pointers_bool_count + BoolOffset + PointersContext
		vm.pointers_bool_count++ // increment counter
		return Address(addr), nil
	}
	return Address(-1), errors.New("Type Error: Constant type not defined.")
}
