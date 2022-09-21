$.ajax({
    url: `https://api.nasa.gov/planetary/apod?api_key=1gI9G84ZafKDEnrbydviGknReOGiVK9jqrQBE3et&start_date=${new Date().toLocaleString().slice(6, 10)}-${new Date().toLocaleString().slice(3, 5)}-01&end_date=${new Date().toLocaleString().slice(6,10)}-${new Date().toLocaleString().slice(3, 5)}-${new Date().toLocaleString().slice(0,2)}`,
    type: "GET",
    success: function(data){
        Wallpaper(data)
    },
    error: function(){
        let prev_date = Number(new Date().toLocaleString().slice(0,2))
        this.url = `https://api.nasa.gov/planetary/apod?api_key=1gI9G84ZafKDEnrbydviGknReOGiVK9jqrQBE3et&start_date=${new Date().toLocaleString().slice(6, 10)}-${new Date().toLocaleString().slice(3, 5)}-01&end_date=${new Date().toLocaleString().slice(6,10)}-${new Date().toLocaleString().slice(3, 5)}-${prev_date-1}`,
        $.ajax({
            url: this.url,
            type: "GET",
            success: function(data){
                Wallpaper(data)
            }
        })
    }
});

const myModal = new bootstrap.Modal('#Wallpaper', {
    keyboard: false
})
myModal.show()

function Wallpaper(data) {
    $(".preloader").fadeOut(0)
    for (let i = 0; i < data.length; i++) {
        if (data[i].media_type === "image") {
            $(".header-row").append(`
                    <div class="card">
                        <a data-bs-toggle="modal" href="#WallpaperModal" role="button">
                            <img src="${data[i].url}" alt="${data[i].title}" class="card-img-top" idi="${i+1}">
                        </a>
                    </div>
            `)
        } else {
            $(".header-row").append(`
                    <div class="card">
                        <a data-bs-toggle="modal" href="#WallpaperModal" role="button">
                            <img src="https://img.youtube.com/vi/${data[i].url.slice(30,41)}/maxresdefault.jpg" alt="${data[i].title}" class="card-img-top" idi="${i+1}">
                        </a>
                    </div>
            `)
        }
        let buttons = document.querySelector(".header-row")
        buttons.addEventListener("click", function (event) {
            if (event.target.tagName === "IMG") {
                let id = event.target.getAttribute("idi")
                let img = document.querySelector(".modal-body")
                let title = document.querySelector(".w-modal-title")

                let button = document.querySelector(".modal-footer")
                button.innerHTML = `
                    <button type="button" class="btn btn-primary" id="setWallpaper">Set Wallpaper</button>
                `
                let setWallpaper = document.querySelector("#setWallpaper")

                if(data[id-1].media_type === "image") {
                    img.innerHTML = `
                        <img src="${data[id - 1].url}" alt="${data[id - 1].title}" class="card-img">
                        <p>Author: <strong>${data[id - 1].copyright}</strong></p>
                        <p>Date of shooting: <strong>${data[id - 1].date}</strong></p>
                        <p>Explanation: <strong>${data[id - 1].explanation}</strong></p>
                    `.replace("undefined", "Unknown")
                    title.innerHTML = `
                        <h5 class="modal-title">${data[id - 1].title}</h5>
                    `

                    setWallpaper.addEventListener("click", function () {
                        $.ajax({
                            url: `http://localhost:8080/api/update/wallpaper`,
                            type: "POST",
                            data: {
                                url: data[id - 1].hdurl
                            },
                            success: function () {
                                $(".toast-body").text("The wallpaper has been installed")
                                let toastLiveExample = document.getElementById('liveToast')
                                let toast = new bootstrap.Toast(toastLiveExample)
                                toast.show()
                            },
                            error: function (err) {
                                console.error(err);
                            }
                        })
                    })
                } else {
                    img.innerHTML = `
                        <iframe width="460" height="315" src="${data[id - 1].url}" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
                        <p>Author: <strong>${data[id - 1].copyright}</strong></p>
                        <p>Date of shooting: <strong>${data[id - 1].date}</strong></p>
                        <p>Explanation: <strong>${data[id - 1].explanation}</strong></p>
                    `.replace("undefined", "Unknown")
                    title.innerHTML = `
                        <h5 class="modal-title">${data[id].title}</h5>
                    `

                    setWallpaper.addEventListener("click", function () {
                        $.ajax({
                            url: `http://localhost:8080/api/update/wallpaper`,
                            type: "POST",
                            data: {
                                url: `https://img.youtube.com/vi/${data[id - 1].url.slice(30,41)}/maxresdefault.jpg`
                            },
                            success: function () {
                                $(".toast-body").text("The wallpaper has been installed")
                                let toastLiveExample = document.getElementById('liveToast')
                                let toast = new bootstrap.Toast(toastLiveExample)
                                toast.show()
                            },
                            error: function (err) {
                                console.error(err);
                            }
                        })
                    })
                }
            }
        })
    }
}