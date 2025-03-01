package managers

import (
	"github.com/domainos-archeology/apollofs/pkg/fs"
	"github.com/domainos-archeology/apollofs/pkg/uid"
	"golang.org/x/exp/slices"
)

type VTOCManager struct {
	lvol *fs.LogicalVolume
}

func NewVTOCManager(lvol *fs.LogicalVolume) *VTOCManager {
	return &VTOCManager{
		lvol,
	}
}

func (vm *VTOCManager) AllocateEntry(header fs.VTOCEHeader) (fs.VTOCX, error) {
	// hash the uid, locate the appropriate vtoc block, checking for a free vtoce there.
	// hashValue := vm.hashUID(header.ObjectUID)

	// startingBlockNumber := hashValue >> 2
	// idx := hashValue & 0x3

	// if there isn't a free vtoce, call BATManager.AllocateBlock to create a new vtoc extension block and add it to the chain.

	// now that we have a vtoc block with a vtoce there, fill it in and return a vtocx referring to it.

	notImplemented()
	return 0, nil
}

func (vm *VTOCManager) LookupEntry(u uid.UID) (*fs.VTOCE, error) {
	notImplemented()
	return nil, nil
}

func (vm *VTOCManager) GetEntry(vtocx fs.VTOCX) (*fs.VTOCE, error) {
	vtoceAddr := vtocx.BlockDAddr()
	vtoceIndex := vtocx.Index()

	// read the block
	block, err := vm.lvol.ReadBlock(vtoceAddr)
	if err != nil {
		return nil, err
	}

	// parse the block
	var vtocBlock fs.VTOCBlock
	err = block.ReadInto(&vtocBlock)
	if err != nil {
		return nil, err
	}

	// return the entry
	return &vtocBlock.Entries[vtoceIndex], nil
}

func (vm *VTOCManager) GetEntryForUID(u uid.UID) (*fs.VTOCE, error) {
	vtocx, err := vm.GetIndexForUID(u)
	if err != nil {
		return nil, err
	}

	// fmt.Println("index = ", vtocx)
	return vm.GetEntry(vtocx)
}

func (vm *VTOCManager) GetIndexForUID(u uid.UID) (fs.VTOCX, error) {
	// fmt.Println(1)
	// hash the uid, locate the appropriate vtoc block, then follow the chain
	// until we find the block+index
	hashValue := vm.hashUID(u)

	startingBlockNumber := hashValue >> 2
	idx := hashValue & 0x3

	// XXX assume it's in the first extent for now (this seems likely for
	// filesystems involed in MAME, since there are no bad spots to break up the
	// vtoc)
	if vm.lvol.VTOCMap[1].NumBlocks > 0 || startingBlockNumber > uint32(vm.lvol.VTOCMap[0].NumBlocks) {
		panic("there are more blocks someplace else.  our assumption no longer holds!")
	}

	blockDAddr := fs.DAddr(startingBlockNumber) + vm.lvol.VTOCMap[0].FirstBlockDAddr

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
		var vtocBucketBlock fs.VTOCBucketBlock
		err = block.ReadInto(&vtocBucketBlock)
		if err != nil {
			return 0, err
		}

		// fmt.Println(4)
		// check the entries in bucket `idx`
		for _, entry := range vtocBucketBlock.Buckets[idx].Entries {
			if entry.UID == u {
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

func (vm VTOCManager) GetFMBlocks(e *fs.VTOCE) ([]fs.DAddr, error) {
	var blocks []fs.DAddr

	for _, daddr := range e.FileMap0 {
		if daddr == 0 {
			break
		}
		blocks = append(blocks, daddr)
	}

	if e.FileMap1Ptr != 0 {
		fm1Blocks, err := vm.getFMBlocks(1, e.FileMap1Ptr)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, fm1Blocks...)
	}
	if e.FileMap2Ptr != 0 {
		fm2Blocks, err := vm.getFMBlocks(2, e.FileMap2Ptr)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, fm2Blocks...)
	}
	if e.FileMap3Ptr != 0 {
		fm3Blocks, err := vm.getFMBlocks(3, e.FileMap3Ptr)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, fm3Blocks...)
	}
	return blocks, nil
}

func (vm VTOCManager) getFMBlocks(level int, blockDAddr fs.DAddr) ([]fs.DAddr, error) {
	if level == 0 || level > 3 {
		panic("out of range file map level")
	}

	block, err := vm.lvol.ReadBlock(blockDAddr)
	if err != nil {
		return nil, err
	}

	var daddrs [256]fs.DAddr
	err = block.ReadInto(&daddrs)
	if err != nil {
		return nil, err
	}

	blocks := daddrs[:]
	// truncate at the first 0 daddr
	zeroIdx := slices.Index(blocks, fs.DAddr(0))
	if zeroIdx != -1 {
		blocks = blocks[:zeroIdx]
	}

	if level == 1 {
		// these daddrs are actual file blocks.  no further recursion
		return blocks, nil
	}

	// otherwise we get our list of daddrs and recurse here
	var rvBlocks []fs.DAddr
	for _, daddr := range blocks {
		fmBlocks, err := vm.getFMBlocks(level-1, daddr)
		if err != nil {
			return nil, err
		}
		rvBlocks = append(rvBlocks, fmBlocks...)
	}

	return rvBlocks, nil
}

func (vm *VTOCManager) hashUID(uid uid.UID) uint32 {
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

func notImplemented() {
	panic("not implemented")
}
