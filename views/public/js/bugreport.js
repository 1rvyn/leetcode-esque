function handleBugReport(title, bug){
    console.log("Bug report title: " + title + "\n" + "Bug report: " + bug);

    // will post 

    try {
        fetch("https://api.irvyn.xyz/bugreport", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "Accept": "application/json"
          },
          body: JSON.stringify({
            title: title.value,
            bugReport: bug.value
          }),
          xhrFields: {
            withCredentials: true
          },
          credentials: "include"
        })
          .then(res => {
            console.log(res.status);
            if (res.status === 200) {
              // Show success message
              const successMessage = document.createElement("p");
              successMessage.textContent = "Successfully made a bug report!";
              successMessage.style.color = "green";
              form.appendChild(successMessage);
            } else {
              // Show error message
              const errorMessage = document.createElement("p");
              errorMessage.textContent = "There was an error submitting the bug report.";
              errorMessage.style.color = "red";
              form.appendChild(errorMessage);
            }
          })
          .catch(err => {
            console.log(err.message);
          });
      } catch (err) {
        console.log(err.message);
      }
}