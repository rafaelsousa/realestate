package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/rafaelsousa/realestate/x/realestate/types"
	"github.com/spf13/cobra"
)

func CmdCreateProperty() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-property [address] [city] [state] [country] [zipcode] [latitude] [longitude]",
		Short: "Create a new property",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argCity := args[1]
			argState := args[2]
			argCountry := args[3]
			argZipcode := args[4]
			argLatitude := args[5]
			argLongitude := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProperty(clientCtx.GetFromAddress().String(), argAddress, argCity, argState, argCountry, argZipcode, argLatitude, argLongitude)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateProperty() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-property [id] [address] [city] [state] [country] [zipcode] [latitude] [longitude]",
		Short: "Update a property",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argAddress := args[1]

			argCity := args[2]

			argState := args[3]

			argCountry := args[4]

			argZipcode := args[5]

			argLatitude := args[6]

			argLongitude := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateProperty(clientCtx.GetFromAddress().String(), id, argAddress, argCity, argState, argCountry, argZipcode, argLatitude, argLongitude)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteProperty() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-property [id]",
		Short: "Delete a property by id",
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

			msg := types.NewMsgDeleteProperty(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
