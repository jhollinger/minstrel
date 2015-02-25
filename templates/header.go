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
      .pane {
        float: left;
        width: 50%;
      }

      #player {
        width: 100%;
      }

      .list {
        margin: 0px;
        padding: 0px;
      }

      .list > li {
        list-style-type: none;
      }

      .list > li > a {
        color: #444;
        text-decoration: none;
      }

      #playlist {
        background: #eee;
      }

      #playlist > li > a {
        padding: 5px;
        display: block;
      }

      #playlist > li > a:hover {
        background-color: #ddd;
      }
    </style>
  </head>
  <body>
`
