package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers property-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/property/properties/{id}", getPropertyHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/property/properties", listPropertyHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/property/create", createPropertyHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/property/update/{id}", updatePropertyHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/property/delete/{id}", deletePropertyHandler(clientCtx)).Methods("POST")

}
