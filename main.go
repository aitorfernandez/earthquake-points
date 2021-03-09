package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8010, "server port")
}

// Handler handles "/" requests.
func Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	z := parseInt(vars["z"])
	x := parseInt(vars["x"])
	y := parseInt(vars["y"])

	fmt.Println(">>>", z, x, y)
}

func main() {
	getAddr := func() string {
		return fmt.Sprintf(":%d", port)
	}

	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/{z:\\d+}/{x:\\d+}/{y:\\d+}.png", Handler)
	srv := &http.Server{
		Addr:    getAddr(),
		Handler: r,
	}
	log.Print("starting server on" + getAddr())
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server.ListenAndServe %v", err)
	}
}
