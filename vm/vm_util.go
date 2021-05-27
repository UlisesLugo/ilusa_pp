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
	if vm.constants == nil {
		return errors.New("Constants map empty in VM.")
	}

	fmt.Println("Constants in Load", vm.constants)
	for key, val := range vm.constants {
		addr := memory.Address(val)

		// insert value in constants memory
		switch a := addr; {
		case a >= 16000 && a < 17000:
			int_val, _ := strconv.Atoi(key)
			vm.mm.mem_constant.memlist[addr] = int_val
		case a >= 17000 && a < 18000:
			flt_val, _ := strconv.ParseFloat(key, 64)
			vm.mm.mem_constant.memlist[addr] = flt_val
		case addr >= 18000 && addr < 19000:
			char_val := key[0]
			vm.mm.mem_constant.memlist[addr] = rune(char_val)
		}
	}

	return nil
}
