import { GoogleOAuthProvider, GoogleLogin } from "@react-oauth/google";

const clientId = "825468007612-o1e2kp9d6dedh7l6c2mgem4bqf2fjnpn.apps.googleusercontent.com"

const onSuccess = (res) => {
    console.log("Success! Current User: ", res)



}

const onFailure = (res) => {
    console.log("Login Failed", res)
}



function SignInWithGoogleButton({ onSuccess, onFailure }) {
    return (
      <GoogleOAuthProvider clientId={clientId}>
        <GoogleLogin
          onSuccess={onSuccess}
          onError={onFailure}
          useOneTap
        />
      </GoogleOAuthProvider>
    );
  }



export default SignInWithGoogleButton;