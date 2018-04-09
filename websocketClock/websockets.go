// websockets.go
package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			alarmtime := ""
			if string(msg) != "start" {
				alarmtime = string(msg)
			}

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}

			go func(a string) {
				for {
					t := time.Now().Local()
					s := t.Format("03:04:05")

					if s == a {
						// Write message back to browser
						if err = conn.WriteMessage(websocket.TextMessage, []byte("alarm!")); err != nil {
							return
						}
					} else {
						// Write message back to browser
						if err = conn.WriteMessage(websocket.TextMessage, []byte(s)); err != nil {
							return
						}
					}

					time.Sleep(1 * time.Second)
				}
			}(alarmtime)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
