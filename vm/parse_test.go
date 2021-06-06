package vm

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/uliseslugo/ilusa_pp/ast"
	"github.com/uliseslugo/ilusa_pp/gocc/lexer"
	"github.com/uliseslugo/ilusa_pp/gocc/parser"
)

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func TestDuck(t *testing.T) {
	p := parser.NewParser()

	tests := []string{
		"../tests/fact_rec_1.isa",
	}

	for _, test := range tests {
		input, err := readFile(test)

		if err != nil {
			t.Fatalf("Error reading file %s", test)
		}

		s := lexer.NewLexer(input)

		st, errtest := p.Parse(s)

		if errtest != nil {
			t.Errorf("Error: %v", errtest)
		}

		prog, ok := st.(*ast.Program)

		// TODO: create obj_test

		if !ok {
			t.Error("Program failed")
		} else {
			f, err := os.Create("../tests/encoding.obj")
			if err != nil {
				panic(err)
			}
			enc := json.NewEncoder(f)

			obj_map := make(map[string]interface{}) // map of key: json object

			// set key for Quads
			obj_map["Quads"] = prog.Quads()

			// set key for Constants Table
			obj_map["Consts"] = prog.Consts()

			// set key for Functable
			funcDir := prog.FuncTable()

			// encodigin map
			enc.Encode(obj_map)

			// for cuad := range prog.Quads() {
			// 	fmt.Println(prog.Quads()[cuad])
			// }

			newVM := NewVirtualMachine()

			// testing vm
			newVM.ReadJSON()

			newVM.funcTable = funcDir

			newVM.RunMachine()

			fmt.Printf("%s\n", "Running Machine Test Passed")
		}
	}

}
