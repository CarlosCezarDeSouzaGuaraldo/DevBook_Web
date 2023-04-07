$('#unfollow').on('click', unfollowUser);
$('#follow').on('click', followUser);

function unfollowUser() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/unfollow`,
        method: 'POST'
    }).done(() => {
        window.location = `/users/${userId}`;
    }).fail(() => {
        Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Something went wrong!'
        });
        $(this).prop('disabled', false);
    });
}

function followUser() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/follow`,
        method: 'POST'
    }).done(() => {
        window.location = `/users/${userId}`;
    }).fail(() => {
        Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Something went wrong!'
        });
        $(this).prop('disabled', false);
    });
}