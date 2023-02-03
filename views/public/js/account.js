console.log("this should print if we are on the account page");

// send get request to api.irvyn.xyz/account
// display the account information

// example response from my postman testing
// {
//     "id": 27,
//     "name": "irvyn",
//     "email": "gg@crypto.com",
//     "password": "mYu96jgcvsA5A+MqEy8iMv/dksFiW2uCjvMH9ZQErbs=",
//     "created_at": "2023-01-29T14:03:07.115309Z",
//     "UserRole": 1
// }


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
    accountInfoHtml += "<p>Password: " + content.password + "</p>";
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

