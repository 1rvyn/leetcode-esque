const form = document.querySelector('form');

form.addEventListener('submit', event => {
    event.preventDefault();

    const formData = new FormData(form);
    const data = Object.fromEntries(formData);

    fetch('https://api.irvyn.xyz/add', {
        method: 'POST',
        headers: {"Content-Type": "application/json",
            "Accept": "application/json"},
        body: JSON.stringify(data),
        xhrFields: {
            withCredentials: true
        },
        credentials: "include",
    })
        .then(response => response.json())
        .then(data => console.log(data))
        .catch(error => console.error(error));
});