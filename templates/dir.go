package templates

const Dir string = `
<div class="pane">
  <ul class="list">
  {{range .Dirs}}
    <li><a href="{{.Path}}">{{.Name}}</a></li>
  {{end}}
  </ul>
</div>

{{if .Tracks}}
<div class="pane">
  <h2>{{.Name}}</h2>
  <audio id="player" src="{{.FirstTrack.StreamPath}}" controls></audio>
  <ul id="playlist" class="list">
  {{range $index, $element := .Tracks}}
    <li><a href="#{{$index}}" class="track" data-track="{{$index}}">{{.Name}}</a></li>
  {{end}}
  </ul>
</div>

<script>
  $(function() {
    var player = new Player($('#player').get(0), {{.StreamPaths}});
    $('.track').click(function() {
      player.play(parseInt($(this).attr('data-track')));
      return false;
    });
  });
</script>
{{end}}
`
