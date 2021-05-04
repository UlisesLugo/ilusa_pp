package semantic

import (
	"errors"

	"github.com/uliseslugo/ilusa_pp/types"
)

func GetReturnType(op Operation, type1, type2 types.CoreType) (types.CoreType, error) {
	result := SemanticCube[op][type1][type2]
	if result == types.Null {
		return result, errors.New("Unsupported operation")
	}
	return result, nil
}
