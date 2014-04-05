define(["utils"], 
function(utils) {

    var id;
    var sid;

    function getId() {
        return id;
    }

    function getSid() {
        return sid;
    }

    function registerCallback(data) {

        if (data == null) {
            $("#server-answer").text("Сервер не отвечает.").css("color", "red");

        } else if (data.result === "ok") {
            $("#server-answer").text("Регистрация прошла успешно.").css("color", "green");
            jsonHandle("login", loginCallback);

        } else if (data.result === "loginExists") {
            $("#server-answer").text("Такой логин уже существует.").css("color", "red");
            $("#password").val("");

        } else if (data.result === "badLogin") {
            $("#server-answer").text("Логин может содержать буквы и/или "
                + "цифры и иметь длину от 2 до 36 символов.").css("color", "red");
            $("#password").val("");

        } else if (data.result === "badPassword") {
            $("#server-answer").text("Пароль должен иметь длину от 2 "
                + "до 36 символов.").css("color", "red");
            $("#password").val("");
        }
    };

    function loginCallback(data) {
        if (data.result === "ok") {
            $("#server-answer").text("Авторизация прошла успешно.").css("color", "green");
            $("#logout-btn, #cabinet-btn").css("visibility", "visible");
            $("#form-register").css("visibility", "hidden");
            id = data.id;
            sid = data.sid;

        } else if (data.result === "invalidCredentials") {
            $("#server-answer").text("Неверный логин или пароль.").css("color", "red");
        }
    };

    function logoutCallback(data) {
        if (data.result === "ok") {
            $("#logout-btn, #cabinet-btn, #form-private, #title").css("visibility", "hidden");
            $("#password").val("");
            $("#server-answer").text("Вы вышли.").css("color", "green").css("visibility", "visible");
            $("#form-register").css("visibility", "visible");
        } else if (data.result === "badSid") {
            $("#server-answer").text("Invalid session ID.").css("color", "red");
        }
    };

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
            };
        }

        utils.postRequest(js, callback, "/handler");
    };

    return {
        getId: getId,
        getSid: getSid,
        registerCallback: registerCallback,
        loginCallback: loginCallback,
        logoutCallback: logoutCallback,
        jsonHandle: jsonHandle
    };

});