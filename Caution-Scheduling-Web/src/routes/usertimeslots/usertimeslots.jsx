import { Link, useSearchParams } from "react-router-dom";
import React, { useEffect, useState } from "react";
import useFetch, { CachePolicies } from "use-http";

const userTimeslots = () => {
  const [params] = useSearchParams();
  const [hourId, setHourId] = useState(params.get("hourId"))
  const [tutorId, setTutorId] = useState(params.get("tutorId"))
  const [dayOfWeek, setDayOfWeek] = useState(params.get("dayOfWeek"))
  const [debounce, setDebounce] = useState(true);

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

  const resetSearch = async (event) => {
    setDayOfWeek('');
    setHourId('');
    setTutorId('');
    setDebounce(!debounce);
  };

  if (usersLoading || timeLoading) {
    return (<><center><div className="loader"></div></center></>)
  }


  return (
    <>
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


        <button
          type="button"
          onClick={resetSearch}
          className="timeslots-search-button"
        >
          Reset Search
        </button>
      </div>

      <div>
        <FilterUserTimeSlots
          FLhourId={hourId}
          FLtutorId={tutorId}
          FLdayOfWeek={dayOfWeek}
          debounce={debounce}

        />
      </div>



    </>
  )
}

function FilterUserTimeSlots({ FLhourId, FLtutorId, FLdayOfWeek, debounce }) {






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
                <td><button>Delete</button></td>

              </tr>
            ))

          }
        </tbody>
      </table>
    </>
  )
}

export default userTimeslots

