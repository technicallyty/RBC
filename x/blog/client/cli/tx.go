package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/regen-network/bec/x/blog"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        blog.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", blog.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreatePost())

	return cmd
}

// CmdCreatePost - lint TODO remove comment
func CmdCreatePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-post [author] [title] [body]",
		Short: "Creates a new post",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			argsTitle := string(args[1])
			argsBody := string(args[2])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := blog.MsgCreatePostRequest{
				Author: clientCtx.GetFromAddress().String(),
				Title:  argsTitle,
				Body:   argsBody,
			}
			svcMsgClientConn := &ServiceMsgClientConn{}
			msgClient := blog.NewMsgClient(svcMsgClientConn)
			_, err = msgClient.CreatePost(cmd.Context(), &msg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.Msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdCreateComment - Creates a comment on a post. Comments need to know what post is being commented on, the content, and who wrote it.
// this is basically copied from the above.
func CmdCreateComment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-comment [author] [post-id] [body]",
		Short: "Creates a new comment on a post",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			argsPostID := string(args[1])
			argsBody := string(args[2])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := blog.MsgCreateCommentRequest{
				Author: clientCtx.GetFromAddress().String(),
				PostID: argsPostID,
				Body:   argsBody,
			}
			svcMsgClientConn := &ServiceMsgClientConn{}
			msgClient := blog.NewMsgClient(svcMsgClientConn)
			_, err = msgClient.CreateComment(cmd.Context(), &msg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.Msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
