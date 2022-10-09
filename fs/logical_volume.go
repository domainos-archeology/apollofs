package fs

type lvLabel struct {
	Version         int16
	Ignore1         int16
	Name            [32]byte
	ID              int64
	BATHeader       BATHeader
	VTOCHeader      VTOCHeader
	LabelWritten    int32 // time LV label writtern
	Ignore2         int16
	LastMountedNode int16
	BootTime        int32
	DismountedTime  int32
}
