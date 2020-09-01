function request() {
    var name = document.getElementById("fname");

    axios.get("/request/" + name.value)
    .then(function (reponse){changeText(reponse.data.payload);})
    //.catch
}

function changeText(json) {
    var text = document.getElementById("text");
    text.innerHTML = json;
}

function help() {
    window.open("https://github.com/Naexys/golfram/blob/master/README.md")
}

function contact() {
    window.open("https://naexys.netlify.app/english/contact.html")
}