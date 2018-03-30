package router

import (
	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
	"learning-golang/api.example.com/pkg/types/routes"
	V1SubRoutes "learning-golang/api.example.com/src/controllers/v1/router"
)

type Router struct {
	Router *mux.Router
}

func (r *Router) Init(db *xorm.Engine) {
	r.Router.Use(Middleware)

	baseRoutes := GetRoutes(db)
	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	v1SubRoutes := V1SubRoutes.GetRoutes(db)
	for name, pack := range v1SubRoutes {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}
}

func (r *Router) AttachSubRouterWithMiddleware(path string, subroutes routes.Routes, middleware mux.MiddlewareFunc) (SubRouter *mux.Router) {

	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return
}

func NewRouter() (r Router) {
	r.Router = mux.NewRouter().StrictSlash(true)

	return
}
