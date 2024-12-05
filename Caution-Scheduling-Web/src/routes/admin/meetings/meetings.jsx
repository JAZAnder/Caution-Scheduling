import { Outlet, Link } from "react-router-dom";
import Background from "../../../background";
import "./meetings.css";
import useFetch from "use-http";
import React, { useEffect, useState } from "react";
import MeetingDetailsButton from "./details/details";

function AdminMeetings() {
  const [loading, setLoading] = useState(false);
  const [userData, setUserData] = useState();
  const [filtering, setFiltering] = useState(false);
  const [tutorName, setTutorName] = useState("");
  const [studentName, setStudentName] = useState("");
  const [endTime, setEndTime] = useState("");
  const [startTime, setStartTime] = useState("");
  const [topicId, setTopicId] = useState("");
  const [dayOfWeek, setDayOfWeek] = useState("");
  const [date, setDate] = useState("");
  const [debounce, setDebounce] = useState(true);
  const [currentTime, setCurrentTime] = useState(new Date());

  useEffect(() => {
    setDebounce(!debounce);
  }, []);

  useEffect(() => {
    const getData = setTimeout(() => {
      setDebounce(!debounce);
      console.log("Should Be Refresh");
    }, 300);
    return () => clearTimeout(getData);
  }, [tutorName, studentName, endTime, startTime, topicId, dayOfWeek, date]);

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentTime(new Date());
    }, 60000); 

    return () => clearInterval(timer);
  }, []);

  const resetSearch = async (event) => {
    setTutorName("");
    setStudentName("");
    setEndTime("");
    setStartTime("");
    setTopicId("");
    setDayOfWeek("");
    setDate("");

    console.log("Should Be Refresh");
    setDebounce(!debounce);
  };

  return (
    <>
      <Background />
      <div className="admin-meetings-container">
        <div className="admin-meetings-page">
          <h2 className="admin-meetings-h2">Admin Meetings</h2>
          <div id="filterOnBar">
            <form>
              <input
                id="tutorName"
                className="admin-meetings-input"
                value={tutorName}
                onChange={(e) => setTutorName(e.target.value)}
                type="text"
                placeholder="Tutor Name"
              />
              <input
                id="studentName"
                className="admin-meetings-input"
                value={studentName}
                onChange={(e) => setStudentName(e.target.value)}
                type="text"
                placeholder="Student Name"
              />
              <input
                id="startTime"
                className="admin-meetings-input"
                value={startTime}
                onChange={(e) => setStartTime(e.target.value)}
                type="text"
                placeholder="Start Time (e.g., 10:00 AM)"
              />
              <input
                id="endTime"
                className="admin-meetings-input"
                value={endTime}
                onChange={(e) => setEndTime(e.target.value)}
                type="text"
                placeholder="End Time (e.g., 11:00 AM)"
              />
              <select
                name="topicId"
                id="topicId"
                className="admin-meetings-select"
                onChange={(e) => setTopicId(e.target.value)}
                value={topicId}
              >
                <option value="">Topic</option>
              </select>

              <select
                name="dayOfWeek"
                id="dayOfWeek"
                className="admin-meetings-select"
                onChange={(e) => setDayOfWeek(e.target.value)}
                value={dayOfWeek}
              >
                <option value="">Day</option>
                <option value="1">Monday</option>
                <option value="2">Tuesday</option>
                <option value="3">Wednesday</option>
                <option value="4">Thursday</option>
                <option value="5">Friday</option>
                <option value="6">Saturday</option>
                <option value="0">Sunday</option>
              </select>

              <input
                id="date"
                className="admin-meetings-input"
                value={date}
                onChange={(e) => setDate(e.target.value)}
                type="text"
                placeholder="Date MMDDYYYY"
              />

              <button
                type="button"
                className="admin-meetings-button"
                disabled={loading}
                onClick={resetSearch}
              >
                {loading ? "Waiting" : "Reset Search"}
              </button>
            </form>
          </div>

          <ListFilteredMeetings
            FLtutor={tutorName}
            FLstudent={studentName}
            FLstartTime={startTime}
            FLendTime={endTime}
            FLtopicId={topicId}
            FLdate={date}
            FLdayOfWeek={dayOfWeek}
            debounce={debounce}
            currentTime={currentTime}
          />
        </div>
      </div>
    </>
  );
}

