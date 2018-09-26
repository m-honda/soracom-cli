package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersRegisterCmdGroupId holds value of 'groupId' option
var SubscribersRegisterCmdGroupId string

// SubscribersRegisterCmdImsi holds value of 'imsi' option
var SubscribersRegisterCmdImsi string

// SubscribersRegisterCmdRegistrationSecret holds value of 'registrationSecret' option
var SubscribersRegisterCmdRegistrationSecret string

// SubscribersRegisterCmdBody holds contents of request body to be sent
var SubscribersRegisterCmdBody string

func init() {
	SubscribersRegisterCmd.Flags().StringVar(&SubscribersRegisterCmdGroupId, "group-id", "", TRAPI(""))

	SubscribersRegisterCmd.Flags().StringVar(&SubscribersRegisterCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersRegisterCmd.Flags().StringVar(&SubscribersRegisterCmdRegistrationSecret, "registration-secret", "", TRAPI(""))

	SubscribersRegisterCmd.Flags().StringVar(&SubscribersRegisterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SubscribersCmd.AddCommand(SubscribersRegisterCmd)
}

// SubscribersRegisterCmd defines 'register' subcommand
var SubscribersRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/subscribers/{imsi}/register:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/register:post:description`),
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

		param, err := collectSubscribersRegisterCmdParams(ac)
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

func collectSubscribersRegisterCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSubscribersRegisterCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersRegisterCmd("/subscribers/{imsi}/register"),
		query:       buildQueryForSubscribersRegisterCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSubscribersRegisterCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersRegisterCmdImsi, -1)

	return path
}

func buildQueryForSubscribersRegisterCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSubscribersRegisterCmd() (string, error) {
	var result map[string]interface{}

	if SubscribersRegisterCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SubscribersRegisterCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersRegisterCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SubscribersRegisterCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersRegisterCmdBody)
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

	if SubscribersRegisterCmdGroupId != "" {
		result["groupId"] = SubscribersRegisterCmdGroupId
	}

	if SubscribersRegisterCmdRegistrationSecret != "" {
		result["registrationSecret"] = SubscribersRegisterCmdRegistrationSecret
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
