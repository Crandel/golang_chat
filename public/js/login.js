$(document).ready(function(){

    $('#sign').click(function(){
        window.location.href = "sign"
    });

    function showError(error){
        $('#error').html(error)
    }

    $('#submit').click(function(){
        var login = $('#login').val(),
            password = $('#password').val();
        $('#error').empty();
        if(login && password){
            $.post('login', {'login': login, 'password': password}, function(data){
                if (data.error){
                    showError(data.error)
                }else{
                    window.location.href = '/'
                }
            });
        }else{
            showError('Please fill all fields')
        }
    });
});
