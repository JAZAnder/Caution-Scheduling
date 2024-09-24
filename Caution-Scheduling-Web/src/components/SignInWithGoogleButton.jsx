import { GoogleOAuthProvider, GoogleLogin } from "@react-oauth/google";

const clientId = "825468007612-o1e2kp9d6dedh7l6c2mgem4bqf2fjnpn.apps.googleusercontent.com"

const onSuccess = (res) => {
    console.log("Success! Current User: ", res)



}

const onFailure = (res) => {
    console.log("Login Failed", res)
}



function SignInWithGoogleButton() {
    return (
        <>
            <GoogleOAuthProvider clientId="825468007612-o1e2kp9d6dedh7l6c2mgem4bqf2fjnpn.apps.googleusercontent.com">
                <GoogleLogin
                    buttonText="Login"
                    onSuccess={onSuccess}
                    onFailure={onFailure}
                    cookiePolicy={'single_host_origin'}
                    isSignedIn={true}
                />
            </GoogleOAuthProvider>

        </>
    )
}



export default SignInWithGoogleButton;