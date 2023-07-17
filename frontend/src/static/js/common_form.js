function onSuccess(response) {
    if (response.status === 'success') {
        alert(response.message);
        if (response.redirect_url) {
            location.href = response.redirect_url;
        } else {
            location.reload();
        }
    } else {
        alert(response.message);
        location.reload();
    }
}

function onError(response) {
    alert(response.status)
    alert("Error: " + response.message);
    location.reload();
}

function submitForm(formName, url, successCallback, errorCallback) {
    var form = $('form[name="' + formName + '"]');
    var data = form.serialize();

    $.ajax({
        type: 'POST',
        url: url,
        data: data,
        success: successCallback,
        error: errorCallback
    });
}

$(document).ready(function () {});

