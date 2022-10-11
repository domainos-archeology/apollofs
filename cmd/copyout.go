package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CopyOut(diskImage string, paths []string) error {
	return fmt.Errorf("command not implemented")
}

func init() {
	rootCmd.AddCommand(copyOutCommand)
}

var copyOutCommand = &cobra.Command{
	Use:   "copyout",
	Short: "Copy a file from the disk image to the host",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
