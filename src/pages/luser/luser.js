async function checkLogin(){
    key = getCookie("key")
    if(!key){
        window.location.replace("./login");
    }
    
    const user = new luser(await user_whoami())
    document.getElementById('username').innerHTML = user.firstName
}


async function loadNavBar(){
    const user = new luser(await user_whoami())
    if(user.isAdmin == true){
        importElements("navBar", "./admin-zone/header.html")
    }
    if(user.isAdmin == false){
        importElements("navBar", "./user-zone/header.html")
    }
}


async function logout(){
    data = await user_logout()
    if(data['error']){
        console.log("Logout Error :"+ data['error'])
    }else{
        console.log("Logged Out")
        window.location.reload()
    }
}

async function createUser(){
    const userName = document.getElementById("userName").value
    const firstName = document.getElementById("firstName").value
    const lastName = document.getElementById("lastName").value
    const email = document.getElementById("email").value
    const password = document.getElementById("password").value
    var isAdmin = false;
    var ele = document.getElementsByName('isAdmin');
    for (i = 0; i < ele.length; i++) {
        if (ele[i].checked)
           isAdmin = ele[i].value
    }
    data = await user_create(userName, firstName, lastName, email, password, isAdmin);
    if(data['error']){
        console.log("Logout Error :"+ data['error'])
        document.getElementById("error").innerHTML = data['error']
    }else{
        alert("New User Created, refresh to populate")
        document.getElementById("userName").value = ""
        document.getElementById("firstName").value = ""
        document.getElementById("lastName").value = ""
        document.getElementById("email").value = ""
        document.getElementById("password").value = ""

    }
} 

async function createHour(){
    const StartTime = document.getElementById("StartTime").value
    const EndTime = document.getElementById("EndTime").value
    const dayOfWeek = document.getElementById("dayOfWeek-Select").value

    data = await hour_create(StartTime, EndTime, dayOfWeek)

    if(data['error']){
        console.log("Logout Error :"+ data['error'])
        document.getElementById("error").innerHTML = data['error']
    }else{
        alert("Time Create, refresh to populate")
        document.getElementById("StartTime").value = ""
        document.getElementById("EndTime").value = ""
    }
}











//START TIMESLOTS PAGE =======================================================
async function populateTime(){
    data = await hour_getAll()
    
    data.forEach(populateTimeForEach)
}

function populateTimeForEach(item){
    const rowInfo = new hour(item)

    var table = document.getElementById("time-table");
    var row = table.insertRow(1);

    var id = row.insertCell(0);
    var dayOfWeek = row.insertCell(1);
    var startTime = row.insertCell(2);
    var endTime = row.insertCell(3);


    id.innerHTML = rowInfo.id
    startTime.innerHTML = rowInfo.startTime
    endTime.innerHTML = rowInfo.endTime
    switch(rowInfo.dayOfWeek){
        case 0:
            dayOfWeek.innerHTML = "Sunday"
            break;
        case 1:
            dayOfWeek.innerHTML = "Monday"
            break;
        case 2:
            dayOfWeek.innerHTML = "Tuesday"
            break;
        case 3:
            dayOfWeek.innerHTML = "Wednesday"
            break;
        case 4:
            dayOfWeek.innerHTML = "Thursday"
            break;
        case 5:
            dayOfWeek.innerHTML = "Friday"
            break;
        case 6:
            dayOfWeek.innerHTML = "Saturday"
            break;
    }
}
//END TIMESLOTS PAGE =======================================================








//START MEETINGS PAGE =======================================================
async function populateMeetings(){
    data = await meeting_GetAll()
    
    data.forEach(populateMeetingsForEach)
}

async function populateMeetingsForEach(item){
    const rowInfo = new meeting(item)

    var table = document.getElementById("meeting-table");
    var row = table.insertRow(1);

    var id = row.insertCell(0);
    var studentName = row.insertCell(1);
    var studentEmail = row.insertCell(2);
    var tutorName = row.insertCell(3);
    var time = row.insertCell(4);
    var location = row.insertCell(5);

    var itemLab = new lab(await lab_getById(item.labId)) 
    var itemUserHour = new userHour(await userhour_GetById(item.userHourId))
    var itemTime = new hour(await hour_getById(itemUserHour.hourId))

    id.innerHTML = rowInfo.id
    studentName.innerHTML = rowInfo.studentName
    studentEmail.innerHTML = rowInfo.studentEmail
    tutorName.innerHTML = itemUserHour.tutor
    time.innerHTML = itemTime.startTime + " - " + itemTime.endTime
    location.innerHTML = itemLab.name
}

async function loadMeetingById(location){
    id = document.getElementById('MeetingId').value
    var mymeeting  = new meeting(await meeting_GetById(id))
    var itemUserHour = new userHour(await userhour_GetById(mymeeting.tutorHourId))
    

    document.getElementById(location).innerHTML = "Student: "+mymeeting.studentName + " | Email: " + mymeeting.studentEmail + " | Tutor: "+ itemUserHour.tutor


}
//END MEETINGS PAGE =======================================================








