package router

import (
	"learning-golang/api.example.com/pkg/types/routes"
	HomeHandler "learning-golang/api.example.com/src/controllers/home"
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() routes.Routes {

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
	}
}
