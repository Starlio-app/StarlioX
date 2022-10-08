let prev_date = Number(new Date().toLocaleString().slice(0,2))-1;
let today = new Date();
let endDate = new Date();

let ids = [];
let id = 0;

const apiKEY = "1gI9G84ZafKDEnrbydviGknReOGiVK9jqrQBE3et";

today.setDate(today.getDate() - 17);
$(document).ready(function() {
    $.ajax({
        url: "https://api.nasa.gov/planetary/apod",
        type: "GET",
        data: {
            api_key: apiKEY,
            start_date: `${today.getFullYear()}-${today.getMonth() + 1}-${today.getDate()}`,
            end_date: `${new Date().toLocaleString().slice(6, 10)}-${new Date().toLocaleString().slice(3, 5)}-${new Date().toLocaleString().slice(0, 2)}`,
        },
        success: function(data) {
            wallpaper(data);
        },
        error: function() {
            $.ajax({
                url: "https://api.nasa.gov/planetary/apod",
                data: {
                    api_key: apiKEY,
                    start_date: `${today.getFullYear()}-${today.getMonth()+1}-${today.getDate()}`,
                    end_date: `${new Date().toLocaleString().slice(6,10)}-${new Date().toLocaleString().slice(3, 5)}-${prev_date}`,
                },
                success: function (data) {
                    wallpaper(data);
                },
                error: function (data) {
                    console.error(data);
                },
            });
        },
    });
});

function wallpaper(data) {
    $(".preloader").hide();
    data = data.reverse();

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
                        <img src="${ids[id]['url']}" alt="${ids[id]['title']}" class="card-img-top" id="${id}">
                    </a>
                </div>
        `);
        } else {
            $(".header-row").append(`
            <div class="card">
                <a data-bs-toggle="modal" href="#WallpaperModal" role="button">
                    <img src="http://localhost:4662/static/image/placeholder.png" 
                        data-src="https://img.youtube.com/vi/${ids[id]['url'].slice(30, 41)}/maxresdefault.jpg"
                        alt="${ids[id]['title']}" 
                        class="card-img-top" 
                        id="${id}"> 
                </a>
            </div>
        `);
        }
    }


    let button_modal = document.querySelector(".header-row");
    button_modal.addEventListener("click", function (event) {
        if (event.target === button_modal) {
            return;
        }

        let id = event.target.getAttribute("id");
        let img = document.querySelector(".modal-body");
        let title = document.querySelector(".w-modal-title");

        let button = document.querySelector(".modal-footer");
        button.innerHTML = `<button type="button" class="btn btn-primary" id="setWallpaper">Set Wallpaper</button>`;
        let setWallpaper = document.querySelector("#setWallpaper");

        ids[id]['copyright'] = ids[id]['copyright'] === undefined ? "NASA" : ids[id]['copyright'];

        if (ids[id]['media_type'] === "image") {
            title.innerHTML = `<h5 class="modal-title">${ids[id]['title']}</h5>`;
            img.innerHTML = `
                <img src="${ids[id]['url']}" alt="${ids[id]['title']}" class="card-img">
                <p><strong>Author:</strong> ${ids[id]['copyright']}</p>
                <p><strong>Date of publication:</strong> ${ids[id]['date']}</p>
                <p><strong>Explanation:</strong> ${ids[id]['explanation']}</p>
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
                <p><strong>Explanation:</strong> ${ids[id]['explanation']}</p>
            `;

            setWallpaper.addEventListener("click", function () {
                wallpaperUpdate(`https://img.youtube.com/vi/${ids[id]['url'].slice(30, 41)}/maxresdefault.jpg`);
            });
        }
    });

    $(window).scroll(function () {
        if (($(window).scrollTop() > $(document).height() - $(window).height() - 100)) {
            $(".preloader").show();
            $(window).off("scroll");

            today.setDate(today.getDate() - 1);

            endDate.setDate(endDate.getDate() - 15);
            endDate.setMonth(endDate.getMonth() - 1);

            $.ajax({
                url: "https://api.nasa.gov/planetary/apod",
                data: {
                    api_key: apiKEY,
                    start_date: `${endDate.getFullYear()}-${endDate.getMonth() + 1}-${endDate.getDate()}`,
                    end_date: `${today.getFullYear()}-${today.getMonth() + 1}-${today.getDate()}`,
                },
                success: function (data) {
                    wallpaper(data);
                    $(window).on("scroll");
                },
            });
        }
    });
}

function wallpaperUpdate(url) {
    $.ajax({
        url: "http://localhost:8080/api/update/wallpaper",
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