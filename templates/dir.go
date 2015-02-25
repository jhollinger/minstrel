package templates

const Dir string = `
{{range .Dirs}}
<p><a href="{{.Path}}">{{.Name}}</a></p>
{{end}}

{{if .Tracks}}
<h2>{{.Name}}</h2>
<audio id="player" src="{{.FirstTrack.StreamPath}}" controls></audio>
<ul>
{{range $index, $element := .Tracks}}
  <li><a href="#{{$index}}" class="track" data-track="{{$index}}">{{.Name}}</a></li>
{{end}}
</ul>

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
