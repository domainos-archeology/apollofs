package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CopyIn(diskImage string, dest string, paths []string) error {
	return fmt.Errorf("command not implemented")
}

func init() {
	rootCmd.AddCommand(copyInCommand)
}

var copyInCommand = &cobra.Command{
	Use:    "copyin",
	Short:  "Copy a file from the host to the disk image",
	Long:   ``,
	PreRun: toggleDebug,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
