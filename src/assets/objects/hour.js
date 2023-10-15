const hour_create_Url = baseUrl+"/api/hour"
const hour_get_Url = baseUrl+"/api/hour/"
const hour_getAll_Url = baseUrl+"/api/hours"

function hour(hour){
    this.id = hour['id']
    this.startTime = hour['startTime']
    this.endTime = hour['endTime']
}

async function hour_create(startTime, EndTime){
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

    var urlencoded = new URLSearchParams();
    urlencoded.append("startTime", startTime);
    urlencoded.append("endTime", EndTime);

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