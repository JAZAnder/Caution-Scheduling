const meeting_GetAll_Url = baseUrl+"/api/meetings"
const meeting_Create_Url = baseUrl+"/api/meeting"
const meeting_GetMine_Url = baseUrl+"/api/meetings/mine"

function meeting(meeting){
    this.id = meeting['id']
    this.tutorHourId = meeting['tutorHourId']
    this.labId = meeting['labId']
    this.studentName = meeting['studentName']
    this.studentEmail = meeting['studentEmail']
}

async function meeting_create(tutorHourId, labId, studentName, studentEmail){
    var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

  var urlencoded = new URLSearchParams();
  urlencoded.append("userHourId", tutorHourId);
  urlencoded.append("labId", labId);
  urlencoded.append("studentName", studentName);
  urlencoded.append("studentEmail", studentEmail);

  var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: urlencoded,
      redirect: 'follow'
  };

  const result = await fetch(meeting_Create_Url, requestOptions)
  const data = await result.json();
  console.log(data)
  return data
}
async function meeting_GetAll(){
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      };
    
      const result = await fetch(meeting_GetAll_Url, requestOptions)
      const data = await result.json();
      console.log(data)
      return data
}

async function meeting_GetMine(){
    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      };
    
      const result = await fetch(meeting_GetMine_Url, requestOptions)
      const data = await result.json();
      console.log(data)
      return data
}