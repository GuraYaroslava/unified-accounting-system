define(function() {

    function postRequest(data, callback, url) {
        $.ajax({
            method: "post",
            type: "post",
            dataType: "json",
            url: url,
            data: JSON.stringify(data),
            ContentType: "application/json; charset=utf-8",
            success: function(data) {
                console.log(data);
                callback(data);
            },
            error: function(ajaxRequest, ajaxOptions, thrownError) {
                console.log(thrownError);
            }
        });
    };

    return {
        postRequest: postRequest
    };

});