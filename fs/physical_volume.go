package fs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const MaxLogicalVolumes = 10
const BlockSize = 1024

type BlockHeader struct {
	UID              int64
	PageWithinObject int32
	LastWritten      int32

	BlockTypeSystemTypeEtc int32
	Ignore1                int32
	Ignore2                int16
	DataChecksum           int16
	BlockDAddr             int32
}

type PhysicalVolume struct {
	Label pvLabel

	file *os.File
}

type pvLabel struct {
	Version             int16
	APOLLO              [6]byte
	Name                [32]byte
	UID                 uint64
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
	return nil
}

func (l pvLabel) Info() {
	fmt.Println("PV Label:")
	fmt.Printf("Version: %d\n", l.Version)
	fmt.Printf("Name: %s\n", string(l.Name[:]))
	fmt.Printf("UID: %d\n", l.UID) // need a better way to format these
	fmt.Printf("DriveType: %d\n", l.DriveType)
	fmt.Printf("TotalBlocksInVolume: %d\n", l.TotalBlocksInVolume)
	fmt.Printf("BlocksPerTrack: %d\n", l.BlocksPerTrack)
	fmt.Printf("TracksPerCylinder: %d\n", l.TracksPerCylinder)
	fmt.Printf("SectorStart: %d\n", l.SectorStart)
	fmt.Printf("SectorSize: %d\n", l.SectorSize)
	fmt.Printf("PreComp: %d\n", l.PreComp)
	fmt.Printf("Logical Volumes:\n")
	for i := 0; i < MaxLogicalVolumes; i++ {
		fmt.Printf("  LV%d: block %d / %d (alt)\n", i, l.LVDAddr[i], l.AltLVLabelDAddr[i])
	}
}

func Mount(diskImage string) (*PhysicalVolume, error) {
	file, err := os.Open(diskImage)
	if err != nil {
		return nil, err
	}

	pvol := &PhysicalVolume{file: file}
	_ /*header0*/, block0, err := pvol.readBlock(0)
	if err != nil {
		return nil, err
	}

	// XXX validate the block header?
	err = binary.Read(bytes.NewReader(block0), binary.BigEndian, &pvol.Label)
	if err != nil {
		return nil, err
	}

	err = pvol.Label.validate()
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

func (pvol *PhysicalVolume) readBlock(blockNum int) (*BlockHeader, []byte, error) {
	_, err := pvol.file.Seek(int64(blockNum*(BlockSize+32)), io.SeekStart)
	if err != nil {
		return nil, nil, err
	}

	var header BlockHeader
	err = binary.Read(pvol.file, binary.BigEndian, &header)
	if err != nil {
		return nil, nil, err
	}

	block := make([]byte, BlockSize)
	n, err := pvol.file.Read(block)
	if err != nil {
		panic("failed read")
	}
	if n != BlockSize {
		panic("short read")
	}

	return &header, block, nil
}
