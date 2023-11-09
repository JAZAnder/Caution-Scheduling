const userHour_getById_Url = baseUrl + "/api/tutor/timeslot/whois/"

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