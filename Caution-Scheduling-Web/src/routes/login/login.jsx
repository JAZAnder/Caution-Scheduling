import { useState } from "react"
import useFetch from 'use-http'
import { useNavigate } from "react-router-dom";
import { Button, Form } from "react-bootstrap";
import SignInWithGoogleButton from "../../components/SignInWithGoogleButton"

export default function Login() {
    const navigate = useNavigate();
    const [userName, setUserName] = useState("")
    const [password, setPassword] = useState("")
    const [error, setError] = useState("");


    const { post, loading } = useFetch("/api/luser/login", {
        method: "post",
        headers:{"Content-Type":  'application/x-www-form-urlencoded' },
        
        onNewData: (_, x) => {
            if ('userName' in x) {

                console.log("Logged in as: ");
                console.log(x);
                navigate("/");

            } else {
               // setError(x);
            }
        },
    });



    return (
        <>
        <div>
        <form className="text-center" onSubmit={handleSubmit}>
                <Form.Group>
                    <Form.Label><strong>Email</strong></Form.Label>
                    <Form.Control
                        id="email"
                        value={userName}
                        onChange={(e) => setUserName(e.target.value)}
                        type="text"
                        autoComplete="email"
                        placeholder="Email"
                        required
                    />
                </Form.Group>
                <Form.Group>
                    <Form.Label><strong>Password</strong></Form.Label>
                    <Form.Control
                        id="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        type="password"
                        autoComplete="password"
                        placeholder='Password'
                        required
                    />
                </Form.Group>
                
                <div style={{ height: "5%" }}>
                {loading ? "Checking Login..." : null}
                </div>
                {error ? error : null}
                <Button
                    variant="secondary"
                    className="background-1"
                    type="submit" disabled={loading}

                >Sign In</Button>


            </form>

            <SignInWithGoogleButton/>
        </div>
            

        </>

    );

    async function handleSubmit(event) {

        event.preventDefault();

        if (loading) {
            return;
        }

        post({
            userName: userName,
            password: password,
        });

    }

}