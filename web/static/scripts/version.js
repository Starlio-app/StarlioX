$.ajax({
    url: "http://localhost:8080/api/get/version",
    type: "GET",
    success: function(data){
        $("#program-version").text(data.message.replace("v", ""));
    }
})