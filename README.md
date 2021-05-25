# ILUSA++ (ilusa_pp)

### _A new programming language made in go_

**This is the project for the Compilers class**

### To update the parser/lexer:

1. Modify `bar.bnf` file
2. Navigate to gocc directory `cd gocc`
3. Run `gocc -a -v -p "github.com/uliseslugo/ilusa_pp/gocc" ../bar.bnf`

### To run the tests

1. Run `go test -v` (from root directory)

## Real Progress - May 24, 2021
✅ Definition of tokens and diagrams  
✅ Contex Free Grammar
✅ Implemented gocc
✅ Tables for Functions and Variables
✅ Semantic Considerations Cube
✅ Implemented Stack Structure
✅ IC to generate quadruples in arithmetic expressions
✅ IC to generate quadruples in logical and relational expressions
✅ IC to generate quadruples for print and assignments
🤔 IC to generate quadruples for non linear statements (for,if,while)
🤔 IC to generate quadruples for return statement
🤔 IC for function definition
🤔 IC for function calling
✅ Virtual Memory Map for virtual machine
✅ Integrated virtual memory with quadruples
✅ Run-time memory map for virtual machine
🤔 Activation Records for function
🤔 Handle recursion in functions
🤔 Create obj from quadruples
🤔 Virtual machine actions (switch)
🤔 Implement Linked List for array indexing
🤔 IC for arrays
🤔 IC to generate quadruples in function calls
🤔 Memory handle for declared functions and variables
🤔 Update documentation and diagramas
