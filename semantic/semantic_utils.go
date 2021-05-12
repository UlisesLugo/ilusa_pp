// Semantic Utils
package semantic

import (
	// go packages
	"errors"
	"strconv"

	// internal packages
	"github.com/uliseslugo/ilusa_pp/types"
)

/*
	GetReturnType
	op Operation string
	type1 type of left operando
	type2 type of right operando
	returns result of operation
*/
func GetReturnType(op Operation, type1, type2 types.CoreType) (types.CoreType, error) {
	semantic_cube := NewSemanticCube() // create semantic cube

	// swap operands to respect keys in semnatic cube
	if int(type1) > int(type2) {
		temp := type1
		type1 = type2
		type2 = temp
	}

	// build key of operation
	cube_key := string(op) + strconv.Itoa(int(type1)) + strconv.Itoa(int(type2))
	// get result in semantic cube
	result, ok := semantic_cube.operations[cube_key]
	if !ok {
		return result, errors.New("unsupported operation")
	}
	return result, nil
}
