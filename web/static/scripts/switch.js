$.ajax({
    url: "http://localhost:8080/api/get/settings",
    type: "GET",
    success: function(data){
        if(data["autochangewallpaper"] === 1) {
            $("#autosetWallpaperSwitch").attr("checked", "true");
            $("#autosetWallpaperText").text("On");
        }
        if(data["startup"] === 1) {
            $("#startupSwitch").attr("checked", "true");
            $("#startupText").text("On");
        }
    }
});

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
                    success: function(data){
                        if(data.status) {
                            $("#autosetWallpaperSwitch").removeAttr("checked");
                            $("#autosetWallpaperText").text("Off");

                            $(".toast-body").text(data.message);
                            let toastLiveExample = document.getElementById('liveToast');
                            let toast = new bootstrap.Toast(toastLiveExample);
                            toast.show();
                        }
                    }
                });
            } else {
                $.ajax({
                    url: "http://localhost:8080/api/update/settings",
                    type: "POST",
                    data: {
                        autochangewallpaper: 1
                    },
                    success: function(data){
                        if(data.status) {
                            $("#autosetWallpaperSwitch").attr("checked", "true");
                            $("#autosetWallpaperText").text("On");

                            $(".toast-body").text(data.message);
                            let toastLiveExample = document.getElementById('liveToast');
                            let toast = new bootstrap.Toast(toastLiveExample);
                            toast.show();
                        } else {
                            $(".toast-body").text("Could not remove the program from autorun.");
                            let toastLiveExample = document.getElementById('liveToast');
                            let toast = new bootstrap.Toast(toastLiveExample);
                            toast.show();
                        }
                    }
                });
            }
        }
    });
});

$("#startupSwitch").click(function() {
    $.ajax({
        url: "http://localhost:8080/api/get/settings",
        type: "GET",
        success: function (data) {
            if (data["startup"] === 1) {
                $.ajax({
                    url: "http://localhost:8080/api/update/settings",
                    type: "POST",
                    data: {
                        startup: 0
                    },
                    success: function () {
                        $.ajax({
                            url: "http://localhost:8080/api/update/del/startup",
                            type: "POST",
                            success: function(data){
                                if(data.status) {
                                    $("#startupSwitch").removeAttr("checked");
                                    $("#startupText").text("Off");

                                    $(".toast-body").text(data.message);
                                    let toastLiveExample = document.getElementById('liveToast');
                                    let toast = new bootstrap.Toast(toastLiveExample);
                                    toast.show();
                                } else {
                                    $(".toast-body").text("Failed to apply settings.");
                                    let toastLiveExample = document.getElementById('liveToast');
                                    let toast = new bootstrap.Toast(toastLiveExample);
                                    toast.show();
                                }
                            }
                        });
                    }
                });
            } else {
                $.ajax({
                    url: "http://localhost:8080/api/update/settings",
                    type: "POST",
                    data: {
                        startup: 1
                    },
                    success: function () {
                        $.ajax({
                          url: "http://localhost:8080/api/update/set/startup",
                          type: "POST",
                          success: function(data){
                              if(data.status) {
                                  $("#startupSwitch").attr("checked", "true");
                                  $("#startupText").text("On");

                                  $(".toast-body").text(data.message);
                                  let toastLiveExample = document.getElementById('liveToast');
                                  let toast = new bootstrap.Toast(toastLiveExample);
                                  toast.show();
                              } else {
                                    $(".toast-body").text("Failed to apply settings.");
                                    let toastLiveExample = document.getElementById('liveToast');
                                    let toast = new bootstrap.Toast(toastLiveExample);
                                    toast.show();
                              }
                          }
                        });
                    }
                });
            }
        }
    });
});