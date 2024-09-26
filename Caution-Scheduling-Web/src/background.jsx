import { Outlet } from "react-router-dom";
import { Link } from "react-router-dom";
import BackgroundImage from "./assets/leavesbackground.jpg";

export default function Background() {
  return (
    <>
      <img
        src={BackgroundImage}
        alt="Placeholder"
        className="lab-schedule-image2"
      />
    </>
  );
}
