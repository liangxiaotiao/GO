package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{Addr: ":8890", WriteTimeout: 2 * time.Second}
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)
	server.Handler = mux
	log.Println("Starting server ....嘻嘻")
	log.Fatal(server.ListenAndServe())
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello,I love all for you 龙骚骚" + r.URL.String()))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye ,龙骚骚 I don't love you anymore" + r.URL.String()))
}
