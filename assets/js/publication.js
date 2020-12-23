$('#form-new-publication').on('submit', publication);
$(document).on('click', '.like-publication', likePublication)
$(document).on('click', '.unlike-publication', unlikePublication)
$('#update-publication').on('click', editPublication);
$('.delete-publication').on('click', deletePublication);

function publication(event) {
    event.preventDefault();

    $.ajax({
        url: "/publications",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        },
    }).done(function(data){
        window.location = "/home";
    }).fail(function(error){
        console.log(error)
        Swal.fire( "Ops!", "Falha ao cadastrar a publicação!", "error" );
    });
}

function deletePublication(event) {
    event.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Confirma a exclusão? Esta ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmation){
        if(!confirmation.value) return;

        const clickedElement = $(event.target)
        const publication = clickedElement.closest('div')
        const publicationId = publication.data('publication-id');

        $.ajax({
            url: `/publications/${publicationId}`,
            method: "DELETE"
        }).done(function(data){
            publication.fadeOut("slow", function(){
                $(this).remove();
            });
        }).fail(function(error){
            console.log(error)
            Swal.fire( "Ops!", "Falha ao excluir a publicação!", "error" );
        });
    });
}

function editPublication(event) {
    $('#update-publication').prop('disabled', true);

    const publicationId = $(this).data('publication-id')

    $.ajax({
        url: `/publications/${publicationId}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        },
    }).done(function(data){
        Swal.fire(
            'Sucesso!',
            'Publicação editada com sucesso!',
            'success'
        ).then(function(){
            window.location = "/home";
        });
    }).fail(function(error){
        Swal.fire( "Ops!", "Falha ao editar a publicação!", "error" );
    }).always(function(){
        $('#update-publication').prop('disabled', false);
    });
}

function likePublication(event) {
    event.preventDefault();

    const clickedElement = $(event.target)
    const publicationId = clickedElement.closest('div').data('publication-id');
    clickedElement.prop('disabled', true);

    $.ajax({
        url: `/publications/${publicationId}/like`,
        method: "POST"
    }).done(function(){

        const likesCounter = clickedElement.next('span');
        const likesCount = parseInt(likesCounter.text());

        likesCounter.text(likesCount+1)

        clickedElement.addClass('unlike-publication')
        clickedElement.addClass('text-danger')
        clickedElement.removeClass('like-publication')

    }).fail(function(error){
        Swal.fire( "Ops!", "Falha ao curtir a publicação!", "error" );
    }).always(function(){
        clickedElement.prop('disabled', false);
    });
}

function unlikePublication(event) {
    event.preventDefault();

    const clickedElement = $(event.target)
    const publicationId = clickedElement.closest('div').data('publication-id');
    clickedElement.prop('disabled', true);

    $.ajax({
        url: `/publications/${publicationId}/unlike`,
        method: "POST"
    }).done(function(){

        const likesCounter = clickedElement.next('span');
        const likesCount = parseInt(likesCounter.text());

        likesCounter.text(likesCount-1)

        clickedElement.addClass('like-publication')
        clickedElement.removeClass('text-danger')
        clickedElement.removeClass('unlike-publication')

    }).fail(function(error){
        Swal.fire( "Ops!", "Falha ao descurtir a publicação!", "error" );
    }).always(function(){
        clickedElement.prop('disabled', false);
    });
}