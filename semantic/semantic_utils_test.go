// Semantic Utils Test
package semantic

import (
	"testing"

	"github.com/uliseslugo/ilusa_pp/types"
)

// TestIntFloatOperations
// @param *testing.T
// tests function of GetReturnType(Op, int, float)
func TestIntFloatOperations(t *testing.T) {
	semantic_cube := NewSemanticCube()
	// get value of operation result
	_, err := semantic_cube.GetReturnType(Add, types.Integer, types.Float)
	if err == nil { // unsupported operation
		t.Fatalf("Integer and Float should not be added")
	}
	t.Log("Int and Float ops are unsupported")
}
