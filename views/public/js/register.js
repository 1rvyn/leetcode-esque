$(document).ready(function(){
    // Get and set element 'form'
    const form = document.querySelector(".register-form");
// Get and set element 'email'
    const email = document.querySelector("#email");
// Get and set element 'password'
    const password = document.querySelector("#password");
    const name = document.querySelector("#name");
// Add an event listener to the register button and use the api to process regster event
    form.addEventListener("submit", async (e) => {
        e.preventDefault();
        // display.textContent = "";
        console.log(password.value + email.value)
        try {
            const res = await fetch("https://irvyn.dev/api/register", {
                // const res = await fetch("/api/register", {
                method: "POST",
                headers: {"Content-Type": "application/json",
                            "Accept": "application/json",
                            "Access-Control-Allow-Origin": "*"
                        },
                body: JSON.stringify({
                    email: email.value,
                    password: password.value,
                    name: name.value,
                }),
                xhrFields: {
                    withCredentials: true
                },
                credentials: "include",
            });
            const content = await res.json();
            console.log(content);
            console.log(res.status);
            if (res.status === 400 || res.status === 401) {
                console.log("there was an issue")
                // append message to page
            }
            else if (res.status === 200){
                console.log("the register has a success response code good job :)")
                // print out the cookie 
                console.log(res, content)
                // window.location.href = "/";

            }
        } catch (err) {
            console.log(err.message);
        }
    });
})
