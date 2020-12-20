$('#login').on('submit', login);

function login(event) {
    event.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val()
        },
    }).done(function(data){
        window.location = "/home"
    }).fail(function(error){
        console.log(error)
        alert("Falha no login!")
    });
}