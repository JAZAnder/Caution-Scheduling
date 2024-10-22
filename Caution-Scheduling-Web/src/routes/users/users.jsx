import { Outlet, Link } from "react-router-dom";
import Background from "../../background";
import "./users.css"

const users = () => {
  return (
    <>
      <Background />
      <div className="users-body">
        <h1>Hello, this is a page</h1>
      </div>
    </>
  );
};

export default users;
