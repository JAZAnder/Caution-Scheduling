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

async function addLab(){
    var options = document.getElementById("tutorId-Select")
    var option = document.createElement("option")
    data = await lab_getall()
    if(data['error']){
        document.getElementById("error").innerHTML = data['error']
        console.log("error : " + data['error'])
    }
    data.forEach(addLabForEach)

}
function addLabForEach(item){
    const optionInfo = new lab(item)
    var options = document.getElementById("tutorId-Select")
    var option = document.createElement("option")

    option.text = optionInfo.name
    option.value = optionInfo.Id

    options.add(option)
}

// Load the readOnlyTable.html into the 'readOnlyTable' div
importElements('readOnlyTable', './assets/elements/readOnlyTable.html');

