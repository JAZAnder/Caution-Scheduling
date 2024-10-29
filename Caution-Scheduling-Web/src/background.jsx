import { Outlet } from "react-router-dom";
import { Link } from "react-router-dom";
import BackgroundImage from "./assets/backgroundimage.webp";

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
