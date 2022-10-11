package fs

import "fmt"

type VTOCHeader struct {
	Version             int16
	VTOCSizeInBlocks    int16
	VTOCBlocksUsed      int32
	NetworkRootDirVTOCX VTOCX
	DiskEntryDirVTOCX   VTOCX
	OSPagingFileVTOCX   VTOCX
	SysbootVTOCX        VTOCX
	VTOCMapData         VTOCMapData
	Unused              [28]byte
}
type VTOCMapData [48]byte

// the next two are the parsed form of VTOCMapData. check logical_volumn.go
type VTOCMap []VTOCMapExtent
type VTOCMapExtent struct {
	FirstBlockDAddr int
	NumBlocks       int
}

type VTOCBlock struct {
	NextBlockDAddr int32
	Entries        [5]VTOCE
}

type VTOCEHeader struct {
	Info             int32 // "VTOC & Object Info" in the docs
	ObjectUID        UID
	ObjectTypeDefUID UID
	ObjectACLUID     UID
	CurrentLength    int32
	BlocksUsed       int32
	LastUsedTime     int32
	LastModifiedTime int32
	DirectoryUID     UID   // "UID of Directory in which object catalogued" in the docs.  containing dir?
	MoreStuff        int32 // "DTM Ext. / Unused / Ref. Count" in the docs
	ObjectLockKey    int32
	Pad3             [4]byte
}

type VTOCE struct {
	Header      VTOCEHeader
	FileMap0    [32]DAddr // daddrs for the first 32 pages of an object
	FileMap1Ptr DAddr
	FileMap2Ptr DAddr
	FileMap3Ptr DAddr
}

type VTOCX int32

func (x VTOCX) Index() int {
	return int(x) & 0xF
}

func (x VTOCX) BlockDAddr() int32 {
	return int32(x) >> 4
}

func (x VTOCX) String() string {
	return fmt.Sprintf("vtoc blk %d, index %d", x.BlockDAddr(), x.Index())
}

func (h VTOCHeader) Print() {
	fmt.Println("VTOCHeader:")
	fmt.Println("  Version:", h.Version)
	fmt.Println("  VTOCSizeInBlocks:", h.VTOCSizeInBlocks)
	fmt.Println("  VTOCBlocksUsed:", h.VTOCBlocksUsed)
	fmt.Println("  NetworkRootDirVTOCX:", h.NetworkRootDirVTOCX)
	fmt.Println("  DiskEntryDirVTOCX:", h.DiskEntryDirVTOCX)
	fmt.Println("  OSPagingFileVTOCX:", h.OSPagingFileVTOCX)
	fmt.Println("  SysbootVTOCX:", h.SysbootVTOCX)
}
