package memory

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/uliseslugo/ilusa_pp/types"
)

/**
	InsertConstant
	inserts new constant in map
	@param cte constant token
	@param t type of constant
	returns address of constant
**/
func (vm *VirtualMemory) InsertConstant(cte string, t types.CoreType) (Address, error) {
	fmt.Println("Constant="+cte, "Type "+strconv.Itoa(int(t)))
	if !vm.FindConstant(cte) { // new constant
		nextAvailable, next_err := vm.NextConst(t)
		if next_err != nil {
			// add constant
			return Address(-1), errors.New("Couldn't find next available address for new constant.")
		}
		vm.constants_map[cte] = int(nextAvailable)
		return Address(nextAvailable), nil
	}
	addr := Address(vm.constants_map[cte]) // constant already in map
	return addr, nil
}

/**
	FindConstant
	@param cte constant token
	return boolean
**/
func (vm *VirtualMemory) FindConstant(cte string) bool {
	_, ok := vm.constants_map[cte]
	return ok
}

/**
	ConstantAddress
	@param cte constant token
	returns address of constant in virtual memory
**/
func (vm *VirtualMemory) ConstantAddress(cte string) (Address, error) {
	addr, ok := vm.constants_map[cte]
	if !ok {
		return Address(-1), errors.New("Constant not found in memory.")
	}

	return Address(addr), nil
}

/**
	ConstantsMap
	returns constatns map
**/
func (vm *VirtualMemory) ConstantMap() map[string]int {
	return vm.constants_map
}
