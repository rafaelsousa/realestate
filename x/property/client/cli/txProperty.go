package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/rafaelsousa/realestate/x/property/types"
)

func CmdCreateProperty() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-property [address] [city] [state] [zip] [country] [latitude] [longitude] [owner_addr]",
		Short: "Creates a new property",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAddress := string(args[0])
			argsCity := string(args[1])
			argsState := string(args[2])
			argsZip := string(args[3])
			argsCountry := string(args[4])
			argsLatitude := string(args[5])
			argsLongitude := string(args[6])
			argsOwner_addr := string(args[7])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProperty(clientCtx.GetFromAddress().String(), string(argsAddress), string(argsCity), string(argsState), string(argsZip), string(argsCountry), string(argsLatitude), string(argsLongitude), string(argsOwner_addr))
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
		Use:   "update-property [id] [address] [city] [state] [zip] [country] [latitude] [longitude] [owner_addr]",
		Short: "Update a property",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsAddress := string(args[1])
			argsCity := string(args[2])
			argsState := string(args[3])
			argsZip := string(args[4])
			argsCountry := string(args[5])
			argsLatitude := string(args[6])
			argsLongitude := string(args[7])
			argsOwner_addr := string(args[8])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateProperty(clientCtx.GetFromAddress().String(), id, string(argsAddress), string(argsCity), string(argsState), string(argsZip), string(argsCountry), string(argsLatitude), string(argsLongitude), string(argsOwner_addr))
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
		Use:   "delete-property [id] [address] [city] [state] [zip] [country] [latitude] [longitude] [owner_addr]",
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
