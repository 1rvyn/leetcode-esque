console.log("this should print if we are on the account page");

window.onload = async function() {
    try {
        const res = await fetch("https://api.irvyn.xyz/account", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json"
        },
            xhrFields: {
            withCredentials: true
            },
    credentials: "include"
    });
        const content = await res.json();
        console.log(content);
        console.log(res.status);
    if (res.status === 200) {
        let formattedCreatedAt = formatDate(content[i].created_at);
        const accountInfoDiv = document.querySelector(".account-info");
        let accountInfoHtml = "";
        accountInfoHtml += "<p>ID: " + content.id + "</p>";
        accountInfoHtml += "<p>Name: " + content.name + "</p>";
        accountInfoHtml += "<p>Email: " + content.email + "</p>";
        accountInfoHtml += "<p>Created At: " + formattedCreatedAt+ "</p>";
        accountInfoHtml += "<p>User Role: " + content.UserRole + "</p>";
        accountInfoDiv.innerHTML = accountInfoHtml;
    } else {
        console.log("There was an error retrieving data from the API");
    }
    } catch (err) {
        console.log(err.message);
    }
};

function formatDate(dateString) {
    const date = new Date(dateString);
    const options = {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        timeZoneName: 'short',
    };
    return date.toLocaleString('en-US', options);
}


    
      
    async function getSubmissions() {
        console.log("clicked get submissions button");
        try{
            const res = await fetch("https://api.irvyn.xyz/submissions", {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                    "Accept": "application/json"
                },
                xhrFields: {
                    withCredentials: true
                },
                credentials: "include"
            });
            const content = await res.json();
            console.log(content);
            console.log(res.status);
            if (res.status === 200) {
                const submissionsDiv = document.querySelector(".submissions");
                let submissionsHtml = "";
                for (let i = 0; i < content.length; i++) {
                    const formattedCreatedAt = formatDate(content[i].created_at);

                    submissionsHtml += `<div class='submission bg-gray-100 my-4 p-4 rounded-lg shadow-md'>`;
                    submissionsHtml += `<p class='font-bold text-gray-700'>Submission ID: <span class='font-normal'>${content[i].id}</span></p>`;
                    submissionsHtml += `<p class='font-bold text-gray-700'>Submission Code: <span class='font-normal'>${content[i].code}</span></p>`;
                    submissionsHtml += `<p class='font-bold text-gray-700'>Submission Output: <span class='font-normal'>${content[i].output}</span></p>`;
                    submissionsHtml += `<p class='font-bold text-gray-700'>Submission Error: <span class='font-normal'>${content[i].error}</span></p>`;
                    submissionsHtml += `<p class='font-bold text-gray-700'>Submission Created At: <span class='font-normal'>${formattedCreatedAt}</span></p>`;
                    submissionsHtml += `</div>`;
                }
                submissionsDiv.innerHTML = submissionsHtml;
            } else {
                console.log("There was an error retrieving data from the API");
            }

        }
        catch (err) {
            console.log(err.message);
        }
    }



