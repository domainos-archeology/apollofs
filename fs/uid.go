package fs

import (
	"fmt"
)

type UID struct {
	Hi uint32
	Lo uint32
}

// some canned UIDs from the docs that we recognize further down in String()
var (
	UIDpvlabel = UID{0x0200, 0x0000}
	UIDlvlabel = UID{0x0201, 0x0000}
	UIDvtoc    = UID{0x0202, 0x0000}
	UIDbat     = UID{0x0203, 0x0000}
)

func (u UID) String() string {
	var trailer string
	switch {
	case u == UIDpvlabel:
		trailer = " (pv_label_$uid)"
	case u == UIDlvlabel:
		trailer = " (lv_label_$uid)"
	case u == UIDvtoc:
		trailer = " (vtoc_$uid)"
	case u == UIDbat:
		trailer = " (bat_$uid)"
	}
	return fmt.Sprintf("%04x.%04x%s", u.Lo, u.Hi, trailer)
}
