$(document).ready(function(){
    // Get and set element 'form'
    const form = document.querySelector(".register-form");
    const email = document.querySelector("#email");
    const password = document.querySelector("#password");
    const name = document.querySelector("#name");
    const display = document.querySelector(".display-message");

    form.addEventListener("submit", async (e) => {
        e.preventDefault();
        console.log(password.value + email.value)
        try {
            const res = await fetch("https://api.irvyn.xyz/register", {
                method: "POST",
                headers: {"Content-Type": "application/json",
                    "Accept": "application/json",
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
            if (!content.success){
                displayErrorMessage("Error registering user. Please try again.");
            }
            if (res.status === 400 || res.status === 401) {
                console.log("there was an issue")

            }
            else if (res.status === 200){
                console.log("the register has a success response code good job :)")
                console.log(res, content)
                if (content.success) {
                    displaySuccessMessage("Successfully registered user. Please verify your email.");
                }
            }
        } catch (err) {
            console.log(err.message);
        }
    });

    function displaySuccessMessage(message) {
        const messageElement = document.createElement("div");
        messageElement.textContent = message;
        messageElement.classList.add('bg-green-100', 'border', 'border-green-400', 'text-green-700', 'px-4', 'py-3', 'rounded', 'mb-4');
        form.appendChild(messageElement);
    }

    function displayErrorMessage(message) {
        const messageElement = document.createElement("div");
        messageElement.textContent = message;
        messageElement.classList.add('bg-red-100', 'border', 'border-red-400', 'text-red-700', 'px-4', 'py-3', 'rounded', 'mb-4');
        form.appendChild(messageElement);
    }
})
