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
	@param op Operation string
	@param type1 type of left operando
	@param type2 type of right operando
	returns result of operation
*/
func GetReturnType(op Operation, type1, type2 types.CoreType, semantic_cube *SemanticCube) (types.CoreType, error) {
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

/*
	GetOperatorHierarchy
	@param op Operation string
	returns level of hierarchy of operator
*/
func GetOperatorHierarchy(op Operation) (int, error) {
	operatorsDict := NewOperatorsDict()
	level, l_ok := operatorsDict.Op_hierarchy[string(op)]
	if !l_ok {
		return level, errors.New("operator not found in Operators Dictionary")
	}
	return level, nil
}
