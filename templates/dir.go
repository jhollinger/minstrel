package templates

const Dir string = `
<div id="player">
  <audio id="controls" src="{{if .Tracks}}{{.FirstTrack.StreamPath}}{{end}}" controls></audio>
</div>

<div id="library">
  <ul id="breadcrumbs">
    <li><a href="/">All</a></li>
  </ul>

  <ul id="list">
  {{range .Dirs}}
    <li><a href="{{.Path}}">{{.Name}}</a></li>
  {{end}}
  {{range $index, $element := .Tracks}}
    <li><a href="#{{$index}}" class="track" data-track="{{$index}}">{{.Name}}</a></li>
  {{end}}
  </ul>
</div>

<script>
  $(function() {
    var player = new Player($('#controls').get(0), {{.StreamPaths}});
    $('.track').click(function() {
      player.play(parseInt($(this).attr('data-track')));
      return false;
    });
  });
</script>
`
