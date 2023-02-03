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


const request = new XMLHttpRequest();
const url = 'https://api.irvyn.xyz/account';
const div = document.querySelector('.account-info');

request.open('POST', url, true);
request.onload = function() {
  if (request.status >= 200 && request.status < 400) {
    const data = JSON.parse(request.responseText);
    const content = `
      <p><b>ID:</b> ${data.id}</p>
      <p><b>Name:</b> ${data.name}</p>
      <p><b>Email:</b> ${data.email}</p>
      <p><b>Created At:</b> ${data.created_at}</p>
      <p><b>User Role:</b> ${data.UserRole}</p>
    `;

    div.innerHTML = content;
    div.style.backgroundColor = '#f2f2f2';
    div.style.padding = '20px';
    div.style.borderRadius = '5px';
  } else {
    console.error('Error retrieving data from API');
  }
};

request.send();


