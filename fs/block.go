package fs

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

const BlockDataSize = 1024
const BlockHeaderSize = 32
const BlockSize = BlockDataSize + BlockHeaderSize

type DAddr uint32
type BlockHeader struct {
	ObjectUID        UID
	PageWithinObject int32
	LastWritten      int32

	BlockSystemTypes int32
	Ignore1          int32
	Ignore2          int16
	DataChecksum     int16
	BlockDAddr       DAddr
}

func (bh *BlockHeader) BlockType() int {
	// XXX this is wrong
	return int(bh.BlockSystemTypes >> 28)
}

func (bh *BlockHeader) SystemType() int {
	// XXX this is wrong too
	return int(bh.BlockSystemTypes) >> 24 & 0xF
}

func (h *BlockHeader) Print() {
	fmt.Printf("BlockHeader:\n")
	fmt.Printf("  object uid: %s\n", h.ObjectUID)
	fmt.Printf("  Page within object: %d\n", h.PageWithinObject)
	fmt.Printf("  Block type: %d\n", h.BlockType())
	fmt.Printf("  System type: %d\n", h.SystemType())
	fmt.Printf("  Checksum: %d\n", h.DataChecksum)
	fmt.Printf("  Block DAddr: %d\n", h.BlockDAddr)
}

type Block struct {
	Header BlockHeader
	Data   [1024]byte // must match BlockDataSize above
}

func (b *Block) ReadInto(data any) error {
	return binary.Read(bytes.NewReader(b.Data[:]), binary.BigEndian, data)
}

func (b *Block) Print(includeContents bool) {
	b.Header.Print()

	if includeContents {
		fmt.Printf("Block data:\n")

		if b.Header.ObjectUID == UIDpvlabel {
			var pvlabel PVLabel
			err := b.ReadInto(&pvlabel)
			if err != nil {
				panic(err)
			}
			pvlabel.Print()
		} else if b.Header.ObjectUID == UIDlvlabel {
			var lvlabel LVLabel
			err := b.ReadInto(&lvlabel)
			if err != nil {
				panic(err)
			}
			lvlabel.Print()
		} else {
			fmt.Println(hex.Dump(b.Data[:]))
		}
	}
}
