package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/rafaelsousa/realestate/x/realestate/types"
	"github.com/spf13/cobra"
)

func CmdCreateCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-certificate [property] [surveyor] [certifiation-date] [certificate-text] [property-map]",
		Short: "Create a new certificate",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argProperty, _ := strconv.Atoi(args[0])
			argSurveyor := args[1]
			argCertifiationDate := args[2]
			argCertificateText := args[3]
			argPropertyMap := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCertificate(clientCtx.GetFromAddress().String(), uint64(argProperty), argSurveyor, argCertifiationDate, argCertificateText, argPropertyMap)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-certificate [id] [property] [surveyor] [certifiation-date] [certificate-text] [property-map]",
		Short: "Update a certificate",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argProperty, _ := strconv.Atoi(args[1])

			argSurveyor := args[2]

			argCertifiationDate := args[3]

			argCertificateText := args[4]

			argPropertyMap := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateCertificate(clientCtx.GetFromAddress().String(), id, uint64(argProperty), argSurveyor, argCertifiationDate, argCertificateText, argPropertyMap)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-certificate [id]",
		Short: "Delete a certificate by id",
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

			msg := types.NewMsgDeleteCertificate(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
