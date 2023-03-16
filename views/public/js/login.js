// Get the value of the message query parameter from the URL
const urlParams = new URLSearchParams(window.location.search);
const message = urlParams.get('message');

// If a message was sent, append it to the page
if (message) {
    const messageElement = document.createElement('div');
    messageElement.innerText = message;
    messageElement.classList.add('bg-green-100', 'border', 'border-green-400', 'text-green-700', 'px-4', 'py-3', 'rounded', 'mb-4');
    const mainElement = document.querySelector('.login-message');
    if (mainElement) {
        mainElement.appendChild(messageElement);
    }
}


$(document).ready(function(){
    // Get and set element 'form'
    const form = document.querySelector(".login-form");
// Get and set element 'email'
    const email = document.querySelector("#email");
// Get and set element 'password'
    const password = document.querySelector("#password");
// Get and set element 'username' (error display)
// Add an event listener to the login button and use the api to process login event
    form.addEventListener("submit", async (e) => {
        e.preventDefault();
        // display.textContent = "";
        console.log(password.value + email.value)
        try {
            const res = await fetch("https://api.irvyn.xyz/login", {
                // const res = await fetch("/api/login", {
                method: "POST",
                headers: {"Content-Type": "application/json",
                            "Accept": "application/json"},
                body: JSON.stringify({
                    email: email.value,
                    password: password.value,
                }),
                xhrFields: {
                    withCredentials: true
                },
                credentials: "include",
            });
            const content = await res.json();
            console.log(content);
            console.log(res.status);
            if (res.status === 400 || res.status === 401 || res.status === 404) {
                console.log("there was an issue")
                // append message to "login-message" div

                document.getElementById("login-error").innerHTML = content.message;

            }
            else if (res.status === 200){
                console.log("the login has a success response code good job :)")
                // // print out the cookie 
                // console.log("cookie is: ", document.cookie)
                // console.log(res, content)
                window.location.href = "/";

            }
        } catch (err) {
            console.log(err.message);
        }
    });
})
