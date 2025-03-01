package drives

import "errors"

type DriveType struct {
	DType          int16
	Name           string
	Cylinders      int16
	Heads          int16
	BlocksPerTrack int16
}

var driveTypeList = []DriveType{
	// Winchester Dtype Class 000 -- 14" Ring/disk PRIAM interface
	{0x001, "PRIAM 3350", 561, 3, 18},
	{0x006, "PRIAM 6650", 1121, 3, 18},
	{0x007, "PRIAM 15450", 1121, 7, 18},

	// Winchester Dtype Class 100 -- 8" ANSI Interface
	{0x103, "Micropolis 1203", 580, 5, 13}, // software only uses 525 cylinders and 12 blocks/track
	{0x104, "PRIAM 3450", 525, 5, 12},
	{0x105, "PRIAM 7050", 1049, 5, 12},

	// Wincherster Dtype Class 200 -- 8" Ring/disk SMD Interface
	{0x201, "NEC D2257", 1024, 8, 18},
	{0x202, "NEC D2246", 687, 6, 18},

	// Winchester Dtype Class 300 -- 5 1/4" ST412 Interface
	{0x301, "Micropolis 50MB", 830, 6, 8},
	{0x302, "Micropolis 86MB", 1024, 8, 8},
	{0x303, "Fujitsu 86MB", 754, 11, 8},
	{0x304, "Maxtor 140MB", 918, 15, 8},
	{0x305, "Maxtor 190MB", 1224, 15, 8},
	{0x306, "Vertex 86MB", 1166, 7, 8},

	// Winchester Dtype Class 400 -- 5 1/4" ST412 Interface
	{0x401, "Vertex 50MB", 987, 5, 9},
	{0x402, "Vertex 86MB", 1166, 7, 9},
	{0x405, "Micropolis 50MB", 1024, 5, 9},
	{0x406, "Micropolis 86MB", 1024, 8, 9},
	{0x410, "Micropolis 50MB", 830, 6, 9},
	{0x411, "Maxtor 190MB", 1224, 15, 9},

	// Winchester Dtype Class 500 -- 5 1/4" ESDI Interface
	{0x503, "Priam/Maxtor 170MB", 1224, 7, 18},
	{0x504, "Priam/Maxtor 380MB", 1224, 15, 18},
	{0x507, "Micropolis 170MB", 1024, 8, 18},

	// There's more in EH87, but I'm not going to bother with them for now
}

var driveTypes map[int16]DriveType

func getDriveTypes() map[int16]DriveType {
	if driveTypes != nil {
		return driveTypes
	}

	driveTypes = make(map[int16]DriveType)
	for _, dt := range driveTypeList {
		driveTypes[dt.DType] = dt
	}

	return driveTypes
}

func GetDriveType(dtype int16) (DriveType, error) {
	dt, ok := getDriveTypes()[dtype]
	if !ok {
		return DriveType{}, errors.New("unknown drive type")
	}
	return dt, nil
}

func (dt DriveType) TotalBlocks() int32 {
	return int32(dt.Cylinders) * int32(dt.Heads) * int32(dt.BlocksPerTrack)
}

func GetDriveTypes() []DriveType {
	return driveTypeList[:]
}
