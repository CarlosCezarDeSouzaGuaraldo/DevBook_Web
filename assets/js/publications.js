$('#new-publication').on('submit', createPublication);

$(document).on('click', '.like-publication', likePublication);
$(document).on('click', '.unlike-publication', unlikePublication);

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
        alert("Error liking publication");
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
        alert("Error liking publication");
    }).always(() => {
        element.prop('disabled', false);
    });
}