// src/Home.jsx
import { useContext } from 'react';
import { Link } from 'react-router-dom';
import labVideo from './assets/LabVideo.mp4';
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import { AuthContext } from './context/AuthContext'; 
function Home() {
  const { user } = useContext(AuthContext);

  return (
    <>
      <main id="root">
        <video autoPlay loop muted>
          <source src={labVideo} type="video/mp4" />
        </video>
        <div className="button-container">
          {user ? (
            <Link to="" className="cta-button">
              You are logged in
            </Link>
          ) : (
            <Link to="/login" className="cta-button">
              Login
            </Link>
          )}
        </div>
      </main>

      <footer className="footer">
        <p>
          This Project is available for download on{" "}
          <a href="https://github.com/JAZAnder/Caution-Scheduling">Github</a>
        </p>
      </footer>
    </>
  );
}

export default Home;
