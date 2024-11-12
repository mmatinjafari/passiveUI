function gatherUrls() {
    const domain = document.getElementById('domain').value;
    const output = document.getElementById('output');
    const loading = document.getElementById('loading');

    if (!domain) {
        alert("Please enter a domain.");
        return;
    }

    output.textContent = "";
    loading.style.display = "block";

    fetch(`/gather-urls?domain=${encodeURIComponent(domain)}`)
        .then(response => response.text())
        .then(data => {
            loading.style.display = "none";
            output.textContent = data;
        })
        .catch(error => {
            loading.style.display = "none";
            output.textContent = `Error: ${error}`;
        });
}
