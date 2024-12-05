import { Outlet, Link } from "react-router-dom";
import Background from "../../background";
import './users.css';
import useFetch from "use-http";
import React, { useEffect, useState } from "react";
import UserDetailsButton from "./details/details";
import NewUserButton from "./create/create";
import { Modal, Button, Form } from "react-bootstrap";

function ManageUsers() {
  const [loading, setLoading] = useState(false);
  const [userData, setUserData] = useState();
  const [filtering, setFiltering] = useState(false);
  const [userName, setUserName] = useState('');
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [role, setRole] = useState('');
  const [debounce, setDebounce] = useState(true);

  useEffect(() => {
    setDebounce(!debounce);
  }, []);

  useEffect(() => {
    const getData = setTimeout(() => {
      setDebounce(!debounce);
      console.log("Should Be Refresh");
    }, 100);
    return () => clearTimeout(getData);
  }, [userName, firstName, lastName, email, role]);

  const resetSearch = async (event) => {
    setUserName('');
    setFirstName('');
    setLastName('');
    setEmail('');
    setRole('');
    
    console.log("Should Be Refresh");
    setDebounce(!debounce);
  };

  return (
    <>
      <Background />
      <div className="manage-users-container1">
        <div className="manage-users-page">
          <div id="filterOnBar">
            <form>
              <input
                id="username"
                value={userName}
                onChange={(e) => setUserName(e.target.value)}
                type="text"
                autoComplete="username"
                placeholder="Username"
              />
              <input
                id="FirstName"
                value={firstName}
                onChange={(e) => setFirstName(e.target.value)}
                type="text"
                placeholder="First Name"
              />
              <input
                id="LastName"
                value={lastName}
                onChange={(e) => setLastName(e.target.value)}
                type="text"
                placeholder="Last Name"
              />
              <input
                id="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                type="text"
                placeholder="Email"
              />
              <select
                name="role"
                id="role"
                onChange={(e) => setRole(e.target.value)}
              >
                <option value=""> role </option>
                <option value="1"> Student</option>
                <option value="2"> Tutors</option>
                <option value="3"> Supervisors</option>
                <option value="4"> Administrators</option>
              </select>

              <button type="button" disabled={loading} onClick={resetSearch}>
                {loading ? 'Waiting' : 'Reset Search'}
              </button>
            </form>
            <NewUserButton />
          </div>
          
          <ListFilteredUser
            FLuserName={userName}
            FLfirstName={firstName}
            FLlastName={lastName}
            FLemail={email}
            FLrole={role}
            debounce={debounce}
          />
        </div>
      </div>
    </>
  );
}

function ListFilteredUser({ FLuserName, FLfirstName, FLlastName, FLemail, FLrole, debounce }) {
  const { data: usersInfo, loading, error } = useFetch(
    `/api/lusers/filter?userName=${FLuserName}&firstName=${FLfirstName}&lastName=${FLlastName}&email=${FLemail}&role=${FLrole}`,
    { method: "get" },
    [debounce]
  );

  const [show, setShow] = useState(false);
  const [selectedUserName, setSelectedUserName] = useState("");
  const [newPassword, setNewPassword] = useState("");

  const handleClose = () => {
    setShow(false);
    setSelectedUserName("");
    setNewPassword("");
  };

  const handleShow = (userName) => {
    setSelectedUserName(userName);
    setShow(true);
  };

  const handleResetPassword = async () => {
    try {
      const response = await fetch('/api/luser/resetpasswd', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({
          UserName: selectedUserName,
          password: newPassword,
        }),
      });
  
      const responseData = await response.json();
      console.log("API Response:", responseData); 
  
      if (response.ok) {
        alert(`Password reset successfully for ${selectedUserName}`);
      } else {
        alert(`Error resetting password: ${responseData.message || JSON.stringify(responseData) || 'Unknown error'}`);
      }
    } catch (error) {
      console.error("Error resetting password:", error);
      alert("Error resetting password. Please try again later.");
    } finally {
      handleClose();
    }
  };
  

  if (loading) {
    return (
      <center>
        <div className="loader"></div>
      </center>
    );
  }

  return (
    <div id="userNameTable">
      <table className="table-with-bordered">
        <thead>
          <tr>
            <th>User Id</th>
            <th>Username</th>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Email</th>
            <th>Role</th>
            <th>Details</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {usersInfo &&
            Object.keys(usersInfo).map((user, i) => (
              <tr key={i}>
                <td>{usersInfo[user].userId}</td>
                <td>{usersInfo[user].userName}</td>
                <td>{usersInfo[user].firstName}</td>
                <td>{usersInfo[user].lastName}</td>
                <td>{usersInfo[user].email}</td>
                <td>{usersInfo[user].role}</td>
                <td>
                  <UserDetailsButton user={usersInfo[user]} />
                </td>
                <td>
                  <Button variant="secondary"
                    onClick={() => handleShow(usersInfo[user].userName)}
                    className="reset-password-button"
                  >
                    Reset Password
                  </Button>
                </td>
              </tr>
            ))}
        </tbody>
      </table>

      <Modal show={show} onHide={handleClose} backdrop="static">
        <Modal.Header closeButton>
          <Modal.Title>Reset Password for {selectedUserName}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>New Password</Form.Label>
            <Form.Control
              type="password"
              value={newPassword}
              onChange={(e) => setNewPassword(e.target.value)}
              placeholder="Enter new password"
            />
          </Form.Group>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Cancel
          </Button>
          <Button
            variant="primary"
            onClick={handleResetPassword}
            disabled={!newPassword}
          >
            Reset Password
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
}

export default ManageUsers;
