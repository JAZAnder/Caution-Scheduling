async function login(){
    username = document.getElementById("username").value
    password = document.getElementById("password").value
    data = await user_login(username,password)
    if(data['error']){
        document.getElementById("password").value = ""
        document.getElementById("error").innerHTML = data['error']
        console.log("Login Fail")
    }else{
        document.getElementById("username").value = ""
        document.getElementById("password").value = ""
        console.log("Logged In")
        window.location.assign("../")
    }
    
}
