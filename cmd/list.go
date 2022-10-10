package cmd

import (
	"fmt"
	"strings"

	"github.com/domainos-archeology/apollofs/fs"
	"github.com/spf13/cobra"
)

func List(diskImage string, paths []string) error {
	pvol, err := fs.Mount(diskImage)
	if err != nil {
		return err
	}
	defer pvol.Unmount()

	// let's just list the root directory for now
	lvol := pvol.LV

	vtocx := lvol.Label.VTOCHeader.SysbootVTOCX
	block, err := lvol.ReadBlock(vtocx.BlockDAddr())
	if err != nil {
		return err
	}

	fmt.Printf("block daddr: %d, vtocx daddr: %d\n", block.Header.BlockDAddr, vtocx.BlockDAddr())
	fmt.Printf("vtoc block uid: %s\n", block.Header.ObjectUID)
	fmt.Printf("page within object: %d\n", block.Header.PageWithinObject)
	fmt.Printf("block/system types: %d\n", block.Header.BlockSystemTypes)
	fmt.Printf("ignore1: %d\n", block.Header.Ignore1)
	fmt.Printf("ignore2: %d\n", block.Header.Ignore2)
	fmt.Printf("checksum: %d\n", block.Header.DataChecksum)
	var vtocBlock fs.VTOCBlock
	err = block.ReadInto(&vtocBlock)
	if err != nil {
		return err
	}

	fmt.Println("VTOC Block:")
	fmt.Printf("  next block daddr: %d\n", vtocBlock.NextBlockDAddr)
	vtoce := vtocBlock.Entries[vtocx.Index()]
	fmt.Printf("  vtoc entry[%d]:\n", vtocx.Index())
	fmt.Printf("    current length: %d\n", vtoce.Header.CurrentLength)
	fmt.Printf("    blocks used: %d\n", vtoce.Header.BlocksUsed)
	fmt.Printf("    object uid: %s\n", vtoce.Header.ObjectUID)
	fmt.Printf("    object typedef uid: %s\n", vtoce.Header.ObjectTypeDefUID)
	fmt.Printf("    object acl uid: %s\n", vtoce.Header.ObjectACLUID)
	fmt.Printf("    directory uid: %s\n", vtoce.Header.DirectoryUID)
	for i, daddr := range vtoce.FileMap0 {
		fmt.Printf("    FileMap0[%d]: %d\n", i, daddr)
	}
	fmt.Printf("    FileMap1Ptr: %d\n", vtoce.FileMap1Ptr)
	fmt.Printf("    FileMap2Ptr: %d\n", vtoce.FileMap2Ptr)
	fmt.Printf("    FileMap3Ptr: %d\n", vtoce.FileMap3Ptr)

	return pvol.Unmount()
}

func init() {
	rootCmd.AddCommand(listCommand)
}

var listCommand = &cobra.Command{
	Use:   "list path...",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("List: " + strings.Join(args, " "))
		return List(diskImage, args)
	},
}
