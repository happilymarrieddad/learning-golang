package router

import (
	"github.com/go-xorm/xorm"
	"learning-golang/api.example.com/pkg/types/routes"
	HomeHandler "learning-golang/api.example.com/src/controllers/home"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) routes.Routes {

	HomeHandler.Init(db)

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
	}
}
