package httpd

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Fatal(err)
			return
		}
	}
}

func WsEndpoint(upgrader *websocket.Upgrader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println(err)
		}

    log.Println("Client Successfull Connected")
    
    reader(ws);
	}
}
