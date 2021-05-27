package vm

import (
	"fmt"
	"testing"
)

func TestReadJSON(t *testing.T) {
	vm := NewVirtualMachine()

	vm.ReadJSON()

	fmt.Println("Constants now are", vm.constants)
	fmt.Println("Quads in runtime memory:")
	for _, q := range vm.quads {
		fmt.Println(q)
	}

	fmt.Println("Constants in virtual machine:")
	for key, element := range vm.constants {
		fmt.Println(key, element)
	}

	vm.RunMachine()

	fmt.Printf("%s\n", "Running Machine Test Passed")

}
