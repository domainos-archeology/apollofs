package uid

import (
	"errors"
	"fmt"
)

type UID struct {
	Hi uint32
	Lo uint32
}

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
