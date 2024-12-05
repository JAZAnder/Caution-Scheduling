import React, { useState, useEffect } from "react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import Background from "../../background";
import "./schedulemeeting.css";
import axios from "axios";
import { toast, ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import qs from 'qs';

const ScheduleMeeting = ({ isAdmin }) => {
  const [selectedDate, setSelectedDate] = useState(null);
  const [startTime, setStartTime] = useState("");
  const [endTime, setEndTime] = useState("");
  const [availableTutors, setAvailableTutors] = useState([]);
  const [selectedTutor, setSelectedTutor] = useState("");
  const [tutorAvailability, setTutorAvailability] = useState([]);
  const [scheduledMeetings, setScheduledMeetings] = useState([]);
  const [notificationMessage, setNotificationMessage] = useState('');

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
        const uniqueTutors = Array.from(
          new Set(response.data.map((tutor) => tutor.tutor))
        ).map((id) => response.data.find((tutor) => tutor.tutor === id));
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
        setTutorAvailability(filteredSlots);
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
        const formattedDate = formatSelectedDate();
        const response = await axios.get(
          `/api/scheduledMeetings/${selectedTutor}/${formattedDate}`
        );
        const meetings = response.data;
        setScheduledMeetings(meetings);
      } catch (error) {
        console.error("Error fetching scheduled meetings:", error);
      }
    };

    fetchScheduledMeetings();
  }, [selectedTutor, selectedDate]);

  const parseTime = (timeStr) => {
    const [time, modifier] = timeStr.split(" ");
    let [hours, minutes] = time.split(":").map(Number);

    if (modifier.toUpperCase() === "PM" && hours !== 12) {
      hours += 12;
    }
    if (modifier.toUpperCase() === "AM" && hours === 12) {
      hours = 0;
    }
    return hours * 60 + minutes;
  };

  const formatTime = (totalMinutes) => {
    let hours = Math.floor(totalMinutes / 60);
    const minutes = totalMinutes % 60;
    const ampm = hours >= 12 ? "PM" : "AM";
    hours = hours % 12 || 12;
    return `${hours}:${minutes.toString().padStart(2, "0")} ${ampm}`;
  };

  const isTimeSlotAvailable = (startMinutes, endMinutes) => {
    for (let time = startMinutes; time < endMinutes; time += 15) {
      const timeStr = formatTime(time);
      const slot = tutorAvailability.find(
        (slot) => slot.startTime === timeStr
      );
      if (!slot) {
        return false;
      }
      const isBooked = scheduledMeetings.some(
        (meeting) => meeting.tutorHourId === slot.id
      );
      if (isBooked) {
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
      const timeStr = formatTime(slotStart);
      const isBooked = scheduledMeetings.some(
        (meeting) =>
          meeting.tutorHourId === slot.id && slot.startTime === timeStr
      );
      if (!isBooked) {
        startTimes.add(timeStr);
      }
    });

    let startTimesArray = Array.from(startTimes).sort(
      (a, b) => parseTime(a) - parseTime(b)
    );

    if (selectedDate) {
      const today = new Date();
      const selected = new Date(selectedDate);
      if (
        selected.getFullYear() === today.getFullYear() &&
        selected.getMonth() === today.getMonth() &&
        selected.getDate() === today.getDate()
      ) {
        const currentMinutes =
          today.getHours() * 60 + today.getMinutes() + 15;
        startTimesArray = startTimesArray.filter(
          (time) => parseTime(time) >= currentMinutes
        );
      }
    }

    return startTimesArray;
  };

  const getAvailableEndTimes = () => {
    if (!startTime || !tutorAvailability.length) return [];

    const selectedStartTime = parseTime(startTime);

    const currentSlot = tutorAvailability.find((slot) => {
      const start = parseTime(slot.startTime);
      return selectedStartTime === start;
    });

    if (!currentSlot) return [];

    let endTimes = [];

    for (let increment = 15; increment <= 45; increment += 15) {
      const proposedEndTime = selectedStartTime + increment;
      if (isTimeSlotAvailable(selectedStartTime, proposedEndTime)) {
        endTimes.push(formatTime(proposedEndTime));
      } else {
        break;
      }
    }

    return endTimes;
  };

  const handleScheduleMeeting = () => {
    if (!selectedTutor || !selectedDate || !startTime || !endTime) {
      toast.error("Please select a tutor, date, start time, and end time.", {
        position: "top-right",
        icon: "❌",
      });
      return;
    }

    const today = new Date();
    const selected = new Date(selectedDate);
    const isToday =
      selected.getFullYear() === today.getFullYear() &&
      selected.getMonth() === today.getMonth() &&
      selected.getDate() === today.getDate();

    if (isToday) {
      const currentMinutes = today.getHours() * 60 + today.getMinutes();
      const selectedStartMinutes = parseTime(startTime);
      if (selectedStartMinutes < currentMinutes) {
        toast.error("Cannot schedule a meeting in the past.", {
          position: "top-right",
          icon: "❌",
        });
        return;
      }
    }

    const formattedDate = formatSelectedDate();
    const startMinutes = parseTime(startTime);
    const endMinutes = parseTime(endTime);

    if (!isTimeSlotAvailable(startMinutes, endMinutes)) {
      toast.error("Selected time slot is no longer available.", {
        position: "top-right",
        icon: "❌",
      });
      return;
    }

    const timeSlots = [];
    for (let time = startMinutes; time < endMinutes; time += 15) {
      timeSlots.push(time);
    }

    const tutorHourIds = [];
    for (let time of timeSlots) {
      const timeStr = formatTime(time);
      const slot = tutorAvailability.find(
        (slot) => slot.startTime === timeStr
      );
      if (slot) {
        tutorHourIds.push(slot.id);
      } else {
        setNotificationMessage(`No availability for time ${timeStr}`);
        return;
      }
    }

    const saveMeeting = async () => {
      try {
        for (let tutorHourId of tutorHourIds) {
          const newMeeting = {
            userHourId: tutorHourId.toString(),
            date: formattedDate,
          };
          console.log("Sending data to API:", newMeeting);

          const response = await axios.post(
            '/api/meeting',
            qs.stringify(newMeeting),
            {
              headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
              },
            }
          );
          console.log("API response:", response.data);
        }

        toast.success("Meeting successfully scheduled!", {
          position: "top-right",
          icon: "✅",
        });
        setScheduledMeetings([
          ...scheduledMeetings,
          ...tutorHourIds.map((id) => ({
            tutorHourId: id,
            date: formattedDate,
          })),
        ]);
        setStartTime("");
        setEndTime("");
      } catch (error) {
        console.error("Error scheduling meeting:", error);
        toast.error("Failed to schedule meeting. Please try again.", {
          position: "top-right",
          icon: "❌",
        });
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
          minDate={new Date()}
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
          disabled={
            !selectedTutor ||
            !selectedDate ||
            !startTime ||
            !endTime ||
            (selectedDate && (() => {
              const today = new Date();
              const selected = new Date(selectedDate);
              selected.setHours(0, 0, 0, 0);
              today.setHours(0, 0, 0, 0);
              return selected < today;
            })())
          }
        >
          Schedule Meeting
        </button>
      </div>
      <ToastContainer />
    </>
  );  
};

export default ScheduleMeeting;
