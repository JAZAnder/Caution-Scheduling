import { Outlet, Link } from "react-router-dom";
import Background from "../../background";
import "./my-meetings.css";

const mymeetings = () => {
  return (
    <>
      <Background />

      <div className="mymeetings-body">
        <h1>Hello, this is a page</h1>
      </div>
    </>
  );
};

export default mymeetings;
