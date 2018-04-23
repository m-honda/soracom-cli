package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersSendSmsCmdImsi holds value of 'imsi' option
var SubscribersSendSmsCmdImsi string

// SubscribersSendSmsCmdPayload holds value of 'payload' option
var SubscribersSendSmsCmdPayload string

// SubscribersSendSmsCmdEncodingType holds value of 'encodingType' option
var SubscribersSendSmsCmdEncodingType int64

// SubscribersSendSmsCmdBody holds contents of request body to be sent
var SubscribersSendSmsCmdBody string

func init() {
	SubscribersSendSmsCmd.Flags().StringVar(&SubscribersSendSmsCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersSendSmsCmd.Flags().StringVar(&SubscribersSendSmsCmdPayload, "payload", "", TRAPI(""))

	SubscribersSendSmsCmd.Flags().Int64Var(&SubscribersSendSmsCmdEncodingType, "encoding-type", 0, TRAPI(""))

	SubscribersSendSmsCmd.Flags().StringVar(&SubscribersSendSmsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SubscribersCmd.AddCommand(SubscribersSendSmsCmd)
}

// SubscribersSendSmsCmd defines 'send-sms' subcommand
var SubscribersSendSmsCmd = &cobra.Command{
	Use:   "send-sms",
	Short: TRAPI("/subscribers/{imsi}/send_sms:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/send_sms:post:description`),
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

		param, err := collectSubscribersSendSmsCmdParams(ac)
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

func collectSubscribersSendSmsCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSubscribersSendSmsCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersSendSmsCmd("/subscribers/{imsi}/send_sms"),
		query:       buildQueryForSubscribersSendSmsCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSubscribersSendSmsCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersSendSmsCmdImsi, -1)

	return path
}

func buildQueryForSubscribersSendSmsCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSubscribersSendSmsCmd() (string, error) {
	if SubscribersSendSmsCmdBody != "" {
		if strings.HasPrefix(SubscribersSendSmsCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersSendSmsCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if SubscribersSendSmsCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return SubscribersSendSmsCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if SubscribersSendSmsCmdPayload != "" {
		result["payload"] = SubscribersSendSmsCmdPayload
	}

	if SubscribersSendSmsCmdEncodingType != 0 {
		result["encodingType"] = SubscribersSendSmsCmdEncodingType
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
