//const baseUrl = "http://local.techwall.xyz"

const user_whoami_Url = baseUrl+"/api/luser/whoami"
const user_login_Url = baseUrl+"/api/luser/login"
const user_logout_Url = baseUrl+"/api/luser/logout"
const user_create_Url = baseUrl+"/api/luser"
const user_getall_Url = baseUrl+"/api/lusers"
const user_resetpassword_Url = baseUrl+"/api/luser/resetmypasswd"


function luser(user){
  this.username = user['userName']
  this.firstName = user['firstName']
  this.lastName = user['lastName']
  this.email = user['email']
  this.isAdmin = user['isAdmin']
}

async function user_whoami(){

  var requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };

  const result = await fetch(user_whoami_Url, requestOptions)
  const data = await result.json();
  console.log(data)
  return data
}

async function user_login(username, password){
  // var contents = ""

  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

  var urlencoded = new URLSearchParams();
  urlencoded.append("userName", username);
  urlencoded.append("password", password);

  var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: urlencoded,
      redirect: 'follow'
  };

  const result = await fetch(user_login_Url, requestOptions)
  const data = await result.json();
  console.log(data)
  return data
}

async function user_logout(){

  var requestOptions = {
    method: 'DELETE',
    redirect: 'follow'
  };

  const result = await fetch(user_logout_Url, requestOptions)
  const data = await result.json();
  console.log(data)
  return data
}

async function user_getall(){
  var requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };

  const result = await fetch(user_getall_Url, requestOptions)
  const data = await result.json();
  console.log(data)
  return data
}

async function user_create(userName, firstName, lastName, email, password, isAdmin){
  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

  var urlencoded = new URLSearchParams();
  urlencoded.append("userName", userName);
  urlencoded.append("firstName", firstName);
  urlencoded.append("lastName", lastName);
  urlencoded.append("email", email);
  urlencoded.append("password", password);
  urlencoded.append("isAdmin", isAdmin);

  var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: urlencoded,
      redirect: 'follow'
  };

  const result = await fetch(user_create_Url, requestOptions)
  const data = await result.json();
  console.log(data)
  return data
}

async function user_resetpassword(oldPassword, newPassword){
  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

  var urlencoded = new URLSearchParams();
  urlencoded.append("oldPassword", oldPassword);
  urlencoded.append("newPassword", newPassword);
  var requestOptions = {
      method: 'PUT',
      headers: myHeaders,
      body: urlencoded,
      redirect: 'follow'
  };

  const result = await fetch(user_resetpassword_Url, requestOptions)
  const data = await result.json();
  console.log(data)
  return data
}