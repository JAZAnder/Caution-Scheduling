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

async function scheduleMeeting(userHourId, labId, studentName, studentEmail, date){
    
    

    data = await meeting_create(userHourId, labId, studentName, studentEmail, date)
    if(data['error']){
        alert("Error : Could not Create Meeting")
    }else{
        alert("Meeting Sceduled!!")
        document.getElementById("Name").value = ""
        document.getElementById("email").value = ""
    }
}
// Load the readOnlyTable.html into the 'readOnlyTable' div
importElements('readOnlyTable', './assets/elements/readOnlyTable.html');

async function addLab(){
    data = await lab_getall()
    if(data['error']){
        document.getElementById("error").innerHTML = data['error']
        console.log("error : " + data['error'])
    }
    data.forEach(addLabForEach)

}
function addLabForEach(item){
    const optionInfo = new lab(item)
    var options = document.getElementById("Lab-Id-Select")
    var option = document.createElement("option")

    option.text = optionInfo.name + "(" + optionInfo.labLocation + ")"
    option.value = optionInfo.Id

    options.add(option)
}

async function loadTime(){
    userName = document.getElementById("tutorId-Select").value
    data = await userHour_GetMine(userName)
    var options = document.getElementById("hourId-Select")
    objectlength = options.length
    for (var i=options.length; i>=0; i--) {
        options.remove(i);
    }

    data.forEach(loadTimeForEach)
}

async function loadTimeForEach(item){
    const optionInfo = new userHour(item)
    var options = document.getElementById("hourId-Select")
    var option = document.createElement("option")

    var hourItem = new hour(await hour_getById(optionInfo.hourId)) 

    option.text = hourItem.startTime + " - " + hourItem.endTime
    option.value = optionInfo.id

    options.add(option)

}


// NEW CODES

async function loadTimesFromDate(dayOfWeek){
    data = await hour_getByDay(dayOfWeek)
    var options = document.getElementById("hourId-Select")
    for (var i=options.length; i>=0; i--) {
        options.remove(i);
    }
    var options = document.getElementById("tutorId-Select")
    for (var i=options.length; i>=0; i--) {
        options.remove(i);
    }
    data.forEach(hourOptionsForEach)
}

async function userOptionsByHour(hourId){
    var options = document.getElementById("tutorId-Select")
    userHours = await userhour_GetTutorByHour(hourId)
    if(data['error']){
        document.getElementById("error").innerHTML = data['error']
        console.log("error : " + data['error'])
    }

    for (var i=options.length; i>=0; i--) {
        options.remove(i);
    }
    userHours.forEach(userHoursForEach)
}

async function userHoursForEach(item){
    const uh = new userHour(item) 
    const optionInfo =  new luser(await user_getInfo(uh.tutor))

    var options = document.getElementById("tutorId-Select")
    var option = document.createElement("option")

    option.text = optionInfo.firstName + "  " + optionInfo.lastName
    option.value = uh.id

    options.add(option)
}



                    