$('#new-publication').on('submit', createPublication);

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
        alert("Publication created successfully");
        window.location = "/home";
    }).fail(() => {
        alert("Error creating publication");
    });
}