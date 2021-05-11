package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
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

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/property/owners/{id}", getOwnerHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/property/owners", listOwnerHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/property/properties/{id}", getPropertyHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/property/properties", listPropertyHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/property/owners", createOwnerHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/property/owners/{id}", updateOwnerHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/property/owners/{id}", deleteOwnerHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/property/properties", createPropertyHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/property/properties/{id}", updatePropertyHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/property/properties/{id}", deletePropertyHandler(clientCtx)).Methods("POST")

}
