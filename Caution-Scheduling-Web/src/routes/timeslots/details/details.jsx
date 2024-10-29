import { useEffect } from "react";
import { useState } from "react";
import { Button, Form, Modal } from "react-bootstrap";
import "../timeslots.css"

function TimeslotDetailsButton(timeslot) {
    const [show, setShow] = useState(false);

    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);

    const [startTime, setStartTime] = useState(timeslot.timeslot.startTime)
    const [endTime, setEndTime] = useState(timeslot.timeslot.endTime)
    const [dayOfWeek, setDayOfWeek] = useState(timeslot.timeslot.dayOfWeek)


    return (
        <>
            <button
                className="timeslots-details-button"
                onClick={handleShow}
            >
                Details
            </button>

            <Modal show={show} onHide={handleClose} backdrop="static">
                <Modal.Header closeButton>
                    <Modal.Title>Timeslot Information for Timeslot #{timeslot.timeslot.id}</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Start Time</strong>
                        </Form.Label>
                        <Form.Control
                            type="text"
                            value={startTime}
                            onChange={(e) => setStartTime(e.target.value)}
                            placeholder="startTime"
                            disabled={true}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>End Time</strong>
                        </Form.Label>
                        <Form.Control
                            type="text"
                            value={endTime}
                            onChange={(e) => setEndTime(e.target.value)}
                            placeholder="endTime"
                            disabled={true}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Day Of Week</strong>
                        </Form.Label>
                        <Form.Control
                            type="text"
                            value={dayOfWeek}
                            onChange={(e) => setDayOfWeek(e.target.value)}
                            placeholder="Day of Week"
                            disabled={true}
                        />
                    </Form.Group>
                </Modal.Body>
                <Modal.Footer>
                <Button variant="secondary" onClick={handleClose}>
                        Close
                    </Button>
                    <Button variant="primary" onClick={handleClose}>
                        Tutors
                    </Button>
                    
                    <Button variant="danger" onClick={handleClose}>
                        Delete Timeslot
                    </Button>
                </Modal.Footer>
            </Modal>
        </>    
        )

}


export default TimeslotDetailsButton