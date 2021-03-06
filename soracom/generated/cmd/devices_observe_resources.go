// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DevicesObserveResourcesCmdDeviceId holds value of 'device_id' option
var DevicesObserveResourcesCmdDeviceId string

// DevicesObserveResourcesCmdInstance holds value of 'instance' option
var DevicesObserveResourcesCmdInstance string

// DevicesObserveResourcesCmdObject holds value of 'object' option
var DevicesObserveResourcesCmdObject string

// DevicesObserveResourcesCmdModel holds value of 'model' option
var DevicesObserveResourcesCmdModel bool

func init() {
	DevicesObserveResourcesCmd.Flags().StringVar(&DevicesObserveResourcesCmdDeviceId, "device-id", "", TRAPI("Target device"))

	DevicesObserveResourcesCmd.Flags().StringVar(&DevicesObserveResourcesCmdInstance, "instance", "", TRAPI("Instance ID"))

	DevicesObserveResourcesCmd.Flags().StringVar(&DevicesObserveResourcesCmdObject, "object", "", TRAPI("Object ID"))

	DevicesObserveResourcesCmd.Flags().BoolVar(&DevicesObserveResourcesCmdModel, "model", false, TRAPI("Whether or not to add model information"))
	DevicesCmd.AddCommand(DevicesObserveResourcesCmd)
}

// DevicesObserveResourcesCmd defines 'observe-resources' subcommand
var DevicesObserveResourcesCmd = &cobra.Command{
	Use:   "observe-resources",
	Short: TRAPI("/devices/{device_id}/{object}/{instance}/observe:post:summary"),
	Long:  TRAPI(`/devices/{device_id}/{object}/{instance}/observe:post:description`),
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

		param, err := collectDevicesObserveResourcesCmdParams(ac)
		if err != nil {
			return err
		}

		body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectDevicesObserveResourcesCmdParams(ac *apiClient) (*apiParams, error) {
	if DevicesObserveResourcesCmdDeviceId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "device-id")
	}

	if DevicesObserveResourcesCmdInstance == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "instance")
	}

	if DevicesObserveResourcesCmdObject == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "object")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForDevicesObserveResourcesCmd("/devices/{device_id}/{object}/{instance}/observe"),
		query:  buildQueryForDevicesObserveResourcesCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesObserveResourcesCmd(path string) string {

	escapedDeviceId := url.PathEscape(DevicesObserveResourcesCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	escapedInstance := url.PathEscape(DevicesObserveResourcesCmdInstance)

	path = strReplace(path, "{"+"instance"+"}", escapedInstance, -1)

	escapedObject := url.PathEscape(DevicesObserveResourcesCmdObject)

	path = strReplace(path, "{"+"object"+"}", escapedObject, -1)

	return path
}

func buildQueryForDevicesObserveResourcesCmd() url.Values {
	result := url.Values{}

	if DevicesObserveResourcesCmdModel != false {
		result.Add("model", sprintf("%t", DevicesObserveResourcesCmdModel))
	}

	return result
}
