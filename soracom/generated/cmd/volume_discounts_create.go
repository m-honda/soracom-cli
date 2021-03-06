// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"fmt"

	"io/ioutil"

	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// VolumeDiscountsCreateCmdStartDate holds value of 'startDate' option
var VolumeDiscountsCreateCmdStartDate string

// VolumeDiscountsCreateCmdVolumeDiscountPaymentType holds value of 'volumeDiscountPaymentType' option
var VolumeDiscountsCreateCmdVolumeDiscountPaymentType string

// VolumeDiscountsCreateCmdVolumeDiscountType holds value of 'volumeDiscountType' option
var VolumeDiscountsCreateCmdVolumeDiscountType string

// VolumeDiscountsCreateCmdContractTermMonth holds value of 'contractTermMonth' option
var VolumeDiscountsCreateCmdContractTermMonth int64

// VolumeDiscountsCreateCmdQuantity holds value of 'quantity' option
var VolumeDiscountsCreateCmdQuantity int64

// VolumeDiscountsCreateCmdBody holds contents of request body to be sent
var VolumeDiscountsCreateCmdBody string

func init() {
	VolumeDiscountsCreateCmd.Flags().StringVar(&VolumeDiscountsCreateCmdStartDate, "start-date", "", TRAPI(""))

	VolumeDiscountsCreateCmd.Flags().StringVar(&VolumeDiscountsCreateCmdVolumeDiscountPaymentType, "volume-discount-payment-type", "", TRAPI(""))

	VolumeDiscountsCreateCmd.Flags().StringVar(&VolumeDiscountsCreateCmdVolumeDiscountType, "volume-discount-type", "", TRAPI(""))

	VolumeDiscountsCreateCmd.Flags().Int64Var(&VolumeDiscountsCreateCmdContractTermMonth, "contract-term-month", 12, TRAPI(""))

	VolumeDiscountsCreateCmd.Flags().Int64Var(&VolumeDiscountsCreateCmdQuantity, "quantity", 0, TRAPI(""))

	VolumeDiscountsCreateCmd.Flags().StringVar(&VolumeDiscountsCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	VolumeDiscountsCmd.AddCommand(VolumeDiscountsCreateCmd)
}

// VolumeDiscountsCreateCmd defines 'create' subcommand
var VolumeDiscountsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/volume_discounts:post:summary"),
	Long:  TRAPI(`/volume_discounts:post:description`),
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

		param, err := collectVolumeDiscountsCreateCmdParams(ac)
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

func collectVolumeDiscountsCreateCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForVolumeDiscountsCreateCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if VolumeDiscountsCreateCmdVolumeDiscountPaymentType == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "volume-discount-payment-type")
		}

	}

	if VolumeDiscountsCreateCmdVolumeDiscountType == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "volume-discount-type")
		}

	}

	if VolumeDiscountsCreateCmdQuantity == 0 {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "quantity")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForVolumeDiscountsCreateCmd("/volume_discounts"),
		query:       buildQueryForVolumeDiscountsCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVolumeDiscountsCreateCmd(path string) string {

	return path
}

func buildQueryForVolumeDiscountsCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForVolumeDiscountsCreateCmd() (string, error) {
	var result map[string]interface{}

	if VolumeDiscountsCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(VolumeDiscountsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(VolumeDiscountsCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if VolumeDiscountsCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(VolumeDiscountsCreateCmdBody)
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

	if VolumeDiscountsCreateCmdStartDate != "" {
		result["startDate"] = VolumeDiscountsCreateCmdStartDate
	}

	if VolumeDiscountsCreateCmdVolumeDiscountPaymentType != "" {
		result["volumeDiscountPaymentType"] = VolumeDiscountsCreateCmdVolumeDiscountPaymentType
	}

	if VolumeDiscountsCreateCmdVolumeDiscountType != "" {
		result["volumeDiscountType"] = VolumeDiscountsCreateCmdVolumeDiscountType
	}

	if VolumeDiscountsCreateCmdContractTermMonth != 12 {
		result["contractTermMonth"] = VolumeDiscountsCreateCmdContractTermMonth
	}

	if VolumeDiscountsCreateCmdQuantity != 0 {
		result["quantity"] = VolumeDiscountsCreateCmdQuantity
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
