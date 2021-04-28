package tables

import "testing"

func TestAddVarRowId(t *testing.T) {
	var row *VarRow
	row = new(VarRow)
	row.SetId("testing_var_id")
	if(row.id_ != row.Id()){
		t.Errorf("Error: id was not correctly set");
	}
	t.Logf(row.Id());
}

// TODO(ulises) Add tests for rest of vartable.go and 
// Signal tested function with flag at the beginning