// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// PaymentStatementsExportCmdExportMode holds value of 'export_mode' option
var PaymentStatementsExportCmdExportMode string

// PaymentStatementsExportCmdPaymentStatementId holds value of 'payment_statement_id' option
var PaymentStatementsExportCmdPaymentStatementId string

func init() {
	PaymentStatementsExportCmd.Flags().StringVar(&PaymentStatementsExportCmdExportMode, "export-mode", "", TRAPI("Export mode (async, sync)"))

	PaymentStatementsExportCmd.Flags().StringVar(&PaymentStatementsExportCmdPaymentStatementId, "payment-statement-id", "", TRAPI("Payment statement ID"))

	PaymentStatementsExportCmd.MarkFlagRequired("payment-statement-id")

	PaymentStatementsCmd.AddCommand(PaymentStatementsExportCmd)
}

// PaymentStatementsExportCmd defines 'export' subcommand
var PaymentStatementsExportCmd = &cobra.Command{
	Use:   "export",
	Short: TRAPI("/payment_statements/{payment_statement_id}/export:post:summary"),
	Long:  TRAPI(`/payment_statements/{payment_statement_id}/export:post:description`),
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

		param, err := collectPaymentStatementsExportCmdParams(ac)
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

		return prettyPrintStringAsJSON(body)

	},
}

func collectPaymentStatementsExportCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForPaymentStatementsExportCmd("/payment_statements/{payment_statement_id}/export"),
		query:  buildQueryForPaymentStatementsExportCmd(),
	}, nil
}

func buildPathForPaymentStatementsExportCmd(path string) string {

	escapedPaymentStatementId := url.PathEscape(PaymentStatementsExportCmdPaymentStatementId)

	path = strReplace(path, "{"+"payment_statement_id"+"}", escapedPaymentStatementId, -1)

	return path
}

func buildQueryForPaymentStatementsExportCmd() url.Values {
	result := url.Values{}

	if PaymentStatementsExportCmdExportMode != "" {
		result.Add("export_mode", PaymentStatementsExportCmdExportMode)
	}

	return result
}
