package tables

import (
	"github.com/uliseslugo/ilusa_pp/gocc/token"
	"github.com/uliseslugo/ilusa_pp/types"
	"testing"
)

func TestAddRow(t *testing.T) {
	// structure for tests entries
	tests := []struct {
		fr *FuncRow
		ft *FuncTable
		want FuncDirectory
	}

	// try with examples
}