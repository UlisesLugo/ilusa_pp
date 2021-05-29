package vm

import (
	"fmt"

	"github.com/uliseslugo/ilusa_pp/memory"
)

func (vm *VirtualMachine) Add(left, right, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)

	fmt.Println("lv", left_val)
	if err_left != nil {
		return err_left
	}

	right_val, err_right := vm.mm.GetValue(right)
	if err_right != nil {
		return err_right
	}
	fmt.Println("rv", right_val)

	//	get Num
	left_num, err_ln := getNum(left_val)
	left_flt, _ := getFloat(left_val)

	right_num, err_rn := getNum(right_val)
	right_flt, _ := getFloat(right_val)

	if err_ln == nil && err_rn == nil {
		result := left_num + right_num
		fmt.Println("res", result)
		err_res := vm.mm.SetValue(res, result)
		if err_res != nil {
			return err_res
		}
		return nil
	}

	result := left_flt + right_flt
	fmt.Println("res", result)
	err_res := vm.mm.SetValue(res, result)
	if err_res != nil {
		return err_res
	}
	return nil
}
