package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(emulateCmd)
}

var emulateCmd = &cobra.Command{
	Use:   "emulate",
	Short: "Emulate a device",
	Long: `Will emulate functions of the device by sending fake sensors data,
	 and respond to watering functions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Device emulation unavailable")
	},
}
