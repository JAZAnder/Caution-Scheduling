import { useState } from "react";
import schedulingLogo from "./assets/CautionSchedulingLogoUpdate.png";
import labVideo from "./assets/LabVideo.mp4";
import "./App.css";

function App() {
  return (
    <>
      <header className="topheader">
        <a href="/" className="header-logo-link">
          <img
            src={schedulingLogo}
            alt="Caution Scheduling Logo"
            className="header-logo"
          />
        </a>
        <span className="center-text">Caution Scheduling</span>
        <nav>
          <ul className="nav-list">
            <li>
              <a href="/">Home</a>
            </li>
            <li>
              <a href="#employee">Employee Login</a>
            </li>
            <li>
              <a href="#otherlink">Other Link</a>
            </li>
          </ul>
        </nav>
      </header>

      <header className="bottomheader">
        <nav>
          <ul className="nav-list">
            <li>
              <a href="#schedule">Lab Schedule</a>
            </li>
            <li>
              <a href="#meeting">Schedule a Meeting</a>
            </li>
            <li>
              <a href="#intolab">Sign into Lab</a>
            </li>
            <li>
              <a href="#virtual">Join Virtually</a>
            </li>
          </ul>
        </nav>
      </header>

      <main id="root">
        <video autoPlay loop muted>
          <source src={labVideo} type="video/mp4" />
        </video>
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

export default App;
