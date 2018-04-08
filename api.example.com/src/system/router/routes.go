package router

import (
	"github.com/go-xorm/xorm"
	"learning-golang/api.example.com/pkg/types/routes"
	AuthHandler "learning-golang/api.example.com/src/controllers/auth"
	HomeHandler "learning-golang/api.example.com/src/controllers/home"
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside main middleware.")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) routes.Routes {

	AuthHandler.Init(db)
	HomeHandler.Init(db)

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
		routes.Route{"AuthStore", "POST", "/auth/login", AuthHandler.Login},
		routes.Route{"AuthCheck", "GET", "/auth/check", AuthHandler.Check},
	}
}
