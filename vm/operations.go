package vm

import (
	"fmt"

	"github.com/uliseslugo/ilusa_pp/memory"
)

// Arithmetic operations

func (vm *VirtualMachine) Add(left, right, res memory.Address) error {
	fmt.Println("left", left)
	left_val, err_left := vm.mm.GetValue(left)

	if err_left != nil {
		return err_left
	}

	right_val, err_right := vm.mm.GetValue(right)
	if err_right != nil {
		return err_right
	}

	// TODO: Check if operation is valid

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

func (vm *VirtualMachine) Sub(left, right, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)
	if err_left != nil {
		return err_left
	}

	right_val, err_right := vm.mm.GetValue(right)
	if err_right != nil {
		return err_right
	}

	// TODO: Check if operation is valid

	//	get Num
	left_num, err_ln := getNum(left_val)
	left_flt, _ := getFloat(left_val)

	right_num, err_rn := getNum(right_val)
	right_flt, _ := getFloat(right_val)

	if err_ln == nil && err_rn == nil {
		result := left_num - right_num
		fmt.Println("res int", result)
		err_res := vm.mm.SetValue(res, result)
		if err_res != nil {
			return err_res
		}
		return nil
	}

	result := left_flt - right_flt
	fmt.Println("res flt", result)
	err_res := vm.mm.SetValue(res, result)
	if err_res != nil {
		return err_res
	}
	return nil
}

func (vm *VirtualMachine) Mult(left, right, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)

	if err_left != nil {
		return err_left
	}

	right_val, err_right := vm.mm.GetValue(right)
	if err_right != nil {
		return err_right
	}

	// TODO: Check if operation is valid

	//	get Num
	left_num, err_ln := getNum(left_val)
	left_flt, _ := getFloat(left_val)

	right_num, err_rn := getNum(right_val)
	right_flt, _ := getFloat(right_val)

	if err_ln == nil && err_rn == nil {
		result := left_num * right_num
		fmt.Println("res", result)
		fmt.Println("Result Address", res)
		err_res := vm.mm.SetValue(res, result)
		if err_res != nil {
			return err_res
		}
		return nil
	}

	result := left_flt * right_flt
	fmt.Println("res", result)
	err_res := vm.mm.SetValue(res, result)
	if err_res != nil {
		return err_res
	}
	return nil
}

func (vm *VirtualMachine) Div(left, right, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)

	if err_left != nil {
		return err_left
	}

	right_val, err_right := vm.mm.GetValue(right)
	if err_right != nil {
		return err_right
	}

	// TODO: Check if operation is valid

	//	get Num
	left_num, err_ln := getNum(left_val)
	left_flt, _ := getFloat(left_val)

	right_num, err_rn := getNum(right_val)
	right_flt, _ := getFloat(right_val)

	if err_ln == nil && err_rn == nil {
		result := left_num / right_num
		fmt.Println("res", result)
		err_res := vm.mm.SetValue(res, result)
		if err_res != nil {
			return err_res
		}
		return nil
	}

	result := left_flt / right_flt
	fmt.Println("res", result)
	err_res := vm.mm.SetValue(res, result)
	if err_res != nil {
		return err_res
	}
	return nil
}

func (vm *VirtualMachine) Assign(left, right, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)

	if err_left != nil {
		return err_left
	}

	err_set := vm.mm.SetValue(res, left_val)
	if err_set != nil {
		return err_set
	}
	return nil
}

// Relational operations

func (vm *VirtualMachine) GreaterT(left, right, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)

	if err_left != nil {
		return err_left
	}

	right_val, err_right := vm.mm.GetValue(right)
	if err_right != nil {
		return err_right
	}

	left_num, err_ln := getNum(left_val)
	left_flt, _ := getFloat(left_val)

	right_num, err_rn := getNum(right_val)
	right_flt, _ := getFloat(right_val)

	var result bool

	if err_ln == nil && err_rn == nil {
		result = left_num > right_num
	} else {
		result = left_flt > right_flt
	}

	// store integer instead of boolean
	var int_res int

	if result {
		int_res = 1
	} else {
		int_res = 0
	}

	fmt.Println("res", int_res)
	err_res := vm.mm.SetValue(res, int_res)
	if err_res != nil {
		return err_res
	}
	return nil
}

