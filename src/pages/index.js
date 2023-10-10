function getNavbar(id){
    document.getElementById(id).innerHTML = "Hello"
    console.log("Navbar Imported")
}

function importElements(elementId, url) {
    const element = document.getElementById(elementId);
    
    if (element) {
        fetch(url)
            .then(response => response.text())
            .then(html => {
                element.innerHTML = html;
            })
            .catch(error => {
                console.error(`Error loading ${url}:`, error);
            });
    }
}

// Load the readOnlyTable.html into the 'readOnlyTable' div
importElements('readOnlyTable', './assets/elements/readOnlyTable.html');

