$('#new-publication').on('submit', createPublication);

$(document).on('click', '.like-publication', likePublication);
$(document).on('click', '.unlike-publication', unlikePublication);

$('#update-publication').on('click', updatePublication);
$('.delete-publication').on('click', deletePublication);

function createPublication(event) {
    event.preventDefault();

    const title = $('#title').val();
    const content = $('#content').val();

    $.ajax({
        url: "/publications",
        method: "POST",
        data: {
            title: title,
            content: content,
        }
    }).done(() => {
        window.location = "/home";
    }).fail((err) => {
        if (err && err.status >= 400) {
            Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: 'Something went wrong!'
            });
        } else {
            window.location = "/home";
        }
    });
}

function likePublication(event) {
    event.preventDefault();

    const element = $(event.target);
    const publicationId = element.closest('div').data('publication-id');

    element.prop('disabled', true);
    $.ajax({
        url: `/publications/${publicationId}/like`,
        method: "POST",
    }).done(() => {
        const countingLikes = element.next('span');
        const qtdLikes = parseInt(countingLikes.text());

        countingLikes.text(qtdLikes + 1);
        element.addClass('unlike-publication');
        element.addClass('text-danger');
        element.removeClass('like-publication');
    }).fail(() => {
        Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Something went wrong!'
        });
    }).always(() => {
        element.prop('disabled', false);
    });
}

function unlikePublication(event) {
    event.preventDefault();

    const element = $(event.target);
    const publicationId = element.closest('div').data('publication-id');

    element.prop('disabled', true);
    $.ajax({
        url: `/publications/${publicationId}/unlike`,
        method: "POST",
    }).done(() => {
        const countingLikes = element.next('span');
        const qtdLikes = parseInt(countingLikes.text());

        countingLikes.text(qtdLikes - 1);
        element.removeClass('unlike-publication');
        element.removeClass('text-danger');
        element.addClass('like-publication');
    }).fail(() => {
        Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Something went wrong!'
        });
    }).always(() => {
        element.prop('disabled', false);
    });
}

function updatePublication() {

    const publicationId = $(this).data('publication-id');
    const title = $('#title').val();
    const content = $('#content').val();

    $(this).prop('disabled', true);
    $.ajax({
        url: `/publications/${publicationId}`,
        method: "PUT",
        data: {
            title: title,
            content: content,
        },
    }).done(() => {
        Swal.fire({
            title: 'Success!',
            text: 'Publication updated successfully!',
            icon: 'success'
        }).then(() => {
            window.location = '/home';
        });
    }).fail(() => {
        Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Something went wrong!'
        });
    }).always(() => {
        $('#update-publication').prop('disabled', false);
    });
}

function deletePublication(event) {
    event.preventDefault();

    Swal.fire({
        icon: 'warning',
        title: 'Attention!',
        text: 'Are you sure you want to delete this post? This action is irreversible!',
        showCancelButton: true,
        cancelButtonText: "Cancel"
    }).then((confirmation) => {
        if (!confirmation.value) return;

        const element = $(event.target);
        const publication = element.closest('div');
        const publicationId = publication.data('publication-id');

        element.prop('disabled', true);
        $.ajax({
            url: `/publications/${publicationId}`,
            method: "DELETE",
        }).done(() => {
            publication.fadeOut('slow', () => {
                $(this).remove();
            });
        }).fail(() => {
            Swal.fire({
                icon: 'error',
                title: 'Oops...',
                text: 'Something went wrong!'
            });
        }).always(() => {
            element.prop('disabled', false);
        });
    });
}