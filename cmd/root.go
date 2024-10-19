package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var vers bool

// RootCmd is the root command for EventCOM
var RootCmd = &cobra.Command{
	Use:   "EventCOM",
	Short: "EventCOM is a demo project",
	Long:  "EventCOM is a demo project",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("no flag found")
	},
}
