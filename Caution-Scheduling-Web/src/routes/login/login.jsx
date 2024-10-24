import { useState, useContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../../context/AuthContext';
import SignInWithGoogleButton from '../../components/SignInWithGoogleButton';
import './login.css';
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
      <div className="login-page-container">
        <div className="login-page-box">
          <div className="login-page-form">
            <h2>Sign In</h2>
            <form onSubmit={handleSubmit}>
              <input
                id="username"
                value={userName}
                onChange={(e) => setUserName(e.target.value)}
                type="text"
                autoComplete="username"
                placeholder="Username"
                required
              />
              <input
                id="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                type="password"
                autoComplete="current-password"
                placeholder="Password"
                required
              />
              <button type="submit" disabled={loading}>
                {loading ? 'Checking Login...' : 'Sign In'}
              </button>
              {error && <div className="login-page-error-message">{error}</div>}
              <span>or sign in with</span>
              <SignInWithGoogleButton />
            </form>
          </div>
          <div className="login-page-intro">
            <h2>Welcome back!</h2>
            <p>Welcome to Caution Scheduling Tutoring! Where all your tutoring needs can be found here!</p>
          </div>
        </div>
      </div>
    </>
  );
}
