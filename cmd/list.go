package cmd

import (
	"fmt"
	"strings"

	"github.com/domainos-archeology/apollofs/fs"
	"github.com/spf13/cobra"
)

func List(diskImage string, paths []string) error {
	pvol, err := fs.Mount(diskImage)
	if err != nil {
		return err
	}
	defer pvol.Unmount()

	// todo

	return pvol.Unmount()
}

func init() {
	rootCmd.AddCommand(listCommand)
}

var listCommand = &cobra.Command{
	Use:   "list path...",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("List: " + strings.Join(args, " "))
		return List(diskImage, args)
	},
}
