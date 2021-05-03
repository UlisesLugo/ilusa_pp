package tables

import (
	"github.com/uliseslugo/ilusa_pp/types"
)

// function row struct
type FuncRow struct {
	id         string
	return_val types.CoreType
	params     []types.CoreType
	// scope int
	// local_vars *VarTable // ask elda
}

// fuction table struct
type FuncTable struct {
	table map[string]*FuncRow
}

// Getter for id
func (fr *FuncRow) Id() string {
	return fr.id
}

// Setter for id
func (fr *FuncRow) SetId(curr_id string) {
	fr.id = curr_id
}

// Getter for return value
func (fr *FuncRow) ReturnValue() types.CoreType {
	return fr.return_val
}

// Setter for return value
func (fr *FuncRow) SetReturnValue(curr_type types.CoreType) {
	fr.return_val = curr_type
}

// Getter for params
func (fr *FuncRow) Params() []types.CoreType {
	return fr.params
}

// Setter for params
func (fr *FuncRow) SetParams(curr_params []types.CoreType) {
	fr.params = curr_params
}

// Getter for Function Table
func (ft *FuncTable) Table() map[string]*FuncRow {
	return ft.table
}

// Add Function Row to Table
func (ft *FuncTable) AddRow(row *FuncRow) bool {
	row.SetId("newFunction")
	return true
}
