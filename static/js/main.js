require(["jquery", "auth"], function($, auth) {

    $("#registerBtn").click(function() {
        console.log("registerBtn");
        auth.jsonHandle("register", auth.registerCallback);
    });

    $("#loginBtn").click(function() {
        auth.jsonHandle("login", auth.loginCallback);
    });

    $("#logoutBtn").click(function() {
        auth.jsonHandle("logout", auth.logoutCallback);
    });

});