import React, { useState } from "react";
import { Outlet, Link } from "react-router-dom";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import Background from "../../background";
import "./ScheduleMeeting.css";

const ScheduleMeeting = ({ isAdmin }) => {
  const [selectedDate, setSelectedDate] = useState(null);
  const [startTime, setStartTime] = useState("");
  const [endTimeOptions, setEndTimeOptions] = useState([]);
  const [endTime, setEndTime] = useState("");

  const isWeekday = (date) => {
    const day = date.getDay();
    return day >= 1 && day <= 4; // Monday (1) to Thursday (4)
  };

  const generateTimeSlots = () => {
    const times = [];
    let start = new Date();
    start.setHours(9, 30, 0, 0); // Start time at 9:30 AM

    let end = new Date();
    end.setHours(21, 0, 0, 0); // End time at 9:00 PM

    while (start <= end) {
      times.push(new Date(start));
      start.setMinutes(start.getMinutes() + 15); // Increment by 15 minutes
    }
    return times;
  };

  const availableStartTimes = generateTimeSlots().map((time) =>
    time.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
  );

  const handleTimeSlotChange = (event) => {
    const selectedStartTime = event.target.value;
    setStartTime(selectedStartTime);

    // Convert start time to a Date object
    const [hours, minutes, period] = selectedStartTime.split(/[: ]/);
    let startHour = parseInt(hours, 10);
    if (period === "PM" && startHour !== 12) {
      startHour += 12; // Convert PM times to 24-hour format
    } else if (period === "AM" && startHour === 12) {
      startHour = 0; // Convert 12 AM to 0 hours
    }

    const start = new Date();
    start.setHours(startHour, parseInt(minutes, 10));

    const newEndOptions = [];
    let nextSlot = new Date(start);
    for (let i = 15; i <= 60; i += 15) {
      nextSlot = new Date(start);
      nextSlot.setMinutes(start.getMinutes() + i);
      if (nextSlot.getHours() > 21 || (nextSlot.getHours() === 21 && nextSlot.getMinutes() > 0)) break;
      newEndOptions.push(
        nextSlot.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
      );
    }

    setEndTimeOptions(newEndOptions);
    setEndTime(newEndOptions[0] || ""); // Set default end time
  };

  return (
    <>
      <Background />
      <h1 style={{ color: 'white' }}>Schedule a Meeting</h1>
      <div className="container">
        {isAdmin && (
          <button className="button admin-button">
            Make a Meeting
          </button>
        )}
        <DatePicker
          selected={selectedDate}
          onChange={(date) => setSelectedDate(date)}
          filterDate={isWeekday}
          placeholderText="Select a date"
          className="input-field"
        />
        <select
          value={startTime}
          onChange={handleTimeSlotChange}
          className="input-field"
        >
          <option value="" disabled>
            Select start time
          </option>
          {availableStartTimes.map((time, index) => (
            <option key={index} value={time}>
              {time}
            </option>
          ))}
        </select>
        <select
          value={endTime}
          onChange={(e) => setEndTime(e.target.value)}
          className="input-field"
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
        <input type="text" placeholder="Tutor" className="input-field" />
        <button className="button schedule-button">Schedule Meeting</button>
      </div>
    </>
  );
};

export default ScheduleMeeting;