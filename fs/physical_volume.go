package fs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const MaxLogicalVolumes = 10

type PhysicalVolume struct {
	Label pvLabel
	LV    *LogicalVolume

	file *os.File
}

type pvLabel struct {
	Version             int16
	APOLLO              [6]byte
	Name                [32]byte
	UID                 UID
	Ignore1             int16
	DriveType           int16
	TotalBlocksInVolume int32
	BlocksPerTrack      int16
	TracksPerCylinder   int16
	LVDAddr             [10]int32
	AltLVLabelDAddr     [10]int32

	SectorStart int16
	SectorSize  int16
	PreComp     int16
}

func (l pvLabel) validate() error {
	if !bytes.Equal(l.APOLLO[:], []byte("APOLLO")) {
		return fmt.Errorf("expected 'APOLLO', got '%s'", string(l.APOLLO[:]))
	}

	// the first lv daddr/alt-daddr should be non-zero, but the rest should be zero
	if l.LVDAddr[0] == 0 {
		return fmt.Errorf("expected LVDAddr[0] to be non-zero")
	}
	if l.AltLVLabelDAddr[0] == 0 {
		return fmt.Errorf("expected AltLVLabelDAddr[0] to be non-zero")
	}

	for i := 1; i < MaxLogicalVolumes; i++ {
		if l.LVDAddr[i] != 0 {
			return fmt.Errorf("expected LVDAddr[%d] to be zero", i)
		}
		if l.AltLVLabelDAddr[i] != 0 {
			return fmt.Errorf("expected AltLVLabelDAddr[%d] to be zero", i)
		}
	}

	return nil
}

func Mount(diskImage string) (*PhysicalVolume, error) {
	file, err := os.Open(diskImage)
	if err != nil {
		return nil, err
	}

	pvol := &PhysicalVolume{file: file}
	block0, err := pvol.ReadBlock(0)
	if err != nil {
		return nil, err
	}

	// XXX validate the block header?
	err = block0.ReadInto(&pvol.Label)
	if err != nil {
		return nil, err
	}

	err = pvol.Label.validate()
	if err != nil {
		return nil, err
	}

	pvol.LV, err = NewLogicalVolume(pvol, pvol.Label.LVDAddr[0])
	if err != nil {
		return nil, err
	}

	return pvol, nil
}

func (pvol *PhysicalVolume) Unmount() error {
	if pvol.file != nil {
		return pvol.file.Close()
	}
	return nil
}

func (pvol *PhysicalVolume) PrintLabel() {
	fmt.Println("PV Label:")
	fmt.Printf("Version: %d\n", pvol.Label.Version)
	fmt.Printf("Name: %s\n", string(pvol.Label.Name[:]))
	fmt.Printf("UID: %s\n", pvol.Label.UID.String())
	fmt.Printf("DriveType: %d\n", pvol.Label.DriveType)
	fmt.Printf("TotalBlocksInVolume: %d\n", pvol.Label.TotalBlocksInVolume)
	fmt.Printf("BlocksPerTrack: %d\n", pvol.Label.BlocksPerTrack)
	fmt.Printf("TracksPerCylinder: %d\n", pvol.Label.TracksPerCylinder)
	fmt.Printf("SectorStart: %d\n", pvol.Label.SectorStart)
	fmt.Printf("SectorSize: %d\n", pvol.Label.SectorSize)
	fmt.Printf("PreComp: %d\n", pvol.Label.PreComp)
	fmt.Printf("Logical Volumes:\n")
	for i := 0; i < MaxLogicalVolumes; i++ {
		fmt.Printf("  LV%d: block %d / %d (alt)\n", i, pvol.Label.LVDAddr[i], pvol.Label.AltLVLabelDAddr[i])
	}
}

func (pvol *PhysicalVolume) LogicalVolumes() []int {
	var lvs []int
	for i := 0; i < MaxLogicalVolumes; i++ {
		if pvol.Label.LVDAddr[i] != 0 {
			lvs = append(lvs, i)
		}
	}
	return lvs
}

func (pvol *PhysicalVolume) ReadBlock(blockNum int32) (*Block, error) {
	_, err := pvol.file.Seek(int64(blockNum*BlockSize), io.SeekStart)
	if err != nil {
		return nil, err
	}

	var block Block
	err = binary.Read(pvol.file, binary.BigEndian, &block)
	if err != nil {
		return nil, err
	}

	return &block, nil
}
