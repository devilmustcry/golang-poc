package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	u := url.URL{
		Scheme:   "ws",
		Host:     "localhost:8080",
		Path:     "/ws",
		RawQuery: fmt.Sprintf("name=%s", strings.TrimSuffix(name, "\n")),
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	done := make(chan struct{})
	msg := make(chan string)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		defer close(msg)
		for {
			fmt.Print("Enter msg: ")
			text, _ := reader.ReadString('\n')
			msg <- text
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case <-done:
			fmt.Printf("DONE")
			return
		case <-ticker.C:
			c.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println(err)
				return
			}
		case t := <-msg:
			fmt.Printf("Sending: %s", t)
			err := c.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
