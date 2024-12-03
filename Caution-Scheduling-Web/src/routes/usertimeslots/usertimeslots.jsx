import { Link, useSearchParams, useNavigate } from "react-router-dom";
import React, { useEffect, useState } from "react";
import useFetch, { CachePolicies } from "use-http";
import { Container, Row, Col, Button, Form, Table } from 'react-bootstrap';
import Background from "../../background";

const userTimeslots = () => {
  
  const [hourId, setHourId] = useState('null')
  const [tutorId, setTutorId] = useState('null')
  const [dayOfWeek, setDayOfWeek] = useState('null')
  const [debounce, setDebounce] = useState(true);
  const navigate = useNavigate();
  const { data: timeslots, timeLoading } = useFetch(
    `/api/timeslots/codes`,
    { method: "get" },
    [debounce]
  );

  const { data: users, usersLoading } = useFetch(
    `/api/lusers`,
    { method: "get" },
    [debounce]
  );

  useEffect(() => {
    setDebounce(!debounce);
  }, []);

  React.useEffect(() => {
    const getData = setTimeout(() => {
      setDebounce(!debounce);
    }, 100);
    return () => clearTimeout(getData);
  }, [dayOfWeek, hourId, tutorId]);

  const Search = async (event) => {
    navigate("/user-timeslots?hourId="+hourId+"&tutorId="+tutorId+"&dayOfWeek="+dayOfWeek)
    
  };

  if (usersLoading || timeLoading) {
    return (<><center><div className="loader"></div></center></>)
  }


  return (
    <>
    <Background />
      <div className="underNavBarSpacer"></div>

      <div className="timeslots-filter">
        <label className="timeslots-filter-label">Filter on:</label>
        <input
          type="text"
          placeholder="Id"
          className="timeslots-filter-input"
        />

        <select
          id='timeCode'
          type="text"
          placeholder="Tutor"
          className="timeslots-filter-input"
        >
          <option value="null"> Pick a Time </option>
          {users &&
            Object.keys(timeslots).map((timeCode, i) => (
              <option value={timeslots[timeCode].timeCode}> {timeslots[timeCode].startTime + " - " + timeslots[timeCode].endTime} </option>
            ))

          }

        </select>


        <select
          id='tutorId'
          onChange={(e) => setTutorId(e.target.value)}
          type="text"
          placeholder="Tutor"
          className="timeslots-filter-input"
        >
          <option value="null"> Pick a Tutor </option>
          {users &&
            Object.keys(users).map((user, i) => (
              <option value={users[user].userId}> {users[user].firstName + " " + users[user].lastName} </option>
            ))

          }

        </select>

        <select
          className="timeslots-filter-input"
          name="dayOfWeek"
          id="dayOfWeek"
          onChange={(e) => setDayOfWeek(e.target.value)}
        >
          <option value=""> Day of Week </option>
          <option value="1"> Monday</option>
          <option value="2"> Tuesday</option>
          <option value="3"> Wednesday</option>
          <option value="4"> Thursday</option>
          <option value="5"> Friday</option>
          <option value="6"> Saturday</option>
          <option value="0"> Sunday</option>

        </select>

<Button variant="primary" onClick={Search}>Search</Button>

      </div>

      <div>
        <FilterUserTimeSlots
          debounce={debounce}
        />
      </div>



    </>
  )
}

function FilterUserTimeSlots({ debounce }) {
  const [params] = useSearchParams();

  const FLhourId = params.get("hourId")
  const FLtutorId = params.get("tutorId")
  const FLdayOfWeek = params.get("dayOfWeek")


  const { data: userTimeslots, loading, error } = useFetch(
    `/api/availability?tutorId=` + FLtutorId + `&hourId=` + FLhourId + `&dayOfWeek=` + FLdayOfWeek,
    { method: "get" },
    [debounce]
  );


  if (loading) {
    return (
      <center>
        <div className="loader"></div>
      </center>
    );
  }

  return (
    <>
      <table>
        <thead>
          <tr>
            <th>Id</th>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Start Time</th>
            <th>EndTime</th>
            <th>Day Of The Week</th>
            <th>Delete</th>
          </tr>
        </thead>
        <tbody>
          {userTimeslots &&
            Object.keys(userTimeslots).map((timeslot, i) => (
              <tr key={i}>
                <td>{userTimeslots[timeslot].id}</td>
                <td>{userTimeslots[timeslot].firstName}</td>
                <td>{userTimeslots[timeslot].lastName}</td>
                <td>{userTimeslots[timeslot].startTime}</td>
                <td>{userTimeslots[timeslot].endTime}</td>
                <td>{userTimeslots[timeslot].dayOfWeek}</td>
                <td><DeleteButton debounce={debounce} id={userTimeslots[timeslot].id}/></td>

              </tr>
            ))

          }
        </tbody>
      </table>
    </>
  )
}

function DeleteButton({id, debounce }){
  const [deleted, setDeleted] = useState(false)
  const requestOptions = {
    method: 'DELETE',
    redirect: 'follow',
  };

  const deleteTimeSlot = async (event) => {
   setDeleted(true)
   fetch('/api/luser/admin/timeslot/'+id, requestOptions)
    
  };

  return(
  <>
  {deleted ? (

    <Button disabled="disabled" variant="danger">Deleted</Button>

  ) : (

    <Button variant="danger" onClick={deleteTimeSlot}>Delete</Button>

  )}
  
  </>
  )

  

}

export default userTimeslots

