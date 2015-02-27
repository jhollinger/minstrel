package main

import (
	"os"
	"path/filepath"
	"strings"
)

type Directory struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Dirs   []Node `json:"dirs"`
	Tracks []Node `json:"tracks"`
}

type Node struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func NewDirectory(path string) *Directory {
	children, _ := filepath.Glob(filepath.Join(root, path, "*"))
	dirs, files := []Node{}, []Node{}
	for i := 0; i < len(children); i++ {
		childName := strings.Replace(filepath.Base(children[i]), filepath.Ext(children[i]), "", 1)
		fi, _ := os.Stat(children[i])
		if fi.IsDir() {
			childPath := strings.Replace(children[i], root, "", 1)
			dirs = append(dirs, Node{childName, childPath})
		} else {
			childPath := strings.Replace(children[i], root, "/file", 1)
			files = append(files, Node{childName, childPath})
		}
	}
	return &Directory{filepath.Base(path), path, dirs, files}
}
