{{define "title" }}
Main Chat page
{{end}}

{{define "head" }}
    <link href="/static/css/chat.css" rel="stylesheet">
{{end}}
{{define "content" }}
<div id='subscribe'>
    {{range .}}
    <p>[{{.CreatedAt.Format "15:04:05"}}] {{.User.Login}}: {{.Message}}</p>
    {{end}}
</div>
    <p>
        <input  type="text" id="message" autocomplete="off">
        <input class="btn btn-primary" type="submit" id='submit' value="Send">
        <input class="btn btn-default" type="button", id="signout", value="Sign Out">
    </p>
{{end}}

{{define "footer" }}
    <script src="/static/js/main.js"></script>
{{end}}
