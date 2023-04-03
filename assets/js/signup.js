$('#signup').on('submit', createUser);

function createUser(event) {
    event.preventDefault();

    const name = $('#name').val();
    const email = $('#email').val();
    const nick = $('#nick').val();
    const password = $('#password').val();
    const confirmPassword = $('#confirmPassword').val();

    if (password !== confirmPassword) {
        alert("The password are diferent");
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
        alert("User signup successfully")
    }).fail((error) => {
        console.log(error);
        alert("Error")
    });
};