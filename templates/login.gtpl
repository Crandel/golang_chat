{{define "title"}}Login{{end}}
{{define "content"}}
<div class="page-header">
  <h1>{{template "title" .}}</h1>
</div>

<form autocomplete="on" method="post">
    <div id="error" style="color: red;"></div>
    <div class="form-group">
        <label>Insert login
            <input type="text" class="form-control" pattern="[A-Za-zА-Яа-я0-9]{3,}" placeholder="Login" id="login" required>
        </label>
    </div>

    <div class="form-group">
        <label>Enter password
            <input type="password" class="form-control" pattern="[A-Za-zА-Яа-я0-9]{8,}" placeholder="Password", id="password" required>
        </label>
    </div>
    <p>
        <input class="btn btn-primary" id="submit" title="Send" type="button" value="Send">
        <input class="btn btn-default" id="sign" title="Sign In" type="button" value="Sign In">
    </p>
</form>
{{end}}
{{define "footer"}}
<script src="/static/js/login.js"></script>
{{end}}
