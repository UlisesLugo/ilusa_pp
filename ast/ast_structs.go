package ast

import (

	// internal packages
	"github.com/uliseslugo/ilusa_pp/semantic"
)

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
