package main

import (
	"http01-project/internal/readfilehelper"
	"http01-project/internal/registerdurhelper"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	http.Handle("/metrics", promhttp.Handler())
	//var sites []string = readfilehelper.ReadLineOfFile("urllist")
	sites, _ := readfilehelper.ReadLineOfFile("urllist")
	go registerdurhelper.Register(sites)
	//http.ListenAndServe(":8081", nil)
	//log.Fatalln(http.ListenAndServe(":8081", nil))
	log.Fatal(http.ListenAndServe(":8081", nil))
	//handler := http.HandlerFunc(pla)
}
