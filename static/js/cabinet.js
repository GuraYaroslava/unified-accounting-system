define(["jquery"], function($, utils) {

    function User() {
        this.keys = ["fname", "lname", "pname", "login", "email", "phone", "address"];
    };

    function drawCabinet(data) {
        if (data != null) {
            $("#server-answer").css("visibility", "hidden");
            $("#title").text("Личный кабинет");
            var user = new User();
            $("<form/>", {
                id: "form-private"
            }).appendTo("#content");

            for (var i in user.keys) {
                var div = $("<div/>", {
                    class: "form-row"
                });

                var lable = $("<label/>", {
                    for: user.keys[i],
                    id: "label-"+user.keys[i]
                });

                var input = $("<input/>", {
                    id: user.keys[i],
                    name: user.keys[i],
                    type: "text"
                });

                $(input).val(data[user.keys[i]]);
                $(div).append(lable);
                $(div).append(input);
                $("#form-private").append(div);
                $("#label-"+user.keys[i]).text(data.tableData.Fields[user.keys[i]].Caption);
            }

            var input = $("<input/>", {
                id: "save-btn",
                name: "save-btn",
                type: "button",
                value: "Сохранить изменения"
            }).appendTo("#form-private");

        } else {
            $("#server-answer").text("Провал.").css("color", "red");
        }
    };

    return {
        drawCabinet: drawCabinet
    };

});