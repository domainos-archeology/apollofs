package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CPboot(diskImage string) error {
	return fmt.Errorf("command not implemented")
}

func init() {
	rootCmd.AddCommand(cpbootCommand)
}

var cpbootCommand = &cobra.Command{
	Use:   "cpboot",
	Short: "Make the disk image bootable (by copying sysboot)",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
