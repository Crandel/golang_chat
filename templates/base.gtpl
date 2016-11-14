<!DOCTYPE html>
<html lang="en">
  <head>
    <title>{{block "title" .}}Simple chat{{end}}</title>
    <meta charset="UTF-8">
    <link rel="icon" type="image/png" href="/static/img/favicon.png">
    <link href="/static/css/reset.css" rel="stylesheet">
    <link href="/static/css/vendor/bootstrap.min.css" rel="stylesheet">
    <script src="/static/js/vendor/jquery-1.12.0.min.js"></script>
    <script src="/static/js/vendor/bootstrap.min.js"></script>
    {{block "head" .}}{{end}}
  </head>
  <body>
    {{block "content" .}}{{end}}
    <footer>
      <span>Property of Vitaliy Drevenchuk 2016</span>
      {{block "footer" .}}{{end}}
    </footer>
  </body>
</html>
