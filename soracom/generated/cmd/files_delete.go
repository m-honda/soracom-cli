// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// FilesDeleteCmdPath holds value of 'path' option
var FilesDeleteCmdPath string

// FilesDeleteCmdScope holds value of 'scope' option
var FilesDeleteCmdScope string

func init() {
	FilesDeleteCmd.Flags().StringVar(&FilesDeleteCmdPath, "path", "", TRAPI("Target path"))

	FilesDeleteCmd.MarkFlagRequired("path")

	FilesDeleteCmd.Flags().StringVar(&FilesDeleteCmdScope, "scope", "", TRAPI("Scope of the request"))

	FilesDeleteCmd.MarkFlagRequired("scope")

	FilesCmd.AddCommand(FilesDeleteCmd)
}

// FilesDeleteCmd defines 'delete' subcommand
var FilesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/files/{scope}/{path}:delete:summary"),
	Long:  TRAPI(`/files/{scope}/{path}:delete:description`),
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

		param, err := collectFilesDeleteCmdParams(ac)
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

func collectFilesDeleteCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForFilesDeleteCmd("/files/{scope}/{path}"),
		query:  buildQueryForFilesDeleteCmd(),
	}, nil
}

func buildPathForFilesDeleteCmd(path string) string {

	escapedPath := harvestFilesPathEscape(FilesDeleteCmdPath)

	path = strReplace(path, "{"+"path"+"}", escapedPath, -1)

	escapedScope := url.PathEscape(FilesDeleteCmdScope)

	path = strReplace(path, "{"+"scope"+"}", escapedScope, -1)

	return path
}

func buildQueryForFilesDeleteCmd() url.Values {
	result := url.Values{}

	return result
}
