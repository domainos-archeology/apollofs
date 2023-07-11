package cmd

import (
	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	diskImage string
	debug     bool

	rootCmd = &cobra.Command{
		Use:   "apollofs",
		Short: "A tool for interacting with Apollo filesystems",
		Long:  "",
	}
)

// Execute executes the root command
func Execute() error {
	level := logrus.WarnLevel
	if debug {
		level = logrus.DebugLevel
	}
	logrus.SetLevel(level)

	return rootCmd.Execute()
}

func toggleDebug(cmd *cobra.Command, args []string) {
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringVarP(&diskImage, "diskImage", "i", "", "Path to disk image (required)")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug output")
}
