// use this file to add the logout() function 
// this function is called when the user clicks the logout button
// it will send a request to the backend logout route where the session will be destroyed
// then it will redirect the user to the home page 
console.log("wtf js has been loaded, and live updates are working?")


function logout() {
    console.log("logout function has been called")
    // send a request to the backend logout route (api.irvyn.xyz/logout)
    // redirect the user to the home page

    $.ajax({
        url: "https://api.irvyn.xyz/logout",
        type: "POST",
        xhrFields: {
            withCredentials: true
        },
        credentials: "include",
        success: function (response) {
            console.log(response)
            window.location.href = "/";
        }
    });
    

}