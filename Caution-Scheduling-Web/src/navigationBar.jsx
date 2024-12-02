import { useContext } from 'react';
import { Link, useNavigate } from 'react-router-dom';
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

  const handleLogout = async (e) => {
    e.preventDefault();
    await logout(); 
    navigate('/'); 
  };

  const getCenterTextClass = () => {
    if (!user) return 'cs-center-text-default';
    switch (user.role) {
      case 'Administrator':
        return 'cs-center-text-admin';
      case 'Supervisor':
        return 'cs-center-text-supervisor';
      case 'Tutor':
        return 'cs-center-text-tutor';
      case 'Student':
        return 'cs-center-text-student';
      default:
        return 'cs-center-text-default';
    }
  };

  return (
    <div className="cs-navigation-bar">
      {!isMobile ? (
        <>
          <header className="cs-topheader">
            <a href="/" className="cs-header-logo-link">
              <img
                src={schedulingLogo}
                alt="Caution Scheduling Logo"
                className="cs-header-logo"
              />
            </a>
            <span className={getCenterTextClass()}>Caution Scheduling</span>
            <nav>
              <ul className="cs-nav-list">
                <li>
                  <Link to={'/'}>Home</Link>
                </li>
                {user ? (
                  <>
                    <li>
                      <Link to={'/MyProfile'}>My Profile</Link>
                    </li>
                    <li>
                      <a href="#" onClick={handleLogout}>Log Out</a>
                    </li>
                  </>
                ) : (
                  <li>
                    <Link to={'/login'}>Login</Link>
                  </li>
                )}
              </ul>
            </nav>
          </header>

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
                  <Link to={"/aboutus"}>About the Team</Link>
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
                {user ? (
                  <>
                    <LinkContainer to="/MyProfile">
                      <Nav.Link className="cs-nav-link">My Profile</Nav.Link>
                    </LinkContainer>
                    <Nav.Link className="cs-nav-link" onClick={handleLogout} href="#">
                      Log Out
                    </Nav.Link>
                  </>
                ) : (
                  <LinkContainer to="/login">
                    <Nav.Link className="cs-nav-link">Login</Nav.Link>
                  </LinkContainer>
                )}
                <LinkContainer to="/labschedule">
                  <Nav.Link className="cs-nav-link">Lab Schedule</Nav.Link>
                </LinkContainer>
                <LinkContainer to="/schedulemeeting">
                  <Nav.Link className="cs-nav-link">Schedule a Meeting</Nav.Link>
                </LinkContainer>
                <LinkContainer to="/aboutus">
                  <Nav.Link className="cs-nav-link">About the Team</Nav.Link>
                </LinkContainer>
                {user && user.role === 'Administrator' && (
                  <>
                    <LinkContainer to="/meetings">
                      <Nav.Link className="cs-nav-link">Meetings</Nav.Link>
                    </LinkContainer>
                    <LinkContainer to="/admin/adminmeetings">
                      <Nav.Link className="cs-nav-link">Admin Meetings</Nav.Link>
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
