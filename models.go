package main

import (
	"os"
	"path/filepath"
	"strings"
)

type Directory struct {
	Path   string
	Dirs   []Node
	Tracks []Node
}

func NewDirectory(path string) *Directory {
	children, _ := filepath.Glob(filepath.Join(root, path, "*"))
	dirs, files := []Node{}, []Node{}
	for i := 0; i < len(children); i++ {
		child := strings.Replace(children[i], root, "", 1)
		fi, _ := os.Stat(children[i])
		if fi.IsDir() {
			dirs = append(dirs, Node{child})
		} else {
			files = append(files, Node{child})
		}
	}
	return &Directory{path, dirs, files}
}

func (this Directory) Name() string {
	return filepath.Base(this.Path)
}

func (this Directory) FirstTrack() *Node {
	return &this.Tracks[0]
}

func (this Directory) StreamPaths() []string {
	paths := []string{}
	for _, track := range this.Tracks {
		paths = append(paths, track.StreamPath())
	}
	return paths
}

type Node struct {
	Path string
}

func (this Node) StreamPath() string {
	return "/stream" + this.Path
}

func (this Node) Name() string {
	name := filepath.Base(this.Path)
	return strings.Replace(name, filepath.Ext(name), "", 1)
}
