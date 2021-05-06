package semantic

import (
	"testing"

	"github.com/uliseslugo/ilusa_pp/types"
)

func TestIntFloatOperations(t *testing.T) {
	_, err := GetReturnType(Add, types.Integer, types.Float)
	if err == nil {
		t.Fatalf("Inteer and Float should not be added")
	}
	t.Log("Int and Float ops are unsupported")
}
