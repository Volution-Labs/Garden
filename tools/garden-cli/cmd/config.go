package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/go-ocf/go-coap"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure on device settings.",
	Long: `Configure device setting via network (Coap) or serial connections.
First device found will be used, or use 'garden-cli device' to select device.`,
	Run: func(cmd *cobra.Command, args []string) {
		//ip, _ := cmd.Flags().GetString("serverIP")
		//interval, _ := cmd.Flags().GetString("interval")
		cmd.Flags().Visit(runThroughFlags)
		if cmd.Flags().NFlag() == 0 {
			cmd.Usage()
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringP("serverIP", "s", "", "Set the serving ip address: [ipaddress, auto]")
	configCmd.Flags().StringP("interval", "i", "", "Set the rate sensor data is sent: [milliseconds]")

}

func runThroughFlags(f *pflag.Flag) {
	stringValue := f.Value.String()
	if f.Name == "serverIP" && stringValue == "auto" {
		stringValue = findLocalIP()
	}
	setConfig(f.Name, stringValue)
}

func setConfig(key string, value string) {
	// Marshal to json
	m := make(map[string]string)
	m[key] = value
	b, err := json.Marshal(m)

	// Send payload
	co, err := coap.Dial("udp", "garden.local:5683")
	if err != nil {
		fmt.Printf("Error dialing: %v", err)
	}
	resp, err := co.Post("/config", coap.MediaType(50), bytes.NewBuffer(b))
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
	}
	if string(resp.Payload()) == "ok" {
		fmt.Printf("--%v | %v set successfully!\n", key, value)
	} else {
		fmt.Printf("--%v | Error setting %v. Response: %v\n", key, value, string(resp.Payload()))
	}
}

func findLocalIP() (address string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				address = ipnet.IP.String()
			}
		}
	}
	fmt.Printf("--serverIP | Auto: Found and using %v.\n", address)
	return
}
