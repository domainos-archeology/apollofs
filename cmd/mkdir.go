package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Mkdir(diskImage string, path string) error {
	return nil
}

func init() {
	rootCmd.AddCommand(mkdirCommand)
}

var mkdirCommand = &cobra.Command{
	Use:   "mkdir path",
	Short: "Create a directory in the disk image (similar to 'mkdir' in the host)",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Mkdir:", args[0])
		return Mkdir(diskImage, args[0])
	},
}
