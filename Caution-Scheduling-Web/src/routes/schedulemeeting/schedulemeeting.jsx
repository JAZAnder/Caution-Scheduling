import React, { useState, useEffect } from "react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import Background from "../../background";
import "./ScheduleMeeting.css";
import axios from "axios";

const ScheduleMeeting = ({ isAdmin }) => {
  const [selectedDate, setSelectedDate] = useState(null);
  const [startTime, setStartTime] = useState(null);
  const [endTimeOptions, setEndTimeOptions] = useState([]);
  const [endTime, setEndTime] = useState("");
  const [tutors, setTutors] = useState([]);
  const [selectedTutor, setSelectedTutor] = useState("");
  const [tutorAvailability, setTutorAvailability] = useState([]);
  const [hours, setHours] = useState([]);

  useEffect(() => {
    const fetchTutors = async () => {
      try {
        const response = await axios.get("/api/lusers/tutors");
        setTutors(response.data);
      } catch (error) {
        console.error("Error fetching tutors:", error);
      }
    };
    fetchTutors();
  }, []);

  useEffect(() => {
    const fetchHours = async () => {
      try {
        const response = await axios.get("/api/hours");
        console.log("Hours Data:", response.data); 
        setHours(response.data);
      } catch (error) {
        console.error("Error fetching hours:", error);
      }
    };
    fetchHours();
  }, []);

  const isWeekday = (date) => {
    const day = date.getDay();
    return day >= 1 && day <= 4;
  };

  const getAvailableHoursForTutor = () => {
    if (!selectedTutor) {
      return [];
    }

    const availableHourIds = tutorAvailability
      .filter((slot) => slot.available)
      .map((slot) => slot.hourId);

    console.log("Available Hour IDs:", availableHourIds);

    const availableHours = hours.filter((hour) => availableHourIds.includes(hour.id));

    console.log("Available Hours:", availableHours); 

    return availableHours;
  };

  const filteredHours = getAvailableHoursForTutor();

  const handleTutorChange = async (e) => {
    const selectedTutorId = e.target.value;
    setSelectedTutor(selectedTutorId);

    try {
      const response = await axios.get(`/api/availability/${selectedTutorId}`);
      console.log("Tutor Availability Data:", response.data);
      setTutorAvailability(response.data);
    } catch (error) {
      console.error("Error fetching tutor availability:", error);
    }
  };

  const handleTimeSlotChange = (event) => {
    const selectedHourId = parseInt(event.target.value, 10);
    setStartTime(selectedHourId);

    const selectedHour = filteredHours.find((hour) => hour.id === selectedHourId);

    if (selectedHour) {
      const startMinutes = parseTime(selectedHour.startTime);
      const endMinutes = parseTime(selectedHour.endTime);
      const increment = 15;

      const newEndOptions = [];
      for (
        let time = startMinutes + increment;
        time <= endMinutes;
        time += increment
      ) {
        newEndOptions.push(formatTime(time));
      }

      setEndTimeOptions(newEndOptions);
      setEndTime(newEndOptions[0] || "");
    } else {
      setEndTimeOptions([]);
      setEndTime("");
    }
  };

  const parseTime = (timeStr) => {
    const [time, modifier] = timeStr.split(" ");
    let [hours, minutes] = time.split(":").map(Number);

    if (modifier === "PM" && hours !== 12) {
      hours += 12;
    }
    if (modifier === "AM" && hours === 12) {
      hours = 0;
    }

    return hours * 60 + minutes;
  };

  const formatTime = (totalMinutes) => {
    let hours = Math.floor(totalMinutes / 60);
    const minutes = totalMinutes % 60;
    const ampm = hours >= 12 ? "PM" : "AM";
    hours = hours % 12 || 12;

    return `${hours < 10 ? "0" : ""}${hours}:${
      minutes < 10 ? "0" : ""
    }${minutes} ${ampm}`;
  };

  return (
    <>
    <div style={{ minHeight: "230px" }}> Black Space?</div>
      <Background />
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
        <select
          value={selectedTutor}
          onChange={handleTutorChange}
          className="input-field"
        >
          <option value="" disabled>
            Select a Tutor
          </option>
          {tutors.map((tutor) => (
            <option key={tutor.userId} value={tutor.userId}>
              {tutor.fullName}
            </option>
          ))}
        </select>
        <div className="time-selection">
          <select
            value={startTime || ""}
            onChange={handleTimeSlotChange}
            className="input-field time-select"
            disabled={!selectedTutor || !tutorAvailability.length}
          >
            <option value="" disabled>
              Select start time
            </option>
            {filteredHours.map((hour) => {
              console.log("Rendering Hour:", hour);
              return (
                <option key={hour.id} value={hour.id}>
                  {hour.startTime}
                </option>
              );
            })}
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
        <button className="button schedule-button">Schedule Meeting</button>
      </div>
    </>
  );
};

export default ScheduleMeeting;
