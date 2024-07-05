package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	. "jilt.com/m/config"
	"jilt.com/m/pkg/models"
	. "jilt.com/m/pkg/routes"
)

func StartWebServer() {
	r := NewRouter()

	http.Handle("/", r)

	log.Println("Starting HTTP service at http://" + ViperConfig.App.Address)
	err := http.ListenAndServe(ViperConfig.App.Address, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at " + ViperConfig.App.Address)
		log.Println("Error: " + err.Error())
	}

	ticker := time.NewTicker(1 * time.Minute)
    go func() {
        for range ticker.C {
            if err := models.DeleteExpiredSessions(); err != nil {
                log.Println("Erreur lors de la suppression des sessions expirées :", err)
				fmt.Println("Erreur lors de la suppression des sessions expirées :", err)
            } else {
				log.Println("Suppression des sessions expirées effectuée avec succès")
				fmt.Println("Suppression des sessions expirées effectuée avec succès")
			}
        }
    }()
}
