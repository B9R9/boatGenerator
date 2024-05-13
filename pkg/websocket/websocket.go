package websocket

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	clients = make(map[*websocket.Conn]bool) // Carte pour stocker les connexions des clients
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Configurer CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if r.Method == "OPTIONS" {
		// Répondre aux requêtes OPTIONS pour CORS
		w.WriteHeader(http.StatusOK)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Enregistrer la connexion du client dans la carte
	clients[conn] = true
}

func StartServer(port int) {
	http.HandleFunc("/echo", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func SendMessageToAll(message []byte) {
	for conn := range clients {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
			conn.Close() // Fermer la connexion si une erreur se produit
			delete(clients, conn) // Supprimer la connexion de la carte
		}
	}
}