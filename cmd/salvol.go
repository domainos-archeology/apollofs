package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Salvol(diskImage string, dest string, paths []string) error {
	return fmt.Errorf("command not implemented")
}

func init() {
	rootCmd.AddCommand(salvolCommand)
}

var salvolCommand = &cobra.Command{
	Use:    "salvol",
	Short:  "Salvage a disk image",
	Long:   ``,
	PreRun: toggleDebug,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
