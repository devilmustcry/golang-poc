package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"golang-poc/websocket/hub"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{}

func serveWs(h *hub.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &hub.Client{Hub: h, Conn: conn, Send: make(chan []byte, 256), Name: r.URL.Query().Get("name")}
	client.Hub.Register <- client

	go client.WaitingForMessage()
}

func main() {
	flag.Parse()
	hub := hub.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
