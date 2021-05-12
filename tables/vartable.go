package tables

import (
	"github.com/uliseslugo/ilusa_pp/gocc/token"
	"github.com/uliseslugo/ilusa_pp/types"
)

type VarRow struct {
	id_ string
	type_ types.CoreType
	token_ *token.Token
	dim1_ int
	dim2_ int
}

// TESTED
func (vr *VarRow) Id() string {
	return vr.id_
}

func (vr *VarRow) Type() types.CoreType {
	return vr.type_
}

func (vr *VarRow) Token() *token.Token {
	return vr.token_
}

func (vr *VarRow) Dim1() int {
	return vr.dim1_
}

func (vr *VarRow) Dim2() int {
	return vr.dim2_
}

// TESTED
func (vr *VarRow) SetId(id string) () {
	vr.id_ = id
}

func (vr *VarRow) SetType(curr_type types.CoreType) {
	vr.type_ = curr_type
}

func (vr *VarRow) SetToken(token *token.Token) {
	vr.token_ = token
}

func (vr *VarRow) SetDim1(dim1 int) {
	vr.dim1_ = dim1
}

func (vr *VarRow) SetDim2(dim2 int) {
	vr.dim2_ = dim2
}
type VarTable struct {
	table map[string]*VarRow
	parent *VarTable
}

func (vt *VarTable) Table() map[string]*VarRow {
	return vt.table;
}

func (vt *VarTable) Parent() *VarTable {
	return vt.parent;
}

func (vt *VarTable) AddRow(id string, curr_type types.CoreType, token *token.Token) error {
	// Testing adding row
	var row *VarRow
	row = new(VarRow)
	row.SetId("ulises")
	return nil
}