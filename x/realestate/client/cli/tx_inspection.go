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

func CmdCreateInspection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-inspection [inspector] [property] [fees] [inspection-results]",
		Short: "Create a new inspection",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argInspector := args[0]
			argProperty, _ := strconv.Atoi(args[1])
			argFees := args[2]
			argInspectionResults := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fees, err := sdk.ParseCoinNormalized(argFees)
			if err != nil {
				return err
			}
			msg := types.NewMsgCreateInspection(clientCtx.GetFromAddress().String(), argInspector, uint64(argProperty), fees, argInspectionResults)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateInspection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-inspection [id] [inspector] [property] [fees] [inspection-results]",
		Short: "Update a inspection",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argInspector := args[1]

			argProperty, _ := strconv.Atoi(args[2])

			argFees := args[3]
			fees, err := sdk.ParseCoinNormalized(argFees)
			if err != nil {
				return err
			}

			argInspectionResults := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateInspection(clientCtx.GetFromAddress().String(), id, argInspector, uint64(argProperty), fees, argInspectionResults)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteInspection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-inspection [id]",
		Short: "Delete a inspection by id",
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

			msg := types.NewMsgDeleteInspection(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
