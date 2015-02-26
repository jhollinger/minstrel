package templates

const Header string = `
<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Music</title>
    <script src="//code.jquery.com/jquery-2.1.3.min.js"></script>
    <script>
      function Player(element, tracks) {
        this.audio = element;
        this.tracks = tracks;
        this.currentTrack = 0;
        this.audio.addEventListener('ended', this.next.bind(this));
      }

      Player.prototype.play = function(track) {
        this.currentTrack = track;
        this.audio.src = this.tracks[track];
        this.audio.pause();
        this.audio.load();
        this.audio.play();
      };

      Player.prototype.next = function() {
        if ( this.currentTrack < this.tracks.length-1 ) {
          this.play(this.currentTrack + 1);
        }
      };
    </script>
    <style type="text/css">
      html, body {
        margin: 0px;
        height: 100%;
      }

      #player {
        position: relative;
        height: 35%;
        background-color: #555;
      }

      #controls {
        position: absolute;
        bottom: 0px;
        width: 100%;
      }

      #library {
        position: relative;
        height: 65%;
        background-color: #eee;
      }

      #breadcrumbs {
        position: fixed;
        width: 100%;
        margin: 0px;
        padding: 3px;
        background-color: #999;
        font-size: 0.8em;
      }

      #breadcrumbs > li {
        list-style-type: none;
        display: inline-block;
      }

      #breadcrumbs > li > a {
        color: #eee;
        text-decoration: none;
      }

      #list {
        position: absolute;
        top: 20px;
        right: 0px;
        bottom: 0px;
        left: 0px;
        margin: 0px;
        padding: 0px;
        overflow-y: auto;
      }

      #list > li {
        list-style-type: none;
      }

      #list > li > a {
        display: block;
        padding: 3px;
        color: #444;
        text-decoration: none;
      }

      #list > li:nth-child(even) {
        background-color: #e7e7e7;
      }
    </style>
  </head>
  <body>
`
