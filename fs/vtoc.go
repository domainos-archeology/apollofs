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

// the next two are the parsed form of VTOCMapData. check logical_volume.go
type VTOCMap []VTOCMapExtent
type VTOCMapExtent struct {
	NumBlocks       int16
	FirstBlockDAddr DAddr
}

type VTOCBlock struct {
	NextBlockDAddr DAddr
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

func (e VTOCE) Print() {
	fmt.Println("VTOCE:")
	fmt.Println("  Info:", e.Header.Info)
	fmt.Println("  ObjectUID:", e.Header.ObjectUID)
	fmt.Println("  ObjectTypeDefUID:", e.Header.ObjectTypeDefUID)
	fmt.Println("  ObjectACLUID:", e.Header.ObjectACLUID)
	fmt.Println("  CurrentLength:", e.Header.CurrentLength)
	fmt.Println("  BlocksUsed:", e.Header.BlocksUsed)
	fmt.Println("  LastUsedTime:", e.Header.LastUsedTime)
	fmt.Println("  LastModifiedTime:", e.Header.LastModifiedTime)
	fmt.Println("  DirectoryUID:", e.Header.DirectoryUID)
	fmt.Println("  MoreStuff:", e.Header.MoreStuff)
	fmt.Println("  ObjectLockKey:", e.Header.ObjectLockKey)
	fmt.Println("  FileMap0:")
	for _, daddr := range e.FileMap0 {
		fmt.Println("    ", daddr)
	}
	fmt.Println("  FileMap1Ptr:", e.FileMap1Ptr)
	fmt.Println("  FileMap2Ptr:", e.FileMap2Ptr)
	fmt.Println("  FileMap3Ptr:", e.FileMap3Ptr)
}

func (m VTOCMap) Print() {
	for i, extent := range m {
		fmt.Printf("  extent %d: %d blocks at %d\n", i, extent.NumBlocks, extent.FirstBlockDAddr)
	}
}
