import LoginButton from "../../components/login"
import { useEffect } from "react";
import { GoogleOAuthProvider } from "@react-oauth/google";

const clientId = "825468007612-o1e2kp9d6dedh7l6c2mgem4bqf2fjnpn.apps.googleusercontent.com"


function LoginTestButtons(){
    return(
        <>
            <GoogleOAuthProvider clientId="825468007612-o1e2kp9d6dedh7l6c2mgem4bqf2fjnpn.apps.googleusercontent.com">

                <LoginButton />

            </GoogleOAuthProvider>
        
        
        </>
    )
}

export default  LoginTestButtons;