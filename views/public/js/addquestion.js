// document.addEventListener('DOMContentLoaded', function () {
//     const form = document.getElementById('submit-question-form');
//
//     form.addEventListener('submit', async function (event) {
//         event.preventDefault();
//
//         const formData = new FormData(form);
//         let problem = $(".problem").val()
//         let example_answer = $(".example_answer").val()
//
//         const requestData = {
//             problem: formData.get('problem'),
//             example_answer: formData.get('example_answer'),
//             example_input: formData.get('example_input'),
//             problem_type: formData.get('problem_type'),
//             problem_difficulty: formData.get('problem_difficulty'),
//             template_code: {
//                 python: formData.get('template_code[python]'),
//                 javascript: formData.get('template_code[javascript]'),
//                 go: formData.get('template_code[go]')
//             }
//         };
//         console.log(requestData);
//         try {
//             const res = await fetch('https://api.irvyn.xyz/new', {
//                 method: "POST",
//                 headers: {
//                     'Content-Type': 'application/json',
//                     'Accept': 'application/json'
//                 },
//                 body: JSON.stringify({
//                     "problem": problem,
//                     "example_answer": example_answer
//                 }),
//                 xhrFields: {
//                     withCredentials: true
//             },
//             credentials: "include",
//             });
//
//
//             if (!res.ok) {
//                 throw new Error(`Cannot POST /newquestion: ${res.status} ${res.statusText}`);
//             }
//
//             const content = await res.json();
//
//             if (res.status === 400 || res.status === 401) {
//                 console.log("There was an issue");
//             } else if (res.status === 200) {
//                 console.log("The form submission has a success response code, good job :)");
//                 console.log(content);
//             }
//
//             // Handle the response content here
//             // e.g., display a success message, redirect, or update the page content
//
//         } catch (err) {
//             console.log(err.message);
//         }
//     });
// });


$(document).ready(function(){
    // Get and set element 'form'
    const form = document.querySelector("#submit-question-form");
// Get and set element 'email'
    const problem = document.querySelector("#problem");
// Get and set element 'password'
    const answer = document.querySelector("#example_answer");
// Get and set element 'username' (error display)
// Add an event listener to the login button and use the api to process login event
    form.addEventListener("submit", async (e) => {
        e.preventDefault();
        // display.textContent = "";
        console.log(problem.value + answer.value)
        try {
            const res = await fetch("https://api.irvyn.xyz/new_question", {
                // const res = await fetch("/api/login", {
                method: "POST",
                headers: {"Content-Type": "application/json",
                    "Accept": "application/json"},
                body: JSON.stringify({
                    problem: problem.value,
                    example_answer: answer.value,
                }),
                xhrFields: {
                    withCredentials: true
                },
                credentials: "include",
            });
            const responseText = await res.text();
            console.log("Response Text: ", responseText);
            const content = JSON.parse(responseText);
            console.log(content);
            console.log(res.status);
            if (res.status === 400 || res.status === 401 || res.status === 404) {
                console.log("there was an issue")
                // append message to "login-message" div

                document.getElementById("question-error").innerHTML = content.message;

            }
            else if (res.status === 200){
                console.log("the question submit has a success response code good job :)")
                // // print out the cookie
                // console.log("cookie is: ", document.cookie)
                // console.log(res, content)

            }
        } catch (err) {
            console.log(err.message);
        }
    });
})
