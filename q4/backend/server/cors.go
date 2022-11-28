package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
