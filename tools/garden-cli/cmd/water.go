package cmd

import (
	"fmt"
	"strings"

	"github.com/go-ocf/go-coap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(waterCmd)
	waterCmd.Flags().StringP("valve", "v", "1", "Specify valve: [number]")
}

var waterCmd = &cobra.Command{
	Use:   "water [argument]",
	Short: "Sends a water on, or off command to a valve device.",
	Long: `Sends a water on, or off command to a valve device. 
Argument [time in milliseconds] for on. [off] to cancel watering.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sendWaterCommand(args[0])
	},
}

func sendWaterCommand(state string) {
	co, err := coap.Dial("udp", "garden.local:5683")
	if err != nil {
		fmt.Printf("Error dialing: %v", err)
	}
	resp, err := co.Post("/water", coap.MediaType(50), strings.NewReader(state))
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
	}
	if string(resp.Payload()) == "ok" {
		fmt.Printf("Valve set to %v\n", state)
	} else {
		fmt.Printf("Error while setting %v | Response: %v\n", state, string(resp.Payload()))
	}
}