func (vm *VirtualMachine) LessT(left, right, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)

	if err_left != nil {
		return err_left
	}

	right_val, err_right := vm.mm.GetValue(right)
	if err_right != nil {
		return err_right
	}

	left_num, err_ln := getNum(left_val)
	left_flt, _ := getFloat(left_val)

	right_num, err_rn := getNum(right_val)
	right_flt, _ := getFloat(right_val)

	var result bool

	if err_ln == nil && err_rn == nil {
		result = left_num < right_num
	} else {
		result = left_flt < right_flt
	}

	// store integer instead of boolean
	var int_res int

	if result {
		int_res = 1
	} else {
		int_res = 0
	}

	fmt.Println("res", int_res)
	err_res := vm.mm.SetValue(res, int_res)
	if err_res != nil {
		return err_res
	}
	return nil
}

func (vm *VirtualMachine) EqualT(left, right, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)

	if err_left != nil {
		return err_left
	}

	right_val, err_right := vm.mm.GetValue(right)
	if err_right != nil {
		return err_right
	}

	left_num, err_ln := getNum(left_val)
	if err_ln != nil {
		return err_right
	}

	right_num, err_rn := getNum(right_val)
	if err_rn != nil {
		return err_right
	}

	var result bool
	// assume the values are 1 or 0
	result = left_num == right_num

	// store integer instead of boolean
	var int_res int

	if result {
		int_res = 1
	} else {
		int_res = 0
	}

	fmt.Println("res", int_res)
	err_res := vm.mm.SetValue(res, int_res)
	if err_res != nil {
		return err_res
	}
	return nil
}

func (vm *VirtualMachine) NotEqualT(left, right, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)

	if err_left != nil {
		return err_left
	}

	right_val, err_right := vm.mm.GetValue(right)
	if err_right != nil {
		return err_right
	}

	left_num, err_ln := getNum(left_val)
	if err_ln != nil {
		return err_right
	}

	right_num, err_rn := getNum(right_val)
	if err_rn != nil {
		return err_right
	}

	var result bool
	// assume the values are 1 or 0
	result = left_num != right_num

	// store integer instead of boolean
	var int_res int

	if result {
		int_res = 1
	} else {
		int_res = 0
	}

	fmt.Println("res", int_res)
	err_res := vm.mm.SetValue(res, int_res)
	if err_res != nil {
		return err_res
	}
	return nil
}

// Logical operations
func (vm *VirtualMachine) And(left, right, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)

	if err_left != nil {
		return err_left
	}

	right_val, err_right := vm.mm.GetValue(right)
	if err_right != nil {
		return err_right
	}

	left_num, err_ln := getNum(left_val)
	if err_ln != nil {
		return err_right
	}

	right_num, err_rn := getNum(right_val)
	if err_rn != nil {
		return err_right
	}

	// assume the values are 1 or 0
	// store integer instead of boolean
	var int_res int

	// assign bool values to operators
	if (left_num == 0 && right_num == 0) || (left_num == 1 && right_num == 1) {
		int_res = 1
	} else {
		int_res = 0
	}

	fmt.Println("res", int_res)
	err_res := vm.mm.SetValue(res, int_res)
	if err_res != nil {
		return err_res
	}
	return nil
}

func (vm *VirtualMachine) Or(left, right, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)

	if err_left != nil {
		return err_left
	}

	right_val, err_right := vm.mm.GetValue(right)
	if err_right != nil {
		return err_right
	}

	left_num, err_ln := getNum(left_val)
	if err_ln != nil {
		return err_right
	}

	right_num, err_rn := getNum(right_val)
	if err_rn != nil {
		return err_right
	}

	// assume the values are 1 or 0
	// store integer instead of boolean
	var int_res int

	// assign bool values to operators
	if left_num == 1 || right_num == 1 {
		int_res = 1
	} else {
		int_res = 0
	}

	fmt.Println("res", int_res)
	err_res := vm.mm.SetValue(res, int_res)
	if err_res != nil {
		return err_res
	}
	return nil
}

func (vm *VirtualMachine) Not(left, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)

	if err_left != nil {
		return err_left
	}

	left_num, err_ln := getNum(left_val)
	if err_ln != nil {
		return err_ln
	}
	// assume the values are 1 or 0
	// store integer instead of boolean
	var int_res int

	// assign bool values to operators
	if left_num == 1 {
		int_res = 0
	} else {
		int_res = 1
	}

	fmt.Println("res", int_res)
	err_res := vm.mm.SetValue(res, int_res)
	if err_res != nil {
		return err_res
	}
	return nil
}

// Unary Operations
func (vm *VirtualMachine) Write(res memory.Address) error {
	result, err_res := vm.mm.GetValue(res)
	if err_res != nil {
		fmt.Print("Error in result output")
		return err_res
	}

	// Print results
	fmt.Println("-----------------------------")
	fmt.Println("Output: ", result)
	fmt.Println("-----------------------------")
	return nil
}
