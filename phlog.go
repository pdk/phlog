package main

import (
	"flag"
	"log"
	"net/http"
)

type ServerConfig struct {
	RootPath   string
	ListenAddr string
	ListenPort string
}

func main() {

	conf := ServerConfig{}

	flag.StringVar(&conf.ListenAddr, "listen", "", "ip address to listen for new connections")
	flag.StringVar(&conf.ListenPort, "port", "8888", "port to listen for new connections")
	flag.StringVar(&conf.RootPath, "root", "./docs", "path to directory containing files to serve")
	flag.Parse()

	log.Printf("starting server at %s:%s", conf.ListenAddr, conf.ListenPort)

	root := http.FileServer(http.Dir(conf.RootPath))

	log.Fatal(http.ListenAndServe(conf.ListenAddr+":"+conf.ListenPort, root))
}
