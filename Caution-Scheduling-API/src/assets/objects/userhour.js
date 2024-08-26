const userHour_getById_Url = baseUrl + "/api/tutor/timeslot/whois/"
const userHour_create_Url = baseUrl + "/api/luser/admin/timeslot"
const userHour_GetAll_Url = baseUrl + "/api/tutor/timeslots"
const userHour_MyCreate_Url = baseUrl + "/api/luser/timeslot"
const userHour_GetMine_Url = baseUrl + "/api/tutor/hours/"
const userHour_GetTutorByHour_Url = baseUrl  + "/api/hour/availability/"

function userHour(userHour){
    this.id = userHour['id']
    this.hourId = userHour['hourId']
    this.tutor = userHour['tutor']
    this.available = userHour['available']
}

async function userhour_GetById(id){
    url = userHour_getById_Url + id
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      };
    
      const result = await fetch(url, requestOptions)
      const data = await result.json();
      console.log(data)
      return data
}

async function userHour_Create(tutorId, hourId){
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/x-www-form-urlencoded");
  
    var urlencoded = new URLSearchParams();
    urlencoded.append("hourId", tutorId);
    urlencoded.append("username", hourId);
  
    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: urlencoded,
        redirect: 'follow'
    };

    const result = await fetch(userHour_create_Url, requestOptions)
  const data = await result.json();
  console.log(data)
  return data
  
}
async function userHour_MyCreate(hourId){
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/x-www-form-urlencoded");
  
    var urlencoded = new URLSearchParams();
    urlencoded.append("hourId", hourId);
  
    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: urlencoded,
        redirect: 'follow'
    };

    const result = await fetch(userHour_MyCreate_Url, requestOptions)
  const data = await result.json();
  console.log(data)
  return data
  
}
async function userHour_GetAll(){
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      };
    
      const result = await fetch(userHour_GetAll_Url, requestOptions)
      const data = await result.json();
      console.log(data)
      return data
}

async function userHour_GetMine(userName){
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      };
      
      const result = await fetch(userHour_GetMine_Url+userName, requestOptions)
      const data = await result.json();
      console.log(data)
      return data
}

async function userhour_GetTutorByHour(hourId){
  url = userHour_GetTutorByHour_Url + hourId
  var requestOptions = {
      method: 'GET',
      redirect: 'follow'
    };
  
    const result = await fetch(url, requestOptions)
    const data = await result.json();
    console.log(data)
    return data
}
