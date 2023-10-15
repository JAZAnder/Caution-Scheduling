const lab_getall_Url = baseUrl+"/api/labs"

function lab(lab){
    this.Id = lab['id']
    this.name = lab['name']
    this.location = ['location']
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