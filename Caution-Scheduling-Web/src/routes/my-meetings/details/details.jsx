import { useEffect } from "react";
import { useState } from "react";
import { Button, Form, Modal } from "react-bootstrap";
import "../my-meetings.css"

function MeetingDetailsButton(meeting) {
    const [show, setShow] = useState(false);

    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);




    return (
        <>
            <button
                className="details-details-button"
                onClick={handleShow}
            >
                Details
            </button>

            <Modal show={show} onHide={handleClose} backdrop="static">
                <Modal.Header closeButton>
                    <Modal.Title>User Information for Meeting #{meeting.meeting.id}</Modal.Title>

                </Modal.Header>
                <Modal.Body>
                    <table>
                        <tr>
                            <td>
                                <td>
                                    <Form.Group className="mb-3">
                                        <Form.Label>
                                            <strong>Subject</strong>
                                        </Form.Label>
                                        <Form.Select aria-label="Default select example">
                                            <option value={meeting.meeting.Topic.id}>{meeting.meeting.Topic.description}</option>
                                        </Form.Select>

                                    </Form.Group>
                                </td>
                            </td>
                            <td>
                                <Form.Group className="mb-3">
                                    <Form.Label>
                                        <strong>Date</strong>
                                    </Form.Label>
                                    <Form.Control
                                        type="text"
                                        value={meeting.meeting.date}
                                        placeholder="Last Name"
                                    />
                                </Form.Group>
                            </td>
                        </tr>
                    </table>


                    <Form.Group>
                        <Form.Control
                            type="text"
                            value={meeting.meeting.Hour.startTime + " - " + meeting.meeting.Hour.endTime}
                            placeholder="Last Name"
                            disabled={true}

                        />
                    </Form.Group>





                    <table >
                        <tr>
                            <th><h4>Student Information</h4></th>
                            <th><h4>Tutor Information</h4></th>
                        </tr>
                        <tr>
                            <td>
                                <Form.Group className="mb-3">
                                    <Form.Label>
                                        <strong>Name</strong>
                                    </Form.Label>
                                    <Form.Control
                                        type="text"
                                        value={meeting.meeting.Student.firstName + " " + meeting.meeting.Student.lastName}
                                        placeholder="Student Name"
                                        disabled={true}
                                    />
                                </Form.Group>
                                <Form.Group className="mb-3">
                                    <Form.Label>
                                        <strong>Email</strong>
                                    </Form.Label>
                                    <Form.Control
                                        type="text"
                                        value={meeting.meeting.Student.email}
                                        placeholder="Student Email"
                                        disabled={true}
                                    />
                                </Form.Group>
                            </td>
                            <td>
                                <Form.Group className="mb-3">
                                    <Form.Label>
                                        <strong>Name</strong>
                                    </Form.Label>
                                    <Form.Control
                                        type="text"
                                        value={meeting.meeting.Tutor.firstName + " " + meeting.meeting.Tutor.lastName}
                                        placeholder="Tutor Name"
                                        disabled={true}
                                    />
                                </Form.Group>
                                <Form.Group className="mb-3">
                                    <Form.Label>
                                        <strong>Email</strong>
                                    </Form.Label>
                                    <Form.Control
                                        type="text"
                                        value={meeting.meeting.Tutor.email}
                                        placeholder="Tutor Email"
                                        disabled={true}
                                    />
                                </Form.Group>
                            </td>
                        </tr>
                    </table>


                    {/* Notes Table */}
                    
                    <h4>Meeting Notes</h4>
                    <table>
                        <tr><th>Student's Name</th></tr>
                        <tr><td>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. </td></tr>
                        <tr><th>Tutor's Name</th></tr>
                        <tr><td>UUt enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</td></tr>
                    </table>

                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>New Note</strong>
                        </Form.Label>
                        <Form.Control
                            type="text"
                            placeholder="New Note"
                        />
                    </Form.Group>
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={handleClose}>
                        Close
                    </Button>
                    <Button variant="success" onClick={handleClose}>
                        Add Note
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    )

}


export default MeetingDetailsButton