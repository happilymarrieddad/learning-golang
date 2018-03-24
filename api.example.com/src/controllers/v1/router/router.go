package router

import (
	"learning-golang/api.example.com/pkg/types/routes"
	StatusHandler "learning-golang/api.example.com/src/controllers/v1/status"
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() (SubRoute map[string]routes.SubRoutePackage) {

	/* ROUTES */
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": routes.SubRoutePackage{
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
			},
			Middleware: Middleware,
		},
	}

	return
}
