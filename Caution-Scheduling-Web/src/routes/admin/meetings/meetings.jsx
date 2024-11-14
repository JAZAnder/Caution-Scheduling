import { Outlet, Link } from "react-router-dom";
import Background from "../../../background";
import './meetings.css';
import useFetch from "use-http";
import React, { useEffect, useState } from "react";
import MeetingDetailsButton from "./details/details";

function AdminMeetings() {
    const [loading, setLoading] = useState(false);
    const [userData, setUserData] = useState();
    const [filtering, setFiltering] = useState(false);
    const [tutorName, setTutorName] = useState('');
    const [studentName, setStudentName] = useState('');
    const [endTime, setEndTime] = useState('');
    const [startTime, setStartTime] = useState('');
    const [topicId, setTopicId] = useState('');
    const [dayOfWeek, setDayOfWeek] = useState('');
    const [date, setDate] = useState('');
    const [debounce, setDebounce] = useState(true);

    useEffect(() => {
        setDebounce(!debounce);
    }, []);

    useEffect(() => {
        const getData = setTimeout(() => {
            setDebounce(!debounce);
            console.log("Should Be Refresh");
        }, 100);
        return () => clearTimeout(getData);
    }, [tutorName, studentName, endTime, startTime, topicId, dayOfWeek, date]);

    const resetSearch = async (event) => {
        setTutorName('');
        setStudentName('');
        setEndTime('');
        setStartTime('');
        setTopicId('');
        setDayOfWeek('');
        setDate('');

        console.log("Should Be Refresh");
        setDebounce(!debounce);
    };

    return (
        <>

            <div className="admin-meetings-page">

 <div style={{ minHeight: "150px" }}> Black Space?</div> 
                <div id="filterOnBar">
                    <form>
                        <input
                            id="tutorName"
                            value={tutorName}
                            onChange={(e) => setTutorName(e.target.value)}
                            type="text"
                            placeholder="Tutor Name"
                        />
                        <input
                            id="studentName"
                            value={studentName}
                            onChange={(e) => setStudentName(e.target.value)}
                            type="text"
                            placeholder="Student Name"
                        />
                        <input
                            id="startTime"
                            value={startTime}
                            onChange={(e) => setStartTime(e.target.value)}
                            type="text"
                            placeholder="Start Time"
                        />
                        <input
                            id="endTime"
                            value={endTime}
                            onChange={(e) => setEndTime(e.target.value)}
                            type="text"
                            placeholder="End Time"
                        />
                        <select
                            name="topicId"
                            id="topicId"
                            onChange={(e) => setTopicId(e.target.value)}
                        >
                            <option value=""> Topic </option>

                        </select>
                        {/* <select
                            name="dayOfWeek"
                            id="dayOfWeek"
                            onChange={(e) => setDayOfWeek(e.target.value)}
                        >
                            <option value=""> Day </option>
                            <option value="1"> Monday</option>
                            <option value="2"> Tuesday</option>
                            <option value="3"> Wednesday</option>
                            <option value="4"> Thursday</option>
                            <option value="5"> Friday</option>
                            <option value="6"> Saturday</option>
                            <option value="0"> Sunday</option>
                        </select> */}

                        <input
                            id="date"
                            value={date}
                            onChange={(e) => setDate(e.target.value)}
                            type="text"
                            placeholder="Date MMDDYYYY"
                        />
                        

                        <button type="button" disabled={loading} onClick={resetSearch}>
                            {loading ? 'Waiting' : 'Reset Search'}
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
                />

            </div>
        </>
    );
}

function ListFilteredMeetings({ FLtutor, FLstudent, FLstartTime, FLendTime, FLtopicId, FLdate, FLdayOfWeek, debounce }) {
    const { data: meetings, loading, error } = useFetch(
        `/api/meetings/filter?tutor=${FLtutor}&student=${FLstudent}&startTime=${FLstartTime}&endTime=${FLendTime}&topicId=${FLtopicId}&date=${FLdate}&dayOfWeek=${FLdayOfWeek}`,
        { method: "get" },
        [debounce]
    );

    if (loading) {
        return (
            <center>
                <div className="loader"></div>
            </center>
        );
    }

    return (
        <>
            <Background />
            <div className="mymeetings-body ">
                <div id="meetingsTable">
                    <table className="table-with-bordered">
                        <thead>
                            <tr>
                                <th>Meeting Id</th>
                                <th>Topic</th>
                                <th>Student</th>
                                <th>Tutor</th>
                                <th>Date</th>
                                <th>Time</th>
                                <th>Details</th>
                            </tr>
                        </thead>
                        <tbody>
                            {meetings &&
                                Object.keys(meetings).map((meeting, i) => (
                                    <tr key={i}>
                                        <td>{meetings[meeting].id}</td>
                                        <td>{meetings[meeting].Topic.description}</td>
                                        <td>{meetings[meeting].Student.firstName + " "+ meetings[meeting].Student.lastName}</td>
                                        <td>{meetings[meeting].TutorHour.Tutor.firstName + " "+ meetings[meeting].TutorHour.Tutor.lastName}</td>
                                        <td>{meetings[meeting].date}</td>
                                        <td>{meetings[meeting].TutorHour.Hour.startTime+ " - "+meetings[meeting].TutorHour.Hour.endTime}</td>
                                        <td>
                                            <MeetingDetailsButton meeting={meetings[meeting]} />
                                        </td>
                                    </tr>
                                ))}
                        </tbody>
                    </table>
                </div>


            </div>


        </>
    )
}

export default AdminMeetings;
