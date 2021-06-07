package tables

import (
	"github.com/uliseslugo/ilusa_pp/gocc/token"
	"github.com/uliseslugo/ilusa_pp/memory"
	"github.com/uliseslugo/ilusa_pp/types"
)

// function row struct
type FuncRow struct {
	id         string
	return_val types.CoreType
	params     map[int]types.CoreType
	// scope int
	local_vars     *VarTable
	address        memory.Address // jump
	Return_address memory.Address
}

// fuction table struct
type FuncTable struct {
	table map[string]*FuncRow
}

func NewFuncTable() *FuncTable {
	return &FuncTable{
		map[string]*FuncRow{},
	}
}

// Getter for id
func (fr *FuncRow) Id() string {
	return fr.id
}

// Setter for id
func (fr *FuncRow) SetId(curr_id string) {
	fr.id = curr_id
}

// Getter for id
func (fr *FuncRow) Address() memory.Address {
	return fr.address
}

// Setter for id
func (fr *FuncRow) SetAddress(addr memory.Address) {
	fr.address = addr
}

func (fr *FuncRow) AddRow(id string, curr_type types.CoreType, token *token.Token) error {
	return fr.local_vars.AddRow(id, curr_type, token)
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
func (fr *FuncRow) Params() map[int]types.CoreType {
	return fr.params
}

// Setter for params
func (fr *FuncRow) SetParams(curr_params map[int]types.CoreType) {
	fr.params = curr_params
}

func (fr *FuncRow) LocalVars() *VarTable {
	return fr.local_vars
}

func (fr *FuncRow) SetLocalVars(local_vars_ *VarTable) {
	fr.local_vars = local_vars_
}

// Getter for Function Table
func (ft *FuncTable) Table() map[string]*FuncRow {
	return ft.table
}

// Add Function Row to Table
func (ft *FuncTable) AddRow(row *FuncRow) bool {
	ft.table[row.id] = row
	return true
}
