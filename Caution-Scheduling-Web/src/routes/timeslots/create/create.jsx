import { useEffect } from "react";
import { useState } from "react";
import { Button, Form, Modal } from "react-bootstrap";
import "../timeslots.css"


function NewTimeslotButton() {
    const [show, setShow] = useState(false);

    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);

    const [startTime, setStartTime] = useState('')
    const [endTime, setEndTime] = useState('')
    const [monday, setMonday] = useState(false)
    const [tuesday, setTuesday] = useState(false)
    const [wednesday, setWednesday] = useState(false)
    const [thursday, setThursday] = useState(false)
    const [friday, setFriday] = useState(false)
    const [saturday, setSaturday] = useState(false)
    const [sunday, setSunday] = useState(false)

    const createNewTimeslots = async () => {
        const myHeaders = new Headers();
        myHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
      
        const urlencoded = new URLSearchParams();
        urlencoded.append('startTime', startTime);
        urlencoded.append('endTime', endTime);
        urlencoded.append('Monday', monday);
        urlencoded.append('Tuesday', tuesday);
        urlencoded.append('Wednesday', wednesday);
        urlencoded.append('Thursday', thursday);
        urlencoded.append('Friday', friday);
        urlencoded.append('Saturday', saturday);
        urlencoded.append('Sunday', sunday);
      
        const requestOptions = {
          method: 'POST',
          headers: myHeaders,
          body: urlencoded,
          redirect: 'follow',
        };
      
        try {
          const response = await fetch('/api/hours', requestOptions);
      
          if (!response.ok) {
            throw new Error('Something Went Wrong');
          }
      
          const data = await response.json();
          //return data;

        } catch (error) {
            alert(error)
          console.error('Error:', error);
          throw new Error('Error');
        }
      };



    return (
        <>
            <button
                className="timeslots-add-button"
                onClick={handleShow}
            >
                Add New Timeslots
            </button>

            <Modal show={show} onHide={handleClose} backdrop="static">
                <Modal.Header closeButton>
                    <Modal.Title>Create a New Timeslots</Modal.Title>
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
                            placeholder="Start Time"
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
                            placeholder="End Time"
                        />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>
                            <strong>Days of the Week</strong>
                        </Form.Label>
                        <Form.Check
                            inline
                            label="Monday"
                            name="monday"
                            type="checkbox"
                            id="monday"
                            value={monday}
                            onChange={(e) => setMonday(!monday)}
                        />
                        <Form.Check
                            inline
                            label="Tuesday"
                            name="tuesday"
                            type="checkbox"
                            id="tuesday"
                            value={tuesday}
                            onChange={(e) => setTuesday(!tuesday)}
                        />
                        <Form.Check
                            inline
                            label="Wednesday"
                            name="wednesday"
                            type="checkbox"
                            id="wednesday"
                            value={wednesday}
                            onChange={(e) => setWednesday(!wednesday)}
                        />
                        <Form.Check
                            inline
                            label="Thursday"
                            name="thursday"
                            type="checkbox"
                            id="thursday"
                            value={thursday}
                            onChange={(e) => setThursday(!thursday)}
                        />
                        <Form.Check
                            inline
                            label="Friday"
                            name="friday"
                            type="checkbox"
                            id="friday"
                            value={friday}
                            onChange={(e) => setFriday(!friday)}
                        />
                        <Form.Check
                            inline
                            label="Saturday"
                            name="saturday"
                            type="checkbox"
                            id="saturday"
                            value={saturday}
                            onChange={(e) => setSaturday(!saturday)}
                        />
                        <Form.Check
                            inline
                            label="Sunday"
                            name="sunday"
                            type="checkbox"
                            id="sunday"
                            value={sunday}
                            onChange={(e) => setSunday(!sunday)}
                        />

                    </Form.Group>
                    
                    
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="danger" onClick={handleClose}>
                        Close
                    </Button>
                    <Button variant="success" className="background-1" onClick={function () {createNewTimeslots(); handleClose();  }}>
                        Create new Timeslots
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    )

}


export default NewTimeslotButton