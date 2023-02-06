async function postRequest() {
    const codeitem = ace.edit('code-editor-ace').getValue();
    const res = await fetch("https://api.irvyn.xyz/code", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
            "Cookie": "jwt=" + getCookie("jwt")
        },
        body: JSON.stringify({
            codeitem
        }),
        credentials: "include"
    });
    const content = await res.json();
    console.log(content);
    console.log(res.status);
    if (res.status === 400 || res.status === 401) {
        console.log("There was an issue");
    } else if (res.status === 200) {
        console.log("The request was successful");
        // take the content and put it inside the terminal div in the HTML
        document.querySelector(".terminal").innerHTML += content.output;
        document.querySelector(".terminal").innerHTML += content.error;
    }
}

function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

document.querySelector(".submitcodebutton").addEventListener("click", postRequest);
