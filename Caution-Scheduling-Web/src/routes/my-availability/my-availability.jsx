import { Outlet, Link } from "react-router-dom";
import Background from "../../background";
import "./my-availability.css"

const myavailability = () => {
  return (
    <>
      <Background />
      <div className="myavailability-body">
        <h1>Hello, this is a page</h1>
      </div>
    </>
  );
};

export default myavailability;
