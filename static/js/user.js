define(["jquery", "auth", "utils"],
function($, auth, utils) {

    var fields = [
        "fname",
        "lname",
        "pname",
        "login",
        "email",
        "phone",
        "address"
    ];

    function drawCabinet(data) {
        if (data != null) {
            $("#server-answer").css("visibility", "hidden");
            $("#cabinet, #title").css("visibility", "visible");
            $("#title").text("Личный кабинет");
            $("#cabinet").empty();

            $("<form/>", {
                id: "form-private"
            }).appendTo("#cabinet");

            for (var i in fields) {
                var div = $("<div/>", {
                    class: "form-row"
                });

                var lable = $("<label/>", {
                    for: fields[i],
                    id: "label-"+fields[i]
                });

                var input = $("<input/>", {
                    id: fields[i],
                    name: fields[i],
                    type: "text"
                });

                $(input).val(data[fields[i]]);
                $(div).append(lable).append(input);
                $("#form-private").append(div);
                $("#label-"+fields[i]).text(data.tableData.Fields[fields[i]].Caption);
            }
            var input = $("<input/>", {
                type: "button",
                value: "Сохранить изменения",
                id: "save-btn",
                name: "submit",
            }).appendTo("#form-private");

            $("#save-btn").click(function() {
                var data = {
                    "action": "updateUser",
                    "id": auth.getId(),
                    "data": getPersonData()
                };
                utils.postRequest(data, updateCallback, "/handler");
            });

        } else {
            $("#server-answer").text("Провал.").css("color", "red");
        }
    };

    function getPersonData() {
        var userData = {};
        for (var i in fields) {
            userData[fields[i]] = $("#"+fields[i]).val();
        }
        return userData;
    };

    function updateCallback(data) {
        $("#cabinet, #title").css("visibility", "hidden");
        $("#server-answer").css("visibility", "visible");
        if (data == null) {
            $("#server-answer").text("Сервер не отвечает.").css("color", "red");
        } else if (data.result === "ok") {
            $("#server-answer").text("Изменения сохранены.").css("color", "green");
        }
    };

    return {
        drawCabinet: drawCabinet,
        getPersonData: getPersonData,
        updateCallback: updateCallback
    };
});