package tables

import (
	"testing"
)

func TestAddFuncRow(t *testing.T) {
	// structure for tests entries
	var row *FuncRow
	row = new(FuncRow)
	row.SetId("newFunction")
	if row.id != row.Id() {
		t.Errorf("Error: id of function was not correctly set")
	}
	t.Log("Passed with func name:", row.Id())
}
