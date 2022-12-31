$(document).ready(function() {

    // const form = document.querySelector(".codeform1");
    // const texted = document.querySelector(".code-editor").val();


    const codeItem = $(".code-editor").val(); // get the code as text/string from the textarea


    console.log("CODE ITEM is: ", codeItem)

    // const editor = $("#code").val();

    $(".submitcodebutton").click(function (event) {
        event.preventDefault()

        console.log("clicked submit code button")
        $.ajax({
            type: "POST",
            url: "/api/code",
            async: true,
            data: {
                code: codeItem
            },
            success: function(response) {
                console.log(response);
            }
        });
    })
})



// Add an event listener to the login button and use the api to process login event
//     submitcodebutton.addEventListener("submit", async (e) => {
//         e.preventDefault();
//
//         // log the code to the console
//         // console.log(editor);
//
//         try {
//             const res = await fetch("/api/code", {
//                 method: "POST",
//                 headers: {"Content-Type": "application/json",
//                             "Accept": "application/json"},
//                 body: JSON.stringify({
//                     payloadItem: texted.value,
//                 }),
//             });
//             const content = await res.json();
//             console.log(content);
//             console.log(res.status);
//             if (res.status >= 200 && res.status < 300 && res.headers.get("Content-Type") === "application/json") {
//                 const content = await res.json();
//                 console.log(content);
//                 console.log(res.status);
//             } else {
//                 console.log(res.status + ": " + await res.text());
//             }
//         } catch (err) {
//             console.log(err);
//         }
//     });
//
// });