package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func txRegisterNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-node [remote-url] [gigabyte-prices] [hourly-prices]",
		Short: "Register a new node with a remote URL and pricing details",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			gigabytePrices, err := sdk.ParseDecCoins(args[1])
			if err != nil {
				return err
			}

			hourlyPrices, err := sdk.ParseDecCoins(args[2])
			if err != nil {
				return err
			}

			msg := v3.NewMsgRegisterNodeRequest(
				ctx.FromAddress.Bytes(),
				gigabytePrices,
				hourlyPrices,
				args[0],
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagGigabytePrices, "", "prices for one gigabyte of bandwidth (e.g., 1000token")
	cmd.Flags().String(flagHourlyPrices, "", "prices for one hour of bandwidth (e.g., 500token")
	cmd.Flags().String(flagRemoteURL, "", "remote URL address for the node")
	return cmd
}

func txUpdateNodeDetails() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-node-details",
		Short: "Update the pricing and remote URL details of an existing node",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			gigabytePrices, err := GetGigabytePrices(cmd.Flags())
			if err != nil {
				return err
			}

			hourlyPrices, err := GetHourlyPrices(cmd.Flags())
			if err != nil {
				return err
			}

			remoteURL, err := cmd.Flags().GetString(flagRemoteURL)
			if err != nil {
				return err
			}

			msg := v3.NewMsgUpdateNodeDetailsRequest(
				ctx.FromAddress.Bytes(),
				gigabytePrices,
				hourlyPrices,
				remoteURL,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagGigabytePrices, "", "prices for one gigabyte of bandwidth (e.g., 1000token)")
	cmd.Flags().String(flagHourlyPrices, "", "prices for one hour of bandwidth (e.g., 500token)")
	cmd.Flags().String(flagRemoteURL, "", "remote URL address for the node")

	return cmd
}

func txUpdateNodeStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-node-status [status]",
		Short: "Update the operational status of a node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := v3.NewMsgUpdateNodeStatusRequest(
				ctx.FromAddress.Bytes(),
				v1base.StatusFromString(args[0]),
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txStartSession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-session [node-addr] [gigabytes] [hours] [denom]",
		Short: "Start a session with a node",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			nodeAddr, err := base.NodeAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			gigabytes, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			hours, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := v3.NewMsgStartSessionRequest(
				ctx.FromAddress.Bytes(),
				nodeAddr,
				gigabytes,
				hours,
				args[3],
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
