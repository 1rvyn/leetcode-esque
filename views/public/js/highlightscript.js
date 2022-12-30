document.addEventListener('DOMContentLoaded', function() {
    // Get the textarea element
    var textarea = document.getElementById('code');

    // Add a change event listener to the textarea
    textarea.addEventListener('change', function() {
        // Get the textarea value and apply Prism.js syntax highlighting
        var code = textarea.value;
        var highlightedCode = Prism.highlight(code, Prism.languages.python);

        // Set the textarea value to the highlighted code
        textarea.value = highlightedCode;
    });
});
