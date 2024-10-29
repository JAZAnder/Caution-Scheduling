import React, { useState, useEffect } from "react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import Background from "../../background";
import "./scheduleMeeting.css";
import axios from "axios";

const ScheduleMeeting = ({ isAdmin }) => {
  const [selectedDate, setSelectedDate] = useState(null);
  const [startTime, setStartTime] = useState("");
  const [endTime, setEndTime] = useState("");
  const [tutors, setTutors] = useState([]);
  const [availableTutors, setAvailableTutors] = useState([]);
  const [selectedTutor, setSelectedTutor] = useState("");
  const [tutorAvailability, setTutorAvailability] = useState([]);
  const [tutorsAvailabilityMap, setTutorsAvailabilityMap] = useState({});
  const [hours, setHours] = useState([]);
  const [scheduledMeetings, setScheduledMeetings] = useState([]);

  useEffect(() => {
    const fetchTutorsAndAvailability = async () => {
      try {
        const tutorsResponse = await axios.get("/api/lusers/tutors");
        const tutorsData = tutorsResponse.data;
        setTutors(tutorsData);

        const availabilityPromises = tutorsData.map((tutor) =>
          axios.get(`/api/availability/${tutor.userId}`).then((res) => ({
            tutorId: tutor.userId,
            availability: res.data,
          }))
        );

        const availabilityResults = await Promise.all(availabilityPromises);

        const availabilityMap = {};
        availabilityResults.forEach((result) => {
          availabilityMap[result.tutorId] = result.availability;
        });

        setTutorsAvailabilityMap(availabilityMap);
      } catch (error) {
        console.error("Error fetching tutors or their availability:", error);
      }
    };

    fetchTutorsAndAvailability();
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

  useEffect(() => {
    if (
      selectedDate &&
      tutors.length > 0 &&
      Object.keys(tutorsAvailabilityMap).length > 0 &&
      hours.length > 0
    ) {
      const jsDayOfWeek = selectedDate.getDay();

      const dayOfWeekMap = {
        1: 1, // Monday
        2: 2, // Tuesday
        3: 3, // Wednesday
        4: 4, // Thursday
      };

      const dayOfWeek = dayOfWeekMap[jsDayOfWeek];

      if (!dayOfWeek) {
        setAvailableTutors([]);
        return;
      }

      const tutorsAvailableOnDate = tutors.filter((tutor) => {
        const availability = tutorsAvailabilityMap[tutor.userId];
        if (!availability) return false;

        const availableHourIds = availability
          .filter((slot) => slot.available)
          .map((slot) => slot.hourId);

        const hoursForDay = hours.filter(
          (hour) =>
            hour.dayOfWeek === dayOfWeek && availableHourIds.includes(hour.id)
        );

        return hoursForDay.length > 0;
      });

      setAvailableTutors(tutorsAvailableOnDate);
    } else {
      setAvailableTutors([]);
    }
  }, [selectedDate, tutors, tutorsAvailabilityMap, hours]);

  const isWeekday = (date) => {
    const day = date.getDay();
    return day >= 1 && day <= 4; // Monday to Thursday
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

  const getAvailableStartTimes = () => {
    if (!selectedTutor || !selectedDate) {
      return [];
    }

    const jsDayOfWeek = selectedDate.getDay(); 

    const dayOfWeekMap = {
      1: 1, // Monday
      2: 2, // Tuesday
      3: 3, // Wednesday
      4: 4, // Thursday
    };

    const dayOfWeek = dayOfWeekMap[jsDayOfWeek];

    if (!dayOfWeek) {
      return [];
    }

    const availableHourIds = tutorAvailability
      .filter((slot) => slot.available)
      .map((slot) => slot.hourId);

    const hoursForDay = hours.filter(
      (hour) => hour.dayOfWeek === dayOfWeek && availableHourIds.includes(hour.id)
    );

    let allStartTimes = [];

    hoursForDay.forEach((hour) => {
      const startMinutes = parseTime(hour.startTime);
      const endMinutes = parseTime(hour.endTime) - 15;
      for (let time = startMinutes; time <= endMinutes; time += 15) {
        allStartTimes.push(formatTime(time));
      }
    });

    const formattedDate = selectedDate.toISOString().split("T")[0];
    const bookedTimes = scheduledMeetings
      .filter(
        (meeting) =>
          meeting.tutorId === selectedTutor && meeting.date === formattedDate
      )
      .map((meeting) => ({
        start: parseTime(meeting.startTime),
        end: parseTime(meeting.endTime),
      }));

    const availableStartTimes = allStartTimes.filter((timeStr) => {
      const time = parseTime(timeStr);
      const isBooked = bookedTimes.some(
        (booking) => time >= booking.start && time < booking.end
      );
      return !isBooked;
    });

    return availableStartTimes;
  };

  const getAvailableEndTimes = () => {
    if (!startTime || !selectedTutor || !selectedDate) {
      return [];
    }

    const jsDayOfWeek = selectedDate.getDay();

    const dayOfWeekMap = {
      1: 1,
      2: 2,
      3: 3,
      4: 4,
    };

    const dayOfWeek = dayOfWeekMap[jsDayOfWeek];

    const availableHourIds = tutorAvailability
      .filter((slot) => slot.available)
      .map((slot) => slot.hourId);

    const hoursForDay = hours.filter(
      (hour) => hour.dayOfWeek === dayOfWeek && availableHourIds.includes(hour.id)
    );

    const startMinutes = parseTime(startTime);

    const currentHourBlock = hoursForDay.find((hour) => {
      const hourStart = parseTime(hour.startTime);
      const hourEnd = parseTime(hour.endTime);
      return startMinutes >= hourStart && startMinutes < hourEnd;
    });

    if (!currentHourBlock) {
      return [];
    }

    const hourEndMinutes = parseTime(currentHourBlock.endTime);

    let endTimes = [];
    for (
      let time = startMinutes + 15;
      time <= hourEndMinutes;
      time += 15
    ) {
      endTimes.push(formatTime(time));
    }

    const formattedDate = selectedDate.toISOString().split("T")[0];
    const bookedTimes = scheduledMeetings
      .filter(
        (meeting) =>
          meeting.tutorId === selectedTutor && meeting.date === formattedDate
      )
      .map((meeting) => ({
        start: parseTime(meeting.startTime),
        end: parseTime(meeting.endTime),
      }));

    const availableEndTimes = endTimes.filter((timeStr) => {
      const endTimeMinutes = parseTime(timeStr);
      const isBooked = bookedTimes.some(
        (booking) =>
          startMinutes < booking.end &&
          endTimeMinutes > booking.start &&
          startMinutes < endTimeMinutes
      );
      return !isBooked;
    });

    return availableEndTimes;
  };

  const handleTutorChange = (e) => {
    const selectedTutorId = e.target.value;
    setSelectedTutor(selectedTutorId);

    const availability = tutorsAvailabilityMap[selectedTutorId] || [];
    setTutorAvailability(availability);

    setStartTime("");
    setEndTime("");
  };

  const handleScheduleMeeting = () => {
    if (!selectedTutor || !selectedDate || !startTime || !endTime) {
      alert("Please select a tutor, date, start time, and end time.");
      return;
    }

    const formattedDate = selectedDate.toISOString().split("T")[0];

    const newMeeting = {
      tutorId: selectedTutor,
      date: formattedDate,
      startTime: startTime,
      endTime: endTime,
    };

    setScheduledMeetings([...scheduledMeetings, newMeeting]);

    alert("Meeting scheduled successfully!");

    setStartTime("");
    setEndTime("");
  };

  const availableStartTimes = getAvailableStartTimes();
  const availableEndTimes = getAvailableEndTimes();

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
          onChange={(date) => {
            setSelectedDate(date);
            setSelectedTutor("");
            setStartTime("");
            setEndTime("");
            setTutorAvailability([]);
          }}
          filterDate={isWeekday}
          placeholderText="Select a date"
          className="input-field"
        />
        <select
          value={selectedTutor}
          onChange={handleTutorChange}
          className="input-field"
          disabled={!availableTutors.length}
        >
          <option value="" disabled>
            Select a Tutor
          </option>
          {availableTutors.map((tutor) => (
            <option key={tutor.userId} value={tutor.userId}>
              {tutor.fullName}
            </option>
          ))}
        </select>
        <div className="time-selection">
          <select
            value={startTime}
            onChange={(e) => {
              setStartTime(e.target.value);
              setEndTime("");
            }}
            className="input-field time-select"
            disabled={!availableStartTimes.length}
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
            className="input-field time-select"
            disabled={!availableEndTimes.length}
          >
            <option value="" disabled>
              Select end time
            </option>
            {availableEndTimes.map((time, index) => (
              <option key={index} value={time}>
                {time}
              </option>
            ))}
          </select>
        </div>
        <button
          className="button schedule-button"
          onClick={handleScheduleMeeting}
          disabled={!selectedTutor || !selectedDate || !startTime || !endTime}
        >
          Schedule Meeting
        </button>
      </div>
    </>
  );
};

export default ScheduleMeeting;
