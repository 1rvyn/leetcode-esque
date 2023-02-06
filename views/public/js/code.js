$(document).ready(function(){
    let editor = ace.edit('code-editor-ace');

    // const editor = $("#code").val();

    $(".submitcodebutton").click(async function (event) {
        let codeitem = editor.getValue();
        event.preventDefault();
        console.log("clicked submit code button")
        try {
            const res = await fetch("https://api.irvyn.xyz/code", {
                method: "POST",
                headers: {"Content-Type": "application/json",
                    "Accept": "application/json"},
                body: JSON.stringify({
                    codeitem
                }),
                xhrFields: {
                    withCredentials: true
                },
                credentials: "include",
            });
            const content = await res.json();
            console.log(content);
            // take the content and put it inside the terminal div in the html
            $(".terminal").append(content.output); // ehh
            $(".terminal").append(content.error); // ehh

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