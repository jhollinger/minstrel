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
  </head>
  <body>
`
