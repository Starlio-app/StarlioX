$(document).ready(function() {
    let today = new Date();
    let date = today.setDate(today.getDate() - 31);
    date = new Date(date);

    const apiKEY = "1gI9G84ZafKDEnrbydviGknReOGiVK9jqrQBE3et";

    $.ajax({
        url: "https://api.nasa.gov/planetary/apod",
        type: "GET",
        data: {
            api_key: apiKEY,
            start_date: `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`,
            end_date: `${new Date().toLocaleString().slice(6, 10)}-${new Date().toLocaleString().slice(3, 5)}-${new Date().toLocaleString().slice(0, 2)}`,
        },
        success: function(data) {
            wallpaper(data);
        },
        error: function() {
            let prev_date = Number(new Date().toLocaleString().slice(0,2))-1;
            $.ajax({
               url: "https://api.nasa.gov/planetary/apod",
               data: {
                   api_key: apiKEY,
                   start_date: `${date.getFullYear()}-${date.getMonth()+1}-${date.getDate()}`,
                   end_date: `${new Date().toLocaleString().slice(6,10)}-${new Date().toLocaleString().slice(3, 5)}-${prev_date}`,
               },
               success: function (data) {
                   wallpaper(data);
               }
            });
        },
    });
});

const myModal = new bootstrap.Modal('#Wallpaper', {
    keyboard: false
});
myModal.show();

function wallpaper(data) {
    $(".preloader").hide();
    data = data.reverse();

    for (let i = 0; i < data.length; i++) {
        if (data[i]['media_type'] === "image") {
            $(".header-row").append(`
                <div class="card">
                    <a data-bs-toggle="modal" href="#WallpaperModal" role="button">
                        <img src="${data[i].url}" alt="${data[i]['title']}" class="card-img-top" idi="${i + 1}">
                    </a>
                </div>
        `);
        } else {
            $(".header-row").append(`
                    <div class="card">
                        <a data-bs-toggle="modal" href="#WallpaperModal" role="button">
                            <img src="https://img.youtube.com/vi/${data[i]['url'].slice(30,41)}/maxresdefault.jpg" alt="${data[i]['title']}" class="card-img-top" idi="${i+1}">
                        </a>
                    </div>
            `);
        }
    }
    let button_modal = document.querySelector(".header-row");
    button_modal.addEventListener("click", function (event) {
        let id = event.target.getAttribute("idi") - 1;
        let img = document.querySelector(".modal-body");
        let title = document.querySelector(".w-modal-title");

        let button = document.querySelector(".modal-footer");
        button.innerHTML = `<button type="button" class="btn btn-primary" id="setWallpaper">Set Wallpaper</button>`;
        let setWallpaper = document.querySelector("#setWallpaper");

        data[id]['copyright'] = data[id]['copyright'] === undefined ? "NASA" : data[id]['copyright'];

        if (data[id]['media_type'] === "image") {
            img.innerHTML = `
                <img src="${data[id].url}" alt="${data[id]['title']}" class="card-img">
                <p><strong>Author:</strong> ${data[id]['copyright']}</p>
                <p><strong>Date of publication:</strong> ${data[id]['date']}</p>
                <p><strong>Explanation:</strong> ${data[id]['explanation']}</p>
            `;

            title.innerHTML = `<h5 class="modal-title">${data[id]['title']}</h5>`;
            setWallpaper.addEventListener("click", function () {
                wallpaperUpdate(data[id]['hdurl']);
            });
        } else {
            img.innerHTML = `
                <iframe width="460" height="315" src="${data[id]['url']}" title="YouTube video player" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowFullScreen></iframe>
                <p><strong>Author:</strong> ${data[id]['copyright']}</p>
                <p><strong>Date of publication:</strong> ${data[id]['date']}</p>
                <p><strong>Explanation:</strong> ${data[id]['explanation']}</p>
            `;

            title.innerHTML = `<h5 class="modal-title">${data[id]['title']}</h5>`;
            setWallpaper.addEventListener("click", function () {
                wallpaperUpdate(`https://img.youtube.com/vi/${data[id]['url'].slice(30,41)}/maxresdefault.jpg`);
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