// src/navigationBar.jsx

import { useContext } from 'react';
import { Link } from 'react-router-dom';
import { AuthContext } from './context/AuthContext'; 
import schedulingLogo from './assets/CautionSchedulingLogoUpdate.png';
import useMediaQuery from './context/useMediaQuery';
import { Navbar, Nav, Container } from 'react-bootstrap';
import { LinkContainer } from 'react-router-bootstrap';
import "./App.css";

export default function NavigationBar() {
  const { user, logout } = useContext(AuthContext);
  const isMobile = useMediaQuery('(max-width: 900px)'); // Check if the screen is mobile

  return (
    <>
      {!isMobile ? ( // Render CSS for Desktop
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
                      <Link to={'/login'} onClick={logout}>
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
                  <Link to={"/aboutus"}>About Us</Link>
                </li>
                {user && user.role === 'Administrator' && (
                  <>
                    <li>
                      <Link to={'/meetings'}>Meetings</Link>
                    </li>
                    <li>
                      <Link to={'/users'}>Users</Link>
                    </li>
                    <li>
                      <Link to={'/labs'}>Labs</Link>
                    </li>
                    <li>
                      <Link to={'/timeslots'}>Timeslots</Link>
                    </li>
                  </>
                )}
                {user && user.role === 'Supervisor' && (
                  <li>
                    <Link to={'/meetings'}>Meetings</Link>
                  </li>
                )}
                {user && user.role === 'Tutor' && (
                  <>
                    <li>
                      <Link to={'/my-meetings'}>My Meetings</Link>
                    </li>
                    <li>
                      <Link to={'/my-availability'}>My Availability</Link>
                    </li>
                    <li>
                      <Link to={'/users'}>Users</Link>
                    </li>
                  </>
                )}
                {user && user.role === 'Student' && (
                  <li>
                    <Link to={'/my-meetings'}>My Meetings</Link>
                  </li>
                )}
              </ul>
            </nav> 
          </header>
        </>
      ) : ( // Render Bootstrap elements for mobile
        <Navbar className="topheader" expand="lg" style={{ backgroundColor: '#1a5632' }}>
          <Container fluid>
            <LinkContainer to="/">
              <Navbar.Brand>
                <img
                  src={schedulingLogo}
                  alt="Caution Scheduling Logo"
                  className="header-logo"
                  style={{ height: '50px' }}
                />
              </Navbar.Brand>
            </LinkContainer>
            <Navbar.Toggle aria-controls="navbarNav" />
            <Navbar.Collapse id="navbarNav">
              <Nav className="me-auto">
                <LinkContainer to="/">
                  <Nav.Link style={{ color: 'white' }}>Home</Nav.Link>
                </LinkContainer>
                <LinkContainer to="/otherlink">
                  <Nav.Link style={{ color: 'white' }}>Other Link</Nav.Link>
                </LinkContainer>
                {user ? (
                  <LinkContainer to="/login">
                    <Nav.Link style={{ color: 'white' }} onClick={logout}>Log Out</Nav.Link>
                  </LinkContainer>
                ) : (
                  <LinkContainer to="/login">
                    <Nav.Link style={{ color: 'white' }}>Employee Login</Nav.Link>
                  </LinkContainer>
                )}
                <LinkContainer to="/labschedule">
                  <Nav.Link style={{ color: 'white' }}>Lab Schedule</Nav.Link>
                </LinkContainer>
                <LinkContainer to="/schedulemeeting">
                  <Nav.Link style={{ color: 'white' }}>Schedule a Meeting</Nav.Link>
                </LinkContainer>
                <LinkContainer to="/signinlab">
                  <Nav.Link style={{ color: 'white' }}>Sign into Lab</Nav.Link>
                </LinkContainer>
                <LinkContainer to="/aboutus">
                  <Nav.Link style={{ color: 'white' }}>About Us</Nav.Link>
                </LinkContainer>
                {user && user.role === 'Administrator' && (
                  <>
                    <LinkContainer to="/meetings">
                      <Nav.Link style={{ color: 'white' }}>Meetings</Nav.Link>
                    </LinkContainer>
                    <LinkContainer to="/users">
                      <Nav.Link style={{ color: 'white' }}>Users</Nav.Link>
                    </LinkContainer>
                    <LinkContainer to="/labs">
                      <Nav.Link style={{ color: 'white' }}>Labs</Nav.Link>
                    </LinkContainer>
                    <LinkContainer to="/timeslots">
                      <Nav.Link style={{ color: 'white' }}>Timeslots</Nav.Link>
                    </LinkContainer>
                  </>
                )}
                {user && user.role === 'Supervisor' && (
                  <LinkContainer to="/meetings">
                    <Nav.Link style={{ color: 'white' }}>Meetings</Nav.Link>
                  </LinkContainer>
                )}
                {user && user.role === 'Tutor' && (
                  <>
                    <LinkContainer to="/my-meetings">
                      <Nav.Link style={{ color: 'white' }}>My Meetings</Nav.Link>
                    </LinkContainer>
                    <LinkContainer to="/my-availability">
                      <Nav.Link style={{ color: 'white' }}>My Availability</Nav.Link>
                    </LinkContainer>
                    <LinkContainer to="/users">
                      <Nav.Link style={{ color: 'white' }}>Users</Nav.Link>
                    </LinkContainer>
                  </>
                )}
                {user && user.role === 'Student' && (
                  <LinkContainer to="/my-meetings">
                    <Nav.Link style={{ color: 'white' }}>My Meetings</Nav.Link>
                  </LinkContainer>
                )}
              </Nav>
            </Navbar.Collapse>
          </Container>
        </Navbar>
      )}
    </>
  );
}
