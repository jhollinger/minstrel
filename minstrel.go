package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port string
var root string

func main() {
	flag.Parse()

	http.HandleFunc("/", playerHandler)
	http.HandleFunc("/dir/", dirHandler)
	http.HandleFunc("/file/", fileHandler)

	fmt.Println("Starting on port " + port)
	fmt.Println("Music library " + root)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func init() {
	flag.StringVar(&root, "music", os.Getenv("HOME")+"/Music", "Music library path")
	flag.StringVar(&port, "port", "8080", "Port")
}
