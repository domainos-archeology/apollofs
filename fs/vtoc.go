package fs

import "fmt"

func notImplemented() {
	panic("not implemented")
}

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
type VTOCMapDataSR9 [48]byte
type VTOCMapData [60]byte

// the next two are the parsed form of VTOCMapData. check logical_volume.go
type VTOCMap []VTOCMapExtent
type VTOCMapExtent struct {
	NumBlocks       uint16
	FirstBlockDAddr DAddr
}

type VTOCBlockSR9 struct {
	NextBlockDAddr DAddr
	Entries        [5]VTOCESR9
}

type VTOCBlock struct {
	Padding        uint32 // not sure what this is, but VTOCE_$READ uses 8 bytes to offset to the entries instead of 4
	NextBlockDAddr DAddr
	Entries        [3]VTOCE
}

// this seem to be a struct in SR10 only, and not present in SR9?
type VTOCBucketBlock struct {
	Buckets [3]VTOCBucket
}

type VTOCBucket struct {
	// bucket header (next daddr/index)
	NextBlockDAddr DAddr
	NextIndex      uint16
	Padding        uint16

	// bucket entries
	Entries [20]VTOCBucketEntry
}

type VTOCBucketEntry struct {
	UID   UID
	VTOCX VTOCX
}

type VTOCEHeaderSR9 struct {
	// "VTOCE & Object Info" in the docs
	Version    uint8
	SystemType uint8
	Flags      uint16

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

	Padding [4]byte // Aegis Internals doesn't have this, but EH'87 does.
}

type VTOCEHeader struct {
	Version          uint8
	SystemType       uint8
	Flags            uint16
	ObjectUID        UID
	ObjectTypeDefUID UID
	CurrentLength    uint32
	BlocksUsed       uint32
	LastModifiedTime uint32
	ExtDtm           uint16
	RefCount         uint16
	LastUsedTime     uint32
	U5               uint32
	U6               uint32
	U7               uint32
	U8               uint32
	U9               uint32
	DirectoryUID     UID
	ObjectLockKey    uint32
	ObjectUserUID    UID
	ObjectGroupUID   UID
	ObjectOrgUID     UID
	U19              uint32
	U20              uint32
	U21              uint32
	U22              uint32
	U23              uint32
	U24              uint32
	U25              uint32
	U26              uint32
	U27              uint32
	U28              uint32
	ObjectACLUID     UID
	Padding          [64]byte
}

type VTOCESR9 struct {
	Header      VTOCEHeaderSR9
	FileMap0    [32]DAddr
	FileMap1Ptr DAddr
	FileMap2Ptr DAddr
	FileMap3Ptr DAddr
}

type VTOCE struct {
	Header      VTOCEHeader
	FileMap0    [29]DAddr // daddrs for the first 32 pages of an object
	FileMap1Ptr DAddr
	FileMap2Ptr DAddr
	FileMap3Ptr DAddr
}

type VTOCX uint32

func NewVTOCX(blockDAddr DAddr, index int) VTOCX {
	return VTOCX(blockDAddr<<4 + DAddr(index))
}

func (x VTOCX) Index() int {
	return int(uint32(x) & 0xF)
}

