package cmd

import (
	"github.com/abumahir/salah/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "salah",
	Short: "Display salah times",
	Long:  `Display salah times beautifully from your terminal.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Init()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
