package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "server",
	Short:        "Setting Server",
	SilenceUsage: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to IVY Backend Server!")
	},
}

func init() {
	rootCmd.AddCommand(StartDevCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
