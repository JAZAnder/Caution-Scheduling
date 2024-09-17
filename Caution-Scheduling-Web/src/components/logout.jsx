import {GoogleLogout} from '@react-oauth/google'

const clientId = "825468007612-o1e2kp9d6dedh7l6c2mgem4bqf2fjnpn.apps.googleusercontent.com"


const onSuccess = () => {
    console.log("Logged out")
}

function Logout(){
    return(
        <div id='signInButton'>
            <GoogleLogout 
                clientId={clientId}
                buttonText="Logout"
                onSuccess={onSuccess}
            />
        </div>
    )
}
export default Logout;
