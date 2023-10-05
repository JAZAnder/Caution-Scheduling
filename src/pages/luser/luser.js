function checkLogin(){
    key = getCookie("key")
    if(!key){
        window.location.replace("./login");
    }

    
}

