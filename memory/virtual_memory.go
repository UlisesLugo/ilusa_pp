package memory

// Virtual memory struct with address counters & constants map
type VirtualMemory struct {
	// global segment
	global_int_count  int
	global_char_count int
	global_bool_count int
	global_ids_count  int
	// code segment
	code_int_count  int
	code_char_count int
	code_bool_count int
	code_ids_count  int
	// stack segment
	local_int_count  int
	local_char_count int
	local_bool_count int
	local_ids_count  int
	// constants segment
	const_int_count  int
	const_char_count int
	const_bool_count int
	const_ids_count  int
	// extra segment
	extra_int_count  int
	extra_char_count int
	extra_bool_count int
	extra_ids_count  int

	constants_map map[string]int // address for constants map
}

// TODO: Create new virtual memory function, get next local, temp, constant functions
// TODO: Constants map functions (add, get, check)
