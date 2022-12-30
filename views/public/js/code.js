$(document).ready(function(){

    const form = document.querySelector(".codeform1");
    const texted = document.querySelector("#code");

    console.log("value is: ", texted.value)

    // const editor = $("#code").val();

// Add an event listener to the login button and use the api to process login event
    form.addEventListener("submit", async (e) => {
        e.preventDefault();        

        // log the code to the console
        // console.log(editor);

        try {
            const res = await fetch("/api/code", {
                method: "POST",
                headers: {"Content-Type": "application/json",
                            "Accept": "application/json"},
                body: JSON.stringify({
                    payloadItem: texted.value,
                }),
            });
            const content = await res.json();
            console.log(content);
            console.log(res.status);
            if (res.status >= 200 && res.status < 300 && res.headers.get("Content-Type") === "application/json") {
                const content = await res.json();
                console.log(content);
                console.log(res.status);
            } else {
                console.log(res.status + ": " + await res.text());
            }
        } catch (err) {
            console.log(err);
        }
    });

});