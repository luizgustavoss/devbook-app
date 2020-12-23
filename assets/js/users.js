$('#form-create-user').on('submit', users);
$('#update-user').on('click', updateUser);
$('#change-password').on('click', changePassword);
$('#delete-user').on('click', deleteUser);
$('#follow').on('click', follow);
$('#unfollow').on('click', unfollow);


function users(event) {
    event.preventDefault();

    if ($('#password').val() !== $('#confirm-password').val()){
        Swal.fire( "Ops!", "As senhas não são iguais!", "error" );
        return;
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
        Swal.fire( "Sucesso!", "Usuário cadastrado com sucesso!", "success" )
            .then(function(){
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
                    Swal.fire( "Ops!", "Falha no login!", "error" );
                });
            });
    }).fail(function(error){
        Swal.fire( "Ops!", "Falha ao cadastrar usuário!", "error" );
    });
}

function changePassword(event) {

    if ($('#new-password').val() !== $('#confirm-password').val()){
        Swal.fire( "Ops!", "As senhas não são iguais!", "error" );
        return;
    }

    $.ajax({
        url: "/change-password",
        method: "POST",
        data: {
            oldPassword: $('#old-password').val(),
            newPassword: $('#new-password').val()
        },
    }).done(function(data){
        Swal.fire( "Sucesso!", "Senha alterada com sucesso!", "success" )
            .then(function(){
                window.location = "/profile";
            });
    }).fail(function(error){
        Swal.fire( "Ops!", "Falha ao alterar a senha!", "error" );
    });
}

function updateUser(event) {
    $('#update-user').prop('disabled', true);

    $.ajax({
        url: "/users",
        method: "PUT",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val()
        },
    }).done(function(data){
        Swal.fire(
            'Sucesso!',
            'Usuário editado com sucesso!',
            'success'
        ).then(function(){
            window.location = "/profile";
        });
    }).fail(function(error){
        Swal.fire( "Ops!", "Falha ao editar o usuário!", "error" );
    }).always(function(){
        $('#update-user').prop('disabled', false);
    });
}

function follow() {

    const userId = $(this).data("user-id");
    console.log("userId " + userId);

    $('#follow').prop('disabled', true)

    $.ajax({
        url: `/users/${userId}/follow`,
        method: "POST",
    }).done(function(){
        window.location = `/users/${userId}`;
    }).fail(function(){
        Swal.fire( "Ops!", "Falha ao seguir o usuário!", "error" );
        $('#follow').prop('disabled', false)
    })
}

function unfollow() {

    const userId = $(this).data("user-id");
    $('#unfollow').prop('disabled', true)

    $.ajax({
        url: `/users/${userId}/unfollow`,
        method: "POST",
    }).done(function(){
        window.location = `/users/${userId}`;
    }).fail(function(){
        Swal.fire( "Ops!", "Falha ao deixar de seguir o usuário!", "error" );
        $('#unfollow').prop('disabled', false)
    })
}

function deleteUser(event) {

    Swal.fire({
        title: "Atenção!",
        text: "Confirma a exclusão da conta? Esta ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmation){
        if(!confirmation.value) return;

        $.ajax({
            url: "/delete-account",
            method: "DELETE"
        }).done(function(data){
            Swal.fire("Sucesso!", "Sua conta foi excluída com sucesso!", "success").then(function(){
                window.location = "/logout"
            })
        }).fail(function(error){
            console.log(error)
            Swal.fire( "Ops!", "Falha ao excluir a conta!", "error" );
        });
    });
}