func (x VTOCX) BlockDAddr() DAddr {
	return DAddr(uint32(x) >> 4)
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

func (b VTOCBucketBlock) Print() {
	fmt.Println("VTOCBucketBlock:")
	for i, bucket := range b.Buckets {
		fmt.Printf("  Bucket %d:\n", i)
		fmt.Println("    NextBlockDAddr:", bucket.NextBlockDAddr)
		fmt.Println("    NextIndex:", bucket.NextIndex)

		for j, entry := range bucket.Entries {
			if entry.VTOCX == 0 {
				continue
			}
			fmt.Printf("    Entry %d:\n", j)
			fmt.Println("      UID:", entry.UID)
			fmt.Println("      VTOCX:", entry.VTOCX)
		}
	}
}

func (b VTOCBlock) Print() {
	fmt.Println("VTOCBlock:")
	fmt.Println(" NextBlockDAddr:", b.NextBlockDAddr)
	for i, e := range b.Entries {
		fmt.Printf(" VTOCE %d:\n", i)
		e.Print()
	}
}

func (e VTOCE) IsDirectory() bool {
	return e.Header.ObjectTypeDefUID == UIDdirectory ||
		e.Header.ObjectTypeDefUID == UIDunix_directory ||
		e.Header.SystemType == 2 ||
		e.Header.SystemType == 1
}

func (e VTOCE) Print() {
	fmt.Println("  Info: version", e.Header.Version, "systype", e.Header.SystemType, "flags", e.Header.Flags)
	fmt.Println("  ObjectUID:", e.Header.ObjectUID)
	fmt.Println("  ObjectTypeDefUID:", e.Header.ObjectTypeDefUID)
	fmt.Println("  ObjectACLUID:", e.Header.ObjectACLUID)
	fmt.Println("  CurrentLength:", e.Header.CurrentLength)
	fmt.Println("  BlocksUsed:", e.Header.BlocksUsed)
	fmt.Println("  LastUsedTime:", e.Header.LastUsedTime)
	fmt.Println("  LastModifiedTime:", e.Header.LastModifiedTime)
	fmt.Println("  DirectoryUID:", e.Header.DirectoryUID)
	// fmt.Println("  MoreStuff:", e.Header.MoreStuff)
	fmt.Println("  ObjectLockKey:", e.Header.ObjectLockKey)
	fmt.Println("  FileMap0:")
	for i, daddr := range e.FileMap0 {
		if daddr == 0 {
			continue
		}
		fmt.Println("    map idx:", i, " daddr:", daddr)
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

type VTOCManager struct {
	lvol *LogicalVolume
}

func NewVTOCManager(lvol *LogicalVolume) *VTOCManager {
	return &VTOCManager{
		lvol: lvol,
	}
}

func (vm *VTOCManager) AllocateEntry(uid UID) (VTOCX, error) {
	// hash the uid, locate the appropriate vtoc block, checking for a free vtoce there.

	// if there isn't a free vtoce, call BATManager.AllocateBlock to create a new vtoc extension block and add it to the chain.

	// now that we have a vtoc block with a vtoce there, fill it in and return a vtocx referring to it.

	notImplemented()
	return 0, nil
}

func (vm *VTOCManager) LookupEntry(uid UID) (*VTOCE, error) {
	notImplemented()
	return nil, nil
}

func (vm *VTOCManager) GetEntry(vtocx VTOCX) (*VTOCE, error) {
	vtoceAddr := vtocx.BlockDAddr()
	vtoceIndex := vtocx.Index()

	// read the block
	block, err := vm.lvol.ReadBlock(vtoceAddr)
	if err != nil {
		return nil, err
	}

	// parse the block
	var vtocBlock VTOCBlock
	err = block.ReadInto(&vtocBlock)
	if err != nil {
		return nil, err
	}

	// return the entry
	return &vtocBlock.Entries[vtoceIndex], nil
}

func (vm *VTOCManager) GetEntryForUID(uid UID) (*VTOCE, error) {
	vtocx, err := vm.GetIndexForUID(uid)
	if err != nil {
		return nil, err
	}

	// fmt.Println("index = ", vtocx)
	return vm.GetEntry(vtocx)
}

func (vm *VTOCManager) GetIndexForUID(uid UID) (VTOCX, error) {
	// fmt.Println(1)
	// hash the uid, locate the appropriate vtoc block, then follow the chain
	// until we find the block+index
	hashValue := vm.hashUID(uid)

	startingBlockNumber := hashValue >> 2
	idx := hashValue & 0x3

	// XXX assume it's in the first extent for now (this seems likely for
	// filesystems involed in MAME, since there are no bad spots to break up the
	// vtoc)
	if vm.lvol.VTOCMap[1].NumBlocks > 0 || startingBlockNumber > uint32(vm.lvol.VTOCMap[0].NumBlocks) {
		panic("there are more blocks someplace else.  our assumption no longer holds!")
	}

	blockDAddr := DAddr(startingBlockNumber) + vm.lvol.VTOCMap[0].FirstBlockDAddr

	// fmt.Println("starting search at block", blockDAddr, "and idx", idx)

	for blockDAddr != 0 {
		// fmt.Println("looking in block", blockDAddr)
		// fmt.Println(2)
		// read the vtoc bkt block
		block, err := vm.lvol.ReadBlock(blockDAddr)
		if err != nil {
			return 0, err
		}

		// fmt.Println(3)
		var vtocBucketBlock VTOCBucketBlock
		err = block.ReadInto(&vtocBucketBlock)
		if err != nil {
			return 0, err
		}

		// fmt.Println(4)
		// check the entries in bucket `idx`
		for _, entry := range vtocBucketBlock.Buckets[idx].Entries {
			if entry.UID == uid {
				// fmt.Println(4.1)
				return entry.VTOCX, nil
			}
		}

		// fmt.Println(5)
		// if we didn't find it, follow the chain and keep looking
		blockDAddr = vtocBucketBlock.Buckets[idx].NextBlockDAddr
		idx = uint32(vtocBucketBlock.Buckets[idx].NextIndex)
	}

	return 0, errNotFound
}

func (vm *VTOCManager) hashUID(uid UID) uint32 {
	vtocVersion := vm.lvol.Label.VTOCHeader.Version
	vtocSize := uint32(vm.lvol.Label.VTOCHeader.VTOCSizeInBlocks)

	switch {
	case vtocVersion < 2:
		// below translitered from domain_os decompiled UID_$HASH
		var tmp uint32

		tmp = uid.Hi ^ uid.Lo
		tmp = tmp ^ (tmp >> 16)

		return (tmp/vtocSize)<<16 | (tmp % vtocSize)

	case vtocVersion == 2:
		u := uid.Hi * 2
		if int32(uid.Lo) < 0 {
			u = u + 1
		}
		return (u>>0x10 ^ u&0xffff) % vtocSize

	case vtocVersion == 3:
		return uint32(uint16(uid.Hi>>0x10)^uint16(uid.Hi)) % vtocSize
	default:
		panic("unhandled vtoc version")
	}
}
