function request() {
    var name = document.getElementById("fname");
    var str = name.value.split("/").join("§");
    axios.get("/request/" + str)
    .then(function (reponse){changeText(reponse.data.payload);})
    //.catch
}

function changeText(json) {
    var text = document.getElementById("text");
    text.innerHTML = json;
}

function help() {
    window.open("https://github.com/Flowtter/golfram/blob/master/README.md")
}

function contact() {
    window.open("https://flowtter.netlify.app/english/contact.html")
}