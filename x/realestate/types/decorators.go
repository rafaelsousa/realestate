package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"strconv"
)

const (
	regexLatitude  = "^-?([1-8]?[1-9]|[1-9]0)\\.{1}\\d{1,6}"
	regexLongitude = "^-?([1]?[1-7][1-9]|[1]?[1-8][0]|[1-9]?[0-9])\\.{1}\\d{1,6}"
)

type RealestateDecorator struct {
}

func NewRealestateDecorator() RealestateDecorator {
	return RealestateDecorator{}
}

func (red RealestateDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}
	for _, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *MsgCreateProperty:
			err := validateProperty(msg)
			if err != nil {
				return ctx, err
			}
		}
	}
	return next(ctx, tx, simulate)
}

func validateProperty(msg *MsgCreateProperty) error {
	_, err := strconv.ParseFloat(msg.Latitude, 64)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "Invalid latitude, please use decimal format")
	}
	_, err = strconv.ParseFloat(msg.Latitude, 64)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "Invalid longitude, please use decimal format")
	}
	return nil
}
