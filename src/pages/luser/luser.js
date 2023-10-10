async function checkLogin(){
    key = getCookie("key")
    if(!key){
        window.location.replace("./login");
    }
    
    const user = new luser(await user_whoami())
    document.getElementById('username').innerHTML = user.firstName
}

async function logout(){
    data = await user_logout()
    if(data['error']){
        console.log("Logout Error :"+ data['error'])
    }else{
        console.log("Logged Out")
        window.location.reload()
    }
}