function ListFilteredMeetings({
  FLtutor,
  FLstudent,
  FLstartTime,
  FLendTime,
  FLtopicId,
  FLdate,
  FLdayOfWeek,
  debounce,
  currentTime,
}) {
  const {
    data: meetings,
    loading,
    error,
  } = useFetch(
    `/api/meetings/filter?tutor=${encodeURIComponent(
      FLtutor
    )}&student=${encodeURIComponent(
      FLstudent
    )}&startTime=${encodeURIComponent(
      FLstartTime
    )}&endTime=${encodeURIComponent(
      FLendTime
    )}&topicId=${encodeURIComponent(
      FLtopicId
    )}&date=${encodeURIComponent(
      FLdate
    )}&dayOfWeek=${encodeURIComponent(FLdayOfWeek)}`,
    { method: "get" },
    [debounce]
  );

  if (loading) {
    return (
      <center>
        <div className="admin-meetings-loader"></div>
      </center>
    );
  }

  if (error) {
    return (
      <div className="admin-meetings-error">
        Error fetching meetings: {error.message}
      </div>
    );
  }

  function parseDate(dateStr) {
    if (!dateStr) return null;

    if (typeof dateStr === "number") {
      const dateStrNum = dateStr.toString();
      if (dateStrNum.length === 8 && /^\d+$/.test(dateStrNum)) {
        const month = parseInt(dateStrNum.substring(0, 2), 10);
        const day = parseInt(dateStrNum.substring(2, 4), 10);
        const year = parseInt(dateStrNum.substring(4, 8), 10);
        return new Date(year, month - 1, day);
      }
      console.error("Invalid date number format:", dateStr);
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
    const month = (date.getMonth() + 1).toString().padStart(2, "0");
    const day = date.getDate().toString().padStart(2, "0");
    const year = date.getFullYear();
    return `${month}/${day}/${year}`;
  }

  function timeToMinutes(timeStr) {
    const timeParts = timeStr.match(/(\d+):(\d+)\s*(AM|PM)/i);
    if (!timeParts) return null;
    let hours = parseInt(timeParts[1], 10);
    const minutes = parseInt(timeParts[2], 10);
    const ampm = timeParts[3].toUpperCase();
    if (ampm === "PM" && hours !== 12) hours += 12;
    if (ampm === "AM" && hours === 12) hours = 0;
    return hours * 60 + minutes;
  }

  function minutesToTimeString(totalMinutes) {
    let hours = Math.floor(totalMinutes / 60);
    const minutes = totalMinutes % 60;
    const ampm = hours >= 12 ? "PM" : "AM";
    if (hours > 12) hours -= 12;
    if (hours === 0) hours = 12;
    const displayMinutes = minutes.toString().padStart(2, "0");
    return `${hours}:${displayMinutes} ${ampm}`;
  }

  function getStartTimeMinutes(meeting) {
    if (
      !meeting.TutorHour ||
      !meeting.TutorHour.Hour ||
      !meeting.TutorHour.Hour.startTime
    ) {
      console.error(
        "Error: meeting.TutorHour.Hour.startTime is undefined for meeting:",
        meeting
      );
      return null;
    }
    return timeToMinutes(meeting.TutorHour.Hour.startTime);
  }

  function getEndTimeMinutes(meeting) {
    if (
      !meeting.TutorHour ||
      !meeting.TutorHour.Hour ||
      !meeting.TutorHour.Hour.endTime
    ) {
      console.error(
        "Error: meeting.TutorHour.Hour.endTime is undefined for meeting:",
        meeting
      );
      return null;
    }
    return timeToMinutes(meeting.TutorHour.Hour.endTime);
  }

  let meetingsArray = [];
  if (meetings) {
    meetingsArray = Object.values(meetings);
  }

  meetingsArray.sort((a, b) => {
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
        currentMergedMeeting.TutorHour.Tutor.id ===
          currentMeeting.TutorHour.Tutor.id &&
        currentMergedMeeting.Student.id === currentMeeting.Student.id &&
        currentMergedMeeting.Topic.id === currentMeeting.Topic.id &&
        currStartTime !== null &&
        currEndTime !== null &&
        currStartTime <= currentMergedMeeting.mergedEndTime
      ) {
        currentMergedMeeting.mergedStartTime = Math.min(
          currentMergedMeeting.mergedStartTime,
          currStartTime
        );
        currentMergedMeeting.mergedEndTime = Math.max(
          currentMergedMeeting.mergedEndTime,
          currEndTime
        );
        currentMergedMeeting.mergedIds.push(currentMeeting.id);
      } else {
        currentMergedMeeting.TutorHour.Hour.startTime = minutesToTimeString(
          currentMergedMeeting.mergedStartTime
        );
        currentMergedMeeting.TutorHour.Hour.endTime = minutesToTimeString(
          currentMergedMeeting.mergedEndTime
        );
        mergedMeetings.push(currentMergedMeeting);
        currentMergedMeeting = {
          ...currentMeeting,
          mergedIds: [currentMeeting.id],
          mergedStartTime: currStartTime,
          mergedEndTime: currEndTime,
        };
      }
    }

    currentMergedMeeting.TutorHour.Hour.startTime = minutesToTimeString(
      currentMergedMeeting.mergedStartTime
    );
    currentMergedMeeting.TutorHour.Hour.endTime = minutesToTimeString(
      currentMergedMeeting.mergedEndTime
    );
    mergedMeetings.push(currentMergedMeeting);

    return mergedMeetings;
  }

  const mergedMeetings = mergeConsecutiveMeetings(meetingsArray);

  function hasMeetingEnded(meeting) {
    const meetingDate = parseDate(meeting.date);
    const endTimeMinutes = getEndTimeMinutes(meeting);
    if (!meetingDate || endTimeMinutes === null) return true;
    const endDateTime = new Date(meetingDate);
    endDateTime.setHours(
      Math.floor(endTimeMinutes / 60),
      endTimeMinutes % 60,
      0,
      0
    );
    return endDateTime <= currentTime;
  }

  const upcomingMeetings = mergedMeetings.filter(
    (meeting) => !hasMeetingEnded(meeting)
  );

  return (
    <>
      <div className="admin-meetings-body">
        <div id="meetingsTable">
          {upcomingMeetings.length === 0 ? (
            <div className="admin-meetings-no-data">
              No upcoming meetings found.
            </div>
          ) : (
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
                {upcomingMeetings.map((meeting, i) => (
                  <tr key={i}>
                    <td>{meeting.mergedIds.join(", ")}</td>
                    <td>{meeting.Topic.description}</td>
                    <td>
                      {meeting.Student.firstName} {meeting.Student.lastName}
                    </td>
                    <td>
                      {meeting.TutorHour.Tutor.firstName}{" "}
                      {meeting.TutorHour.Tutor.lastName}
                    </td>
                    <td>{formatDate(meeting.date)}</td>
                    <td>
                      {meeting.TutorHour.Hour.startTime} -{" "}
                      {meeting.TutorHour.Hour.endTime}
                    </td>
                    <td>
                      <MeetingDetailsButton meeting={meeting} />
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          )}
        </div>
      </div>
    </>
  );
}

export default AdminMeetings;
