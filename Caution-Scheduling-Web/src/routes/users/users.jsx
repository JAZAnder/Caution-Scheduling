import { Outlet, Link } from "react-router-dom";
import Background from "../../background";
import './users.css'
import useFetch from "use-http";
import { useState } from "react";

function manageUsers() {
  const [users, setUsers] = useState(null)
  const {
    data: usersInfo,
    loading,
    error,
} = useFetch(
    "/api/lusers",
    {
        method: "get",
    },
    []
);


  return (
    <>
      <Background />
      <button>Add New User</button>
      <div id="filterOnBar">

      </div>
      <div id="userNameTable">
        <table className="table-with-bordered">
          <tr>
            <th> User Id </th>
            <th> Username </th>
            <th> First Name </th>
            <th> Last Name </th>
            <th> Email </th>
            <th> Role </th>
            <th> Details </th>
          </tr>

{

  Object.keys(usersInfo).map(user => (
    <tr>
      <td> {usersInfo[user].userId} </td>
    </tr>
  )

  )
}

        </table>
      </div>
    </>
  );
};

export default manageUsers;
