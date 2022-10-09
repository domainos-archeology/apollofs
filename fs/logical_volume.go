package fs

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/domainos-archeology/apollofs/util"
)

type LogicalVolume struct {
	pvol  *PhysicalVolume
	Label lvLabel
}

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

func NewLogicalVolume(pvol *PhysicalVolume, labelDaddr int32) (*LogicalVolume, error) {
	lvol := &LogicalVolume{pvol: pvol}

	block, err := pvol.readBlock(labelDaddr)
	if err != nil {
		return nil, err
	}

	// XXX validate the block header?
	err = binary.Read(bytes.NewReader(block.Data[:]), binary.BigEndian, &lvol.Label)
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
	fmt.Printf("ID: %d\n", lvol.Label.ID)
	fmt.Printf("Label Written: %s\n", util.FormatTimestamp(lvol.Label.LabelWritten))
	fmt.Printf("Boot Time: %s\n", util.FormatTimestamp(lvol.Label.BootTime))
	fmt.Printf("Dismounted Time: %s\n", util.FormatTimestamp(lvol.Label.DismountedTime))
}
