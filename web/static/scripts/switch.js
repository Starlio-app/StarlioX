$.ajax({
    url: "http://localhost:8080/api/get/settings",
    type: "GET",
    success: function(data){
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
                        $.ajax({
                            url: "http://localhost:8080/api/update/del/startapp",
                            type: "POST",
                            success: function(data){
                                $(".toast-body").text(data.message);
                                let toastLiveExample = document.getElementById('liveToast')
                                let toast = new bootstrap.Toast(toastLiveExample)
                                toast.show()
                                if(data.status) {
                                    $("#autorunSwitch").removeAttr("checked");
                                    $("#autorunText").text("Off");
                                }
                            }
                        })
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
                        $.ajax({
                            url: "http://localhost:8080/api/update/add/startapp",
                            type: "POST",
                            success: function(data){
                                $(".toast-body").text(data.message);
                                let toastLiveExample = document.getElementById('liveToast')
                                let toast = new bootstrap.Toast(toastLiveExample)
                                toast.show()
                                if(data.status) {
                                    $("#autorunSwitch").attr("checked", "true");
                                    $("#autorunText").text("On");
                                }
                            }
                        })
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