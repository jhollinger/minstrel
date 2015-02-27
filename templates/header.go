package templates

const Header string = `
<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Music</title>
    <script src="//code.jquery.com/jquery-2.1.3.min.js"></script>
    <script src="//cdnjs.cloudflare.com/ajax/libs/underscore.js/1.8.2/underscore-min.js"></script>
    <script src="//cdnjs.cloudflare.com/ajax/libs/backbone.js/1.1.2/backbone-min.js"></script>
    <link href='http://fonts.googleapis.com/css?family=Share+Tech+Mono' rel='stylesheet' type='text/css'>
    <style type="text/css">
      html, body {
        margin: 0px;
        height: 100%;
        font-family: sans-serif;
        cursor: default;
      }

      #player {
        position: relative;
        height: 35%;
        background: linear-gradient(180deg, rgb(100,100,100), rgb(70,70,70));
      }

      #player h2 {
        position: absolute;
        top: 50%;
        width: 100%;
        margin: -1em 0px 0px 0px;
        padding: 0px;
        text-align: center;
        font-size: 2em;
        font-weight: normal;
        font-family: 'Share Tech Mono', monospace;
        color: #e5e5e5;
        text-shadow: 2px 2px 2px rgb(30,30,30);
      }

      #controls {
        position: absolute;
        bottom: 0px;
        width: 100%;
        background-color: #151515; /* for Chrome */
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
        font-size: 0.8em;
        background: linear-gradient(180deg, rgb(180,180,180), rgb(150,150,150));
        box-shadow: 0px 0px 3px #555;
      }

      #breadcrumbs > li {
        list-style-type: none;
        display: inline-block;
        color: #f9f9f9;
      }

      #list {
        position: absolute;
        top: 22px;
        right: 0px;
        bottom: 0px;
        left: 0px;
        margin: 0px;
        padding: 0px;
        overflow-y: auto;
        font-size: 0.9em;
      }

      #list > li {
        list-style-type: none;
        padding: 3px;
        color: #444;
        marign-right: 5px;
      }

      #list > li:nth-child(even) {
        background-color: #e7e7e7;
      }

      #list > li:active, #list > li.active {
        background-color: #25b;
        color: #eee;
      }
    </style>
  </head>
  <body>
`
