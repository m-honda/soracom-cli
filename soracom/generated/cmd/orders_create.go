package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OrdersCreateCmdShippingAddressId holds value of 'shippingAddressId' option
var OrdersCreateCmdShippingAddressId string

// OrdersCreateCmdBody holds contents of request body to be sent
var OrdersCreateCmdBody string

func init() {
	OrdersCreateCmd.Flags().StringVar(&OrdersCreateCmdShippingAddressId, "shipping-address-id", "", TRAPI(""))

	OrdersCreateCmd.Flags().StringVar(&OrdersCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OrdersCmd.AddCommand(OrdersCreateCmd)
}

// OrdersCreateCmd defines 'create' subcommand
var OrdersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/orders:post:summary"),
	Long:  TRAPI(`/orders:post:description`),
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

		param, err := collectOrdersCreateCmdParams(ac)
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

func collectOrdersCreateCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForOrdersCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOrdersCreateCmd("/orders"),
		query:       buildQueryForOrdersCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForOrdersCreateCmd(path string) string {

	return path
}

func buildQueryForOrdersCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForOrdersCreateCmd() (string, error) {
	var result map[string]interface{}

	if OrdersCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OrdersCreateCmdBody, "@") {
			fname := strings.TrimPrefix(OrdersCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if OrdersCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(OrdersCreateCmdBody)
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

	if OrdersCreateCmdShippingAddressId != "" {
		result["shippingAddressId"] = OrdersCreateCmdShippingAddressId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
