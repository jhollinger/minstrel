# minstrel

A browser-based music library written in Go and Backbone.js.

I use my Chromebook around the house and lot. Unfortunately there's no DLNA client for ChromeOS, so I can't use it to play music from my home media server. (Writing one may have been the better solution, but since it doesn't already exist I'm assuming it's difficult. Also, I've been wanting to write something useful with Go.) So I threw together a simple music library and player in HTML5, backed by Go.

## Use

    minstrel --music ~/Music --port 8080

Then open your browser to http://hostname:8080. Check the wiki for init script examples.

## Download

Check the [releases page](https://github.com/jhollinger/minstrel/releases) for a pre-built binary.

## Build

    go get github.com/jhollinger/minstrel
    cd $GOPATH/src/github.com/jhollinger/minstrel
    go build
