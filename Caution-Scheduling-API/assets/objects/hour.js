const hour_create_Url = baseUrl+"/api/hour"
const hour_get_Url = baseUrl+"/api/hour/"
const hour_delete_Url = baseUrl+"/api/hour/"
const hour_getAll_Url = baseUrl+"/api/hours"
const hour_getByDay_Url = baseUrl+"/api/hour/day/"

function hour(hour){
    this.id = hour['id']
    this.startTime = hour['startTime']
    this.endTime = hour['endTime']
    this.dayOfWeek = hour['dayOfWeek']
}

async function hour_create(startTime, EndTime, dayOfWeek){
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

    var urlencoded = new URLSearchParams();
    urlencoded.append("startTime", startTime);
    urlencoded.append("endTime", EndTime);
    urlencoded.append("dayOfWeek", dayOfWeek);

    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: urlencoded,
        redirect: 'follow'
    };

    const result = await fetch(hour_create_Url, requestOptions)
    const data = await result.json();
    console.log(data)
    return data
}

async function hour_getAll(){
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      };
    
      const result = await fetch(hour_getAll_Url, requestOptions)
      const data = await result.json();
      console.log(data)
      return data
}

async function hour_getById(Id){
    url = hour_get_Url + Id
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      };
    
      const result = await fetch(url, requestOptions)
      const data = await result.json();
      console.log(data)
      return data
}

async function hour_deleteById(Id){
    url = hour_delete_Url + Id
    var requestOptions = {
        method: 'DELETE',
        redirect: 'follow'
      };
    
      const result = await fetch(url, requestOptions)
      const data = await result.json();
      console.log(data)
      return data
}

async function hour_getByDay(day){
    url = hour_getByDay_Url + day
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      };
    
      const result = await fetch(url, requestOptions)
      const data = await result.json();
      console.log(data)
      return data
}