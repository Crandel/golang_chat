$(document).ready(function(){
    try{
        var sock = new WebSocket('ws://' + window.location.host + '/ws');
    }
    catch(err){
        var sock = new WebSocket('wss://' + window.location.host + '/ws');
    }
    var colors = ['green', 'red', 'yellow', 'blue'];
    var fillCollor = Object.create(null);
    // show message in div#subscribe
    function showMessage(ob) {
        var messageElem = $('#subscribe'),
            height = 0,
            date = new Date(),
            options = {hour12: false},
            username = $('<span/>', {'class': ob.colorname, 'html': ' ' + ob.username + ': '})[0].outerHTML,
            m = $('<div/>', {
                'id': ob.id,
                'data-user-id': ob.user_id,
                'html': '[' + date.toLocaleTimeString('en-US', options) + ']' + username + ob.message});
        messageElem.append(m);
        messageElem.find('div').each(function(i, value){
            height += parseInt($(this).height());
        });

        messageElem.animate({scrollTop: height});
    }

    function getSystemMessage(m){
        return {'username': 'System','id': 0, 'user_id': 0, 'classname':'black', 'message': m};
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
