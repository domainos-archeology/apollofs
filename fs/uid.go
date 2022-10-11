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
	pvlabelUID = UID{0x0200, 0x0000}
	lvlabelUID = UID{0x0201, 0x0000}
	vtocUID    = UID{0x0202, 0x0000}
	batUID     = UID{0x0203, 0x0000}
)

func (u UID) String() string {
	var trailer string
	switch {
	case u == pvlabelUID:
		trailer = " (pv_label_$uid)"
	case u == lvlabelUID:
		trailer = " (lv_label_$uid)"
	case u == vtocUID:
		trailer = " (vtoc_$uid)"
	case u == batUID:
		trailer = " (bat_$uid)"
	}
	return fmt.Sprintf("%04x.%04x%s", u.Lo, u.Hi, trailer)
}
