package templates

const Player string = `
<div id="player">
  <div id="player-info"></div>
  <audio id="controls" controls></audio>
</div>

<div id="library"></div>

<script id="player-info-template" type="text/template">
  <h2><%= name %></h2>
</script>

<script id="library-template" type="text/template">
  <ul id="breadcrumbs">
    <li class="dir" data-path="/">All</li>
    <% _.each(breadcrumbs(), function(crumb) { %>
    <li class="dir" data-path="<%= crumb.path %>">&gt; <%= crumb.name %></li>
    <% }) %>
  </ul>

  <ul id="list">
  <% _.each(get('dirs'), function(dir) { %>
    <li class="dir" data-path="<%= dir.path %>"><%= dir.name %></li>
  <% }) %>
  <% _.each(get('tracks'), function(track, i) { %>
    <li class="track" data-track="<%= i %>" data-path="<%= track.path %>"><%= track.name %></li>
  <% }) %>
  </ul>
</script>

<script>
  $(function() {
    var Directory = Backbone.Model.extend({
      idAttribute: "path",

      urlRoot: '/dir',

      breadcrumbs: function() {
        var path = this.get('path').substring(1);
        var crumbs = path ? path.split('/') : [];
        return crumbs.map(function(crumb, i) {
          return {name: crumb, path: crumbs.slice(0, i+1).join('/')};
        });
      }
    });

    var LibraryView = Backbone.View.extend({
      el: $('#library'),

      events: {
        "click .dir": "openDir",
        "click .track": "playTrack"
      },

      initialize: function() {
        this.model = new Directory({path: '/'});
        this.listenTo(this.model, 'change:dirs', this.render);
        this.model.fetch();
      },

      template: _.template($('#library-template').html()),

      render: function() {
        this.$el.html(this.template(this.model));
      },

      openDir: function(e) {
        this.model.set('path', $(e.target).attr('data-path'));
        this.model.fetch();
      },

      playTrack: function(e) {
        player.tracks = this.model.get('tracks');
        player.play(parseInt($(e.target).attr('data-track')));
      }
    });

    var PlayerInfoView = Backbone.View.extend({
      el: $('#player-info'),

      initialize: function() {
        this.track = {};
      },

      set: function(track) {
        this.track = track;
        this.render();
      },

      template: _.template($('#player-info-template').html()),

      render: function() {
        this.$el.html(this.template(this.track));
      }
    });

    new LibraryView();
    var player = new Player($('#controls').get(0), new PlayerInfoView());
  });

  function Player(element, infoView) {
    this.audio = element;
    this.info = infoView;
    this.tracks = [];
    this.currentTrack = 0;
    this.audio.addEventListener('ended', this.next.bind(this));
  }

  Player.prototype.play = function(i) {
    this.currentTrack = i;
    this.audio.src = this.tracks[i].path;
    this.audio.pause();
    this.audio.load();
    this.audio.play();
    this.info.set(this.tracks[i]);
  };

  Player.prototype.next = function() {
    if ( this.currentTrack < this.tracks.length-1 ) {
      this.play(this.currentTrack + 1);
    }
  };
</script>
`
