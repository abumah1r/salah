package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "salah",
	Short: "Display salah times",
	Long:  `Display salah times beautifully from your terminal.`,
}

func Execute() error {
	return rootCmd.Execute()
}
