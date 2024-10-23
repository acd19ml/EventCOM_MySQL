package cmd

import (
	"fmt"

	"github.com/acd19ml/EventCOM_MySQL/version"
	"github.com/spf13/cobra"
)

var vers bool

// RootCmd is the root command for EventCOM
var RootCmd = &cobra.Command{
	Use:   "EventCOM",
	Short: "EventCOM is a demo project",
	Long:  "EventCOM is a demo project",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "print EventCOM version")
}
