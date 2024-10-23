import { Outlet, Link } from "react-router-dom";
import Background from "../../background";
import './users.css'
import useFetch from "use-http";
import React, { useEffect } from "react";
import { useState } from "react";
import UserDetailsButton from "./details/details";
import NewUserButton from "./create/create"

function manageUsers() {
  const [loading, setLoading] = useState(false)
  const [userData, setUserData] = useState()
  const [filtering, setFiltering] = useState(false)
  const [userName, setUserName] = useState('')
  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')
  const [email, setEmail] = useState('')
  const [role, setRole] = useState('')
  const [debounce, SetDebounce] = useState(true)

  useEffect(() => {
    SetDebounce(!debounce)
  }, [])





  React.useEffect(() => {
    const getData = setTimeout(() => {
      SetDebounce(!debounce)
    }, 100)
    return () => clearTimeout(getData)
  }, [userName, firstName, lastName, email, role])



  const resetSearch = async (event) => {
    setUserName('')
    setFirstName('')
    setLastName('')
    setEmail('')
    setRole('')
    SetDebounce(!debounce)
  }


  return (
    <>
      <div style={{ minHeight: "150px" }}> Black Space?</div>
      <NewUserButton />
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
          <select name="role" id="role" onChange={(e) => setRole(e.target.value)}>
            <option value=""> role </option>
            <option value="1"> Student</option>
            <option value="2" > Tutors</option>
            <option value="3" > Supervisors</option>
            <option value="4" > Administrators</option>
          </select>


          <button type="button" disabled={loading} onClick={resetSearch}>
            {loading ? 'Waiting' : 'Reset Search'}
          </button>
        </form>

      </div>


      <ListFilteredUser FLuserName={userName} FLfirstName={firstName} FLlastName={lastName} FLemail={email} FLrole={role} debounce={debounce} />



    </>
  )
};



function ListFilteredUser({ FLuserName, FLfirstName, FLlastName, FLemail, FLrole, debounce }) {

  const {
    data: usersInfo,
    loading,
    error,
  } = useFetch(
    "/api/lusers/filter?userName=" + FLuserName + "&firstName=" + FLfirstName + "&lastName" + FLlastName + "&email=" + FLemail + "&role=" + FLrole,
    {
      method: "get",
    },
    [debounce]
  );

  if (loading) {
    return (
      <>
        <center> <div className="loader"></div></center>
      </>
    )
  }


  return (
    <>
      <div id="userNameTable">
        <table className="table-with-bordered">
          <thead>
            <tr>
              <th> User Id </th>
              <th> Username </th>
              <th> First Name </th>
              <th> Last Name </th>
              <th> Email </th>
              <th> Role </th>
              <th> Details </th>
            </tr>
          </thead>

          <tbody>
            {



              Object.keys(usersInfo).map((user, i) => (
                <tr key={i}>
                  <td> {usersInfo[user].userId} </td>
                  <td> {usersInfo[user].userName} </td>
                  <td> {usersInfo[user].firstName} </td>
                  <td> {usersInfo[user].lastName} </td>
                  <td> {usersInfo[user].email} </td>
                  <td> {usersInfo[user].role} </td>
                  <td><UserDetailsButton user={usersInfo[user]} /></td>
                </tr>
              )

              )

            }
          </tbody>


        </table>
      </div>
    </>
  )


}






export default manageUsers;
