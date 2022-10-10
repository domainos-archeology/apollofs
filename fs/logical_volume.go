package fs

import (
	"fmt"

	"github.com/domainos-archeology/apollofs/util"
)

type LogicalVolume struct {
	pvol       *PhysicalVolume
	startDAddr int32
	Label      lvLabel
}

type lvLabel struct {
	Version         int16
	Ignore1         int16
	Name            [32]byte
	ID              UID
	BATHeader       BATHeader
	VTOCHeader      VTOCHeader
	LabelWritten    uint32 // time LV label writtern
	Ignore2         int16
	LastMountedNode int16
	BootTime        uint32
	DismountedTime  uint32
}

func NewLogicalVolume(pvol *PhysicalVolume, startDAddr int32) (*LogicalVolume, error) {
	lvol := &LogicalVolume{
		pvol:       pvol,
		startDAddr: startDAddr,
	}

	block, err := pvol.ReadBlock(startDAddr)
	if err != nil {
		return nil, err
	}

	fmt.Printf("lvlabel block uid: %s\n", block.Header.ObjectUID)

	// XXX validate the block header?
	err = block.ReadInto(&lvol.Label)
	if err != nil {
		return nil, err
	}

	// not yet
	// err = lvol.Label.validate()
	// if err != nil {
	// 	return nil, err
	// }

	return lvol, nil
}

func (lvol *LogicalVolume) PrintLabel() {
	fmt.Println("LV Label:")
	fmt.Printf("Version: %d\n", lvol.Label.Version)
	fmt.Printf("Name: %s\n", string(lvol.Label.Name[:]))
	fmt.Printf("ID: %s\n", lvol.Label.ID)
	fmt.Printf("Label Written: %s\n", util.FormatTimestamp(lvol.Label.LabelWritten))
	fmt.Printf("Boot Time: %s\n", util.FormatTimestamp(lvol.Label.BootTime))
	fmt.Printf("Dismounted Time: %s\n", util.FormatTimestamp(lvol.Label.DismountedTime))
	lvol.Label.VTOCHeader.Print()
}

func (lvol *LogicalVolume) ReadBlock(blockNum int32) (*Block, error) {
	return lvol.pvol.ReadBlock(lvol.startDAddr + blockNum)
}
