import { Outlet } from "react-router-dom";
import schedulingLogo from "./assets/CautionSchedulingLogoUpdate.png";
import { Link } from 'react-router-dom';

export default function NavigationBar() {
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
                            <Link to={"/"}>Home</Link>
                        </li>
                        <li>
                            <Link to={"/login"}>Employee Login</Link>
                        </li>
                        <li>
                            <Link to={"/otherLink"}>Other Link</Link>
                        </li>
                    </ul>
                </nav>
            </header>

            <header className="bottomheader">
                <nav>
                    <ul className="nav-list">
                        <li>
                            <Link to={"/schedule"}>Lab Schedule</Link>
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
<Outlet/>
        </>
    )
}