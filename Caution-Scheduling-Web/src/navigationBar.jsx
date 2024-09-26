// src/navigationBar.jsx

import { useContext } from 'react';
import { Link } from 'react-router-dom';
import { AuthContext } from './context/AuthContext'; 
import schedulingLogo from './assets/CautionSchedulingLogoUpdate.png';

export default function NavigationBar() {
  const { user, logout } = useContext(AuthContext);

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
              <Link to={'/'}>Home</Link>
            </li>
            <li>
              <Link to={'/otherlink'}>Other Link</Link>
            </li>
            {user ? (
              <>
                <li>
                  <Link to = {'/login'} onClick={logout}>
                    Log Out
                  </Link>
                </li>
              </>
            ) : (
              <li>
                <Link to={'/login'}>Employee Login</Link>
              </li>
            )}
          </ul>
        </nav>
      </header>

            <header className="bottomheader">
                <nav>
                    <ul className="nav-list">
                        <li>
                            <Link to={"/labschedule"}>Lab Schedule</Link>
                        </li>
                        <li>
                            <Link to={"/schedulemeeting"}>Schedule a Meeting</Link>
                        </li>
                        <li>
                            <Link to={"/signinlab"}>Sign into Lab</Link>
                        </li>
                        <li>
                            <Link to={"/joinvirtually"}>Join Virtually</Link>
                        </li>
                        <li>
                            <Link to={"/aboutus"}>About Us</Link>
                        </li>
                    </ul>
                </nav>
            </header>
        </>
    )
}
