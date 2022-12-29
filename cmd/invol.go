package cmd

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"os"
	"time"

	"github.com/domainos-archeology/apollofs/fs"
	"github.com/domainos-archeology/apollofs/util"
	"github.com/spf13/cobra"
)

var (
	sysbootPath string
	diskSpec    string
)

const (
	blocksPriam3350 = 30294
	blocksPriam6650 = 60534
)

func blocksFromSpec(spec string) (int, error) {
	switch spec {
	case "priam3350":
		return blocksPriam3350, nil
	case "priam6650":
		return blocksPriam6650, nil
	default:
		return 0, errors.New("invalid disk spec")
	}
}

func writeBlockAt(file *os.File, block fs.Block, blockNum int) error {
	_, err := file.Seek(int64(blockNum*fs.BlockSize), io.SeekStart)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.BigEndian, &block)
	if err != nil {
		return err
	}
	return nil
}

func createBlock(header fs.BlockHeader, data any) fs.Block {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, data)
	if err != nil {
		panic(err)
	}

	var blockData [1024]byte
	copy(blockData[:], buf.Bytes())

	return fs.Block{
		Header: header,
		Data:   blockData,
	}
}

func createPVLabelBlock(lvdaddr, altlvdaddr int32) fs.Block {
	return createBlock(
		fs.BlockHeader{
			ObjectUID: fs.UIDpvlabel,
		}, fs.PVLabel{
			Version:             1,
			APOLLO:              [6]byte{'A', 'P', 'O', 'L', 'L', 'O'},
			Name:                [32]byte{'A', 'P', 'O', 'L', 'L', 'O', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			UID:                 fs.UID{0x776175af, 0x10012345}, // copied from my mame image
			DriveType:           1,                              // is this DTYPE from EH87?
			TotalBlocksInVolume: blocksPriam3350,
			BlocksPerTrack:      18, // copied from EH87
			TracksPerCylinder:   3,  // same as number of heads?
			LVDAddr:             [10]int32{lvdaddr, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			AltLVLabelDAddr:     [10]int32{altlvdaddr, 0, 0, 0, 0, 0, 0, 0, 0},

			// next three copied from mame image.  no clue what they should be.
			SectorStart: 5,
			SectorSize:  2260,
			PreComp:     5,
		},
	)
}

func createLVLabelBlock(daddr fs.DAddr) fs.Block {
	return createBlock(
		fs.BlockHeader{
			ObjectUID:        fs.UIDlvlabel,
			PageWithinObject: int32(daddr) - 1,
			BlockDAddr:       daddr,
		},
		fs.LVLabel{
			Version:        1, // 1 == >= SR10
			Name:           [32]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			UID:            fs.UID{0x776175d5, 0x20012345}, // copied from my mame image
			LabelWritten:   util.TimestampToApolloEpoch(time.Now()),
			BootTime:       util.TimestampToApolloEpoch(time.Now()),
			DismountedTime: util.TimestampToApolloEpoch(time.Now()),
		},
	)
}

func copySysboot(file *os.File, sysbootPath string) error {
	sysbootFile, err := os.Open(sysbootPath)
	if err != nil {
		return err
	}
	defer sysbootFile.Close()

	// copy the sysboot file a block at a time
	sysbootBlockNum := 0
	for i := 0; i < 10; i++ {
		var block fs.Block
		block.Header = fs.BlockHeader{
			ObjectUID:        fs.UID{0x776175d5, 0x30012345},
			PageWithinObject: int32(sysbootBlockNum),
			BlockDAddr:       fs.DAddr(2 + sysbootBlockNum),
		}
		err = binary.Read(sysbootFile, binary.BigEndian, &block.Data)
		if err != nil {
			return err
		}
		err = writeBlockAt(file, block, 2+i)
		if err != nil {
			return err
		}
	}

	return nil
}

func invol(diskImage string) error {
	// create a new disk image at that path

	numBlocks, err := blocksFromSpec(diskSpec)
	if err != nil {
		return err
	}

	file, err := os.Create(diskImage)
	if err != nil {
		return err
	}
	defer file.Close()

	// initialize the image to be completely empty
	bytesToWrite := make([]byte, numBlocks*fs.BlockSize)

	_, err = file.Write(bytesToWrite)
	if err != nil {
		return err
	}

	// write our pv label
	writeBlockAt(file, createPVLabelBlock(1, 30293), 0)

	// write our lv labels
	writeBlockAt(file, createLVLabelBlock(1), 1)         // primary
	writeBlockAt(file, createLVLabelBlock(30293), 30293) // alternate

	if sysbootPath != "" {
		err = copySysboot(file, sysbootPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(involCommand)

	involCommand.Flags().StringVarP(&sysbootPath, "cpboot", "b", "", "Path to sysboot file to copy to disk image")
	involCommand.Flags().StringVarP(&diskSpec, "diskSpec", "s", "", "string specification of disk type")
}

var involCommand = &cobra.Command{
	Use:   "invol",
	Short: "Initialize a disk image",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return invol(diskImage)
	},
}
