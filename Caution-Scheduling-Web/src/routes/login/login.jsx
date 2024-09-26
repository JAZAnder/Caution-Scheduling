import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button, Form } from "react-bootstrap";
import SignInWithGoogleButton from "../../components/SignInWithGoogleButton";
import './login.css';

const user_login = async (username, password) => {
  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

  var urlencoded = new URLSearchParams();
  urlencoded.append("userName", username);
  urlencoded.append("password", password);

  var requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: urlencoded,
    redirect: 'follow'
  };

  try {
    const result = await fetch("/api/luser/login", requestOptions);
    const data = await result.json();
    return data;
  } catch (error) {
    console.error("Login failed:", error);
    throw new Error("Login failed. Please try again.");
  }
};

export default function Login() {
  const navigate = useNavigate();
  const [userName, setUserName] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setLoading(true);
    setError("");

    try {
      const data = await user_login(userName, password);
      if (data.userName) {
        console.log("Logged in as: ", data);
        navigate("/");
      } else {
        setError("Invalid login credentials. Please try again.");
      }
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

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
}
