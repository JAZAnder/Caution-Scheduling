import { useState } from "react";
import useFetch from "use-http";
import { useNavigate } from "react-router-dom";
import { Button, Form } from "react-bootstrap";
import SignInWithGoogleButton from "../../components/SignInWithGoogleButton";
import './login.css';

export default function Login() {
  const navigate = useNavigate();
  const [userName, setUserName] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const { post, loading } = useFetch("/api/luser/login", {
    method: "post",
    headers: { "Content-Type": "application/x-www-form-urlencoded" },

    onNewData: (_, x) => {
      if ("userName" in x) {
        console.log("Logged in as: ");
        console.log(x);
        navigate("/");
      } else {
        // setError(x);
      }
    },
  });

  return (
    <div className="container">
      <div className="forms-container">
        <div className="form-control signin-form">
          <form className="text-center" onSubmit={handleSubmit}>
            <h2>Sign In</h2>
            <Form.Group>
              <Form.Control
                id="email"
                value={userName}
                onChange={(e) => setUserName(e.target.value)}
                type="text"
                autoComplete="email"
                placeholder="Username"
                required
              />
            </Form.Group>
            <Form.Group>
              <Form.Control
                id="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                type="password"
                autoComplete="password"
                placeholder="Password"
                required
              />
            </Form.Group>
            <Button variant="secondary" type="submit" disabled={loading}>
              {loading ? "Checking Login..." : "Sign In"}
            </Button>
            <div>{error ? <span>{error}</span> : null}</div>
            <span>or sign in with</span>
            <div className="socials">
              <SignInWithGoogleButton />
            </div>
          </form>
        </div>
      </div>
      <div className="intros-container">
        <div className="intro-control signin-intro">
          <div className="intro-control__inner">
            <h2>Welcome back!</h2>
            <p>
              Welcome to Caution Scheduling Tutoring! Where all your tutoring
              needs can be found here!
            </p>
          </div>
        </div>
      </div>
    </div>
  );

  async function handleSubmit(event) {
    event.preventDefault();
    if (loading) return;

    post({
      userName: userName,
      password: password,
    });
  }
}
