$(document).ready(function(){

    $('#login').click(function(){
        window.location.href = "login";
    });

    function showError(error){
        $('#error').html(error);
    }

    $('#submit').click(function(){
        var login = $('#login_text').val(),
            email = $('#email').val(),
            password = $('#password'),
            password2 = $('#password2');
        $('#error').empty();
        if(password.val() === password2.val()){
            if (login && email && password){
                $.post('sign', {'login': login, 'email': email, 'password': password.val()})
                .done(function(data){
                    window.location.href = '/';
                })
                .fail(function(data){
                    showError(data.responseText);
                });
            }else{
                showError('Please fill all fields');
            }
        }else{
            showError('Passwords must be the same');
        }
    });
});
