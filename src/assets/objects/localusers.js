const baseUrl = "http://local.techwall.xyz"

const whoamiUrl = baseUrl+"/api/luser/whoami"
const loginUrl = baseUrl+"/api/luser/login"

function userwhoami(){
  
  var myHeaders = new Headers();
  myHeaders.append("Cookie", "key=ydvu/XZ8D+agCQ==");

  var requestOptions = {
    method: 'GET',
    headers: myHeaders,
    redirect: 'follow'
  };

  fetch(whoamiUrl, requestOptions)
    .then(response => response.text())
    .then(result => console.log(result))
    .catch(error => console.log('error', error))
  
}

function userlogin(username, password){
  var contents = ""

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

  return fetch(loginUrl, requestOptions)
  .then(response => response.json())
  .then(result => {
    console.log(result)
  })
  .catch(error => console.log('error', error));
}