package tables

import(
	"github.com/uliseslugo/ilusa_pp/types"
	"github.com/uliseslugo/ilusa_pp/gocc/token"
)

type VarRow struct {
	id_ string
	type_ types.CoreType
	token_ *token.Token
}

type VarTable struct {
	table map[string]*VarRow
	parent *VarTable
}