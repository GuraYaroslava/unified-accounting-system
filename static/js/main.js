require(["auth", "utils"],
function(auth, utils) {

    $(document).ready(function() {
        utils.postRequest(
            {
                "action": "select",
                "table": "Contests",
                "fields": ["id", "name"]
            },
            listSubjects,
            "/handler"
        );

        utils.postRequest(
            {
                "action": "getId",
            },
            showCabinet,
            "/handler"
        );
    });

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
        location.href = "/handler/selectById/Users/";
    });

    $("#home-btn").click(function() {
        location.href = "/";
    });

    function listSubjects(data) {
        for (i in data) {
            $("<a/>", {
                text: data[i].name,
                href: "/handler/ShowBlank/"+data[i].id,
                class: "form-row",
            }).appendTo("div#list-contests");
        }
    }

    function showCabinet(data) {
        console.log(data)
        if (data.id != null) {
            $("#logout-btn, #cabinet-btn").css("visibility", "visible");
        } else {
            $("#logout-btn, #cabinet-btn").css("visibility", "hidden");
        }
    }

});