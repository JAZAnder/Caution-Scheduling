import { useContext } from 'react';
import { Link , useNavigate } from 'react-router-dom';
import { AuthContext } from './context/AuthContext'; 
import schedulingLogo from './assets/CautionSchedulingLogoUpdate.png';
import useMediaQuery from './context/useMediaQuery';
import { Navbar, Nav, Container } from 'react-bootstrap';
import { LinkContainer } from 'react-router-bootstrap';
import "./App.css"; 

export default function NavigationBar() {
  const { user, logout } = useContext(AuthContext);
  const isMobile = useMediaQuery('(max-width: 900px)'); 
  const navigate = useNavigate();

  const handleLogout = (e) => {
    e.preventDefault();
    logout(); 
    navigate('/'); 
    window.location.reload(); 
  };

  return (
    <div className="cs-navigation-bar">
      {!isMobile ? ( 
        <>
          {/* Desktop Header */}
          <header className="cs-topheader">
            <a href="/" className="cs-header-logo-link">
              <img
                src={schedulingLogo}
                alt="Caution Scheduling Logo"
                className="cs-header-logo"
              />
            </a>
            <span className="cs-center-text">Caution Scheduling</span>
            <nav>
              <ul className="cs-nav-list">
                <li>
                  <Link to={'/'}>Home</Link>
                </li>
                <li>
                  <Link to={'/otherlink'}>Other Link</Link>
                </li>
                {user ? (
                  <li>
                    <Link to={'/login'} onClick={handleLogout}>Log Out</Link>
                  </li>
                ) : (
                  <li>
                    <Link to={'/login'}>Employee Login</Link>
                  </li>
                )}
              </ul>
            </nav>
          </header>

          {/* Desktop Bottom Header */}
          <header className="cs-bottomheader">
            <nav>
              <ul className="cs-nav-list">
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
                      <Link to={'/admin/adminmeetings'}>Admin Meetings</Link>
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
                    <li>
                      <Link to={'/timeslotmanagement'}>Timeslot Management</Link>
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
                      <Link to={'/meetings'}>My Meetings</Link>
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
                    <Link to={'/meetings'}>My Meetings</Link>
                  </li>
                )}
              </ul>
            </nav> 
          </header>
        </>
      ) : ( 
        // Mobile Navbar using React Bootstrap
        <Navbar className="cs-navbar" expand="lg" style={{ backgroundColor: '#1a5632' }}>
          <Container fluid>
            <LinkContainer to="/">
              <Navbar.Brand>
                <img
                  src={schedulingLogo}
                  alt="Caution Scheduling Logo"
                  className="cs-header-logo"
                  style={{ height: '50px' }}
                />
              </Navbar.Brand>
            </LinkContainer>
            <Navbar.Toggle aria-controls="navbarNav" />
            <Navbar.Collapse id="navbarNav">
              <Nav className="me-auto">
                <LinkContainer to="/">
                  <Nav.Link className="cs-nav-link">Home</Nav.Link>
                </LinkContainer>
                <LinkContainer to="/otherlink">
                  <Nav.Link className="cs-nav-link">Other Link</Nav.Link>
                </LinkContainer>
                {user ? (
                  <LinkContainer to="/login">
                    <Nav.Link className="cs-nav-link" onClick={handleLogout}>Log Out</Nav.Link>
                  </LinkContainer>
                ) : (
                  <LinkContainer to="/login">
                    <Nav.Link className="cs-nav-link">Employee Login</Nav.Link>
                  </LinkContainer>
                )}
                <LinkContainer to="/labschedule">
                  <Nav.Link className="cs-nav-link">Lab Schedule</Nav.Link>
                </LinkContainer>
                <LinkContainer to="/schedulemeeting">
                  <Nav.Link className="cs-nav-link">Schedule a Meeting</Nav.Link>
                </LinkContainer>
                <LinkContainer to="/signinlab">
                  <Nav.Link className="cs-nav-link">Sign into Lab</Nav.Link>
                </LinkContainer>
                <LinkContainer to="/aboutus">
                  <Nav.Link className="cs-nav-link">About Us</Nav.Link>
                </LinkContainer>
                {user && user.role === 'Administrator' && (
                  <>
                    <LinkContainer to="/meetings">
                      <Nav.Link className="cs-nav-link">Meetings</Nav.Link>
                    </LinkContainer>
                    <LinkContainer to="/users">
                      <Nav.Link className="cs-nav-link">Users</Nav.Link>
                    </LinkContainer>
                    <LinkContainer to="/labs">
                      <Nav.Link className="cs-nav-link">Labs</Nav.Link>
                    </LinkContainer>
                    <LinkContainer to="/timeslots">
                      <Nav.Link className="cs-nav-link">Timeslots</Nav.Link>
                    </LinkContainer>
                    <LinkContainer to="/timeslotmanagement">
                      <Nav.Link className="cs-nav-link">Timeslot Management</Nav.Link>
                    </LinkContainer>
                  </>
                )}
                {user && user.role === 'Supervisor' && (
                  <LinkContainer to="/meetings">
                    <Nav.Link className="cs-nav-link">Meetings</Nav.Link>
                  </LinkContainer>
                )}
                {user && user.role === 'Tutor' && (
                  <>
                    <LinkContainer to="/meetings">
                      <Nav.Link className="cs-nav-link">My Meetings</Nav.Link>
                    </LinkContainer>
                    <LinkContainer to="/my-availability">
                      <Nav.Link className="cs-nav-link">My Availability</Nav.Link>
                    </LinkContainer>
                    <LinkContainer to="/users">
                      <Nav.Link className="cs-nav-link">Users</Nav.Link>
                    </LinkContainer>
                  </>
                )}
                {user && user.role === 'Student' && (
                  <LinkContainer to="/meetings">
                    <Nav.Link className="cs-nav-link">My Meetings</Nav.Link>
                  </LinkContainer>
                )}
              </Nav>
            </Navbar.Collapse>
          </Container>
        </Navbar>
      )}
    </div>
  );
}
