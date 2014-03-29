require(["jquery", "auth", "utils", "user"], function($, auth, utils, user) {

    $("#register-btn").click(function() {
        auth.jsonHandle("register", auth.registerCallback);
    });

    $("#login-btn").click(function() {
        auth.jsonHandle("login", auth.loginCallback);
    });

    $("#logout-btn").click(function() {
        auth.jsonHandle("logout", auth.logoutCallback);
    });

    $("#cabinet-btn").click(function() {
        var data = {
            "action": "identification",
            "id": auth.getId()
        }
        utils.postRequest(data, user.drawCabinet, "/handler");
    });

});