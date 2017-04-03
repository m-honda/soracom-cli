package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesGetCmdDeviceId holds value of 'device_id' option
var LoraDevicesGetCmdDeviceId string

func init() {
	LoraDevicesGetCmd.Flags().StringVar(&LoraDevicesGetCmdDeviceId, "device-id", "", TR("lora_devices.get_lora_device.get.parameters.device_id.description"))

	LoraDevicesCmd.AddCommand(LoraDevicesGetCmd)
}

// LoraDevicesGetCmd defines 'get' subcommand
var LoraDevicesGetCmd = &cobra.Command{
	Use:   "get",
	Short: TR("lora_devices.get_lora_device.get.summary"),
	Long:  TR(`lora_devices.get_lora_device.get.description`),
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

		param, err := collectLoraDevicesGetCmdParams()
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

func collectLoraDevicesGetCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraDevicesGetCmd("/lora_devices/{device_id}"),
		query:  buildQueryForLoraDevicesGetCmd(),
	}, nil
}

func buildPathForLoraDevicesGetCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesGetCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
