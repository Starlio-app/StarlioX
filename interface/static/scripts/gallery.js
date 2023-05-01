const startDate = new Date();
const endDate = new Date();

const ids = [];
let id = ids.length;
const apiKEY = "1gI9G84ZafKDEnrbydviGknReOGiVK9jqrQBE3et";

function wallpaperLoad(data) {
    $(".preloader").hide();

    for (let i = 0; id < data.length; i++, id++) {
        ids.push(data[i]);
        if (ids.filter((item) => item['url'] === data[i]['url']).length > 1) continue;

        const image = new Image();
        image.src = ids[id]['media_type'] === "video" ?
            `https://ytproxy.dc09.ru/vi/${ids[id]['url'].slice(30, 41)}/maxresdefault.jpg?host=i.ytimg.com` :
            ids[id]['url'];

        image.onload = function () {
            if (image.width + image.height !== 210) {
                $(`img[data-src="${image.src}"]`)
                    .attr("src", `${image.src}`)
                    .removeAttr("data-src");
            }
        }

        if (ids[id]['media_type'] === "image") {
            $(".header-row").append(`
            <div class=card>
                <a data-bs-toggle="modal" href="#wallpaper" role="button">
                    <img src="${ids[id]['url']}" 
                    title="${ids[id]['title']}" 
                    alt="${ids[id]['title']}"
                    class="card-img-top shimmer"
                    id="${id}">
                </a>
            </div>
        `)
        } else {
            $(".header-row").append(`
            <div class=card>
                <a data-bs-toggle="modal" href="#wallpaper" role="button">
                    <img src="/static/assets/placeholder.png" 
                    data-src="https://ytproxy.dc09.ru/vi/${ids[id]['url'].slice(30, 41)}/maxresdefault.jpg?host=i.ytimg.com"
                    title="${ids[id]['title']}" 
                    alt="${ids[id]['title']}"
                    class="card-img-top shimmer"
                    id="${id}">
                </a>
            </div>
        `)
        }
    }


    $(".card-img-top").on("load", function () {
        $(this).removeClass("shimmer");
    });

    const button_modal_window = document.querySelector(".header-row");

    button_modal_window.addEventListener("click", function (event) {
        if (event.target === button_modal_window) return;

        const card_id = event.target.getAttribute("id");
        const modale_window = document.querySelector(".modal-body");
        const title = document.querySelector(".w-modal-title");
        const button = document.querySelector(".modal-footer");

        button.innerHTML = `<lottie-player id="favorite" 
                                    src="/static/assets/lottie/lf30_rcvcAS.json" 
                                    background="transparent" 
                                    speed="1">
                            </lottie-player>
                            <button type="button" class="btn btn-primary" id="setWallpaper">Download</button>
            `;

        const wallpaper_img = $(`img#${id}.card-img-top`);
        if (wallpaper_img && wallpaper_img.src === "http://localhost:4000/static/assets/placeholder.png") {
            const button_setWallpaper = document.querySelector("#setWallpaper")
            button_setWallpaper.disabled = true;
            button_setWallpaper.style.backgroundColor = "grey";
            button_setWallpaper.style.color = "white;"
            button_setWallpaper.style.border = "none";
        }

        const setWallpaper = document.querySelector("#setWallpaper");
        const favorite = document.querySelector("#favorite");

        title.innerHTML = `<h5 class="modal-title">${ids[card_id]['title']}</h5>`;

        if (ids[card_id]['media_type'] === "image") {
            modale_window.innerHTML = `
                <img src="${ids[card_id]['url']}" alt="${ids[card_id]['title']}" class="card-img">
                <p id="explanation">${ids[card_id]['explanation']}</p>
            `;

            setWallpaper.addEventListener("click", function () {
                updateWallpaper(card_id[card_id]['hdurl']);
            });
        } else {
            const host = new URL(ids[card_id]['url']);

            if (["youtube.com", "youtu.be", "yt.be"].includes(host.hostname.replace("www.", ""))) {
                ids[card_id]['url'] = `https://yt.dc09.ru${host.pathname}`
            }

            modale_window.innerHTML = `
                <iframe style="display: block; margin: auto; border-radius: 13px;"
                width="440" 
                height="250"
                src=${ids[card_id]['url']}
                allowFullScreen></iframe>
                <p id="explanation">${ids[card_id]['explanation']}</p>
            `

            setWallpaper.addEventListener("click", function () {
                updateWallpaper(`https://ytproxy.dc09.ru/vi/${ids[card_id]['url'].substring(25)}/maxresdefault.jpg?host=i.ytimg.com`)
            });
        }

        if (ids[card_id]['copyright'] !== undefined) {
            return modale_window.innerHTML += `<p>© ${ids[card_id]['copyright']}</p>`;
        }
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

function updateWallpaper(url) {
    $.ajax({
        url: "http://localhost:4000/api/update/wallpaper",
        type: "POST",
        data: { url: url },
        success: function (data) {
            if (data.status) notify(data.message)
        },
    });
}

function notify(message) {
    if (!("Notification" in window)) return;

    else if (Notification.permission === "granted") {
        new Notification(message);
    }

    else if (Notification.permission !== "denied") {
        Notification.requestPermission(function (permission) {
            if (permission === "granted") {
                new Notification(message);
            }
        });
    }
}
