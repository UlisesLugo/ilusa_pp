package tables

import (
	"github.com/uliseslugo/ilusa_pp/gocc/token"
	"github.com/uliseslugo/ilusa_pp/types"
)

type VarRow struct {
	id_ string
	type_ types.CoreType
	token_ *token.Token
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