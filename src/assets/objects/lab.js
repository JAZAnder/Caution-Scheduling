const lab_getall_Url = baseUrl+"/api/labs"
const lab_create_Url = baseUrl+"/api/lab"

function lab(lab){
    this.Id = lab['id']
    this.name = lab['name']
    this.labLocation = lab['location']
}

async function lab_getall(){

    var requestOptions = {
        method: 'GET',
        redirect: 'follow'
      }; 

      const result = await fetch(lab_getall_Url, requestOptions)
      const data = await result.json();
      console.log(data)
      return data
}

async function lab_create(name, lablocation){
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/x-www-form-urlencoded");
    var urlencoded = new URLSearchParams();
    urlencoded.append("name", name);
    urlencoded.append("location", lablocation);
    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: urlencoded,
        redirect: 'follow'
    };
    const result = await fetch(lab_create_Url, requestOptions)
      const data = await result.json();
      console.log(data)
      return data
}