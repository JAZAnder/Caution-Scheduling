import { Outlet, Link } from "react-router-dom";
import Background from "../../background";
import './my-meetings.css';
import React, { useEffect, useState } from "react";
import useFetch from "use-http";
import MeetingDetailsButton from "./details/details.jsx";

function MyMeeting() {
  const [meetings, setMeetings] = useState([]);
  const [loading, setLoading] = useState(true);
  const { get, response, error } = useFetch('/api/meetings');
  const [currentTime, setCurrentTime] = useState(new Date());

  useEffect(() => {
    async function fetchMeetings() {
      setLoading(true);
      const fetchedMeetings = await get();
      if (response.ok) {
        setMeetings(fetchedMeetings);
      } else {
        setMeetings([]);
      }
      setLoading(false);
    }
    fetchMeetings();
  }, [get, response]);

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentTime(new Date());
    }, 60000);

    return () => clearInterval(timer);
  }, []);

  if (loading) {
    return (
      <center>
        <div className="loader"></div>
      </center>
    );
  }

  if (!meetings || meetings.length === 0) {
    return (
      <>
        <Background />
        <div className="my-meetings-container">
          <div className="mymeetings-text-center">No meetings found. Try clicking another tab or reload?</div>
        </div>
      </>
    );
  }

  console.log('Meetings Data:', meetings);

  function parseDate(dateStr) {
    if (!dateStr) return null;
  
    if (typeof dateStr === 'number') {
      const dateStrNum = dateStr.toString();
      if (dateStrNum.length === 8 && /^\d+$/.test(dateStrNum)) {
        const month = parseInt(dateStrNum.substring(0, 2), 10);
        const day = parseInt(dateStrNum.substring(2, 4), 10);
        const year = parseInt(dateStrNum.substring(4, 8), 10);
        return new Date(year, month - 1, day); 
      }
      console.error('Invalid date number format:', dateStr);
      return null;
    }

    if (dateStr instanceof Date) {
      return dateStr;
    }
  
    return null;
  }

  function formatDate(dateStr) {
    const date = parseDate(dateStr);
    if (!date) return dateStr;
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const day = date.getDate().toString().padStart(2, '0');
    const year = date.getFullYear();
    return `${month}/${day}/${year}`;
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

  function minutesToTimeString(totalMinutes) {
    let hours = Math.floor(totalMinutes / 60);
    const minutes = totalMinutes % 60;
    const ampm = hours >= 12 ? 'PM' : 'AM';
    if (hours > 12) hours -= 12;
    if (hours === 0) hours = 12;
    const displayMinutes = minutes.toString().padStart(2, '0');
    return `${hours}:${displayMinutes} ${ampm}`;
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
      mergedStartTime: getStartTimeMinutes(meetingsArray[0]),
      mergedEndTime: getEndTimeMinutes(meetingsArray[0]),
    };

    for (let i = 1; i < meetingsArray.length; i++) {
      const currentMeeting = meetingsArray[i];

      const currStartTime = getStartTimeMinutes(currentMeeting);
      const currEndTime = getEndTimeMinutes(currentMeeting);

      if (
        currentMergedMeeting.date === currentMeeting.date &&
        currentMergedMeeting.Tutor.id === currentMeeting.Tutor.id &&
        currentMergedMeeting.Student.id === currentMeeting.Student.id &&
        currentMergedMeeting.Topic.id === currentMeeting.Topic.id &&
        currStartTime !== null &&
        currEndTime !== null &&
        currStartTime <= currentMergedMeeting.mergedEndTime
      ) {
        currentMergedMeeting.mergedStartTime = Math.min(currentMergedMeeting.mergedStartTime, currStartTime);
        currentMergedMeeting.mergedEndTime = Math.max(currentMergedMeeting.mergedEndTime, currEndTime);
        currentMergedMeeting.mergedIds.push(currentMeeting.id);
      } else {
        currentMergedMeeting.Hour.startTime = minutesToTimeString(currentMergedMeeting.mergedStartTime);
        currentMergedMeeting.Hour.endTime = minutesToTimeString(currentMergedMeeting.mergedEndTime);
        mergedMeetings.push(currentMergedMeeting);
        currentMergedMeeting = {
          ...currentMeeting,
          mergedIds: [currentMeeting.id],
          mergedStartTime: currStartTime,
          mergedEndTime: currEndTime,
        };
      }
    }

    currentMergedMeeting.Hour.startTime = minutesToTimeString(currentMergedMeeting.mergedStartTime);
    currentMergedMeeting.Hour.endTime = minutesToTimeString(currentMergedMeeting.mergedEndTime);
    mergedMeetings.push(currentMergedMeeting);

    return mergedMeetings;
  }

  const mergedMeetings = mergeConsecutiveMeetings(meetings);

  console.log('Merged Meetings:', mergedMeetings);

  function hasMeetingEnded(meeting) {
    const meetingDate = parseDate(meeting.date);
    const endTimeMinutes = getEndTimeMinutes(meeting);
    if (!meetingDate || endTimeMinutes === null) return true;
    const endDateTime = new Date(meetingDate);
    endDateTime.setHours(Math.floor(endTimeMinutes / 60), endTimeMinutes % 60, 0, 0);
    return endDateTime <= currentTime;
  }

  const upcomingMeetings = mergedMeetings.filter(meeting => !hasMeetingEnded(meeting));

  return (
    <>
      <Background />
      <div className="my-meetings-container">
        <div className="mymeetings-body">
          <div className="mymeetings-page">
            <h2 className="my-meetings-h2">My Meetings</h2>
            {upcomingMeetings.length === 0 ? (
              <div className="mymeetings-text-center">No upcoming meetings found.</div>
            ) : (
              <table className="mymeetings-table">
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
                  {upcomingMeetings.map((meeting, i) => (
                    <tr key={i}>
                      <td>{meeting.mergedIds.join(', ')}</td>
                      <td>{meeting.Topic.description}</td>
                      <td>
                        {meeting.Student.firstName} {meeting.Student.lastName}
                      </td>
                      <td>
                        {meeting.Tutor.firstName} {meeting.Tutor.lastName}
                      </td>
                      <td>{formatDate(meeting.date)}</td>
                      <td>
                        {meeting.Hour.startTime} - {meeting.Hour.endTime}
                      </td>
                      <td>
                        <MeetingDetailsButton meeting={meeting} className="details-details-button" />
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            )}
          </div>
        </div>
      </div>
    </>
  );
}

export default MyMeeting;
