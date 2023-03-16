$(document).ready(function(){
    let editor = ace.edit('box-2-top');

    // const editor = $("#code").val();

    $("#submitcodebutton").click(async function (event) {
        let codeitem = editor.getValue();
        event.preventDefault();
        console.log("clicked submit code button")
        let language = $("#language-select").val()
        console.log("language is:", $("#language-select").val())
        try {
            const res = await fetch("https://api.irvyn.xyz/code", {
                method: "POST",
                headers: {"Content-Type": "application/json",
                    "Accept": "application/json"},
                body: JSON.stringify({
                    code: codeitem,
                    language: language
                }),
                xhrFields: {
                    withCredentials: true
                },
                credentials: "include",
            });
            const content = await res.json();
            console.log(content);
            // take the content and put it inside the terminal div in the html
            $(".widget").append(content.output); // ehh
            $(".widget").append(content.error); // ehh

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


// send the request to get the new template code
async function updateCodeTemplate(language) {
    try {
      const response = await fetch(`/codetemplate?language=${language}`);
      const data = await response.json();
      console.log(data)
      const codeTemplate = data.Codetemplate;
  
      editor.setValue(codeTemplate);
    } catch (error) {
      console.error("Error fetching code template:", error);
    }
  }


function setEditorLanguage(mode) {
    switch (mode) {
      case 'python':
        editor.session.setMode('ace/mode/python');
        break;
      case 'javascript':
        editor.session.setMode('ace/mode/javascript');
        break;
      case 'go':
        editor.session.setMode('ace/mode/golang');
        break;
      // Add more languages here
      default:
        editor.session.setMode('ace/mode/python');
    }
  }

  // Get the dropdown element
var languageSelect = document.getElementById('language-select');

// Set the initial mode based on the selected language
setEditorLanguage(languageSelect.value);

// Listen for changes in the dropdown and update the editor's mode
languageSelect.addEventListener('change', function () {
  setEditorLanguage(this.value);
  console.log(this.value);
  console.log("changed language");
  updateCodeTemplate(this.value);

});