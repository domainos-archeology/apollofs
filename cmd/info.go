package cmd

import (
	"fmt"
	"strconv"

	"github.com/domainos-archeology/apollofs/fs"
	"github.com/spf13/cobra"
)

func labels(diskImage string) error {
	pvol, err := fs.Mount(diskImage)
	if err != nil {
		return err
	}
	defer pvol.Unmount()

	fmt.Println("Disk image:", diskImage)
	pvol.PrintLabel()
	fmt.Println()
	pvol.LV.PrintLabel()

	return pvol.Unmount()
}

func vtoc(diskImage string) error {
	pvol, err := fs.Mount(diskImage)
	if err != nil {
		return err
	}
	defer pvol.Unmount()

	fmt.Println("Disk image:", diskImage)
	pvol.LV.PrintLabel()

	return pvol.Unmount()
}

func block(diskImage string, physDAddr string) error {
	daddr, err := strconv.Atoi(physDAddr)
	if err != nil {
		return err
	}

	pvol, err := fs.Mount(diskImage)
	if err != nil {
		return err
	}
	defer pvol.Unmount()

	fmt.Println("Disk image:", diskImage)

	block, err := pvol.ReadBlock(int32(daddr))
	if err != nil {
		return err
	}

	block.Print()

	return pvol.Unmount()
}

func init() {
	rootCmd.AddCommand(infoCommand)

	infoCommand.AddCommand(labelsCommand)
	infoCommand.AddCommand(vtocCommand)
	infoCommand.AddCommand(blockCommand)
}

var infoCommand = &cobra.Command{
	Use:   "info [labels|vtoc|block]",
	Short: "Dump information about filesystem structures",
	Long:  "",
}

var labelsCommand = &cobra.Command{
	Use:   "labels",
	Short: "dump info about physical and logical volume labels",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return labels(diskImage)
	},
}

var vtocCommand = &cobra.Command{
	Use:   "vtoc",
	Short: "dump info about logical volume VTOC",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return vtoc(diskImage)
	},
}

var blockCommand = &cobra.Command{
	Use:   "block [physDAddr]",
	Short: "dump info about a disk block, given its physical address",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return block(diskImage, args[0])
	},
}
