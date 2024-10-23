import { Outlet, Link } from "react-router-dom";
import Background from "../../background";
import "./labs.css"

const labs = () => {
  return (
    <>
      <Background />
      <div className="labs-body">
        <h1>Hello, this is a page</h1>
      </div>
    </>
  );
};

export default labs;
