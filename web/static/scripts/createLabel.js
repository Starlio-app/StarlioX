let button = document.querySelector("#createLabelButton");
button.addEventListener("click", function() {
    $.ajax({
       url: "http://localhost:3000/api/create/label",
       type: "POST",
       success: function(data) {
           if(data.status) {
                $(".toast-body").text(data.message);
                let toastLiveExample = document.getElementById('liveToast');
                let toast = new bootstrap.Toast(toastLiveExample);
                toast.show();
           } else {
                $(".toast-body").text("An error occurred while creating the label.");
                let toastLiveExample = document.getElementById('liveToast');
                let toast = new bootstrap.Toast(toastLiveExample);
                toast.show();
           }
       }
    });
});