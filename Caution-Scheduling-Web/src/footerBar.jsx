import { Outlet } from "react-router-dom";
import { Link } from "react-router-dom";

export default function FooterBar() {
  return (
    <>
      <footer className="cs-footer">
        <p>
          This Project is available for download on{" "}
          <a href="https://github.com/JAZAnder/Caution-Scheduling">Github</a>
        </p>
      </footer>
    </>
  );
}
