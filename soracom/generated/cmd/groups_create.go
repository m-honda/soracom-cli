package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GroupsCreateCmdBody holds contents of request body to be sent
var GroupsCreateCmdBody string

func init() {

	GroupsCreateCmd.Flags().StringVar(&GroupsCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	GroupsCmd.AddCommand(GroupsCreateCmd)
}

// GroupsCreateCmd defines 'create' subcommand
var GroupsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/groups:post:summary"),
	Long:  TRAPI(`/groups:post:description`),
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

		param, err := collectGroupsCreateCmdParams(ac)
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

func collectGroupsCreateCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForGroupsCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForGroupsCreateCmd("/groups"),
		query:       buildQueryForGroupsCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForGroupsCreateCmd(path string) string {

	return path
}

func buildQueryForGroupsCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForGroupsCreateCmd() (string, error) {
	var result map[string]interface{}

	if GroupsCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(GroupsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(GroupsCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if GroupsCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(GroupsCreateCmdBody)
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

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
