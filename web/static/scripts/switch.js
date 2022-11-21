$(document).ready(async function() {
    const $startupSwitch = $("#settings_startupSwitch");
    const $startupSwitchTogglerName = $("#settings_startupTogglerName");

    const $wallpaperSwitch = $("#settings_autoSetWallpaperSwitch");
    const $wallpaperSwitchTogglerName = $("#settings_autoSetWallpaperTogglerName");

    const $loggingSwitch = $("#settings_saveLoggSwitch");
    const $loggingSwitchTogglerName = $("#settings_saveLoggTogglerName");

    getSettings().then((data) => {
       if (data["wallpaper"] === 1) {
           $wallpaperSwitch.attr("checked", "true");
           $wallpaperSwitchTogglerName.text("On");
       } else if (data["startup"] === 1) {
           $startupSwitch.attr("checked", "true");
           $startupSwitchTogglerName.text("On");
       } else if (data["save_logg"] === 1) {
           $loggingSwitch.attr("checked", "true");
           $loggingSwitchTogglerName.text("On");
       }
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

    $loggingSwitch.click(async function() {
        $.ajax({
            url: "http://localhost:3000/api/get/settings",
            type: "GET",
            success: function (data) {
                if (data["save_logg"] === 1) {
                    $.ajax({
                        url: "http://localhost:3000/api/update/settings",
                        type: "POST",
                        data: {
                            "save_logg": 0,
                        },
                        success: function (data) {
                            if(data["status"]) {
                                $loggingSwitchTogglerName.text("Off");
                                $loggingSwitch.removeAttr("checked");

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
                            "save_logg": 1,
                        },
                        success: function (data) {
                            if(data["status"]) {
                                $loggingSwitchTogglerName.text("On");
                                $loggingSwitch.attr("checked", "true");

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
});

/**
 * @param {String} message
 */

function toast(message) {
    if (message === null) {
        return "Required parameter 'message' is missing.";
    }

    $(".toast-body").text(message);
    let toastLiveExample = document.getElementById('liveToast');
    let toast = new bootstrap.Toast(toastLiveExample);
    toast.show();
}

/**
 * @param {Number} i
 */

function editStartup(i) {
    if (i !== 1 || i !== 0 || i === null) {
        return "Required parameter 'i' is missing.";
    }

    return $.ajax({
        url: "http://localhost:3000/api/update/startup",
        type: "POST",
        data: {
            "startup": i
        },
    });
}

function getSettings() {
    return $.ajax({
        url: "http://localhost:3000/api/get/settings",
        type: "GET",
    });
}