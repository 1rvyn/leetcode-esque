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
    const accountInfoDiv = document.querySelector(".account-info");
    let accountInfoHtml = "";
    accountInfoHtml += "<p>ID: " + content.id + "</p>";
    accountInfoHtml += "<p>Name: " + content.name + "</p>";
    accountInfoHtml += "<p>Email: " + content.email + "</p>";
    accountInfoHtml += "<p>Created At: " + content.created_at + "</p>";
    accountInfoHtml += "<p>User Role: " + content.UserRole + "</p>";
    accountInfoDiv.innerHTML = accountInfoHtml;
    } else {
    console.log("There was an error retrieving data from the API");
    }
    } catch (err) {
    console.log(err.message);
    }
    };

    
      
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
                    submissionsHtml += "<div class='submission'>";
                    submissionsHtml += "<p>Submission ID: " + content[i].id + "</p>";
                    submissionsHtml += "<p>Submission Code: " + content[i].code + "</p>";
                    submissionsHtml += "<p>Submission Output: " + content[i].output + "</p>";
                    submissionsHtml += "<p>Submission Error: " + content[i].error + "</p>";
                    submissionsHtml += "<p>Submission Created At: " + content[i].created_at + "</p>";
                    submissionsHtml += "</div>";
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



