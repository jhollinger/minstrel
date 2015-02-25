package main

import (
	"flag"
	"fmt"
	"github.com/jhollinger/minstrel/templates"
	"html/template"
	"log"
	"net/http"
	"os"
)

var port string
var root string

var headerTmpl *template.Template
var footerTmpl *template.Template
var dirTmpl *template.Template

func main() {
	flag.Parse()

	http.HandleFunc("/", dirHandler)
	http.HandleFunc("/stream/", streamHandler)

	fmt.Println("Starting on port " + port)
	fmt.Println("Music library " + root)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func init() {
	flag.StringVar(&root, "music", os.Getenv("HOME")+"/Music", "Music library path")
	flag.StringVar(&port, "port", "8080", "Port")

	var err error
	headerTmpl, err = template.New("header").Parse(templates.Header)
	if err != nil {
		panic(err)
	}
	footerTmpl, err = template.New("footer").Parse(templates.Footer)
	if err != nil {
		panic(err)
	}
	dirTmpl, err = template.New("dir").Parse(templates.Dir)
	if err != nil {
		panic(err)
	}
}
