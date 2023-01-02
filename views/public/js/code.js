$(document).ready(function() {

    let editor = ace.edit('code-editor-ace');

    // const editor = $("#code").val();

    $(".submitcodebutton").click(async function (event) {
        let codeitem = editor.getValue();
        event.preventDefault();
        console.log("clicked submit code button")
        try {
            const res = await fetch("/api/code", {
                method: "POST",
                headers: {"Content-Type": "application/json",
                    "Accept": "application/json"},
                body: JSON.stringify({
                    codeitem
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
            }
        } catch (err) {
            console.log(err.message);
        }
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