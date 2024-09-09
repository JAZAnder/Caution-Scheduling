import { useState } from "react";
import schedulingLogo from "./assets/CautionSchedulingLogoUpdate.png";
import { Link } from 'react-router-dom';
import labVideo from "./assets/LabVideo.mp4";
import 'bootstrap/dist/css/bootstrap.min.css';
import "./App.css";

function Home() {
  return (
    <>
      <main id="root">
        <video autoPlay loop muted>
          <source src={labVideo} type="video/mp4" />
        </video>
        <div className="button-container">
          <a href="#register" className="cta-button">
            Register
          </a>
          <a href="#login" className="cta-button">
            Login
          </a>
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
