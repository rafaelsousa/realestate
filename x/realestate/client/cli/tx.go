package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateProperty())
	cmd.AddCommand(CmdUpdateProperty())
	cmd.AddCommand(CmdDeleteProperty())
	cmd.AddCommand(CmdCreateCertificate())
	cmd.AddCommand(CmdUpdateCertificate())
	cmd.AddCommand(CmdDeleteCertificate())
	cmd.AddCommand(CmdCreateLocking())
	cmd.AddCommand(CmdUpdateLocking())
	cmd.AddCommand(CmdDeleteLocking())
	cmd.AddCommand(CmdCreateInspection())
	cmd.AddCommand(CmdUpdateInspection())
	cmd.AddCommand(CmdDeleteInspection())
	cmd.AddCommand(CmdCreateTransference())
	cmd.AddCommand(CmdUpdateTransference())
	cmd.AddCommand(CmdDeleteTransference())
	cmd.AddCommand(CmdCreateHouse())
	cmd.AddCommand(CmdUpdateHouse())
	cmd.AddCommand(CmdDeleteHouse())
	cmd.AddCommand(CmdRequestCertification())
	cmd.AddCommand(CmdEmitCertificate())
	cmd.AddCommand(CmdLockProperty())
	cmd.AddCommand(CmdUnlockProperty())
	cmd.AddCommand(CmdUnlockAssets())
	// this line is used by starport scaffolding # 1

	return cmd
}
