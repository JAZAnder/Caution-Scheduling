//const baseUrl = "http://local.techwall.xyz"

const user_whoami_Url = baseUrl+"/api/luser/whoami"
const user_login_Url = baseUrl+"/api/luser/login"
const user_logout_Url = baseUrl+"/api/luser/logout"
const user_create_Url = baseUrl+"/api/luser"
const user_getall_Url = baseUrl+"/api/lusers"


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