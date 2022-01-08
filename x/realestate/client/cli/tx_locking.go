package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/rafaelsousa/realestate/x/realestate/types"
	"github.com/spf13/cobra"
)

func CmdCreateLocking() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-locking [owner] [date-locking] [date-unlocking] [unlock-fees] [property]",
		Short: "Create a new locking",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOwner := args[0]
			argDateLocking := args[1]
			argDateUnlocking := args[2]
			argUnlockFees, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}
			argProperty, _ := strconv.Atoi(args[4])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateLocking(clientCtx.GetFromAddress().String(), argOwner, argDateLocking, argDateUnlocking, &argUnlockFees, uint64(argProperty))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateLocking() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-locking [id] [owner] [date-locking] [date-unlocking] [unlock-fees] [property]",
		Short: "Update a locking",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argOwner := args[1]

			argDateLocking := args[2]

			argDateUnlocking := args[3]

			argUnlockFees, _ := sdk.ParseCoinNormalized(args[4])

			argProperty, _ := strconv.Atoi(args[5])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateLocking(clientCtx.GetFromAddress().String(), id, argOwner, argDateLocking, argDateUnlocking, argUnlockFees, uint64(argProperty))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteLocking() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-locking [id]",
		Short: "Delete a locking by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteLocking(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
