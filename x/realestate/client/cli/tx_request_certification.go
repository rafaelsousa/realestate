package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/rafaelsousa/realestate/x/realestate/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRequestCertification() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-certification [property] [fees] [surveyor]",
		Short: "Broadcast message RequestCertification",
		Args:  cobra.ExactArgs(3),
		Long: `
Initiates the flow to certificate a house. 
Params :
- Property : Id of the minted property 
- Fees : funds that will be allocated to the surveyour to execute the task
- Surveyor : address of the surveyour in charge of certificating the house
The --from argument is mandatory, as it represents the owner (or the fee payer) of the certification.
		`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argProperty := args[0]
			argFees := args[1]
			argSurveyor := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			property, err := strconv.Atoi(argProperty)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "property should be a number", err)
			}
			fees, err := sdk.ParseCoinNormalized(argFees)
			if err != nil {
				return err
			}
			msg := types.NewMsgRequestCertification(
				clientCtx.GetFromAddress().String(),
				uint64(property),
				fees,
				argSurveyor,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
