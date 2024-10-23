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
      <main className="cs-root">
        <video autoPlay loop muted className="cs-root video">
          <source src={labVideo} type="video/mp4" />
        </video>
        <div className="cs-button-container">
          {user ? (
            <Link to="" className="cs-cta-button">
              You are logged in
            </Link>
          ) : (
            <Link to="/login" className="cs-cta-button">
              Login
            </Link>
          )}
        </div>
      </main>

      <footer className="cs-footer">
        <p>
          This Project is available for download on{" "}
          <a href="https://github.com/JAZAnder/Caution-Scheduling">Github</a>
        </p>
      </footer>
    </>
  );
}

export default Home;
