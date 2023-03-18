document.addEventListener('DOMContentLoaded', function () {
    const form = document.getElementById('submit-question-form');

    form.addEventListener('submit', async function (event) {
        event.preventDefault();

        const formData = new FormData(form);
        const requestData = {
            problem: formData.get('problem'),
            example_answer: formData.get('example_answer'),
            example_input: formData.get('example_input'),
            problem_type: formData.get('problem_type'),
            problem_difficulty: formData.get('problem_difficulty'),
            template_code: {
                python: formData.get('template_code[python]'),
                javascript: formData.get('template_code[javascript]'),
                go: formData.get('template_code[go]')
            }
        };
        console.log(requestData);
        try {
            const res = await fetch('https://api.irvyn.xyz/newquestion', {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json'
                },
                body: JSON.stringify(requestData),
                xhrFields: {
                    withCredentials: true
            },
            credentials: "include",
            });


            if (!res.ok) {
                throw new Error(`Cannot POST /newquestion: ${res.status} ${res.statusText}`);
            }

            const content = await res.json();

            if (res.status === 400 || res.status === 401) {
                console.log("There was an issue");
            } else if (res.status === 200) {
                console.log("The form submission has a success response code, good job :)");
                console.log(content);
            }

            // Handle the response content here
            // e.g., display a success message, redirect, or update the page content

        } catch (err) {
            console.log(err.message);
        }
    });
});
