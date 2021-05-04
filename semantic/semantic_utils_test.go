package semantic

import (
	"testing"

	"github.com/uliseslugo/ilusa_pp/types"
)

func TestIntFloatOperations(t *testing.T) {
	result, err := GetReturnType(Add, types.Integer, types.Float)
	if err != nil {
		t.Errorf("Integer and Float cannot be added")
	}
	if result != types.Float {
		t.Errorf("Integer and Float sum (Add) does not return a float")
	}

	t.Log("Operation suported with result type:", result)
}
