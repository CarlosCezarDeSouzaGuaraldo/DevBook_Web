$('#unfollow').on('click', unfollowUser);
$('#follow').on('click', followUser);
$('#edit-user').on('submit', editUser);
$('#update-password').on('submit', updatePassword);

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

function editUser(event) {
    event.preventDefault();

    const name = $('#name').val();
    const email = $('#email').val();
    const nick = $('#nick').val();

    $.ajax({
        url: "/edit-user",
        method: "PUT",
        data: {
            name: name,
            email: email,
            nick: nick,
        }
    }).done(() => {
        Swal.fire("Success!", "User updated successfully!", "success")
            .then(function () {
                window.location = "/profile";
            });
    }).fail(() => {
        Swal.fire("Ops...", "Something went wrong!", "error");
    });
}

function updatePassword(event) {
    event.preventDefault();

    const password = $('#new-password').val();
    const current = $('#currently-password').val();
    const confirmPassword = $('#confirm-password').val();

    if (password !== confirmPassword) {
        Swal.fire({
            icon: 'warning',
            title: 'Oops...',
            text: 'The password are diferent!'
        });
        return;
    }

    $.ajax({
        url: "/update-password",
        method: "POST",
        data: {
            new: password,
            current: current
        }
    }).done(() => {
        Swal.fire({
            title: 'Success!',
            text: 'Password updated successfully!',
            icon: 'success'
        }).then(() => {
            window.location = '/profile';
        });
    }).fail(function () {
        Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Something went wrong!'
        });
    });

}