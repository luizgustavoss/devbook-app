$('#form-create-user').on('submit', createUser);

function createUser(event) {
    event.preventDefault();

    if ($('#password').val() !== $('#confirm-password').val()){
        alert("As senhas não são iguais")
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $('#name').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            password: $('#password').val()
        },
    }).done(function(data){
        alert("Usuário cadastrado com sucesso!" + data)
    }).fail(function(error){
        console.log(error)
        alert("Falha ao cadastrar usuário!")
    });
}
