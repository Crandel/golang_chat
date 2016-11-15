{{define "title" }}
Main Chat page
{{end}}

{{define "head" }}
<link href="/static/css/chat.css" rel="stylesheet">
{{end}}
{{define "content" }}
<div class="wrapper">
  <h1>Chat</h1>
  <div id='subscribe'>

    {{range .Mess}}
    <!-- message -->
    <div id="{{.ID}}" data-user-id="{{.User.ID}}"  class="chat-div-line clearfix">
      <div class="text-user-inline ">
        <time>[{{.CreatedAt.Format "15:04:05"}}]</time>
        <span class="userlogin black"> {{.User.Login}}:</span>
        <span>{{.Message}}</span>
      </div>
      {{if eq $.AuthUser .User.ID}}
      <div class='button-right'>
        <button type='button' class='edit btn btn-warning btn-sm'>
          <span class='glyphicon glyphicon-pencil' aria-hidden="true"></span> Edit
        </button>
        <button type='button' class='remove btn btn-danger btn-sm'>
          <span class='glyphicon glyphicon-remove' aria-hidden="true"></span> Remove
        </button>
      </div>
      {{end}}
    </div>
    {{end}}
  </div>
  <div class="wrapper-input clearfix">
    <div class="signout-div">
      <input type="text" id="message" autocomplete="off" data-id={{.AuthUser}}>
      <input class="btn btn-primary" type="submit" id='submit' value="Send">
      <input class="btn btn-default" type="button", id="signout", value="Sign Out">
    </div>
  </div>
  <div id="editModal" class="modal">
    <!-- Modal content -->
    <div class="modal-content">
      <h4>Please edit message</h4>
      <textarea name="message" id="textarea-edit" rows="2" cols="50"></textarea></br>
      <button id="saveButton" type='button' class='btn btn-success btn-sm'>
        <span class='glyphicon glyphicon-ok' aria-hidden="true"></span> Save
      </button>
    </div>
  </div>
</div>
{{end}}

{{define "footer" }}
    <script src="/static/js/main.js"></script>
{{end}}
