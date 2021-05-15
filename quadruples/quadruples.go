package quadruples

import (
	"fmt"

	"github.com/uliseslugo/ilusa_pp/semantic"
)

/*
	Quadruple struct
	op: action value
	addr1: reference to first operator
	addr2: reference to second operator
	addr3: reference to result variable
*/
type Cuadruplo struct {
	Op   semantic.Operation
	Var1 string
	Var2 string
	Res  string
}

// TODO: Getters and setters for Cuadruplo
func (c Cuadruplo) String() string {
	return fmt.Sprintf("%v %s %s %s", c.Op, c.Var1, c.Var2, c.Res)
}
