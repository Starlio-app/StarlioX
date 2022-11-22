import { initializeApp } from "https://www.gstatic.com/firebasejs/9.14.0/firebase-app.js";
import { getAnalytics } from "https://www.gstatic.com/firebasejs/9.14.0/firebase-analytics.js";

$.ajax({
    url: 'http://localhost:3000/api/get/settings',
    type: 'GET',
    dataType: 'json',
    success: function (data) {
        if (data["analytics"] === 1) {
            const firebaseConfig = {
                apiKey: "AIzaSyCeHtV4wmB9xJY4vfcpt7wX-WvlV-5S6v4",
                authDomain: "everynasa-181a1.firebaseapp.com",
                databaseURL: "https://everynasa-181a1-default-rtdb.firebaseio.com",
                projectId: "everynasa-181a1",
                storageBucket: "everynasa-181a1.appspot.com",
                messagingSenderId: "369869513900",
                appId: "1:369869513900:web:2ff68e57f95a36bf87ab09",
                measurementId: "G-JN83RYFK56"
            };

            const app = initializeApp(firebaseConfig);
            getAnalytics(app);
        } else {
            $("#analytics").remove();
        }
    }
})