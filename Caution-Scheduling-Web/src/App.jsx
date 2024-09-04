import { useState } from 'react';
import reactLogo from './assets/react.svg';
import viteLogo from '/vite.svg';
import './App.css';

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <header className="topheader">
        <a href="/" className="header-logo-link">
          <img
            src="https://www.southeastern.edu/wp-content/uploads/2023/11/southeastern-logo.svg"
            alt="Southeastern Louisiana University Logo"
            className="header-logo"
          />
        </a>
        <span className="center-text">Caution Scheduling</span>
        <nav>
          <ul className="nav-list">
            <li><a href="#home">Home</a></li>
            <li><a href="#employee">Employee Login</a></li>
            <li><a href="#otherlink">Other Link</a></li>
          </ul>
        </nav>
      </header>

      <header className="bottomheader">
        <nav>
          <ul className="nav-list">
            <li><a href="#schedule">Lab Schedule</a></li>
            <li><a href="#meeting">Schedule a Meeting</a></li>
            <li><a href="#intolab">Sign into Lab</a></li>
            <li><a href="#virtual">Join Virtually</a></li>
          </ul>
        </nav>
      </header>

      <main id="root">
        <div>
          <a href="https://vitejs.dev" target="_blank">
            <img src={viteLogo} className="logo" alt="Vite logo" />
          </a>
          <a href="https://react.dev" target="_blank">
            <img src={reactLogo} className="logo react" alt="React logo" />
          </a>
        </div>
        <div className="button-container">
          <a href="#register" className="cta-button">
            Register
          </a>
          <a href="#login" className="cta-button">
            Login
          </a>
        </div>
        <h1>Vite + React</h1>
        <div className="card">
          <button onClick={() => setCount((count) => count + 1)}>
            count is {count}
          </button>
          <p>
            Edit <code>src/App.jsx</code> and save to test HMR
          </p>
        </div>
        <p className="read-the-docs">
          Click on the Vite and React logos to learn more
        </p>
      </main>
      <footer className="footer">
        <p>This Project is available for download on <a href="https://github.com/JAZAnder/Caution-Scheduling">Github</a></p>
      </footer>
    </>
  );
}

export default App;
