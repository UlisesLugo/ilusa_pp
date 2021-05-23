package memory

import (
	"fmt"
)

type Address int

// Memory integer that start new context
const GlobalContext = 0        // global context + global temps
const LocalContext = 8000      // local context + local temps
const ConstantsContext = 16000 // constants context
const PointersContext = 20000  // pointers segment for array accessing

const Scopestart = 30000 // main scope

// Segment size for segment
const segmentSize = 1000

// Offset for constants
const IntOffset = 0
const FloatOffset = 1000
const CharOffset = 2000
const BoolOffset = 3000
const IdOffset = 4000 // Ask Ulises (functions)

// Offset for temporals
const TempIntOffset = 4000
const TempFloatOffset = 5000
const TempCharOffset = 6000
const TempBoolOffset = 7000

// const TempIdOffset = 8000

// Returns -1 if Address is negative
func (a Address) String() string {
	if a < 0 {
		return "-1"
	}

	return fmt.Sprintf("%d", a)
}
