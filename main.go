package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"time"
	"bufio"
)

func main() {
	server4()
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

// responds with Hello World
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		// ...
		w.WriteHeader(http.StatusNoContent)
		return
	}
	response := helloWorldResponse{Message: "Hello World"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("panicking")
	}
	fmt.Fprint(w, string(data))
}

// tcp server that echoes every message from client
func server1() {
	ln,  _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		conn.Write([]byte(msg + "\n"))
	}
}

// tcp server that sends random bytes to client every half second
func server4() {
	ln, _ := net.Listen("tcp", ":8082")
	conn, _ := ln.Accept()
	for {
		time.Sleep(500 * time.Millisecond)
		randBytes := make([]byte, 4)
		rand.Read(randBytes)
		conn.Write(randBytes)
	}
}

