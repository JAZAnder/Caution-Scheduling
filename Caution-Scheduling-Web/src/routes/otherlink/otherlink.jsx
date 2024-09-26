import { Outlet, Link } from "react-router-dom";
import Background from "../../background";

const otherLink = () => {
  return (
    <>
      <Background />
      <h1 style={{ color: 'white' }}> Hello this is a other link page</h1>
    </>
  );
};

export default otherLink;
