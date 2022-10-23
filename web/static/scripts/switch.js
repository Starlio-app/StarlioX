$(document).ready(async function() {
    const $startupSwitch = $("#settings_startupSwitch");
    const $startupSwitchTogglerName = $("#settings_startupTogglerName");

    const $wallpaperSwitch = $("#settings_autoSetWallpaperSwitch");
    const $wallpaperSwitchTogglerName = $("#settings_autoSetWallpaperTogglerName");

    $.ajax({
        url: "http://localhost:3000/api/get/settings",
        type: "GET",
        success: function(data) {
            console.log(data);
            if (data["wallpaper"] === 1) {
                $wallpaperSwitch.attr("checked", "true");
                $wallpaperSwitchTogglerName.text("On");
            }
            if (data["startup"] === 1) {
                $startupSwitch.attr("checked", "true");
                $startupSwitchTogglerName.text("On");
            }
        },
    });

    $wallpaperSwitch.click(async function() {
        $.ajax({
            url: "http://localhost:3000/api/get/settings",
            type: "GET",
            success: function (data) {
                if (data["wallpaper"] === 1) {
                    $.ajax({
                        url: "http://localhost:3000/api/update/settings",
                        type: "POST",
                        data: {
                            "wallpaper": 0,
                        },
                        success: function (data) {
                            if(data["status"]) {
                                $wallpaperSwitchTogglerName.text("Off");
                                $wallpaperSwitch.removeAttr("checked");

                                toast(data.message);
                            } else {
                                toast("Failed to apply settings.");
                            }
                        },
                    });
                } else {
                    $.ajax({
                        url: "http://localhost:3000/api/update/settings",
                        type: "POST",
                        data: {
                            "wallpaper": 1,
                        },
                        success: function (data) {
                            if(data["status"]) {
                                $wallpaperSwitchTogglerName.text("On");
                                $wallpaperSwitch.attr("checked", "true");

                                toast(data.message);
                            } else {
                                toast("Failed to apply settings.");
                            }
                        },
                    });
                }
            },
        })
    });


    $startupSwitch.click(async function() {
        $.ajax({
            url: "http://localhost:3000/api/get/settings",
            type: "GET",
            success: function (data) {
                if (data["startup"] === 1) {
                    $.ajax({
                        url: "http://localhost:3000/api/update/settings",
                        type: "POST",
                        data: {
                            "startup": 0,
                        },
                        success: async function (data) {
                            if (data["status"]) {
                                await editStartup(0);

                                $startupSwitchTogglerName.text("Off");
                                $startupSwitch.removeAttr("checked");

                                toast(data.message);
                            } else {
                                toast("Failed to apply settings.");
                            }
                        },
                    });
                } else {
                    $.ajax({
                        url: "http://localhost:3000/api/update/settings",
                        type: "POST",
                        data: {
                            "startup": 1,
                        },
                        success: async function (data) {
                            if (data["status"]) {
                                await editStartup(1);

                                $startupSwitchTogglerName.text("On");
                                $startupSwitch.attr("checked", "true");

                                toast(data.message);
                            } else {
                                toast("Failed to apply settings.");
                            }
                        },
                    });
                }
            },
        });
    });
});

function toast(message) {
    $(".toast-body").text(message);
    let toastLiveExample = document.getElementById('liveToast');
    let toast = new bootstrap.Toast(toastLiveExample);
    toast.show();
}

function editStartup(i) {
    return $.ajax({
        url: "http://localhost:3000/api/update/startup",
        type: "POST",
        data: {
            "startup": i
        },
    });
}