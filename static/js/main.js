require(["auth", "utils"], function(auth, utils) {

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
        //utils.postRequest(data, user.drawCabinet, "/handler/selectById/Users/" + auth.getId());
        location.href = "/handler/selectById/Users/" + auth.getId();
    });

});