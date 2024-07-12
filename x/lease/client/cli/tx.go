package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

func txEndLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "end-lease [id]",
		Short: "End an existing lease",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := v1.NewMsgEndLeaseRequest(
				ctx.FromAddress.Bytes(),
				id,
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

func txRenewLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew-lease [id] [hours] [denom]",
		Short: "Renew an existing lease for a specified duration",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			hours, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := v1.NewMsgRenewLeaseRequest(
				ctx.FromAddress.Bytes(),
				id,
				hours,
				args[2],
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

func txStartLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-lease [node-addr] [hours] [denom] [renewable]",
		Short: "Start a lease with a node for the specified duration",
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

			hours, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			renewable, err := strconv.ParseBool(args[3])
			if err != nil {
				return err
			}

			msg := v1.NewMsgStartLeaseRequest(
				ctx.FromAddress.Bytes(),
				nodeAddr,
				hours,
				args[2],
				renewable,
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

func txUpdateLease() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-lease [id] [renewable]",
		Short: "Update the details of an existing lease",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			renewable, err := strconv.ParseBool(args[1])
			if err != nil {
				return err
			}

			msg := v1.NewMsgUpdateLeaseRequest(
				ctx.FromAddress.Bytes(),
				id,
				renewable,
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
