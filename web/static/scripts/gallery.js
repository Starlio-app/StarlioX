$.ajax({
    url: `https://api.nasa.gov/planetary/apod?api_key=1gI9G84ZafKDEnrbydviGknReOGiVK9jqrQBE3et&start_date=${new Date().toLocaleString().slice(6, 10)}-${new Date().toLocaleString().slice(3, 5)}-01&end_date=${new Date().toLocaleString().slice(6,10)}-${new Date().toLocaleString().slice(3, 5)}-${new Date().toLocaleString().slice(3, 5)}`,
    type: "GET",
    success: function(data){
        for (let i = 0; i < data.length; i++) {
            if (data[i].media_type == "image") {
                $(".header-row").append(`
                <span>
                    <div class="card">
                        <img src="${data[i].url}" alt="${data[i].title}" class="card-img-top">
                    </div>
                </span>
            `)
            }
        }
    }
});