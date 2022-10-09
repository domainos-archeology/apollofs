package cmd

import (
	"fmt"

	"github.com/domainos-archeology/apollofs/fs"
	"github.com/spf13/cobra"
)

func info(diskImage string) error {
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

func init() {
	rootCmd.AddCommand(infoCommand)
}

var infoCommand = &cobra.Command{
	Use:   "info",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return info(diskImage)
	},
}
