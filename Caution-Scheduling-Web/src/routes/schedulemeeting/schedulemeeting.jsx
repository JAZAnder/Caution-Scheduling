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
  const [availableTutors, setAvailableTutors] = useState([]);
  const [selectedTutor, setSelectedTutor] = useState("");
  const [tutorAvailability, setTutorAvailability] = useState([]);
  const [tutorsAvailabilityMap, setTutorsAvailabilityMap] = useState({});
  const [hours, setHours] = useState([]);

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

        const availableHours = hours.filter(
          (hour) => availableHourIds.includes(hour.id) && hour.dayOfWeek === dayOfWeek
        );

        return availableHours.length > 0;
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

  const getAvailableHoursForTutor = () => {
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

    const hoursForDay = hours.filter((hour) => hour.dayOfWeek === dayOfWeek);

    const availableHourIds = tutorAvailability
      .filter((slot) => slot.available)
      .map((slot) => slot.hourId);

    console.log("Available Hour IDs:", availableHourIds);

    const availableHours = hoursForDay.filter((hour) =>
      availableHourIds.includes(hour.id)
    );

    console.log("Available Hours:", availableHours);

    return availableHours;
  };

  const filteredHours = getAvailableHoursForTutor();

  const handleTutorChange = (e) => {
    const selectedTutorId = e.target.value;
    setSelectedTutor(selectedTutorId);

    const availability = tutorsAvailabilityMap[selectedTutorId] || [];
    setTutorAvailability(availability);
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
          onChange={(date) => {
            setSelectedDate(date);
            setSelectedTutor("");
            setStartTime(null);
            setEndTime("");
            setEndTimeOptions([]);
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
            value={startTime || ""}
            onChange={handleTimeSlotChange}
            className="input-field time-select"
            disabled={!selectedTutor || !tutorAvailability.length || !selectedDate}
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
