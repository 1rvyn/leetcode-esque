// Get form values
const form = $(".register-form");
const email = $("#email").val();
const name = $("#name").val();
const password = $("#password").val();

// Set up request body
const requestBody = {
  email: email,
  name: name,
  password: password
};

// Set up headers
const headers = {
  "Content-Type": "application/json",
  "Access-Control-Allow-Origin": "*"
};

// Send POST request to API
$.ajax({
  type: "POST",
  url: "https://irvyn.dev/api/register",
  headers: headers,
  data: JSON.stringify(requestBody),
  success: function(data) {
    console.log(data);
  },
  error: function(error) {
    console.error("There was a problem with the request:", error);
  }
});

// Prevent form from refreshing the page
form.submit(function(event) {
  event.preventDefault();
});
