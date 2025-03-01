package uid

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
