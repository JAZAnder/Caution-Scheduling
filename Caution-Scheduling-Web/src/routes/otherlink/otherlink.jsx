import { Outlet, Link } from "react-router-dom";
import Background from "../../background";
import './otherlink.css';


const otherLink = () => {
  return (
    <>
      <Background />
      <div className="otherlink-body">
        <h1>Hello, this is a page</h1>
      </div>
    </>
  );
};

export default otherLink;