package semantic

import (
	"errors"
	"strconv"

	"github.com/uliseslugo/ilusa_pp/types"
)

func GetReturnType(op Operation, type1, type2 types.CoreType) (types.CoreType, error) {
	semantic_cube := NewSemanticCube()

	if int(type1) > int(type2) {
		temp := type1
		type1 = type2
		type2 = temp
	}

	cube_key := string(op) + strconv.Itoa(int(type1)) + strconv.Itoa(int(type2))
	result, ok := semantic_cube.operations[cube_key]
	if !ok {
		return result, errors.New("unsupported operation")
	}
	return result, nil
}
