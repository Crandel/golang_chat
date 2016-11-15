$(document).ready(function(){
    try{
        var sock = new WebSocket('ws://' + window.location.host + '/ws');
    }
    catch(err){
        var sock = new WebSocket('wss://' + window.location.host + '/ws');
    }
    var colors = ['green', 'red', 'yellow', 'blue'];
    var fillCollor = Object.create(null);
    function createMessageBox(ob){
        var date = new Date(),
            options = {hour12: false},
            message = [
            "<div id=" + ob.id + " data-user-id=" + ob.user_id + " class='chat-div-line clearfix'>",
            " <div class='text-user-inline'>",
            "  <time>[" + date.toLocaleTimeString('en-US', options) + "]</time>",
            "  <span class='userlogin " + ob.colorname + "'>" + ob.username + "</span>",
            "  <span>" + ob.message + "</span>",
            " </div>"],
            buttons = [
            " <div class='button-right'>",
            "  <button type='button' class='edit btn btn-warning btn-sm'>",
            "   <span class='glyphicon glyphicon-pencil' aria-hidden=\"true\"></span> Edit",
            "  </button>",
            "  <button type='button' class='remove btn btn-danger btn-sm'>",
            "   <span class='glyphicon glyphicon-remove' aria-hidden=\"true\"></span> Remove",
            "  </button>",
            " </div>",
            "</div>"
            ],
            array = message;
        if (ob.id != 0){
            array = $.merge(array, buttons);
        }
        var m = $(array.join('\n'));
        return m;
    }

    // show message in div#subscribe
    function showMessage(ob) {
        var messageElem = $('#subscribe'),
            height = 0,
            m = createMessageBox(ob);
        messageElem.append(m);
        messageElem.find('div').each(function(i, value){
            height += parseInt($(this).height());
        });

        messageElem.animate({scrollTop: height});
    }

    function getSystemMessage(m){
        return {'username': 'System','id': 0, 'user_id': 0, 'colorname':'black', 'message': m};
    }
    function sendMessage(){
        var msg = $('#message'),
            obj = {'message': msg.val(), 'user_id': msg.data('id')};
        sock.send(JSON.stringify(obj));
        msg.val('').focus();
    }

    sock.onopen = function(){
        showMessage(getSystemMessage('Connection to server started'));
    };

    // send message from form
    $('#submit').click(function() {
        sendMessage();
    });

    $('#message').keyup(function(e){
        if(e.keyCode == 13){
            sendMessage();
        }
    });

    $('body').on('click', '.edit', function(){
        var par = $(this).parent().parent(),
            id = par[0].id;
        console.log(par);
        console.log(id, "id");
        var modal = $('#editModal');
        modal.data("id", id);
        modal.css('display', 'block');
    });

    $('body').on('click', '.remove', function(){
        var par = $(this).parent().parent(),
            id = par[0].id;
        console.log(id, "id");
        $.delete("/message/" + id)
        .done(function(data){
            console.log(data, 'success');
        })
        .fail(function(data){
            console.log(data, 'fail');
        });
    });

    $('#saveButton').click(function(){
        var modal = $('#editModal'),
            id = modal.data()['id'],
            message=$(this).prev().prev().val();
        $.post("/message/" + id, {'message': message})
            .done(function(data) {
                console.log(data, 'edit success');
                modal.css('display', 'none');
            })
            .fail(function(data){
                console.log(data, 'edit fail');
            });
    });
    // income message handler
    sock.onmessage = function(event) {
        var obj = JSON.parse(event.data);
        if (obj.user_id in fillCollor){
            obj.colorname = fillCollor[obj.user_id];
        }else{
            var color = colors[Math.floor(Math.random() * colors.length)];
            obj.colorname = color;
            fillCollor[obj.user_id] = color;
        }
        showMessage(obj);
    };

    $('#signout').click(function(){
        window.location.href = "signout";
    });

    sock.onclose = function(event){
        if(event.wasClean){
            showMessage(getSystemMessage('Clean connection end'));
        }else{
            showMessage(getSystemMessage('Connection broken'));
        }
    };

    sock.onerror = function(error){
        showMessage(getSystemMessage(error));
    };
});
