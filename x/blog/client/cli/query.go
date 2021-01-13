package cli

import (
	"fmt"

	// "strings"

	"github.com/regen-network/bec/x/blog"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group glob queries under a subcommand
	cmd := &cobra.Command{
		Use:                        blog.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", blog.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdAllPosts())
	cmd.AddCommand(CmdAllComments())

	return cmd
}

func CmdAllPosts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-post",
		Short: "list all post",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := blog.NewQueryClient(clientCtx)

			params := &blog.QueryAllPostsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllPosts(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "blog")

	return cmd
}

// CmdAllComments - query all comments on a given post ID
func CmdAllComments() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-comment",
		Short: "list all comments on a post",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// the post's ID we are interested in getting comments from
			argsPostID := string(args[0])
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := blog.NewQueryClient(clientCtx)

			params := &blog.QueryAllCommentsRequest{
				PostID: argsPostID,
			}

			res, err := queryClient.AllComments(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "blog")

	return cmd
}
