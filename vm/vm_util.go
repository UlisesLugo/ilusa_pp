package vm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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

func (vm *VirtualMachine) ReadJSON() {
	jsonFile, err := os.Open("../encoding.obj")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened OBJ")
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