//START MY ACCOUNT PAGE==================================================
async function fillInMyAccount(){
    const user = new luser(await user_whoami())
    document.getElementById('headName').innerHTML = user.firstName
    document.getElementById('userName').innerHTML = user.username
    document.getElementById('firstName').innerHTML = user.firstName
    document.getElementById('lastName').innerHTML = user.lastName
    document.getElementById('emailAddress').innerHTML = user.email
}

async function resetpassword(){
    oldPassword = document.getElementById('oldPW').value
    newPassword = document.getElementById('newPW').value
    data = await user_resetpassword(oldPassword, newPassword);
    if(data['error']){
        console.log("Logout Error :"+ data['error'])
        alert("Failed to Reset Password")
        document.getElementById('oldPW').value = ""
        document.getElementById('newPW').value = ""
    }else{
        alert("Password Reset")
        document.getElementById('oldPW').value = ""
        document.getElementById('newPW').value = ""
    }
}
//END MY ACCOUNT PAGE==================================================








//START LAB PAGE =======================================================
async function newLab(){
    labName = document.getElementById('labName').value
    lablocation = document.getElementById('lablocation').value
    data = await lab_create(labName, lablocation);
    if(data['error']){
        console.log("Creation Errer :"+ data['error'])
        alert("Failed to Create New Lab")
    }else{
        alert("New Lab Created")
        document.getElementById('labName').value = ""
        document.getElementById('lablocation').value = ""
    }
}


async function populateLabs(){
    data = await lab_getall()
    
    data.forEach(populatelabsForEach)
}

function populatelabsForEach(item){
    const rowInfo = new lab(item)

    var table = document.getElementById("lab-table");
    var row = table.insertRow(1);

    var id = row.insertCell(0);
    var labName = row.insertCell(1);
    var labLocation = row.insertCell(2);


    id.innerHTML = rowInfo.Id
    labName.innerHTML = rowInfo.name
    labLocation.innerHTML = rowInfo.labLocation
}

async function populateLabHours(){
    data = await labHour_getall()

    data.forEach(populateLabHoursForEach)
}
async function populateLabHoursForEach(item){
    const rowInfo = new labHour(item)

    var table = document.getElementById("labhour-table");
    var row = table.insertRow(1);

    var id = row.insertCell(0);
    var labName = row.insertCell(1);
    var dayOfWeek = row.insertCell(2);
    var StartTime = row.insertCell(3);
    var EndTime = row.insertCell(4);
    var Tutor = row.insertCell(5);

    var itemLab = new lab(await lab_getById(item.labId)) 
    var itemTime = new hour(await hour_getById(item.hourId))
    var itemUser = new userHour(await userhour_GetById(item.userHourId))

    id.innerHTML = rowInfo.Id
    labName.innerHTML = itemLab.name
    StartTime.innerHTML = itemTime.startTime
    EndTime.innerHTML = itemTime.endTime
    Tutor.innerHTML = itemUser.tutor

    switch(itemTime.dayOfWeek){
        case 0:
            dayOfWeek.innerHTML = "Sunday"
            break;
        case 1:
            dayOfWeek.innerHTML = "Monday"
            break;
        case 2:
            dayOfWeek.innerHTML = "Tuesday"
            break;
        case 3:
            dayOfWeek.innerHTML = "Wednesday"
            break;
        case 4:
            dayOfWeek.innerHTML = "Thursday"
            break;
        case 5:
            dayOfWeek.innerHTML = "Friday"
            break;
        case 6:
            dayOfWeek.innerHTML = "Saturday"
            break;
    }
}

async function addLab(){
    var options = document.getElementById("Lab-Id-Select")
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
    var options = document.getElementById("Lab-Id-Select")
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

    switch(optionInfo.dayOfWeek){
        case 0:
            dayOfWeek = "Sunday"
            break;
        case 1:
            dayOfWeek = "Monday"
            break;
        case 2:
            dayOfWeek = "Tuesday"
            break;
        case 3:
            dayOfWeek = "Wednesday"
            break;
        case 4:
            dayOfWeek = "Thursday"
            break;
        case 5:
            dayOfWeek = "Friday"
            break;
        case 6:
            dayOfWeek = "Saturday"
            break;
    }

    option.text = dayOfWeek +": "+ optionInfo.startTime + " - " + optionInfo.endTime
    option.value = optionInfo.id

    options.add(option)
}

async function createlabHour(){
    const labId = document.getElementById("Lab-Id-Select").value
    const timeSlotId = document.getElementById("hourId-Select").value
    const userName = document.getElementById("tutorId-Select").value
    data = await labHour_Create(labId, timeSlotId, userName)
    if(data['error']){
        console.log("Logout Error :"+ data['error'])

        alert("Failed to Create a new Timeslot\n"+data['error'])
    }else{
        alert("New Lab Hours Created, refresh to populate")
    }

}
//START END PAGE =======================================================





//START EMPLOYEE PAGE ======================================== 
async function createUserHour(){
    const timeSlotId = document.getElementById("hourId-Select").value
    const userName = document.getElementById("tutorId-Select").value
    data = await userHour_Create(timeSlotId, userName)
    if(data['error']){
        console.log("Logout Error :"+ data['error'])

        alert("Failed to Create a new Timeslot\n"+data['error'])
    }else{
        alert("New Availability Created, refresh to populate")
    }
}

