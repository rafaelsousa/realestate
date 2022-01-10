package main

import (
	"github.com/rafaelsousa/realestate/ipfs"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/rafaelsousa/realestate/app"
	"github.com/tendermint/spm/cosmoscmd"
)

func main() {
	ipfs.SetupIPFS()
	// Initializes the IPFS

	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(Completion())
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
