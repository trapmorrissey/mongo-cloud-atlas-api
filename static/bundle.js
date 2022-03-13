

function showMessage(input, message, type) {
    // get the small element and set the message
    const msg = input.parentNode.querySelector("small");
    msg.innerText = message;
    // update the class for the input
    input.className = type ? "success" : "error";
    return type;
}

function showError(input, message) {
    return showMessage(input, message, false);
}

function showSuccess(input) {
    return showMessage(input, "", true);
}

function hasValue(input, message) {
    if (input.value.trim() === "") {
        return showError(input, message);
    }
    return showSuccess(input);
}


const form = document.querySelector("#signup");

const NAME_REQUIRED = "Please enter a title!";
const INFO_REQUIRED = "Please enter some information to send!";

form.addEventListener("submit", function (event) {
    // stop form submission
    event.preventDefault();

    // getting current date and time
    let today = new Date();
    let date = today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate();
    let time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
    let dateTime = date+' '+time;

    // validate the form
    let nameValid = hasValue(form.elements["name"], NAME_REQUIRED);
    let emailValid = hasValue(form.elements["info1"], INFO_REQUIRED);
    // if valid, submit the form.
    if (nameValid && emailValid) {
        let title = form.elements["name"];
        let info1 = form.elements["info1"];
        let info2 = form.elements["info2"];
        let info3 = form.elements["info3"];
        let info4 = form.elements["info4"];

        const data = JSON.stringify({
            "isbn": dateTime,
            "title": title.value,
            "info": {
                "info1": info1.value,
                "info2": info2.value,
                "info3": info3.value,
                "info4": info4.value
            }
        });

        const xhr = new XMLHttpRequest();
        xhr.withCredentials = true;

        xhr.addEventListener("readystatechange", function () {
            if (this.readyState === this.DONE) {
                console.log(this.responseText);
            }
        });

        xhr.open("POST", "http://YOUR_POST_API_URL");
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.setRequestHeader("X-Requested-With", "XMLHttpRequest",);
        xhr.setRequestHeader("Access-Control-Allow-Origin", "*",);

        xhr.send(data);
        alert("Successfully sent data to API!")
    }

        //alert("Demo only. No form was posted.")
});
