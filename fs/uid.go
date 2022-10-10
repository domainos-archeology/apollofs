package fs

import (
	"fmt"
)

// from the apollo docs:

// UIDs are 64-bit unique identifiers for objects. They are composed of:
// • A 36 bit creation record
// • 8 bits reserved for future use
// • A 20 bit node ID

type UID struct {
	Hi uint32
	Lo uint32
}

func (u UID) String() string {
	creationTime := uint64(u.Hi)<<4 | uint64(u.Lo)>>28
	nodeID := u.Lo & 0x1fffff
	return fmt.Sprintf("%04x.%04x", nodeID, creationTime)
}
