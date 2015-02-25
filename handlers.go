package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func dirHandler(w http.ResponseWriter, req *http.Request) {
	path := filepath.Join(root, req.URL.Path)
	if _, err := os.Stat(path); err == nil {
		dir := NewDirectory(req.URL.Path)
		headerTmpl.Execute(w, &dir)
		dirTmpl.Execute(w, &dir)
		footerTmpl.Execute(w, &dir)
	} else {
		w.WriteHeader(404)
		io.WriteString(w, "<h1>Not Found</h1>")
	}
}

func streamHandler(w http.ResponseWriter, req *http.Request) {
	path := filepath.Join(root, req.URL.Path[7:])
	if f, err := os.Open(path); err == nil {
		defer f.Close()
		w.Header().Set("Content-Type", "application/octet-stream")
		buffer := make([]byte, 100)
		for {
			n, err := f.Read(buffer)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}
			w.Write(buffer)
		}
	} else {
		w.WriteHeader(404)
		io.WriteString(w, err.Error())
	}
}
