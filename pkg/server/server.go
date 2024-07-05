package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	. "jilt.com/m/config"
	"jilt.com/m/pkg/models"
	. "jilt.com/m/pkg/routes"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
	 	return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func StartWebServer() {
	r := NewRouter()

	assets := http.FileServer(http.Dir(ViperConfig.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)

	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

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

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			delete(clients, conn)
			return
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
	
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}	