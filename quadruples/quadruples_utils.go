package quadruples

import (
	"fmt"	
	"strconv"
)


func ParseQuadruples(quads_list []Cuadruplo) {
	loc_stack := make([]int,0)
	for i, quad := range quads_list[1:] {
		switch quad.Op {
			case "START_GO":
				loc_stack = append(loc_stack, i+2)
			case "END_GO":
				loc_stack = loc_stack[:len(loc_stack)-1]
			case "GOTO":
				curr_pos,_ := strconv.Atoi(quad.Res)
				next_pos := curr_pos + loc_stack[len(loc_stack)-1]
				quads_list[i+1].Res= fmt.Sprint(next_pos)
			case "GOTOF":
				curr_pos,_ := strconv.Atoi(quad.Res)
				next_pos := curr_pos + loc_stack[len(loc_stack)-1]
				quads_list[i+1].Res= fmt.Sprint(next_pos)


		}
	}
}
