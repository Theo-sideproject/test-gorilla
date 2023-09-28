package main

import (
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("Client connect: ", s.ID())
		return nil
	})

	server.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
		log.Printf("Message reçu de %s: %s\n", s.ID(), msg)
		// Vous pouvez envoyer le message à d'autres clients ici
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Printf(": %s - raison: %s\n", s.ID(), reason)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public"))) // Servez le contenu statique depuis le dossier "public"

	log.Println("Serveur Socket.IO écoutant sur :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
