$('#signup').on('submit', createUser);

function createUser(event) {
    event.preventDefault();

    const name = $('#name').val();
    const email = $('#email').val();
    const nick = $('#nick').val();
    const password = $('#password').val();
    const confirmPassword = $('#confirmPassword').val();

    if (password !== confirmPassword) {
        Swal.fire({
            icon: 'warning',
            title: 'Oops...',
            text: 'The password are diferent!'
        });
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: name,
            nick: nick,
            email: email,
            password: password,
        }
    }).done(() => {
        Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'User created successfully!'
        }).then(() => {
            $.ajax({
                url: "/login",
                method: 'POST',
                data: {
                    email: email,
                    password: password,
                }
            }).done(() => {
                window.location = '/home';
            }).fail(() => {
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: 'Something went wrong!'
                });
            });
        });
    }).fail(() => {
        Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Something went wrong!'
        });
    });
};