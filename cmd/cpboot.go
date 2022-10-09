package cmd

import "github.com/spf13/cobra"

func CPboot(diskImage string) error {
	// todo
	return nil
}

func init() {
	rootCmd.AddCommand(cpbootCommand)
}

var cpbootCommand = &cobra.Command{
	Use:   "cpboot",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
