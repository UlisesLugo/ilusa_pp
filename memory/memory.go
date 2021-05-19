package memory

import (
	"fmt"
)

type Address int

// Memory integer that start new context
const DataSegment = 0 // global scope
const CodeSegment = 4000
const StackSegment = 8000
const ConstantsSegment = 12000
const ExtraSegment = 16000

// const Scopestart = 20000 // main scope

// Segment size for segment
const segmentSize = 1000

// Offset for constants
const IntOffset = 0
const FloatOffset = 1000
const CharOffset = 2000
const BoolOffset = 3000
const IdOffset = 4000 // Ask Ulises (functions)

// Returns -1 if Address is negative
func (a Address) String() string {
	if a < 0 {
		return "-1"
	}

	return fmt.Sprintf("%d", a)
}
