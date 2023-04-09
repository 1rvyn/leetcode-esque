$(document).ready(function(){
    let editor = ace.edit('box-2-top');
    let questionID = document.getElementById('questionID').value; // get the QuestionID value


    $("#submitcodebutton").click(async function (event) {
        let codeitem = editor.getValue();
        event.preventDefault();
        console.log("clicked submit code button");
        let language = $("#language-select").val()
        console.log("language is:", $("#language-select").val());
        try {
            const res = await fetch("https://api.irvyn.xyz/code", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Accept": "application/json"
                },
                body: JSON.stringify({
                    code: codeitem,
                    language: language,
                    QuestionID: questionID // include the QuestionID value in the body
                }),
                xhrFields: {
                    withCredentials: true
                },
                credentials: "include",
            });
            const content = await res.json();
            console.log(content.result);// log the results of marking


            console.log(res.status);
            if (res.status === 400 || res.status === 401) {
                console.log("there was an issue");
            } else if (res.status === 200) {
                console.log("the login has a success response code good job :)");
                updateTestResultsLights(content.result); // Call the function to update the test result lights
            }
        } catch (err) {
            console.log(err.message);
        }

    })
})

// a light system based on the test results
function updateTestResultsLights(testResults) {
    console.log('Received testResults:', testResults);

    const container = document.querySelector('.test-results-container');
    container.innerHTML = ''; // Clear the container

    const isArrayOfObjects = Array.isArray(testResults) && typeof testResults[0] === 'object';
    const isArrayOfBooleans = Array.isArray(testResults) && (typeof testResults[0] === 'boolean' || testResults[0] === 'true' || testResults[0] === 'false');

    console.log('isArrayOfObjects:', isArrayOfObjects);
    console.log('isArrayOfBooleans:', isArrayOfBooleans);

    if (isArrayOfObjects || isArrayOfBooleans) {
        testResults.forEach((result, index) => {
            const light = document.createElement('div');
            light.style.width = '20px';
            light.style.height = '20px';
            light.style.borderRadius = '50%';
            light.style.marginRight = '20px';

            if (isArrayOfObjects) {
                // Handle the current object structure
                light.style.backgroundColor = result.success ? 'green' : 'red';
            } else {
                // Handle the new array of booleans or strings
                const success = result === true || result === 'true';
                light.style.backgroundColor = success ? 'green' : 'red';
            }

            container.appendChild(light);
        });
    } else {
        console.error('Unknown testResults format:', testResults);
    }
}






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
        updateCodeTemplate(mode).then(r => console.log("updated code template"));
        break;
      case 'javascript':
        editor.session.setMode('ace/mode/javascript');
        updateCodeTemplate(mode).then(r => console.log("updated code template"));
          break;
      case 'go':
        editor.session.setMode('ace/mode/golang');
        updateCodeTemplate(mode).then(r => console.log("updated code template"));

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