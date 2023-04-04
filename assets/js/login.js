$('#login').on('submit', doLogin);

function doLogin(event) {
    event.preventDefault();

    const email = $('#email').val();
    const password = $('#password').val();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: email,
            password: password,
        }
    }).done(() => {
        window.location = "/home";
    }).fail(() => {
        alert("Invalid user or password")
    });
}