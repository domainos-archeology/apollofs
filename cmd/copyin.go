package cmd

import (
	"github.com/spf13/cobra"
)

func CopyIn(diskImage string, dest string, paths []string) error {
	// todo
	return nil
}

func init() {
	rootCmd.AddCommand(copyInCommand)
}

var copyInCommand = &cobra.Command{
	Use:   "copyin",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
