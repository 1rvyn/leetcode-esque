// $(document).ready(function() {
//     console.log("jquery loaded")
//     $('.login-form').submit(function() {
//         // Get all the forms elements and their values in one step
//
//
//     });
// })
//
//
//
//
//
//
//

$(document).ready(function(){
    // Get and set element 'form'
    const form = document.querySelector(".login-form");
// Get and set element 'email'
    const email = document.querySelector("#email");
// Get and set element 'password'
    const password = document.querySelector("#password");
// Get and set element 'username' (error display)
    const display = document.querySelector(".error");
// Add an event listener to the login button and use the api to process login event
    form.addEventListener("submit", async (e) => {
        e.preventDefault();
        // display.textContent = "";
        console.log(password.value + email.value)
        try {
            const res = await fetch("/api/login", {
                method: "POST",
                headers: {"Content-Type": "application/json",
                            "Accept": "application/json"},
                body: JSON.stringify({
                    email: email.value,
                    password: password.value,
                }),
            });
            const content = await res.json();
            console.log(content);
            console.log(res.status);
            if (res.status === 400 || res.status === 401) {
                console.log("there was an issue")
            }
            else if (res.status === 200){
                console.log("the login has a success response code good job :)")
                window.location.href = "/account";
            }
        } catch (err) {
            console.log(err.message);
        }
    });
})
