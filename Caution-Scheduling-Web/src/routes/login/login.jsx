import { useState, useContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Button, Form } from 'react-bootstrap';
import { AuthContext } from '../../context/AuthContext';
import SignInWithGoogleButton from '../../components/SignInWithGoogleButton';
import './login.css';  // Ensure this file is present
import Background from "../../background";

const authenticateUser = async (username, password) => {
  const myHeaders = new Headers();
  myHeaders.append('Content-Type', 'application/x-www-form-urlencoded');

  const urlencoded = new URLSearchParams();
  urlencoded.append('userName', username);
  urlencoded.append('password', password);

  const requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: urlencoded,
    redirect: 'follow',
  };

  try {
    const response = await fetch('/api/luser/login', requestOptions);

    if (!response.ok) {
      throw new Error('Invalid login credentials.');
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Login failed:', error);
    throw new Error('Login failed. Please try again.');
  }
};

export default function Login() {
  const navigate = useNavigate();
  const { user, login } = useContext(AuthContext);
  const [userName, setUserName] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (user) {
      navigate('/');
    }
  }, [user, navigate]);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setLoading(true);
    setError('');

    try {
      const userData = await authenticateUser(userName, password);
      login(userData);
      navigate('/');
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <Background />
      <div className="login-page">
        <div className="container">
          <div className="forms-container">
            <div className="form-control signin-form">
              {user ? (
                <div className="logged-in-message">
                  <h2>You are logged in</h2>
                  <p>Hi {user.userName}</p>
                </div>
              ) : (
                <form className="text-center" onSubmit={handleSubmit}>
                  <h2>Sign In</h2>
                  <Form.Group>
                    <Form.Control
                      className="form-input"  // Add this class for custom styling
                      id="username"
                      value={userName}
                      onChange={(e) => setUserName(e.target.value)}
                      type="text"
                      autoComplete="username"
                      placeholder="Username"
                      required
                    />
                  </Form.Group>
                  <Form.Group>
                    <Form.Control
                      className="form-input"  // Add this class for custom styling
                      id="password"
                      value={password}
                      onChange={(e) => setPassword(e.target.value)}
                      type="password"
                      autoComplete="current-password"
                      placeholder="Password"
                      required
                    />
                  </Form.Group>
                  <Button className="submit-button" type="submit" disabled={loading}>
                    {loading ? 'Checking Login...' : 'Sign In'}
                  </Button>
                  {error && <div className="error-message">{error}</div>}

                  <span>or sign in with</span>
                  <div className="socials">
                    <SignInWithGoogleButton />
                  </div>
                </form>
              )}
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
      </div>
    </>
  );
}
