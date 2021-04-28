package tables

// import(
// 	"github.com/uliseslugo/ilusa_pp/types"
// 	"github.com/uliseslugo/ilusa_pp/gocc/token"
// 	"fmt"
// )

// // function row struct
// type FuncRow struct {
// 	id string
// 	return_val types.CoreType
// 	params []types.CoreType
// 	scope int
// 	local_vars *VarTable // ask elda
// }

// // fuction table struct
// type FuncTable struct {
// 	table map[string]*FuncRow
// }

// // Getter for id
// func (fr *FuncRow) Id() string {
// 	return fr.id
// }

// // Setter for id
// func (fr *FuncRow) SetId(curr_id string) {
// 	fr.id = curr_id
// }

// // Getter for return value
// func (fr *FuncRow) ReturnValue() string {
// 	return fr.return_val
// }

// // Setter for return value
// func (fr *FuncRow) ReturnValue(curr_type types.CoreType) string {
// 	fr.return_val = curr_type
// }

// // Getter for params
// func (fr *FuncRow) Params() []types.CoreType {
// 	return fr.params
// }

// // Setter for params
// func (fr *FuncRow) SetParams(curr_params []types.CoreType) {
// 	fr.params = curr_params
// }

// // Getter for scope
// func (fr *FuncRow) Scope() {
// 	return fr.scope
// }

// // Setter for scope
// func (fr *FuncRow) SetScope(curr_scope int) {
// 	fr.scope = curr_scope
// }

// // Getter for local variables entry
// func (fr *FuncRow) LocalVars() {
// 	return fr.local_vars
// }

// // Setter for local variables entry
// func (fr *FuncRow) LocalVars(curr_vars *VarTable) {
// 	fr.local_vars = curr_vars
// }

// // Getter for Function Table
// func (fr *FuncRow) Table() map[string]*FuncRow {
// 	return fr.table
// }

// // Add Function Row to Table
// func (ft *FuncTable) AddRow(row *FuncRow) bool {
// 	// // check if key is inserted in map
// 	// val, ok := ft.table[e.Id()]

// 	// if !ok {
// 	// 	// add key
// 	// 	ft.table[row.Id()] = row
// 	// }

// 	// return !ok

// 	row.SetId("newFunction")
// 	fmt.Println(row.Id())
// 	return true

// }

// TODO(isabel) Add tests for rest of vartable.go
// Signal tested function with flag at the beginning