package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "server address")

func ServeHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "NotFound", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", ServeHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Printf("Starting server: %s", *addr)
}
