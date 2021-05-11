package keeper

import (
	"github.com/rafaelsousa/realestate/x/property/types"
)

var _ types.QueryServer = Keeper{}
