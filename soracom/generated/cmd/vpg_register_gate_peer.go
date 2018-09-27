package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// VpgRegisterGatePeerCmdInnerIpAddress holds value of 'innerIpAddress' option
var VpgRegisterGatePeerCmdInnerIpAddress string

// VpgRegisterGatePeerCmdOuterIpAddress holds value of 'outerIpAddress' option
var VpgRegisterGatePeerCmdOuterIpAddress string

// VpgRegisterGatePeerCmdVpgId holds value of 'vpg_id' option
var VpgRegisterGatePeerCmdVpgId string

// VpgRegisterGatePeerCmdBody holds contents of request body to be sent
var VpgRegisterGatePeerCmdBody string

func init() {
	VpgRegisterGatePeerCmd.Flags().StringVar(&VpgRegisterGatePeerCmdInnerIpAddress, "inner-ip-address", "", TRAPI(""))

	VpgRegisterGatePeerCmd.Flags().StringVar(&VpgRegisterGatePeerCmdOuterIpAddress, "outer-ip-address", "", TRAPI(""))

	VpgRegisterGatePeerCmd.Flags().StringVar(&VpgRegisterGatePeerCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))

	VpgRegisterGatePeerCmd.Flags().StringVar(&VpgRegisterGatePeerCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	VpgCmd.AddCommand(VpgRegisterGatePeerCmd)
}

// VpgRegisterGatePeerCmd defines 'register-gate-peer' subcommand
var VpgRegisterGatePeerCmd = &cobra.Command{
	Use:   "register-gate-peer",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/gate/peers:post:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/gate/peers:post:description`),
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

		param, err := collectVpgRegisterGatePeerCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectVpgRegisterGatePeerCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForVpgRegisterGatePeerCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVpgRegisterGatePeerCmd("/virtual_private_gateways/{vpg_id}/gate/peers"),
		query:       buildQueryForVpgRegisterGatePeerCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForVpgRegisterGatePeerCmd(path string) string {

	path = strings.Replace(path, "{"+"vpg_id"+"}", VpgRegisterGatePeerCmdVpgId, -1)

	return path
}

func buildQueryForVpgRegisterGatePeerCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForVpgRegisterGatePeerCmd() (string, error) {
	var result map[string]interface{}

	if VpgRegisterGatePeerCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VpgRegisterGatePeerCmdBody, "@") {
			fname := strings.TrimPrefix(VpgRegisterGatePeerCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if VpgRegisterGatePeerCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(VpgRegisterGatePeerCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if VpgRegisterGatePeerCmdInnerIpAddress != "" {
		result["innerIpAddress"] = VpgRegisterGatePeerCmdInnerIpAddress
	}

	if VpgRegisterGatePeerCmdOuterIpAddress != "" {
		result["outerIpAddress"] = VpgRegisterGatePeerCmdOuterIpAddress
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
