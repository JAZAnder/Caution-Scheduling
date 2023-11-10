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

async function userOptions(){
    var options = document.getElementById("tutorId-Select")
    var option = document.createElement("option")
    data = await user_getall()
    if(data['error']){
        document.getElementById("error").innerHTML = data['error']
        console.log("error : " + data['error'])
    }
    data.forEach(userOptionsForEach)
}

function userOptionsForEach(item){
    const optionInfo = new luser(item)
    var options = document.getElementById("tutorId-Select")
    var option = document.createElement("option")

    option.text = optionInfo.firstName + "  " + optionInfo.lastName
    option.value = optionInfo.username

    options.add(option)
}

async function hourOptions(){
    var options = document.getElementById("hourId-Select")
    var option = document.createElement("option")
    data = await hour_getAll()
    if(data['error']){
        document.getElementById("error").innerHTML = data['error']
        console.log("error : " + data['error'])
    }
    data.forEach(hourOptionsForEach)
}

function hourOptionsForEach(item){
    const optionInfo = new hour(item)
    var options = document.getElementById("hourId-Select")
    var option = document.createElement("option")

    option.text = optionInfo.startTime + " - " + optionInfo.endTime
    option.value = optionInfo.id

    options.add(option)
}

async function scheduleMeeting(){
    tutorHourId = 0
    labId = document.getElementById("hourId-Select").value
    studentName = document.getElementById("Name").value
    studentEmail = document.getElementById("email").value
    data = await meeting_create(tutorHourId, labId, studentName, studentEmail)
    if(data['error']){
        document.getElementById("error").innerHTML = data['error']
        alert("Error : "+data['error'])
    }else{
        alert("Meeting Sceduled!!")
    }
}
// Load the readOnlyTable.html into the 'readOnlyTable' div
importElements('readOnlyTable', './assets/elements/readOnlyTable.html');