async function loadUserhours(){
    data = await userHour_GetAll()

    data.forEach(loadUserHourForEach)
}

async function loadUserHourForEach(item){
    const rowInfo = new userHour(item)
    var table = document.getElementById("userhour-table");
    var row = table.insertRow(1);

    var id = row.insertCell(0);
    var Employee = row.insertCell(1);
    var dayOfWeek = row.insertCell(2);
    var StartTime = row.insertCell(3);
    var EndTime = row.insertCell(4);

    var itemTime = new hour(await hour_getById(item.hourId))
    
    id.innerHTML = rowInfo.id
    Employee.innerHTML = rowInfo.tutor
    StartTime.innerHTML = itemTime.startTime
    EndTime.innerHTML = itemTime.endTime

    switch(itemTime.dayOfWeek){
        case 0:
            dayOfWeek.innerHTML = "Sunday"
            break;
        case 1:
            dayOfWeek.innerHTML = "Monday"
            break;
        case 2:
            dayOfWeek.innerHTML = "Tuesday"
            break;
        case 3:
            dayOfWeek.innerHTML = "Wednesday"
            break;
        case 4:
            dayOfWeek.innerHTML = "Thursday"
            break;
        case 5:
            dayOfWeek.innerHTML = "Friday"
            break;
        case 6:
            dayOfWeek.innerHTML = "Saturday"
            break;
    }

}

async function populateUsers(){
    data = await user_getall()
    
    data.forEach(populateUsersForEach)
}

function populateUsersForEach(item){
    const rowInfo = new luser(item)

    var table = document.getElementById("users-Table");
    var row = table.insertRow(1);

    var UserId = row.insertCell(0);
    var userName = row.insertCell(1);
    var firstName = row.insertCell(2);
    var lastName = row.insertCell(3);
    var email = row.insertCell(4);
    var isAdmin = row.insertCell(5);

    UserId.innerHTML = "0"
    userName.innerHTML = rowInfo.username
    firstName.innerHTML = rowInfo.firstName
    lastName.innerHTML = rowInfo.lastName
    email.innerHTML = rowInfo.email
    isAdmin.innerHTML = rowInfo.isAdmin
    
}

function resetPasswordModel(){
    var modal = document.getElementById("Password_Reset_Modal");
    modal.style.display = "block";
}

async function userOptionsForReset(){
    data = await user_getall()
    if(data['error']){
        document.getElementById("error").innerHTML = data['error']
        console.log("error : " + data['error'])
    }
    data.forEach(userOptionsForResetForEach)
}

function userOptionsForResetForEach(item){
    const optionInfo = new luser(item)
    var options = document.getElementById("userNameForReset-select")
    var option = document.createElement("option")

    option.text = optionInfo.firstName + "  " + optionInfo.lastName
    option.value = optionInfo.username

    options.add(option)
}

//START USER MEETINGS ======================================== 
async function populateMyMeetings(){
    data = await meeting_GetMine()
    
    data.forEach(populateMyMeetingsForEach)
}

async function populateMyMeetingsForEach(item){
    const rowInfo = new meeting(item)

    var table = document.getElementById("meeting-table");
    var row = table.insertRow(1);

    var id = row.insertCell();
    var studentName = row.insertCell(1);
    var studentEmail = row.insertCell(2);
    var time = row.insertCell(3);
    var location = row.insertCell(4);

    var itemLab = new lab(await lab_getById(item.labId)) 
    var itemUserHour = new userHour(await userhour_GetById(item.userHourId))
    var itemTime = new hour(await hour_getById(itemUserHour.hourId))

    id.innerHTML = rowInfo.id
    studentName.innerHTML = rowInfo.studentName
    studentEmail.innerHTML = rowInfo.studentEmail
    time.innerHTML = itemTime.startTime + " - " + itemTime.endTime
    location.innerHTML = itemLab.name
}

//END USER MEETINGS

//START USER-MEETINGS
async function createUserHour(){
    const timeSlotId = document.getElementById("hourId-Select").value
    console.log(document.getElementById("hourId-Select").value)
    data = await userHour_MyCreate(timeSlotId)
    if(data['error']){
        console.log("Logout Error :"+ data['error'])

        alert("Failed to Create a new Timeslot\n"+data['error'])
    }else{
        alert("New Availability Created, refresh to populate")
    }
}

async function loadMyUserhours(){

    var user = new luser(await user_whoami())

    data = await userHour_GetMine(user.username)

    data.forEach(loadUserHourForEach)
}

async function loadMyUserHourForEach(item){
    const rowInfo = new userHour(item)
    var table = document.getElementById("userhour-table");
    var row = table.insertRow(1);

    var id = row.insertCell(0);
    var StartTime = row.insertCell(2);
    var EndTime = row.insertCell(3);

    var itemTime = new hour(await hour_getById(item.hourId))
    
    id.innerHTML = rowInfo.id
    StartTime.innerHTML = itemTime.startTime
    EndTime.innerHTML = itemTime.endTime

    
}