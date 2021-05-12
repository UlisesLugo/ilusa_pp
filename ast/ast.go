package ast

import (
	"errors"
	"fmt"

	"github.com/uliseslugo/ilusa_pp/gocc/token"
	"github.com/uliseslugo/ilusa_pp/semantic"
	"github.com/uliseslugo/ilusa_pp/tables"
)

type Attrib interface{}

type Program struct {
	nombre      string
	operaciones []Cuadruplo
	id          *token.Token
}

func (p *Program) String() string {
	return p.nombre
}

type Cuadruplo struct {
	operacion semantic.Operation
	var1      string
	var2      string
	res       string
}

// Reads the program name id and returns it as a literal
func NewProgram(id Attrib) (*Program, error) {
	nombre := string(id.(*token.Token).Lit)
	new_id, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("Program " + nombre + "is not valid")
	}
	return &Program{nombre, nil, new_id}, nil
}

func NewClass(id Attrib) (string, error) {
	fmt.Println("In NewClass Func")
	className := string(id.(*token.Token).Lit)
	fmt.Println(className)
	return className, nil
}

func NewFunction(id Attrib) (*tables.FuncRow, error) {
	fmt.Println("In NewFunction Func")
	idName := string(id.(*token.Token).Lit)
	row := new(tables.FuncRow)
	row.SetId(idName)
	fmt.Println("Function:", row.Id())
	return row, nil
}

func NewVariable(id, dim1, dim2 Attrib) (*tables.VarRow, error) {
	tok, tok_ok := id.(*token.Token)
	if !tok_ok {
		return nil, errors.New("Problem in casting id token")
	}
	row := &tables.VarRow{} // TODO Constructor for VarRow
	row.SetId(string(tok.Lit))
	row.SetToken(tok)
	fmt.Println("New var:", row.Id())
	return row, nil
}

func NewExpression(exp1, exp2 Attrib) (*Exp, error) {
	fmt.Println("New expression created")
	new_term, _ := exp1.(Exp)
	new_exp, _ := exp2.(Exp)
	return &Exp{&new_term, &new_exp}, nil
}

func NewOpExpression(op, exp Attrib) (*Op_exp, error) {
	tok, t_ok := op.(*token.Token)
	if !t_ok {
		return nil, errors.New("Problem in casting operator")
	}
	new_op := semantic.Operation(tok.Lit)
	fmt.Println("Operator: ", string(tok.Lit));
	return &Op_exp{new_op, nil}, nil
}

func NewIdConst(id Attrib) (string, error) {
	tok, ok := id.(*token.Token)
	if !ok {
		return "", errors.New("Problem in id constants")
	}
	fmt.Println("New id exp", string(tok.Lit))
	return string(tok.Lit), nil
}

func NewIntConst(value Attrib) (int, error){
	num, ok := value.(*token.Token).Int32Value()
	if ok != nil{
		return -1, errors.New("Problem in integer constants")
	}
	return int(num), nil
}

func NewFloatConst(value Attrib) (float64, error){
	num, ok := value.(*token.Token).Float64Value()
	if ok != nil{
		return -1, errors.New("Problem in float constants")
	}
	return float64(num), nil
}
// TODO? () Move structs into astx for cleanliness?
