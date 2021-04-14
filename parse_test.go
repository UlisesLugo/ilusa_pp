package test

import(
    "github.com/uliseslugo/ilusa_pp/lexer"
    "github.com/uliseslugo/ilusa_pp/parser"
    "os"
    "testing"
    "fmt"
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
    tests := []string {
        "first_test.txt",
        "second_test.txt",
    }

    for _, test := range tests {
        input, err := readFile(test)

        if err != nil {
            t.Fatalf("Error reading file %s", test);
        }

        s := lexer.NewLexer(input);
        
        st, errtest := p.Parse(s);

        if errtest != nil {
            t.Errorf("Error: %v", errtest);
        }

        fmt.Println("Program name:", st);
    }
}