const labhour_create_Url = baseUrl+"/api/lab/timeslot/"
const labhour_getall_Url = baseUrl+"/api/lab/timeslots"


function labHour(labhour){
    this.Id = labhour['id']
    this.labId = labhour['labId']
    this.hourId = labhour['hourId']
    this.userHourId = labhour['userHourId']

}

async function labHour_getall(){
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      }; 

      const result = await fetch(labhour_getall_Url, requestOptions)
      const data = await result.json();
      console.log(data)
      return data
}

async function labHour_Create(labId, hourId, tutorName){
    var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

  var urlencoded = new URLSearchParams();
  urlencoded.append("hourId", hourId);
  urlencoded.append("TutorName", tutorName);
  var requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: urlencoded,
    redirect: 'follow'
};

const result = await fetch(labhour_create_Url+labId, requestOptions)
const data = await result.json();
console.log(data)
return data
}