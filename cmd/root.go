package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "apollofs",
		Short: "A tool for interacting with Apollo filesystems",
		Long:  "",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var diskImage string

func init() {
	cobra.OnInitialize()

	rootCmd.AddCommand(copyInCommand)
	rootCmd.PersistentFlags().StringVarP(&diskImage, "diskImage", "d", "", "Path to disk image (required)")
}
