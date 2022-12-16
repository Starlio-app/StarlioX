const startDate = new Date();
const endDate = new Date();

const ids = [];
let id = 0;

const apiKEY = "1gI9G84ZafKDEnrbydviGknReOGiVK9jqrQBE3et";

function wallpaperLoad(data) {
    $(".preloader").hide();

    for (let i = 0; i < data.length; i++, id++) {
        ids.push(data[i]);

        if (ids.filter((item) => item['url'] === data[i]['url']).length > 1) {
            continue;
        }

        let image = new Image();
        image.src = ids[id]['media_type'] === "video" ?
            `https://img.youtube.com/vi/${ids[id]['url'].slice(30, 41)}/maxresdefault.jpg` :
            ids[id]['url'];

        image.onload = function() {
            if(image.width+image.height !== 210) {
                $(`img[data-src="${image.src}"]`).attr("src", `${image.src}`);
            }
        }

        if (ids[id]['media_type'] === "image") {
            $(".header-row").append(`
                <div class="card">
                    <a data-bs-toggle="modal" href="#WallpaperModal" role="button">
                        <img src="${ids[id]['url']}" alt="${ids[id]['title']}" class="card-img-top shimmer" id="${id}">
                    </a>
                </div>
        `);
        } else {
            $(".header-row").append(`
            <div class="card">
                <a data-bs-toggle="modal" href="#WallpaperModal" role="button">
                    <img src="http://localhost:3000/static/image/placeholder.png" 
                        data-src="https://img.youtube.com/vi/${ids[id]['url'].slice(30, 41)}/maxresdefault.jpg"
                        alt="${ids[id]['title']}" 
                        class="card-img-top shimmer" 
                        id="${id}"> 
                </a>
            </div>
        `);
        }
    }

    $(".card-img-top").on("load", function() {
        $(this).removeClass("shimmer");
    });

    const button_modal = document.querySelector(".header-row");
    button_modal.addEventListener("click", function (event) {
        if (event.target === button_modal) {
            return;
        }

        let id = event.target.getAttribute("id");
        let img = document.querySelector(".modal-body");
        let title = document.querySelector(".w-modal-title");

        let button = document.querySelector(".modal-footer");

        if($(`img#${id}.card-img-top`).attr("src") !== "http://localhost:3000/static/image/placeholder.png") {
            button.innerHTML = `<lottie-player id="favorite" 
                                    src="https://assets2.lottiefiles.com/private_files/lf30_rcvcAS.json" 
                                    background="transparent" 
                                    speed="1">
                                </lottie-player>
                                <button type="button" class="btn btn-primary" id="setWallpaper">Set Wallpaper</button>
            `;
        } else {
            button.innerHTML = `<lottie-player id="favorite" 
                                    src="https://assets2.lottiefiles.com/private_files/lf30_rcvcAS.json" 
                                    background="transparent" 
                                    speed="1">
                                </lottie-player>
                                <button type="button" 
                                    class="btn" 
                                    id="setWallpaper" 
                                    disabled 
                                    style="background-color: grey; color: white;">
                                    Set Wallpaper
                                </button>
            `;
        }

        const setWallpaper = document.querySelector("#setWallpaper");
        const favorite = document.querySelector("#favorite");

        $.ajax({
            url: "http://localhost:3000/api/get/favorites",
            type: "GET",
            data: {
                title: ids[id]['title'],
            },
            success: function (data) {
                if (data !== "No favorites found") {
                    favorite.setDirection(1);
                    favorite.play();
                }
            },
        })

        ids[id]['copyright'] = ids[id]['copyright'] === undefined ? "NASA" : ids[id]['copyright'];

        const explanation = ids[id]['explanation'].length > 200 ? ids[id]['explanation'].slice(0, 200) + "..." : ids[id]['explanation'];
        console.log(explanation);
        if (ids[id]['media_type'] === "image") {
            title.innerHTML = `<h5 class="modal-title">${ids[id]['title']}</h5>`;
            img.innerHTML = `
                <img src="${ids[id]['url']}" alt="${ids[id]['title']}" class="card-img">
                <p><strong>Author:</strong> ${ids[id]['copyright']}</p>
                <p><strong>Date of publication:</strong> ${ids[id]['date']}</p>
                <p><strong>Explanation:</strong> ${explanation}</p>
                
                <button type="button" class="show-more btn btn-primary" id="show-more">Show more</button>
            `;

            setWallpaper.addEventListener("click", function () {
                wallpaperUpdate(ids[id]['hdurl']);
            });
        } else {
            title.innerHTML = `<h5 class="modal-title">${ids[id]['title']}</h5>`;
            img.innerHTML = `
                <iframe width="460" 
                height="315" 
                src="${ids[id]['url']}" 
                title="YouTube video player" 
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" 
                allowFullScreen></iframe>
                
                <p><strong>Author:</strong> ${ids[id]['copyright']}</p>
                <p><strong>Date of publication:</strong> ${ids[id]['date']}</p>
                <p><strong>Explanation:</strong> ${explanation}</p>
                <button type="button" class="show-more btn btn-primary" id="show-more">Show more</button>
            `;

            setWallpaper.addEventListener("click", function () {
                wallpaperUpdate(`https://img.youtube.com/vi/${ids[id]['url'].slice(30, 41)}/maxresdefault.jpg`);
            });
        }

        favorite.addEventListener("click", function() {
            $.ajax({
                url: "http://localhost:3000/api/get/favorites",
                type: "GET",
                data: {
                    title: ids[id]['title'],
                },
                success: function(data) {
                    if (data === "No favorites found") {
                        if (ids[id]['media_type'] === "image") {
                            favoriteAdd(ids[id]['title'],
                                ids[id]['explanation'],
                                ids[id]['copyright'],
                                ids[id]['date'],
                                ids[id]['url'],
                                ids[id]['hdurl'],
                                ids[id]['media_type'],
                            );
                        } else {
                            favoriteAdd(ids[id]['title'],
                                ids[id]['explanation'],
                                ids[id]['copyright'],
                                ids[id]['date'],
                                `https://img.youtube.com/vi/${ids[id]['url'].slice(30, 41)}/0.jpg`,
                                `https://img.youtube.com/vi/${ids[id]['url'].slice(30, 41)}/maxresdefault.jpg`,
                                ids[id]['media_type'],
                            );
                        }

                        favorite.setDirection(1);
                        favorite.play()
                    } else {
                        favoriteRemove(ids[id]['title']);
                        favorite.setDirection(-1);
                        favorite.play();
                    }
                }
            });
        });

        let showMore = document.querySelector("#show-more");
        showMore.addEventListener("click", function () {
            let explanation = document.querySelector(".modal-body p:nth-child(4)");
            if(showMore.innerHTML === "Show more") {
                explanation.innerHTML = `<strong>Explanation:</strong> ${ids[id]['explanation']}`;
                showMore.innerHTML = "Show less";
            } else {
                explanation.innerHTML = `<strong>Explanation:</strong> ${ids[id]['explanation'].length > 200 ? ids[id]['explanation'].slice(0, 200) + "..." : ids[id]['explanation']}`;
                showMore.innerHTML = "Show more";
            }
        });
    });

    $(window).scroll(function () {
        if (($(window).scrollTop() > $(document).height() - $(window).height() - 100)) {
            $(".preloader").show();
            $(window).off("scroll");

            startDate.setDate(startDate.getDate() - 1);

            endDate.setDate(endDate.getDate() - 15);
            endDate.setMonth(endDate.getMonth() - 1);

            $.ajax({
                url: "https://api.nasa.gov/planetary/apod",
                data: {
                    api_key: apiKEY,
                    start_date: `${endDate.getUTCFullYear()}-${endDate.getUTCMonth() + 1}-${endDate.getUTCDate()}`,
                    end_date: `${startDate.getUTCFullYear()}-${startDate.getUTCMonth() + 1}-${startDate.getUTCDate()}`,
                },
                success: function (data) {
                    wallpaperLoad(data);
                    $(window).on("scroll");
                },
            });
        }
    });
}

