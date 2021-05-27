package test

import (
	"fmt"
	"os"
	"testing"

	"encoding/json"

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
<<<<<<< HEAD
		"vars_test_1.txt",
=======
		"expr_test_2.txt",
>>>>>>> 93eb5621e2e24fe71401d6ded6fe97e774cce622
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
			f, err := os.Create("encoding.obj")
			if err != nil {
				panic(err)
			}
			enc := json.NewEncoder(f)
			enc.Encode(prog.Quads())
			for cuad := range prog.Quads() {
				fmt.Println(prog.Quads()[cuad])
			}
		}
	}
}
