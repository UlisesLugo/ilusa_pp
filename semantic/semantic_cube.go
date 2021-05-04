package semantic

import (
	"github.com/uliseslugo/ilusa_pp/types"
)

type Operation int // operation type int

// enumarate opeartion starting in 0
const (
	Add = iota
	Sub
	Mult
	Div
	Mod
	And
	Or
	Not
	LessT
	GreaterT
	Equal
)

// Create semantic cube matrix
// 11 operations
// 4 types for left operator
// 4 types for right operator

var SemanticCube = [][][]types.CoreType{
	{
		// Add
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Sub
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Mult
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Div
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Mod
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// And
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Or
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Not
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// LessT
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// GreaterT
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
	{
		// Equal
		{types.Integer}, // Integer w Integer
		{types.Float},   // Integer w Float
		{types.Null},    // Integer w Char
	},
}
