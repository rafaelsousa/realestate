package keeper

import (
	"github.com/rafaelsousa/realestate/x/realestate/types"
)

var _ types.QueryServer = Keeper{}
