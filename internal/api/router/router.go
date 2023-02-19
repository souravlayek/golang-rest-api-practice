package router

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func RedocHandler() http.Handler {
	return middleware.Redoc(middleware.RedocOpts{
		SpecURL: "/swagger.yaml",
		Path:    "/docs",
	}, nil)
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/docs", RedocHandler())
	CreateBookRoutes(router)
	return router
}
