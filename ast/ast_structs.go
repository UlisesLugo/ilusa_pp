// Structs for AST
package ast

import (

	// internal packages
	"github.com/uliseslugo/ilusa_pp/gocc/token"
	"github.com/uliseslugo/ilusa_pp/memory"
	"github.com/uliseslugo/ilusa_pp/quadruples"
	"github.com/uliseslugo/ilusa_pp/semantic"
	"github.com/uliseslugo/ilusa_pp/types"
)

// Attrib interface
// Empty interface used to cast attributes
type Attrib interface{}

/*
	Program struct
	nombre: program name
	cuads: list of quadruples
	id: program token
*/
type Program struct {
	nombre string
	quads_ []quadruples.Cuadruplo
	id     *token.Token
	constants map[string]int
}

func (p *Program) String() string {
	return p.nombre
}

func (p *Program) Quads() []quadruples.Cuadruplo {
	return p.quads_
}

func (p *Program) Consts() map[string]int {
	return p.constants
}

/*
	Expresison struct
	exp1: Exp
	exp2: Exp
*/
type Exp struct {
	exp1    *Exp
	op_exp_ *Op_exp
	const_ *Constant
}

/*
	Operator Expresison struct
	op: operation
	exp: Expression
*/
type Op_exp struct {
	operation semantic.Operation
	const_    *Constant
}

/*
 Constant struct
 value: literal
*/
type Constant struct {
	value          string
	tok            *token.Token
	type_          types.CoreType
	local_address_ memory.Address
}

func (const_ *Constant) Type() types.CoreType {
	return const_.type_
}

func (const_ *Constant) Token() *token.Token {
	return const_.tok
}

func (const_ *Constant) Value() string {
	return const_.value
}

func (const_ *Constant) Address() memory.Address {
	return const_.local_address_
}
