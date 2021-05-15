package memory

import (
	"fmt"
)

type Address int

// Memory integer that start new context
const Globalstart = 0
const Localstart = 5000
const Tempstart = 10000
const Constantstart = 15000
const Scopestart = 20000 // main scope

// Segment size for context
const segmentSize = 1000

// Offset for constants
const IntOffset = 0
const FloatOffset = IntOffset + segmentSize
const CharOffset = FloatOffset + segmentSize
const BoolOffset = CharOffset + segmentSize
const IdOffset = BoolOffset + segmentSize

// Returns -1 if Address is negative
func (a Address) String() string {
	if a < 0 {
		return "-1"
	}

	return fmt.Sprintf("%d", a)
}
