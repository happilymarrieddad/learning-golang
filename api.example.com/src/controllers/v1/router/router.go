package router

import (
	"github.com/go-xorm/xorm"
	"learning-golang/api.example.com/pkg/types/routes"
	UsersHandler "learning-golang/api.example.com/src/controllers/v1/routes/users"
	StatusHandler "learning-golang/api.example.com/src/controllers/v1/status"
	"log"
	"net/http"
)

var db *xorm.Engine

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")
		if len(token) < 1 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		log.Println("Inside V1 Middleware")

		next.ServeHTTP(w, r)
	})
}

func GetRoutes(DB *xorm.Engine) (SubRoute map[string]routes.SubRoutePackage) {
	db = DB

	StatusHandler.Init(DB)
	UsersHandler.Init(DB)

	/* ROUTES */
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": {
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
				routes.Route{"UsersIndex", "GET", "/users", UsersHandler.Index},
				routes.Route{"UsersStore", "POST", "/users", UsersHandler.Store},
			},
			Middleware: Middleware,
		},
	}

	return
}
