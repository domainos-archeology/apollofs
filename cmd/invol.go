package cmd

import "github.com/spf13/cobra"

func Invol(diskImage string) error {
	// todo
	return nil
}

func init() {
	rootCmd.AddCommand(involCommand)
}

var involCommand = &cobra.Command{
	Use:   "invol",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
