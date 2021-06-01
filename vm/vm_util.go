package vm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"errors"

	"github.com/uliseslugo/ilusa_pp/memory"
	"github.com/uliseslugo/ilusa_pp/quadruples"
)

type Quad struct {
	Op   string `json:"Op"`
	Dir1 string `json:"Var1"`
	Dir2 string `json:"Var2"`
	Dir3 string `json:"Res"`
}

type Quads struct {
	Quads []Quad `json:"Quads"`
}

type Consts struct {
	Consts map[string]int `json:"Consts"`
}

/**
	ReadJSON
	@param vm VirtualMachine
	reads obj JSON encoding
**/
func (vm *VirtualMachine) ReadJSON() {
	jsonFile, err := os.Open("../encoding.obj")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	OBJFile, err_obj := ioutil.ReadFile("../encoding.obj")

	if err_obj != nil {
		fmt.Println(err_obj)
	}

	var quads Quads

	error_ := json.Unmarshal([]byte(OBJFile), &quads)

	if error_ != nil {
		fmt.Println(error_)
	}

	// iterate over quads from OBJ
	for i := 0; i < len(quads.Quads); i++ {
		op := quads.Quads[i].Op
		var1 := quads.Quads[i].Dir1
		var2 := quads.Quads[i].Dir2
		res := quads.Quads[i].Dir3
		// Load quad in vm
		vm.quads = append(vm.quads, quadruples.Cuadruplo{Op: op, Var1: var1, Var2: var2, Res: res})
	}

	var consts Consts

	error_ctes := json.Unmarshal([]byte(OBJFile), &consts)

	if error_ctes != nil {
		fmt.Println(error_ctes)
	}

	vm.constants = consts.Consts
}

/**
	LoadConstants
	@param vm Virtual Machine
	returns error
	loads constants from virtual memory constants map to run time memory
**/
func (vm *VirtualMachine) LoadConstants() error {
	fmt.Println("In LoadConstants of VM Main Memory")
	if vm.constants == nil {
		return errors.New("Constants map empty in VM.")
	}

	for key, val := range vm.constants {
		// insert value in constants memory
		switch a := memory.Address(val) - vm.mm.mem_constant.baseAddr; {
		case a >= 0 && a < 1000:
			type_addr := a - memory.IntOffset
			int_val, _ := strconv.Atoi(key)
			vm.mm.mem_constant.integers[type_addr] = int_val
		case a >= 1000 && a < 2000:
			type_addr := a - memory.FloatOffset
			flt_val, _ := strconv.ParseFloat(key, 64)
			vm.mm.mem_constant.floats[type_addr] = flt_val
		case a >= 2000 && a < 3000:
			char_val := key[0]
			type_addr := a - memory.CharOffset
			vm.mm.mem_constant.chars[type_addr] = rune(char_val)
			// TODO: Add bool constants and ids ??
		}
	}

	return nil
}

// get num
func getNum(val interface{}) (int, error) {
	int_, ok := val.(int)

	if !ok {
		return 0, errors.New("Cannot convert current value to num")
	}

	return int_, nil
}

// getFloat
func getFloat(val interface{}) (float64, error) {
	flt_, ok := val.(float64)

	if !ok {
		return 0, errors.New("Cannot convert current value to num")
	}

	return flt_, nil
}

// TODO: get char

// get bool

// get id
