package quadruples

import (
	"github.com/uliseslugo/ilusa_pp/semantic"
)

/*
	Quadruple struct
	operacion: action value
	addr1: reference to first operator
	addr2: reference to second operator
	addr3: reference to result variable
*/
type Cuadruplo struct {
	operacion semantic.Operation
	var1      string
	var2      string
	res       string
}

// TODO: Getters and setters for Cuadruplo
