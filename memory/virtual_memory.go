package memory

// Virtual memory struct with address counters & constants map
type VirtualMemory struct {
	// global segment
	global_int_count  int
	global_char_count int
	global_bool_count int
	global_ids_count  int

	// global temporal segment
	global_temp_int_count  int
	global_temp_char_count int
	global_temp_bool_count int
	global_temp_ids_count  int

	// local segment
	local_int_count  int
	local_char_count int
	local_bool_count int
	local_ids_count  int

	// local temporal segment
	local_temp_int_count  int
	local_temp_char_count int
	local_temp_bool_count int
	local_temp_ids_count  int

	// constants segment
	const_int_count  int
	const_char_count int
	const_bool_count int
	const_ids_count  int

	// pointers
	pointers_int_count  int
	pointers_char_count int
	pointers_bool_count int
	pointers_ids_count  int

	constants_map map[string]int // address for constants map
}

// TODO: Create new virtual memory function, get next local, temp, constant functions
// TODO: Constants map functions (add, get, check)

/**
	ResetLocalMemory
	resets counters of local scope in Virtual memory
**/
func (vm *VirtualMemory) ResetLocalMemory() {
	vm.local_int_count = 0
	vm.local_char_count = 0
	vm.local_bool_count = 0
	vm.local_ids_count = 0

	vm.local_temp_int_count = 0
	vm.local_temp_char_count = 0
	vm.local_temp_bool_count = 0
	vm.local_temp_ids_count = 0
}
