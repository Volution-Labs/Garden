package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "garden-cli",
	Short: "Cli tool to config and debug garden devices.",
	Long:  `Cli tool to config and debug garden devices.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() {
	// err := doc.GenMarkdownTree(rootCmd, "/tmp")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
