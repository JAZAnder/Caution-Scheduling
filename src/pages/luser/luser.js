async function checkLogin(){
    key = getCookie("key")
    if(!key){
        window.location.replace("./login");
    }
    
    const user = new luser(await user_whoami())
    document.getElementById('username').innerHTML = user.firstName
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

    data = await hour_create(StartTime, EndTime)

    if(data['error']){
        console.log("Logout Error :"+ data['error'])
        document.getElementById("error").innerHTML = data['error']
    }else{
        alert("Time Create, refresh to populate")
        document.getElementById("StartTime").value = ""
        document.getElementById("EndTime").value = ""
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

    var userName = row.insertCell(0);
    var firstName = row.insertCell(1);
    var lastName = row.insertCell(2);
    var email = row.insertCell(3);
    var isAdmin = row.insertCell(4);

    userName.innerHTML = rowInfo.username
    firstName.innerHTML = rowInfo.firstName
    lastName.innerHTML = rowInfo.lastName
    email.innerHTML = rowInfo.email
    isAdmin.innerHTML = rowInfo.isAdmin
}

async function populateTime(){
    data = await hour_getAll()
    
    data.forEach(populateTimeForEach)
}

function populateTimeForEach(item){
    const rowInfo = new hour(item)

    var table = document.getElementById("time-table");
    var row = table.insertRow(1);

    var id = row.insertCell(0);
    var startTime = row.insertCell(1);
    var endTime = row.insertCell(2);


    id.innerHTML = rowInfo.id
    startTime.innerHTML = rowInfo.startTime
    endTime.innerHTML = rowInfo.endTime
}


async function populateMeetings(){
    data = await meeting_GetAll()
    
    data.forEach(populateMeetingsForEach)
}

function populateMeetingsForEach(item){
    const rowInfo = new meeting(item)

    var table = document.getElementById("meeting-table");
    var row = table.insertRow(1);

    var id = row.insertCell(0);
    var studentName = row.insertCell(1);
    var studentEmail = row.insertCell(2);
    var tutorHourId = row.insertCell(3);


    id.innerHTML = rowInfo.id
    studentName.innerHTML = rowInfo.studentName
    studentEmail.innerHTML = rowInfo.studentEmail
    tutorHourId.innerHTML = rowInfo.tutorHourId

}