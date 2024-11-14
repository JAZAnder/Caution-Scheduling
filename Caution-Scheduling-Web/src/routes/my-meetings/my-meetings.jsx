import { Outlet, Link } from "react-router-dom";
import Background from "../../background";
import './my-meetings.css';
import React from "react";
import useFetch from "use-http";
import MeetingDetailsButton from "./details/details.jsx";

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

  if (!meetings || meetings.length === 0) {
    return <div>No meetings found.</div>;
  }

  console.log('Meetings Data:', meetings);

  function parseDate(dateStr) {
    console.log('dateStr:', dateStr, 'Type:', typeof dateStr);
    if (!dateStr) return null;

    if (typeof dateStr === 'string') {
      if (dateStr.length === 8) {
        const month = dateStr.substring(0, 2);
        const day = dateStr.substring(2, 4);
        const year = dateStr.substring(4, 8);
        return new Date(`${year}-${month}-${day}`);
      }

      const parsedDate = new Date(dateStr);
      if (!isNaN(parsedDate)) {
        return parsedDate;
      }
    }

    if (dateStr instanceof Date) {
      return dateStr;
    }

    if (typeof dateStr === 'number') {
      return new Date(dateStr);
    }

    return null;
  }

  function timeToMinutes(timeStr) {
    const timeParts = timeStr.match(/(\d+):(\d+)\s*(AM|PM)/i);
    if (!timeParts) return null;
    let hours = parseInt(timeParts[1], 10);
    const minutes = parseInt(timeParts[2], 10);
    const ampm = timeParts[3].toUpperCase();
    if (ampm === 'PM' && hours !== 12) hours += 12;
    if (ampm === 'AM' && hours === 12) hours = 0;
    return hours * 60 + minutes;
  }

  function getStartTimeMinutes(meeting) {
    if (!meeting.Hour || !meeting.Hour.startTime) {
      console.error('Error: meeting.Hour.startTime is undefined for meeting:', meeting);
      return null;
    }
    return timeToMinutes(meeting.Hour.startTime);
  }

  function getEndTimeMinutes(meeting) {
    if (!meeting.Hour || !meeting.Hour.endTime) {
      console.error('Error: meeting.Hour.endTime is undefined for meeting:', meeting);
      return null;
    }
    return timeToMinutes(meeting.Hour.endTime);
  }

  meetings.sort((a, b) => {
    const dateA = parseDate(a.date);
    const dateB = parseDate(b.date);
    if (dateA && dateB && dateA - dateB !== 0) return dateA - dateB;
    return getStartTimeMinutes(a) - getStartTimeMinutes(b);
  });

  function mergeConsecutiveMeetings(meetingsArray) {
    const mergedMeetings = [];
    if (meetingsArray.length === 0) return mergedMeetings;

    let currentMergedMeeting = {
      ...meetingsArray[0],
      mergedIds: [meetingsArray[0].id],
    };

    for (let i = 1; i < meetingsArray.length; i++) {
      const currentMeeting = meetingsArray[i];

      const prevEndTime = getEndTimeMinutes(currentMergedMeeting);
      const currStartTime = getStartTimeMinutes(currentMeeting);

      if (
        currentMergedMeeting.date === currentMeeting.date &&
        currentMergedMeeting.Tutor.id === currentMeeting.Tutor.id &&
        currentMergedMeeting.Student.id === currentMeeting.Student.id &&
        currentMergedMeeting.Topic.id === currentMeeting.Topic.id &&
        prevEndTime !== null &&
        currStartTime !== null &&
        prevEndTime === currStartTime
      ) {
        currentMergedMeeting.Hour.endTime = currentMeeting.Hour.endTime;
        currentMergedMeeting.mergedIds.push(currentMeeting.id);
      } else {
        mergedMeetings.push(currentMergedMeeting);
        currentMergedMeeting = {
          ...currentMeeting,
          mergedIds: [currentMeeting.id],
        };
      }
    }

    mergedMeetings.push(currentMergedMeeting);
    return mergedMeetings;
  }

  const mergedMeetings = mergeConsecutiveMeetings(meetings);

  console.log('Merged Meetings:', mergedMeetings);

  return (
    <>
      <Background />
      <div className="mymeetings-body">
        <div id="userNameTable">
          <table className="table-with-bordered">
            <thead>
              <tr>
                <th>Meeting Id(s)</th>
                <th>Topic</th>
                <th>Student</th>
                <th>Tutor</th>
                <th>Date</th>
                <th>Time</th>
                <th>Details</th>
              </tr>
            </thead>
            <tbody>
              {mergedMeetings.map((meeting, i) => (
                <tr key={i}>
                  <td>{meeting.mergedIds.join(', ')}</td>
                  <td>{meeting.Topic.description}</td>
                  <td>
                    {meeting.Student.firstName} {meeting.Student.lastName}
                  </td>
                  <td>
                    {meeting.Tutor.firstName} {meeting.Tutor.lastName}
                  </td>
                  <td>{meeting.date}</td>
                  <td>
                    {meeting.Hour.startTime} - {meeting.Hour.endTime}
                  </td>
                  <td>
                    <MeetingDetailsButton meeting={meeting} />
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
