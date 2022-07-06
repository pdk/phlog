package main

import (
	"flag"
	"log"
	"net/http"
)

type ServerConfig struct {
	RootPath string
}

func main() {

	conf := ServerConfig{}

	flag.StringVar(&conf.RootPath, "root", "./docs", "path to directory containing files to serve")
	flag.Parse()

	log.Printf("starting server")

	root := http.FileServer(http.Dir(conf.RootPath))

	log.Fatal(http.ListenAndServe(":8888", root))
}
