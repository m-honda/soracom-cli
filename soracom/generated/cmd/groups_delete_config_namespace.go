// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// GroupsDeleteConfigNamespaceCmdGroupId holds value of 'group_id' option
var GroupsDeleteConfigNamespaceCmdGroupId string

// GroupsDeleteConfigNamespaceCmdNamespace holds value of 'namespace' option
var GroupsDeleteConfigNamespaceCmdNamespace string

func init() {
	GroupsDeleteConfigNamespaceCmd.Flags().StringVar(&GroupsDeleteConfigNamespaceCmdGroupId, "group-id", "", TRAPI("Target group."))

	GroupsDeleteConfigNamespaceCmd.MarkFlagRequired("group-id")

	GroupsDeleteConfigNamespaceCmd.Flags().StringVar(&GroupsDeleteConfigNamespaceCmdNamespace, "namespace", "", TRAPI("Namespace to be deleted."))

	GroupsDeleteConfigNamespaceCmd.MarkFlagRequired("namespace")

	GroupsCmd.AddCommand(GroupsDeleteConfigNamespaceCmd)
}

// GroupsDeleteConfigNamespaceCmd defines 'delete-config-namespace' subcommand
var GroupsDeleteConfigNamespaceCmd = &cobra.Command{
	Use:   "delete-config-namespace",
	Short: TRAPI("/groups/{group_id}/configuration/{namespace}:delete:summary"),
	Long:  TRAPI(`/groups/{group_id}/configuration/{namespace}:delete:description`),
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

		param, err := collectGroupsDeleteConfigNamespaceCmdParams(ac)
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

func collectGroupsDeleteConfigNamespaceCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForGroupsDeleteConfigNamespaceCmd("/groups/{group_id}/configuration/{namespace}"),
		query:  buildQueryForGroupsDeleteConfigNamespaceCmd(),
	}, nil
}

func buildPathForGroupsDeleteConfigNamespaceCmd(path string) string {

	escapedGroupId := url.PathEscape(GroupsDeleteConfigNamespaceCmdGroupId)

	path = strReplace(path, "{"+"group_id"+"}", escapedGroupId, -1)

	escapedNamespace := url.PathEscape(GroupsDeleteConfigNamespaceCmdNamespace)

	path = strReplace(path, "{"+"namespace"+"}", escapedNamespace, -1)

	return path
}

func buildQueryForGroupsDeleteConfigNamespaceCmd() url.Values {
	result := url.Values{}

	return result
}
