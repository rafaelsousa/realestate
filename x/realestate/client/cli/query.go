package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rafaelsousa/realestate/x/realestate/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group realestate queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListProperty())
	cmd.AddCommand(CmdShowProperty())
	cmd.AddCommand(CmdListCertificate())
	cmd.AddCommand(CmdShowCertificate())
	cmd.AddCommand(CmdListLocking())
	cmd.AddCommand(CmdShowLocking())
	cmd.AddCommand(CmdListInspection())
	cmd.AddCommand(CmdShowInspection())
	cmd.AddCommand(CmdListTransference())
	cmd.AddCommand(CmdShowTransference())
	cmd.AddCommand(CmdListHouse())
	cmd.AddCommand(CmdShowHouse())
	// this line is used by starport scaffolding # 1

	return cmd
}
