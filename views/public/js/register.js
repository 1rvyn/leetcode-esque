    console.log("this should print on the register page")

    $(document).ready(function(){
        const form = document.querySelector(".register-form");
        const email = document.querySelector("#email");
        const name = document.querySelector("#name");
        const password = document.querySelector("#password");
    // Add an event listener to the login button and use the api to process login event
        form.addEventListener("submit", async (e) => {
            e.preventDefault();
            // display.textContent = "";
            console.log(password.value + email.value + name.value)
            try {
                const res = await fetch("https://irvyn.dev/api/register", {
                    // const res = await fetch("/api/register", {
                    method: "POST",
                    headers: {"Content-Type": "application/json",
                        "Accept": "application/json",
                   "Access-Control-Allow-Origin": "irvyn.xyz"
                 },
                    body: JSON.stringify({
                        name: name.value,
                        email: email.value,
                        password: password.value,
                    }),
                    xhrFields: {
                        withCredentials: true,

                    },
                    credentials: "include",
                });
                const content = await res.json();
                console.log(content);
                console.log(res.status);
                if (res.status === 400 || res.status === 401) {
                    console.log("there was an issue")
                }
                else if (res.status === 200){
                    console.log("content inside is:", content)

                    console.log("the login has a success response code good job :)")
                    // window.location.href = "/login";
                }
            } catch (err) {
                console.log(err.message);
            }
        });
    })
