package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesTerminateCmdDeviceId holds value of 'device_id' option
var LoraDevicesTerminateCmdDeviceId string

func init() {
	LoraDevicesTerminateCmd.Flags().StringVar(&LoraDevicesTerminateCmdDeviceId, "device-id", "", TR("lora_devices.terminate_lora_device.post.parameters.device_id.description"))

	LoraDevicesCmd.AddCommand(LoraDevicesTerminateCmd)
}

// LoraDevicesTerminateCmd defines 'terminate' subcommand
var LoraDevicesTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TR("lora_devices.terminate_lora_device.post.summary"),
	Long:  TR(`lora_devices.terminate_lora_device.post.description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectLoraDevicesTerminateCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectLoraDevicesTerminateCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraDevicesTerminateCmd("/lora_devices/{device_id}/terminate"),
		query:  buildQueryForLoraDevicesTerminateCmd(),
	}, nil
}

func buildPathForLoraDevicesTerminateCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesTerminateCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesTerminateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
