package routes

import (
	"net/http"
	"strings"

	. "jilt.com/m/config"
)

func NewRouter() *http.ServeMux {
    router := http.NewServeMux()

    for _, route := range webRoutes {
        route := route // Créez une copie de la variable `route` pour chaque itération

        handler := route.HandlerFunc

        switch route.Method {
        case "GET":
            router.HandleFunc(route.Pattern, func(w http.ResponseWriter, r *http.Request) {
                if r.Method == http.MethodGet {
                    if strings.HasPrefix(r.URL.Path, route.Pattern) {
                        handler.ServeHTTP(w, r)
                    } else {
                        http.NotFound(w, r)
                    }
                } else {
                    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
                }
            })
        case "POST":
            router.HandleFunc(route.Pattern, func(w http.ResponseWriter, r *http.Request) {
                if r.Method == http.MethodPost {
                    if strings.HasPrefix(r.URL.Path, route.Pattern) {
                        handler.ServeHTTP(w, r)
                    } else {
                        http.NotFound(w, r)
                    }
                } else {
                    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
                }
            })
        default:
            router.HandleFunc(route.Pattern, func(w http.ResponseWriter, r *http.Request) {
                http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            })
        }
    }

    staticDir := ViperConfig.App.Static
    fs := http.FileServer(http.Dir(staticDir))

    router.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
        http.StripPrefix("/static/", fs).ServeHTTP(w, r)
    })

    return router
}
