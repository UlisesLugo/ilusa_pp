package tables

import (
	// internal packages
	"github.com/uliseslugo/ilusa_pp/gocc/token"
	"github.com/uliseslugo/ilusa_pp/types"
)

/*
	Variable entry structure
	id string variable name
	type_ data type of variable
	token_ token for variable for gocc
	dim1_ int for arrays
	dim2_ int for 2d arrays
*/
type VarRow struct {
	id_    string
	type_  types.CoreType
	token_ *token.Token
	dim1_  int
	dim2_  int
}

// TESTED

// Getter
// Returns variable id
func (vr *VarRow) Id() string {
	return vr.id_
}

// Getter
// Returns var type
func (vr *VarRow) Type() types.CoreType {
	return vr.type_
}

// Getter
// returns var token
func (vr *VarRow) Token() *token.Token {
	return vr.token_
}

// Getter
// returns dm1 of var
func (vr *VarRow) Dim1() int {
	return vr.dim1_
}

// Getter
// returns dim2 of var
func (vr *VarRow) Dim2() int {
	return vr.dim2_
}

// TESTED

// Setter
func (vr *VarRow) SetId(id string) {
	vr.id_ = id
}

// Setter
func (vr *VarRow) SetType(curr_type types.CoreType) {
	vr.type_ = curr_type
}

// Setter
func (vr *VarRow) SetToken(token *token.Token) {
	vr.token_ = token
}

// Setter
func (vr *VarRow) SetDim1(dim1 int) {
	vr.dim1_ = dim1
}

// Setter
func (vr *VarRow) SetDim2(dim2 int) {
	vr.dim2_ = dim2
}

/*
	Variable Table struct
	table map with var rows
	parent reference
*/
type VarTable struct {
	table  map[string]*VarRow
	parent *VarTable
}

// Getter
func (vt *VarTable) Table() map[string]*VarRow {
	return vt.table
}

// Getter
func (vt *VarTable) Parent() *VarTable {
	return vt.parent
}

/*
	AddRow
	@param id
	@param curr_type
	@param token
	adds new variable row
*/
func (vt *VarTable) AddRow(id string, curr_type types.CoreType, token *token.Token) error {
	// Testing adding row
	row := new(VarRow)
	row.SetId("ulises")
	vt.table[id] = row
	// TODO: Add to var table (?)
	return nil
}
