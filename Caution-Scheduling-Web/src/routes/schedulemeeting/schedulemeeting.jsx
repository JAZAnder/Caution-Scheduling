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
  const [availableTutors, setAvailableTutors] = useState([]);
  const [selectedTutor, setSelectedTutor] = useState("");
  const [tutorAvailability, setTutorAvailability] = useState([]);
  const [scheduledMeetings, setScheduledMeetings] = useState([]);

  const formatSelectedDate = () => {
    if (!selectedDate) return "";
    const day = String(selectedDate.getDate()).padStart(2, "0");
    const month = String(selectedDate.getMonth() + 1).padStart(2, "0");
    const year = selectedDate.getFullYear();
    return `${month}${day}${year}`;
  };

  useEffect(() => {
    if (!selectedDate) return;

    const fetchAvailableTutors = async () => {
      try {
        const formattedDate = formatSelectedDate();
        const response = await axios.get(`/api/availability/${formattedDate}`);
        const uniqueTutors = Array.from(new Set(response.data.map(tutor => tutor.tutor)))
          .map(id => response.data.find(tutor => tutor.tutor === id));
        setAvailableTutors(uniqueTutors);
      } catch (error) {
        console.error("Error fetching available tutors for the date:", error);
      }
    };

    fetchAvailableTutors();
    setSelectedTutor("");
    setTutorAvailability([]);
    setScheduledMeetings([]);
    setStartTime("");
    setEndTime("");
  }, [selectedDate]);

  useEffect(() => {
    if (!selectedTutor || !selectedDate) return;

    const fetchTutorAvailability = async () => {
      try {
        const formattedDate = formatSelectedDate();
        const response = await axios.get(
          `/api/availability/${selectedTutor}/${formattedDate}`
        );
        const slots = response.data;
        const selectedDayOfWeek = selectedDate.getDay();
        const filteredSlots = slots.filter(
          (slot) => parseInt(slot.dayOfWeek) === selectedDayOfWeek
        );
        const mergedSlots = mergeAvailabilitySlots(filteredSlots);
        setTutorAvailability(mergedSlots);
      } catch (error) {
        console.error("Error fetching tutor's availability:", error);
      }
    };

    fetchTutorAvailability();
    setStartTime("");
    setEndTime("");
  }, [selectedTutor, selectedDate]);

  useEffect(() => {
    if (!selectedTutor || !selectedDate) return;

    const fetchScheduledMeetings = async () => {
      try {
        const formattedDate = selectedDate.toISOString().split("T")[0];
        const response = await axios.get(
          `/api/scheduledMeetings/${selectedTutor}/${formattedDate}`
        );
        const meetings = response.data;
        const selectedDayOfWeek = selectedDate.getDay();
        const filteredMeetings = meetings.filter(
          (meeting) => {
            const meetingDate = new Date(meeting.date);
            return meetingDate.getDay() === selectedDayOfWeek;
          }
        );

        setScheduledMeetings(filteredMeetings);
      } catch (error) {
        console.error("Error fetching scheduled meetings:", error);
      }
    };

    fetchScheduledMeetings();
  }, [selectedTutor, selectedDate]);

  const parseTime = (timeStr) => {
    timeStr = timeStr.trim();
    let time, modifier;
    const regex = /(\d{1,2}:\d{2})\s*([AP]M)/i;
    const match = timeStr.match(regex);
    if (match) {
      time = match[1];
      modifier = match[2].toUpperCase();
    } else {
      console.error(`Invalid time format: ${timeStr}`);
      return null;
    }
    let [hours, minutes] = time.split(":").map(Number);
    if (modifier === "PM" && hours !== 12) hours += 12;
    if (modifier === "AM" && hours === 12) hours = 0;
    return hours * 60 + minutes;
  };

  const formatTime = (totalMinutes) => {
    let hours = Math.floor(totalMinutes / 60);
    const minutes = totalMinutes % 60;
    const ampm = hours >= 12 ? "PM" : "AM";
    const displayHours = hours % 12 || 12;
    return `${displayHours}:${minutes < 10 ? "0" : ""}${minutes} ${ampm}`;
  };

  const mergeAvailabilitySlots = (slots) => {
    if (!slots.length) return [];
    slots.sort((a, b) => parseTime(a.startTime) - parseTime(b.startTime));
    const mergedSlots = [];
    let currentSlot = { ...slots[0] };

    for (let i = 1; i < slots.length; i++) {
      const nextSlot = slots[i];
      const currentEnd = parseTime(currentSlot.endTime);
      const nextStart = parseTime(nextSlot.startTime);

      if (currentEnd === nextStart) {
        currentSlot.endTime = nextSlot.endTime;
      } else {
        mergedSlots.push(currentSlot);
        currentSlot = { ...nextSlot };
      }
    }

    mergedSlots.push(currentSlot);
    return mergedSlots;
  };

  const isTimeSlotAvailable = (startMinutes, endMinutes) => {
    const formattedDate = selectedDate.toISOString().split("T")[0];

    const meetingsOnDate = scheduledMeetings.filter(
      (meeting) =>
        meeting.tutorId === selectedTutor &&
        meeting.date === formattedDate
    );

    for (let meeting of meetingsOnDate) {
      const meetingStart = parseTime(meeting.startTime);
      const meetingEnd = parseTime(meeting.endTime);
      if (
        startMinutes < meetingEnd && endMinutes > meetingStart
      ) {
        return false;
      }
    }

    return true;
  };

  const getAvailableStartTimes = () => {
    if (!tutorAvailability.length) return [];

    let startTimes = new Set();

    tutorAvailability.forEach((slot) => {
      const slotStart = parseTime(slot.startTime);
      const slotEnd = parseTime(slot.endTime);

      for (let time = slotStart; time <= slotEnd - 15; time += 15) {
        let isAvailable = true;

        for (let meeting of scheduledMeetings) {
          const meetingStart = parseTime(meeting.startTime);
          const meetingEnd = parseTime(meeting.endTime);
          if (
            time < meetingEnd && (time + 15) > meetingStart
          ) {
            isAvailable = false;
            break;
          }
        }

        if (isAvailable) {
          startTimes.add(formatTime(time));
        }
      }
    });

    const startTimesArray = Array.from(startTimes).sort((a, b) => parseTime(a) - parseTime(b));

    return startTimesArray;
  };

  const getAvailableEndTimes = () => {
    if (!startTime || !tutorAvailability.length) return [];

    const selectedStartTime = parseTime(startTime);

    const currentSlot = tutorAvailability.find((slot) => {
      const start = parseTime(slot.startTime);
      const end = parseTime(slot.endTime);
      return selectedStartTime >= start && selectedStartTime < end;
    });

    if (!currentSlot) return [];

    const slotEnd = parseTime(currentSlot.endTime);
    let endTimes = [];
    let proposedEndTime = selectedStartTime + 15;

    while (
      proposedEndTime <= selectedStartTime + 45 &&
      proposedEndTime <= slotEnd
    ) {
      if (isTimeSlotAvailable(selectedStartTime, proposedEndTime)) {
        endTimes.push(formatTime(proposedEndTime));
      } else {
        break;
      }
      proposedEndTime += 15;
    }

    return endTimes;
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

    const startMinutes = parseTime(startTime);
    const endMinutes = parseTime(endTime);
    if (!isTimeSlotAvailable(startMinutes, endMinutes)) {
      alert("Selected time slot is no longer available.");
      return;
    }

    const saveMeeting = async () => {
      try {
        await axios.post('/api/meeting', newMeeting);
        setScheduledMeetings([...scheduledMeetings, newMeeting]);
        alert("Meeting scheduled successfully!");
        setStartTime("");
        setEndTime("");
      } catch (error) {
        console.error("Error scheduling meeting:", error);
        alert("Failed to schedule meeting. Please try again.");
      }
    };

    saveMeeting();
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
            setStartTime("");
            setEndTime("");
          }}
          filterDate={(date) => {
            const day = date.getDay();
            return day >= 1 && day <= 4;
          }}
          placeholderText="Select a date"
          className="input-field"
        />
        <select
          value={selectedTutor}
          onChange={(e) => setSelectedTutor(e.target.value)}
          className="input-field"
          disabled={!availableTutors.length}
        >
          <option value="" disabled>
            Select a Tutor
          </option>
          {availableTutors.map((tutor) => (
            <option key={tutor.tutor} value={tutor.tutor}>
              {tutor.firstName} {tutor.lastName}
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
            disabled={!getAvailableStartTimes().length}
          >
            <option value="" disabled>
              Select start time
            </option>
            {getAvailableStartTimes().map((time, index) => (
              <option key={index} value={time}>
                {time}
              </option>
            ))}
          </select>
          <select
            value={endTime}
            onChange={(e) => setEndTime(e.target.value)}
            className="input-field time-select"
            disabled={!getAvailableEndTimes().length}
          >
            <option value="" disabled>
              Select end time
            </option>
            {getAvailableEndTimes().map((time, index) => (
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
