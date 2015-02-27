package main

import (
	"encoding/json"
	"fmt"
	"github.com/jhollinger/minstrel/templates"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func playerHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", strconv.Itoa(len(templates.Header)+len(templates.Player)+len(templates.Footer)))
	io.WriteString(w, templates.Header)
	io.WriteString(w, templates.Player)
	io.WriteString(w, templates.Footer)
}

func dirHandler(w http.ResponseWriter, req *http.Request) {
	relPath := req.URL.Path[4:]
	fullPath := filepath.Join(root, relPath)
	if _, ferr := os.Stat(fullPath); ferr == nil {
		dir := NewDirectory(relPath)
		if data, jerr := json.Marshal(dir); jerr == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", strconv.Itoa(len(data)))
			w.Write(data)
		} else {
			w.WriteHeader(500)
			io.WriteString(w, jerr.Error())
		}
	} else {
		w.WriteHeader(404)
	}
}

func fileHandler(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open(filepath.Join(root, req.URL.Path[5:]))

	if err != nil {
		w.WriteHeader(404)
		io.WriteString(w, err.Error())
		return
	}

	defer f.Close()
	var startByte, endByte int64
	fi, _ := f.Stat()
	startByte = 0
	endByte = fi.Size() - 1
	statusCode := http.StatusOK

	if r := req.Header.Get("Range"); r != "" {
		byteRange := strings.Split(strings.Replace(r, "bytes=", "", 1), "-")
		startByte, _ = strconv.ParseInt(byteRange[0], 10, 64)
		if len(byteRange) > 1 && byteRange[1] != "" {
			endByte, _ = strconv.ParseInt(byteRange[1], 10, 64)
		}
		statusCode = http.StatusPartialContent
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.FormatInt(endByte-startByte+1, 10))
	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", startByte, endByte, fi.Size()))
	w.WriteHeader(statusCode)

	buffer := make([]byte, endByte-startByte+1)
	n, rErr := f.ReadAt(buffer, startByte)
	if rErr != nil && rErr != io.EOF {
		panic(rErr)
	}
	if n > 0 {
		w.Write(buffer)
	}
}
