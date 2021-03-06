package types

type CoreType int

// CoreType defines the atomical types of ILUSA++
//
// Integer: Whole number (32 bits), positive or negative
// Float: Floating point number with double precision (64 Bits), positive or negative
// Char: ASCI character (8 bits)
// Null: For operations not allowed in ILUSA++
const (
	Integer CoreType = iota
	Float
	Char
	Bool
	Ids
	Null
)
