package keeper

import (
	"github.com/rafaelsousa/realestate/x/blueprints/types"
)

var _ types.QueryServer = Keeper{}
