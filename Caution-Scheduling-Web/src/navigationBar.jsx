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
      {!isMobile ? ( // Render custom CSS for desktop
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
                  <li>
                    <Link to={'/login'} onClick={logout}>
                      Log Out
                    </Link>
                  </li>
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
                  style={{ height: '50px' }} // You can adjust this style if needed
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
              </Nav>
            </Navbar.Collapse>
          </Container>
        </Navbar>
      )}
    </>
  );
}
