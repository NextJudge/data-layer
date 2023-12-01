package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/rs/cors"
	"goji.io"
	"goji.io/pat"
)

var db *Database

func main() {
	var port = flag.String("p", "5000", "port")
	var bind = flag.String("bind", "*", "interface to bind on")

	db, err := NewDatabase()
	if err != nil {
		log.Fatalln("error initializing db")
	}

	mux := goji.NewMux()
	c := cors.New(cors.Options{
		AllowedOrigins: cfg.CORSOrigin,
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowedMethods: []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		MaxAge:         3600,
	})
	mux.Use(c.Handler)

	mux.HandleFunc(pat.Get("/v1/users"), getUsers)

	var addr string
	if *bind == "*" {
		addr = ":" + *port
	} else {
		addr = *bind + ":" + *port
	}
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalln("http listener error")
	}
}
