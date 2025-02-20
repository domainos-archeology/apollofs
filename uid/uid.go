package uid

import (
	"errors"
	"fmt"
)

type UID struct {
	Hi uint32
	Lo uint32
}

// some canned UIDs from the docs that we recognize further down in String()
var (
	Empty = UID{0, 0}

	UIDpvlabel  = UID{0x0200, 0x0000}
	UIDlvlabel  = UID{0x0201, 0x0000}
	UIDvtoc     = UID{0x0202, 0x0000}
	UIDbat      = UID{0x0203, 0x0000}
	UIDvtoc_bkt = UID{0x0204, 0x000}

	UIDrecords        = UID{0x0300, 0x0000}
	UIDhdr_undef      = UID{0x0301, 0x0000}
	UIDobject_file    = UID{0x0302, 0x0000}
	UIDundef          = UID{0x0304, 0x0000}
	UIDpad            = UID{0x0305, 0x0000}
	UIDroot           = UID{0x0308, 0x0000}
	UIDinput_pad      = UID{0x0309, 0x0000}
	UIDsio            = UID{0x030a, 0x0000}
	UIDddf            = UID{0x030b, 0x0000}
	UIDmbx            = UID{0x030c, 0x0000}
	UIDnulldev        = UID{0x030d, 0x0000}
	UIDd3m_area       = UID{0x030e, 0x0000}
	UIDd3m_sch        = UID{0x030f, 0x0000}
	UIDpipe           = UID{0x0310, 0x0000}
	UIDuasc           = UID{0x0311, 0x0000}
	UIDdirectory      = UID{0x0312, 0x0000}
	UIDunix_directory = UID{0x0313, 0x0000}
	UIDmt             = UID{0x0314, 0x0000}
	UIDsysboot        = UID{0x0315, 0x0000}
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
	case u == UIDvtoc_bkt:
		trailer = " (vtoc_bkt_$uid)"

	case u == UIDrecords:
		trailer = " (records_$uid)"
	case u == UIDhdr_undef:
		trailer = " (hdr_undef_$uid)"
	case u == UIDobject_file:
		trailer = " (object_file_$uid)"
	case u == UIDundef:
		trailer = " (UNDEF_$uid)"
	case u == UIDpad:
		trailer = " (pad_$uid)"
	case u == UIDroot:
		trailer = " (NAME_$CANNED_ROOT_UID)"
	case u == UIDinput_pad:
		trailer = " (input_pad_$uid)"
	case u == UIDsio:
		trailer = " (sio_$uid)"
	case u == UIDddf:
		trailer = " (ddf_$uid)"
	case u == UIDmbx:
		trailer = " (mbx_$uid)"
	case u == UIDnulldev:
		trailer = " (nulldev_$uid)"
	case u == UIDd3m_area:
		trailer = " (d3m_area_$uid)"
	case u == UIDd3m_sch:
		trailer = " (d3m_sch_$uid)"
	case u == UIDpipe:
		trailer = " (pipe_$uid)"
	case u == UIDuasc:
		trailer = " (uasc_$uid)"
	case u == UIDdirectory:
		trailer = " (directory_$uid)"
	case u == UIDunix_directory:
		trailer = " (unix_directory_$uid)"
	case u == UIDmt:
		trailer = " (mt_$uid)"
	case u == UIDsysboot:
		trailer = " (sysboot_$uid)"
	}
	return fmt.Sprintf("%04x.%04x%s", u.Lo, u.Hi, trailer)
}

// creationTime is 36 bit, nodeID is 20 bit
func Generate(creationTime uint64, nodeID uint32) (UID, error) {
	// check for bit sizes
	if creationTime > 0xfffffffff {
		return UID{}, errors.New("creationTime too large")
	}
	if nodeID > 0xfffff {
		return UID{}, errors.New("nodeID too large")
	}

	// UID layout:
	// [36-bit creationTime | 8-bit reserved | 20-bit nodeID]
	// Shift creationTime left by 28 (8+20) to put it in the top 36 bits.
	// The reserved 8 bits are left as 0.
	// The nodeID goes in the lower 20 bits.
	uid := (creationTime << 28) | (uint64(nodeID) & 0xfffff)
	return UID{
		Hi: uint32(uid >> 32),
		Lo: uint32(uid),
	}, nil
}
