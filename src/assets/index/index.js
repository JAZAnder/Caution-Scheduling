function getNavbar(id){
    document.getElementById(id).innerHTML = "Hello"
    console.log("Navbar Imported")
}

async function importElements(id, url){
    fetch(url).then(function (response) {
        // The API call was successful!
        return response.text();
    }).then(function (html) {
        // This is the HTML from our response as a text string
        document.getElementById(id).innerHTML = html;      
    }).catch(function (err) {
        // There was an error
        console.warn('Something went wrong.', err);
    })
    
}