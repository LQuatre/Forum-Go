package server

import (
	"log"
	"net/http"

	. "jilt.com/m/config"
	. "jilt.com/m/pkg/routes"
)

func StartWebServer() {
	r := NewRouter()

	assets := http.FileServer(http.Dir(ViperConfig.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)

	log.Println("Starting HTTP service at http://" + ViperConfig.App.Address)
	err := http.ListenAndServe(ViperConfig.App.Address, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at " + ViperConfig.App.Address)
		log.Println("Error: " + err.Error())
	}
}
