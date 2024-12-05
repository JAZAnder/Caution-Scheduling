import { useEffect } from "react";
import { useState } from "react";
import { Button, Form, Modal } from "react-bootstrap";
import useFetch from "use-http";

function NewUserTimeslotButton() {
    const [show, setShow] = useState(false);

    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);

    const [userId, setUserId] = useState('')
    const [timeslotId, setTimeslotId] = useState('')
    const [dayOfWeek, setDayOfWeek] = useState('')
    const [availabilityCreate, setAvailabilityCreate] = useState(false)

    const { data: timeslots } = useFetch(
        `/api/hour/day/`+dayOfWeek,
        { method: "get" },
        [dayOfWeek]
    );

    const createNewUserAvailability = async () => {
        const myHeaders = new Headers();
        myHeaders.append('Content-Type', 'application/x-www-form-urlencoded');

        const urlencoded = new URLSearchParams();
        urlencoded.append('userId', userId);
        urlencoded.append('timeslotId', timeslotId);


        const requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: urlencoded,
            redirect: 'follow',
        };

        try {
            const response = await fetch('/api/availability', requestOptions);

            if (!response.ok) {
                throw new Error('Something Went Wrong');
            }

            const data = await response.json();
            setAvailabilityCreate(true)
            //return data;

        } catch (error) {
            alert(error)
            console.error('Error:', error);
            throw new Error('Error');
        }
    };

    const { data: users, usersLoading } = useFetch(
        `/api/lusers`,
        { method: "get" },
        []
    );
    if (usersLoading) {
        return (<><center><div className="loader"></div></center></>)
    }
    return (
        <>
            <button
                className="add-new-user-availability-button"
                onClick={handleShow}
            >
                Add new availability
            </button>

            <Modal show={show} onHide={handleClose} backdrop="static">
                <Modal.Header closeButton>
                    <Modal.Title>Create a New User Availability</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>User</strong>
                        </Form.Label>
                        <Form.Select
                            onChange={(e) => setUserId(e.target.value)}

                        >
                            <option value="null"> Pick a Tutor </option>
                            {users &&
                                Object.keys(users).map((user, i) => (
                                    <option value={users[user].userId}> {users[user].firstName + " " + users[user].lastName} </option>
                                ))

                            }
                        </Form.Select>
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Day of Week</strong>
                        </Form.Label>
                        <Form.Select aria-label="Default select example" onChange={(e) => setDayOfWeek(e.target.value)}>
                            <option value="">Select Day of Week</option>
                            <option value="1">Monday</option>
                            <option value="2">Tuesday</option>
                            <option value="3">Wednesday</option>
                            <option value="4">Thursday</option>
                            <option value="5">Friday</option>
                            <option value="6">Saturday</option>
                            <option value="0">Sunday</option>
                        </Form.Select>
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Timeslot</strong>
                        </Form.Label>
                        <Form.Select id="timeslotSelectFelid" onChange={(e) => setTimeslotId(e.target.value)}>
                            <option value="">Select Time Slot</option>
                            {timeslots &&
                                Object.keys(timeslots).map((timeslot, i) => (
                                    <option value={timeslots[timeslot].id}> {timeslots[timeslot].startTime + " - " + timeslots[timeslot].endTime} </option>
                                ))

                            }
                        </Form.Select>
                    </Form.Group>

                </Modal.Body>
                <Modal.Footer>
                    <Button variant="danger" onClick={handleClose}>
                        Close
                    </Button>
                    <Button variant="success" className="background-1" onClick={function () { createNewUserAvailability(); handleClose(); }}>
                        Create new user availability
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    )

}


export default NewUserTimeslotButton