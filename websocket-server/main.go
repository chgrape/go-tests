package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	cons map[*websocket.Conn]bool
}

func create_server() *Server {
	return &Server{
		cons: make(map[*websocket.Conn]bool),
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Heiiii")
	io.WriteString(w, "Hello world")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	s := create_server()

	http.HandleFunc("/", test)
	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		(*s).cons[conn] = true
		fmt.Println("Connected")
		fmt.Println(s)
	})
	http.ListenAndServe(":3000", nil)

}
