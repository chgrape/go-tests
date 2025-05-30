package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct{
	cons map[*websocket.Conn]bool
}

func create_server()*Server{
	return &Server{
		cons: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) accept_connection(ws *websocket.Conn){
	s.cons[ws] = true
}

func main(){
	s := create_server()

	http.Handle("/connect", websocket.)


	s.accept_connection()
}
