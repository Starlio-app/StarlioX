$.ajax({
   url: "http://localhost:8080/api/get/version",
   type: "GET",
    success: function(version) {
       $.ajax({
           url: "https://api.github.com/repos/Redume/EveryNasa/tags",
           type: "GET",
           success: function(data) {
                if (version.message !== data[0].name) {
                    let div = document.createElement("div");
                    div.id = "update";
                    div.innerHTML = `
                        <p id="update-title">A new update is available</p>
                        <p id="update-desc">You can download the new version of the program.</p>
                        <button id="update-button" class="btn btn-primary">Download</button>
                    `;
                    document.body.appendChild(div);
                    let updateButton = document.querySelector("#update-button");
                    updateButton.addEventListener("click", function() {
                        window.location.href = `https://github.com/Redume/EveryNasa/releases/download/${data[0].name}/EveryNasa.msi`;
                        $(".toast-body").text("Downloading the update...");
                        let toastLiveExample = document.getElementById('liveToast');
                        let toast = new bootstrap.Toast(toastLiveExample);
                        toast.show();
                    })
                }
           }
       })
    }
});