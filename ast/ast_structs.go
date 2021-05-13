// Structs for AST
package ast

import (

	// internal packages
	"github.com/uliseslugo/ilusa_pp/gocc/token"
	"github.com/uliseslugo/ilusa_pp/semantic"
	"github.com/uliseslugo/ilusa_pp/types"
)

// Attrib interface
// Empty interface used to cast attributes
type Attrib interface{}

/*
	Program struct
	nombre: program name
	operaciones: list of quadruples
	id: program token
*/
type Program struct {
	nombre      string
	operaciones []Cuadruplo
	id          *token.Token
}

func (p *Program) String() string {
	return p.nombre
}

/*
	Quadruple struct
	operacion: action value
	addr1: reference to first operator
	addr2: reference to second operator
	addr3: reference to result variable
*/
type Cuadruplo struct {
	operacion semantic.Operation
	var1      string
	var2      string
	res       string
}

// TODO: Functions for Quadruples

/*
	Expresison struct
	exp1: Exp
	exp2: Exp
*/
type Exp struct {
	exp1 *Exp
	exp2 *Exp
}

/*
	Operator Expresison struct
	op: operation
	exp: Expression
*/
type Op_exp struct {
	operation semantic.Operation
	exp       *Exp
}

/*
 Constant struct
 value: literal
*/
type Constant struct {
	value string
	tok   *token.Token
	type_ *types.CoreType
}

func (const_ *Constant) GetType() *types.CoreType {
	return const_.type_
}

func (const_ *Constant) GetToken() *token.Token {
	return const_.tok
}

func (const_ *Constant) GetValue() string {
	return const_.value
}
