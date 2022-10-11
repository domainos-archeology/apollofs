package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Invol(diskImage string) error {
	return fmt.Errorf("command not implemented")
}

func init() {
	rootCmd.AddCommand(involCommand)
}

var involCommand = &cobra.Command{
	Use:   "invol",
	Short: "Initialize a disk image",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
