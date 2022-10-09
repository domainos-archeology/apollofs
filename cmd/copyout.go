package cmd

import "github.com/spf13/cobra"

func CopyOut(diskImage string, paths []string) error {
	// todo
	return nil
}

func init() {
	rootCmd.AddCommand(copyOutCommand)
}

var copyOutCommand = &cobra.Command{
	Use:   "copyout",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
