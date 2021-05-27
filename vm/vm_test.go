package vm

import (
	"testing"
)

// User struct which contains a name
// a type and a list of social links

// func TestRunProgram(t *testing.T) {
// 	tests := []string{
// 		"tests/mergesort.obj",
// 	}

// 	for _, test := range tests {
// 		vm := NewVirtualMachine()

// 		err := vm.LoadProgram(test)
// 		if err != nil {
// 			t.Fatalf("Could not load program: %v", err)
// 		}

// 		err = vm.Run()
// 		if err != nil {
// 			t.Fatalf("Runtime Error: %v", err)
// 		}

// 		fmt.Printf("%s\n", vm)
// 	}
// }
func TestReadJSON(t *testing.T) {
	vm := NewVirtualMachine()

	vm.ReadJSON()

}
