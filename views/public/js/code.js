document.addEventListener("DOMContentLoaded", function() {
    const submitCodeButton = document.querySelector(".submitcodebutton");
    if (!submitCodeButton) return;
  
    let editor = ace.edit("code-editor-ace");
    
    submitCodeButton.addEventListener("click", async function(event) {
      let codeitem = editor.getValue();
      event.preventDefault();
      console.log("clicked submit code button");
      try {
        const res = await fetch("https://api.irvyn.xyz/code", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Accept: "application/json"
          },
          body: JSON.stringify({
            codeitem
          })
        });
        const content = await res.json();
        console.log(content);
        // take the content and put it inside the terminal div in the html
        $(".terminal").append(content.output);
        $(".terminal").append(content.error);
  
        console.log(res.status);
        if (res.status === 400 || res.status === 401) {
          console.log("there was an issue");
        } else if (res.status === 200) {
          console.log("the login has a success response code good job :)");
        }
      } catch (err) {
        console.log(err.message);
      }
    });
  });
  

// function getCookie(cname) {
//     var name = cname + "=";
//     var decodedCookie = decodeURIComponent(document.cookie);
//     var ca = decodedCookie.split(';');
//     for (var i = 0; i < ca.length; i++) {
//         var c = ca[i];
//         while (c.charAt(0) == ' ') {
//             c = c.substring(1);
//         }
//         if (c.indexOf(name) == 0) {
//             return c.substring(name.length, c.length);
//         }
//     }
//     return "";
// }
// document.querySelector(".submitcodebutton").addEventListener("click", postRequest);
