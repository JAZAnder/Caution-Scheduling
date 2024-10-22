import { Outlet, Link } from "react-router-dom";
import Background from "../../background";
import './schedulemeeting.css';

const scheduleMeeting = () => {
  return (
    <>
      <Background />
      <div className="schedulemeeting-body">
        <h1>Hello, this is a Schedule Meeting page</h1>
      </div>
    </>
  );
};

export default scheduleMeeting;