import { Outlet, Link } from "react-router-dom";
import Background from "../../background";
import './my-meetings.css';
import React, { useEffect, useState } from "react";
import useFetch from "use-http";
import MeetingDetailsButton from "./details/details.jsx"

function MyMeeting() {
  const { data: meetings, loading, error } = useFetch(
    `/api/meetings`,
    { method: "get" },
    []
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
  <Background />
  <div className="mymeetings-body ">
    {/* <div style={{ minHeight: "150px" }}> Black Space?</div> */}
    <div id="userNameTable">
        <table className="table-with-bordered">
          <thead>
            <tr>
              <th>Meeting Id</th>
              <th>Topic</th>
              <th>Student</th>
              <th>Tutor</th>
              <th>Date</th>
              <th>Time</th>
              <th>Details</th>
            </tr>
          </thead>
          <tbody>
            {meetings &&
              Object.keys(meetings).map((meeting, i) => (
                <tr key={i}>
                  <td>{meetings[meeting].id}</td>
                  <td>{meetings[meeting].Topic.description}</td>
                  <td>{meetings[meeting].Student.firstName} {meetings[meeting].Student.lastName}</td>
                  <td>{meetings[meeting].Tutor.firstName} {meetings[meeting].Tutor.lastName}</td>
                  <td>{meetings[meeting].date}</td>
                  <td>{meetings[meeting].Hour.startTime} - {meetings[meeting].Hour.endTime}</td>
                  <td>
                    <MeetingDetailsButton meeting={meetings[meeting]} />
                  </td>
                </tr>
              ))}
          </tbody>
        </table>
      </div>


  </div>
      

    </>
  );
}

export default MyMeeting;
