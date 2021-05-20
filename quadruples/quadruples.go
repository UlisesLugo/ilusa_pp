package quadruples

import (
	"fmt"
)

/*
	Quadruple struct
	op: action value
	addr1: reference to first operator
	addr2: reference to second operator
	addr3: reference to result variable
*/
type Cuadruplo struct {
	Op   int
	Var1 int
	Var2 int
	Res  int
}

// TODO: Getters and setters for Cuadruplo
func (c Cuadruplo) String() string {
	return fmt.Sprintf("%v %v %v %v", c.Op, c.Var1, c.Var2, c.Res)
}
