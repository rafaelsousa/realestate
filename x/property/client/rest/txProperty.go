package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/rafaelsousa/realestate/x/property/types"
)

type createPropertyRequest struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	Creator    string       `json:"creator"`
	Address    string       `json:"address"`
	City       string       `json:"city"`
	State      string       `json:"state"`
	Zip        string       `json:"zip"`
	Country    string       `json:"country"`
	Latitude   string       `json:"latitude"`
	Longitude  string       `json:"longitude"`
	Owner_addr string       `json:"owner_addr"`
}

func createPropertyHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPropertyRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedAddress := req.Address

		parsedCity := req.City

		parsedState := req.State

		parsedZip := req.Zip

		parsedCountry := req.Country

		parsedLatitude := req.Latitude

		parsedLongitude := req.Longitude

		parsedOwner_addr := req.Owner_addr

		msg := types.NewMsgCreateProperty(
			req.Creator,
			parsedAddress,
			parsedCity,
			parsedState,
			parsedZip,
			parsedCountry,
			parsedLatitude,
			parsedLongitude,
			parsedOwner_addr,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updatePropertyRequest struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	Creator    string       `json:"creator"`
	Address    string       `json:"address"`
	City       string       `json:"city"`
	State      string       `json:"state"`
	Zip        string       `json:"zip"`
	Country    string       `json:"country"`
	Latitude   string       `json:"latitude"`
	Longitude  string       `json:"longitude"`
	Owner_addr string       `json:"owner_addr"`
}

func updatePropertyHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req updatePropertyRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err = sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedAddress := req.Address

		parsedCity := req.City

		parsedState := req.State

		parsedZip := req.Zip

		parsedCountry := req.Country

		parsedLatitude := req.Latitude

		parsedLongitude := req.Longitude

		parsedOwner_addr := req.Owner_addr

		msg := types.NewMsgUpdateProperty(
			req.Creator,
			id,
			parsedAddress,
			parsedCity,
			parsedState,
			parsedZip,
			parsedCountry,
			parsedLatitude,
			parsedLongitude,
			parsedOwner_addr,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deletePropertyRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
}

func deletePropertyHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req deletePropertyRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err = sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgDeleteProperty(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
