package ast

import (
	"github.com/uliseslugo/ilusa_pp/semantic")


type Exp struct {
	exp1     *Exp
	exp2     *Exp
}

type Op_exp struct {
	operation semantic.Operation
	exp *Exp
}