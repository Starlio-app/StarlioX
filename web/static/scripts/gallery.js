let today = new Date();
let endDate = new Date();

const ids = [];
let id = 0;

const apiKEY = "1gI9G84ZafKDEnrbydviGknReOGiVK9jqrQBE3et";

today.setDate(today.getDate() - 17);
$(document).ready(function() {
    $.ajax({
        url: "https://api.nasa.gov/planetary/apod",
        type: "GET",
        data: {
            api_key: apiKEY,
            start_date: `${today.getUTCFullYear()}-${today.getUTCMonth() + 1}-${today.getUTCDate()}`,
            end_date: `${endDate.getUTCFullYear()}-${endDate.getUTCMonth() + 1}-${endDate.getUTCDate()}`,
        },
        success: function(data) {
            wallpaper(data);
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

    //if all images are loaded delete class
    $(".card-img-top").on("load", function() {
        $(this).removeClass("shimmer");
    });


    let button_modal = document.querySelector(".header-row");
    button_modal.addEventListener("click", function (event) {
        if (event.target === button_modal) {
            return;
        }

        let id = event.target.getAttribute("id");
        let img = document.querySelector(".modal-body");
        let title = document.querySelector(".w-modal-title");

        let button = document.querySelector(".modal-footer");

        if($(`img#${id}.card-img-top`).attr("src") !== "http://localhost:3000/static/image/placeholder.png") {
            button.innerHTML = `<button type="button" class="btn btn-primary" id="setWallpaper">Set Wallpaper</button>`;
        } else {
            button.innerHTML = `<button type="button" 
                                        class="btn" 
                                        id="setWallpaper" 
                                        disabled 
                                        style="background-color: grey; color: white;">
                                        Set Wallpaper
                                </button>`;
        }

        let setWallpaper = document.querySelector("#setWallpaper");
        ids[id]['copyright'] = ids[id]['copyright'] === undefined ? "NASA" : ids[id]['copyright'];

        let explanation = ids[id]['explanation'].length > 200 ? ids[id]['explanation'].slice(0, 200) + "..." : ids[id]['explanation'];
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


        let showMore = document.querySelector("#show-more");
        showMore.addEventListener("click", function () {
            let explanation = document.querySelector(".modal-body p:nth-child(4)");
            if(showMore.innerHTML === "Show more") {
                explanation.innerHTML = `<strong>Explanation:</strong> ${ids[id]['explanation']}`;
                showMore.innerHTML = "Show less";
            } else {
                explanation.innerHTML = explanation.innerHTML.slice(0, 200) + "...";
                showMore.innerHTML = "Show more";
            }
        });
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
                    start_date: `${endDate.getUTCFullYear()}-${endDate.getUTCMonth() + 1}-${endDate.getUTCDate()}`,
                    end_date: `${today.getUTCFullYear()}-${today.getUTCMonth() + 1}-${today.getUTCDate()}`,
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