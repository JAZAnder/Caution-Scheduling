import React, { useState, useEffect } from "react";
import { Outlet, Link } from "react-router-dom";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import Background from "../../background";
import "./ScheduleMeeting.css";
import axios from "axios";

const ScheduleMeeting = ({ isAdmin }) => {
  const [selectedDate, setSelectedDate] = useState(null);
  const [startTime, setStartTime] = useState("");
  const [endTimeOptions, setEndTimeOptions] = useState([]);
  const [endTime, setEndTime] = useState("");
  const [tutors, setTutors] = useState([]);
  const [selectedTutor, setSelectedTutor] = useState("");
  const [tutorTimeSlots, setTutorTimeSlots] = useState([]);

  useEffect(() => {
    const fetchTutors = async () => {
      try {
        const response = await axios.get("/api/lusers/tutors");
        const tutors = response.data;
        setTutors(tutors);
      } catch (error) {
        console.error("Error fetching tutors:", error);
      }
    };
    fetchTutors();
  }, []);

  const isWeekday = (date) => {
    const day = date.getDay();
    return day >= 1 && day <= 4;
  };

  const generateTimeSlots = () => {
    const times = [];
    let start = new Date();
    start.setHours(9, 30, 0, 0);

    let end = new Date();
    end.setHours(21, 0, 0, 0);

    while (start <= end) {
      times.push(new Date(start));
      start.setMinutes(start.getMinutes() + 15);
    }
    return times;
  };

  const availableStartTimes = generateTimeSlots().map((time) =>
    time.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" })
  );

  const getAvailableStartTimesForTutor = () => {
    if (!selectedTutor) {
      return availableStartTimes;
    }
    return availableStartTimes.filter((time) => {
      const timeSlot = tutorTimeSlots.find(
        (slot) => slot.startTime === time && slot.available
      );
      return timeSlot !== undefined;
    });
  };

  const filteredStartTimes = getAvailableStartTimesForTutor();

  const handleTutorChange = async (e) => {
    const selectedTutor = e.target.value;
    setSelectedTutor(selectedTutor);

    try {
      const response = await axios.get(`/api/tutor/hours/${selectedTutor}`);
      const timeSlots = response.data;
      setTutorTimeSlots(timeSlots);
    } catch (error) {
      console.error("Error fetching time slots:", error);
    }
  };

  const handleTimeSlotChange = (event) => {
    const selectedStartTime = event.target.value;
    setStartTime(selectedStartTime);

    const [hours, minutes, period] = selectedStartTime.split(/[: ]/);
    let startHour = parseInt(hours, 10);
    if (period === "PM" && startHour !== 12) {
      startHour += 12;
    } else if (period === "AM" && startHour === 12) {
      startHour = 0;
    }

    const start = new Date();
    start.setHours(startHour, parseInt(minutes, 10));

    const newEndOptions = [];
    let nextSlot = new Date(start);
    for (let i = 15; i <= 60; i += 15) {
      nextSlot = new Date(start);
      nextSlot.setMinutes(start.getMinutes() + i);
      if (
        nextSlot.getHours() > 21 ||
        (nextSlot.getHours() === 21 && nextSlot.getMinutes() > 0)
      )
        break;
      newEndOptions.push(
        nextSlot.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" })
      );
    }

    setEndTimeOptions(newEndOptions);
    setEndTime(newEndOptions[0] || "");
  };

  return (
    <>
      <Background />
      <h1 style={{ color: "white" }}>Schedule a Meeting</h1>
      <div className="container">
        {isAdmin && (
          <button className="button admin-button">Make a Meeting</button>
        )}
        <DatePicker
          selected={selectedDate}
          onChange={(date) => setSelectedDate(date)}
          filterDate={isWeekday}
          placeholderText="Select a date"
          className="input-field"
        />
        <div className="time-selection">
          <select
            value={startTime}
            onChange={handleTimeSlotChange}
            className="input-field time-select"
            disabled={!selectedTutor} // Disable until a tutor is selected
          >
            <option value="" disabled>
              Select start time
            </option>
            {filteredStartTimes.map((time, index) => (
              <option key={index} value={time}>
                {time}
              </option>
            ))}
          </select>
          <select
            value={endTime}
            onChange={(e) => setEndTime(e.target.value)}
            className="input-field time-select"
            disabled={!endTimeOptions.length}
          >
            <option value="" disabled>
              Select end time
            </option>
            {endTimeOptions.map((time, index) => (
              <option key={index} value={time}>
                {time}
              </option>
            ))}
          </select>
        </div>
        <select
          value={selectedTutor}
          onChange={handleTutorChange}
          className="input-field"
        >
          <option value="" disabled>
            Select a Tutor
          </option>
          {tutors.map((tutor, index) => (
            <option key={tutor.userId || index} value={tutor.userName}>
              {tutor.fullName}
            </option>
          ))}
        </select>
        <button className="button schedule-button">Schedule Meeting</button>
      </div>
    </>
  );
};

export default ScheduleMeeting;