/**
 * @param url {String} Url of the image to be set as wallpaper
 */

function wallpaperUpdate(url) {
    $.ajax({
        url: "http://localhost:3000/api/update/wallpaper",
        type: "POST",
        data: {
            url: url,
        },
        success: function (data) {
            $(".toast-body").text(data.message);
            let toastLiveExample = document.getElementById('liveToast');
            let toast = new bootstrap.Toast(toastLiveExample);
            toast.show();
        },
    });
}

/**
 * @param title {String} Title of the image
 * @param explanation {String} Explanation of the image
 * @param copyright {String} Author of the image
 * @param date {String} Date of publication
 * @param url {String} Url of the image
 * @param hdurl {String} Url of the image in high definition
 * @param media_type {String} Type of media
 */

function favoriteAdd(title, explanation, copyright, date, url, hdurl, media_type) {
    $.ajax({
        url: "http://localhost:3000/api/add/favorite",
        type: "POST",
        data: {
            title: title,
            explanation: explanation,
            date: date,
            url: url,
            hdurl: hdurl,
            media_type: media_type,
        },
        success: function (data) {
            $(".toast-body").text(data.message);
            let toastLiveExample = document.getElementById('liveToast');
            let toast = new bootstrap.Toast(toastLiveExample);
            toast.show();
        },
    });
}

/**
 * @param title {String} Title of the image
 */

function favoriteRemove(title) {
    $.ajax({
        url: "http://localhost:3000/api/del/favorite",
        type: "POST",
        data: {
            title: title,
        },
        success: function (data) {
            $(".toast-body").text(data.message);
            let toastLiveExample = document.getElementById('liveToast');
            let toast = new bootstrap.Toast(toastLiveExample);
            toast.show();
        },
    });
}