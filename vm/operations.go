package vm

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/uliseslugo/ilusa_pp/memory"
	"github.com/uliseslugo/ilusa_pp/tables"
)

// Arithmetic operations

func (vm *VirtualMachine) Add(left, right, res memory.Address) error {
	fmt.Println("left", left)
	left_val, err_left := vm.mm.GetValue(left)

	if err_left != nil {
		fmt.Println("Error getting left value of addition.")
		return err_left
	}

	right_val, err_right := vm.mm.GetValue(right)
	if err_right != nil {
		fmt.Println("Error getting right value of addition.")
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
		fmt.Println("int res", result)
		err_res := vm.mm.SetValue(res, result)
		if err_res != nil {
			fmt.Println("Error setting int value")
			return err_res
		}
		return nil
	}

	result := left_flt + right_flt
	fmt.Println("flt res", result)
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

func (vm *VirtualMachine) Assign(left, res memory.Address) error {
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

// Unary operations
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

func (vm *VirtualMachine) Goto(res int) error {
	if res < 0 {
		return errors.New("Invalid value for instruction pointer")
	}
	// set instruction pointer to res value
	vm.ip = res
	return nil
}

func (vm *VirtualMachine) GotoF(left memory.Address, jump int) error {
	left_val, err_left := vm.mm.GetValue(left)
	if err_left != nil {
		fmt.Println("Error in left value for GOTOF")
		return err_left
	}

	left_num, err_ln := getNum(left_val)
	if err_ln != nil {
		fmt.Println("Error in left value for GOTOF, not a number")
		return err_ln
	}

	fmt.Println("Value found for GOTOF", left_num)

	if left_num == 0 {
		fmt.Println("Falso")
		vm.ip = jump
		fmt.Println("Jump to: ", vm.ip)
	} else {
		vm.ip++
	}

	return nil
}

// Functions operations

func (vm *VirtualMachine) Gosub(funcId string) error {
	var funcR tables.FuncRow
	// Get func row
	for i := range vm.funcTable {
		fr := vm.funcTable[i]

		if fr.Id() == funcId {
			funcR = fr
			break
		}
	}

	// Push to prev functions
	// Create new context
	callStackLen := len(vm.mm.callStack)
	var topCall *MemoryBlock
	if callStackLen > 0 {
		topCall = vm.mm.callStack[callStackLen-1]
		vm.mm.prevFuncsStack = append(vm.mm.prevFuncsStack, topCall)
	}

	// Check if func has return type
	// if funcR.ReturnValue() != 5 { // is not null
	// 	assign_err := vm.Assign()
	// }

	// save current ip
	str_ip := strconv.Itoa(vm.ip)
	vm.jumps = vm.jumps.Push(str_ip)

	// Unconditional jump
	vm.ip = int(funcR.Address())
	return nil
}

func (vm *VirtualMachine) EndFunc() error {
	// Pop from prevFunctions
	prevFuncLen := len(vm.mm.prevFuncsStack)
	if prevFuncLen != 0 {
		vm.mm.prevFuncsStack = vm.mm.prevFuncsStack[:prevFuncLen-1]
	}
	// Pop from call stack
	callStackLen := len(vm.mm.callStack)
	if callStackLen != 0 {
		vm.mm.callStack = vm.mm.callStack[:callStackLen-1]
	}

	// update ip
	top_ip, ok := vm.jumps.Top()
	if !ok {
		return errors.New("Couldn't get top of vm jumps.")
	}
	vm.jumps.Pop()

	top_ip_int, err_ip := strconv.Atoi(top_ip)
	if err_ip != nil {
		return errors.New("Problem casting top ip to integer")
	}
	// Return tu destination
	vm.ip = top_ip_int
	return nil
}

func (vm *VirtualMachine) Era(funcId string) error {
	newMB := NewMemoryBlock(funcId, memory.LocalContext)
	// push mb to stack to save curr context
	vm.mm.callStack = append(vm.mm.callStack, newMB)
	return nil
}

func (vm *VirtualMachine) Param(left, res memory.Address) error {
	left_val, err_left := vm.mm.GetValue(left)
	if err_left != nil {
		return err_left
	}
	fmt.Println("leftV ", left_val)

	callStackLen := len(vm.mm.callStack)

	var funcCall *MemoryBlock

	if callStackLen > 0 {
		funcCall = vm.mm.callStack[callStackLen-1]
	}

	// TODO: Check type of param

	// Set param
	funcCall.SetValue(res, left_val)
	return nil
}

func (vm *VirtualMachine) Return(res memory.Address) error {
	res_val, err_res := vm.mm.GetValue(res)
	if err_res != nil {
		return err_res
	}
	// vm.mm.SetValue(res, res_val)

	// Get func row of curr call stack
	callStackLen := len(vm.mm.callStack)
	var currFunc *MemoryBlock
	if callStackLen != 0 {
		currFunc = vm.mm.callStack[callStackLen-1]
	}

	currFuncId := currFunc.id

	// Get address of global variable of function
	var funcR tables.FuncRow
	// Get func row
	for i := range vm.funcTable {
		fr := vm.funcTable[i]

		if fr.Id() == currFuncId {
			funcR = fr
			break
		}
	}

	vm.mm.SetValue(funcR.Return_address, res_val)
	fmt.Println("Global var:", funcR)

	// End - pop from call stack and prev functions
	// Pop from prevFunctions
	prevFuncLen := len(vm.mm.prevFuncsStack)
	if prevFuncLen != 0 {
		vm.mm.prevFuncsStack = vm.mm.prevFuncsStack[:prevFuncLen-1]
	}
	// Pop from call stack
	if callStackLen != 0 {
		vm.mm.callStack = vm.mm.callStack[:callStackLen-1]
	}

	// update ip
	top_ip, ok := vm.jumps.Top()
	if !ok {
		return errors.New("Couldn't get top of vm jumps.")
	}
	vm.jumps.Pop()

	top_ip_int, err_ip := strconv.Atoi(top_ip)
	if err_ip != nil {
		return errors.New("Problem casting top ip to integer")
	}
	// Return tu destination
	vm.ip = top_ip_int

	return nil
}
