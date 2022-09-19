$.ajax({
    url: "http://localhost:8080/api/get/settings",
    type: "GET",
    success: function(data){
        if(data["autoupdate"] === 1){
            $("#updateSwitch").attr("checked", "true");
            $("#updateText").text("On");
        }
        if(data["autostart"] === 1) {
            $("#autorunSwitch").attr("checked", "true");
            $("#autorunText").text("On");
        }
        if(data["autochangewallpaper"] === 1) {
            $("#autosetWallpaperSwitch").attr("checked", "true");
            $("#autosetWallpaperText").text("On");
        }
    }

})

$("#updateSwitch").click(function(){
    $.ajax({
        url: "http://localhost:8080/api/get/settings",
        type: "GET",
        success: function(data){
            if(data["autoupdate"] === 1){
                $.ajax({
                    url: "http://localhost:8080/api/update/settings",
                    type: "POST",
                    data: {
                        autoupdate: 0
                    },
                    success: function(){
                        $("#updateSwitch").removeAttr("checked");
                        $("#updateText").text("Off");
                    }
                })
            } else {
                $.ajax({
                    url: "http://localhost:8080/api/update/settings",
                    type: "POST",
                    data: {
                        autoupdate: 1
                    },
                    success: function(){
                        $("#updateSwitch").attr("checked", "true");
                        $("#updateText").text("On");
                    }
                })
            }
        }
    });
})


$("#autorunSwitch").click(function(){
    $.ajax({
        url: "http://localhost:8080/api/get/settings",
        type: "GET",
        success: function(data){
            if(data["autostart"] === 1){
                $.ajax({
                    url: "http://localhost:8080/api/update/settings",
                    type: "POST",
                    data: {
                        autostart: 0
                    },
                    success: function(){
                        $("#autorunSwitch").removeAttr("checked");
                        $("#autorunText").text("Off");
                    }
                })
            } else {
                $.ajax({
                    url: "http://localhost:8080/api/update/settings",
                    type: "POST",
                    data: {
                        autostart: 1
                    },
                    success: function(){
                        $("#autorunSwitch").attr("checked", "true");
                        $("#autorunText").text("On");
                    }
                })
            }
        }
    });
})

$("#autosetWallpaperSwitch").click(function(){
    $.ajax({
        url: "http://localhost:8080/api/get/settings",
        type: "GET",
        success: function(data){
            if(data["autochangewallpaper"] === 1){
                $.ajax({
                    url: "http://localhost:8080/api/update/settings",
                    type: "POST",
                    data: {
                        autochangewallpaper: 0
                    },
                    success: function(){
                        $("#autosetWallpaperSwitch").removeAttr("checked");
                        $("#autosetWallpaperText").text("Off");
                    }
                })
            } else {
                $.ajax({
                    url: "http://localhost:8080/api/update/settings",
                    type: "POST",
                    data: {
                        autochangewallpaper: 1
                    },
                    success: function(){
                        $("#autosetWallpaperSwitch").attr("checked", "true");
                        $("#autosetWallpaperText").text("On");
                    }
                })
            }
        }
    });
})