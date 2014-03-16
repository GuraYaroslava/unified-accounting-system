define(["jquery"], function($) {

    var sid;

    function registerCallback(data) {

        var serverAnswer = $("#server-answer");

        if (data == null)
        {
            serverAnswer.text("Сервер не отвечает.").css("color", "red");

        } else if (data.result === "ok") {
            console.log("Регистрация прошла успешно.");
            serverAnswer.text("Registration is successful.").css("color", "green");

        } else if (data.result === "loginExists") {
            serverAnswer.text("Такой логин уже существует.");

        } else if (data.result === "badLogin") {
            serverAnswer.text("Логин может содержать буквы и/или "
                + "цифры и иметь длину от 2 до 36 символов.").css("color", "red");

        } else if (data.result === "badPassword") {
            serverAnswer.text("Пароль должен иметь длину от 2 "
                + "до 36 символов.").css("color", "red");
        }
    }

    function loginCallback(data) {
        if (data.result === "ok") {
            $("#server-answer").text("Авторизация прошла успешно.").css("color", "green");
            $("#logoutBtn").css("visibility", "visible");
            sid = data.sid;

        } else if (data.result === "invalidCredentials") {
            $("#server-answer").text("Неверный логин или пароль.").css("color", "red");
        }
    }

    function logoutCallback(data) {
        if (data.result === "ok") {
            $("#server-answer").text("Вы вышли.").css("color", "green");
            $("#logoutBtn").css("visibility", "hidden");

        } else if (data.result === "badSid") {
            $("#server-answer").text("Invalid session ID.").css("color", "red");
        }
    }

    function jsonHandle(action, callback) {
        if (action == "logout") {
            var js = {
                "action": "logout",
                "sid": sid
            }

        } else {
            var js = {
                "action": action,
                "login": $("#username").val(),
                "password": $("#password").val(),
            }
        }

        console.log(js);
        $.ajax({
            method: "post",
            type: "post",
            dataType: "json",
            url: "/handler",
            data: JSON.stringify(js),
            ContentType: "application/json; charset=utf-8",
            success: function(data) {
                document.getElementById("form-register").reset();
                $("#server-answer").empty();
                callback(data);
            },
            error: function(ajaxRequest, ajaxOptions, thrownError) {
                console.log(thrownError);
            }
        });

    }

    return {
        registerCallback: registerCallback,
        loginCallback: loginCallback,
        logoutCallback: logoutCallback,
        jsonHandle: jsonHandle
    }

});