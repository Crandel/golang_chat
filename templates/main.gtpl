{{define "title" }}
Main Chat page
{{end}}

{{define "head" }}
    <link href="/static/css/chat.css" rel="stylesheet">
{{end}}
{{define "content" }}
<div id='subscribe'>
    {{range .Mess}}
    <div id="{{.ID}}" data-user-id="{{.User.ID}}">[{{.CreatedAt.Format "15:04:05"}}] <span class="black">{{.User.Login}}:</span> {{.Message}}</div>
    {{end}}
</div>
    <p>
        <input type="text" id="message" autocomplete="off" data-id={{.ID}}>
        <input class="btn btn-primary" type="submit" id='submit' value="Send">
        <input class="btn btn-default" type="button", id="signout", value="Sign Out">
    </p>
{{end}}

{{define "footer" }}
    <script src="/static/js/main.js"></script>
{{end}}
