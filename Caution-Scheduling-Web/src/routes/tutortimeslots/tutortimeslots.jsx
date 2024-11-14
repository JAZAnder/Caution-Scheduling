import React, { useState, useEffect } from "react";
import {
  Container,
  Row,
  Col,
  Button,
  Table,
  Form,
  Modal,
} from "react-bootstrap";
import Background from "../../background";
import useFetch from "use-http";
import "./TutorTimeslots.css";

const tutortimeslots = () => {
  const { get, loading, error } = useFetch("/api/timeslots");
  const [showModal, setShowModal] = useState(false);
  const [selectedTimeslot, setSelectedTimeslot] = useState(null);
  const [timeslots, setTimeslots] = useState([]);

  useEffect(() => {
    const fetchTimeslots = async () => {
      const fetchedTimeslots = await get();
      if (fetchedTimeslots) {
        setTimeslots(fetchedTimeslots);
      }
    };
    fetchTimeslots();
  }, [get]);

  const handleEditClick = (timeslot) => {
    setSelectedTimeslot(timeslot);
    setShowModal(true);
  };

  const handleSaveChanges = () => {
    setTimeslots(
      timeslots.map((slot) =>
        slot.id === selectedTimeslot.id ? selectedTimeslot : slot
      )
    );
    setShowModal(false);
    setSelectedTimeslot(null);
  };
  return (
    <>
      <Background />
      <div className="tutor-timeslots-container">
        <div className="tutor-timeslots-container-content">
          <Container className="mt-5 tutor-timeslots-page">
            <h2 className="text-center">Tutor Timeslots Management</h2>
  
            <Row className="mt-4 justify-content-center">
              <Col xs={12}>
                <Table striped bordered hover responsive className="timeslot-table">
                  <thead>
                    <tr>
                      <th>ID</th>
                      <th>Day</th>
                      <th>Start Time</th>
                      <th>End Time</th>
                      <th>Actions</th>
                    </tr>
                  </thead>
                  <tbody>
                    {loading ? (
                      <tr>
                        <td colSpan="5" className="text-center">
                          Loading...
                        </td>
                      </tr>
                    ) : error ? (
                      <tr>
                        <td colSpan="5" className="text-center text-danger">
                          Error loading data
                        </td>
                      </tr>
                    ) : (
                      timeslots.map((slot) => (
                        <tr key={slot.id}>
                          <td>{slot.id}</td>
                          <td>{slot.dayOfWeek}</td>
                          <td>{slot.startTime}</td>
                          <td>{slot.endTime}</td>
                          <td>
                            <Button
                              variant="warning"
                              className="edit-button"
                              onClick={() => handleEditClick(slot)}
                            >
                              Edit
                            </Button>{" "}
                            <Button variant="danger" className="delete-button">
                              Delete
                            </Button>
                          </td>
                        </tr>
                      ))
                    )}
                  </tbody>
                </Table>
              </Col>
            </Row>
  
            <Modal show={showModal} onHide={() => setShowModal(false)} centered>
              <Modal.Header closeButton>
                <Modal.Title>Edit Timeslot</Modal.Title>
              </Modal.Header>
              <Modal.Body>
                {selectedTimeslot && (
                  <Form>
                    <Form.Group controlId="editDaySelect" className="mb-3">
                      <Form.Label>Day of the Week</Form.Label>
                      <Form.Control
                        as="select"
                        value={selectedTimeslot.dayOfWeek}
                        onChange={(e) =>
                          setSelectedTimeslot({
                            ...selectedTimeslot,
                            dayOfWeek: e.target.value,
                          })
                        }
                        className="form-select"
                      >
                        <option value="">Select a day</option>
                        <option value="Monday">Monday</option>
                        <option value="Tuesday">Tuesday</option>
                        <option value="Wednesday">Wednesday</option>
                        <option value="Thursday">Thursday</option>
                        <option value="Friday">Friday</option>
                      </Form.Control>
                    </Form.Group>
  
                    <Form.Group controlId="editStartTime" className="mb-3">
                      <Form.Label>Start Time</Form.Label>
                      <Form.Control
                        type="time"
                        value={selectedTimeslot.startTime}
                        onChange={(e) =>
                          setSelectedTimeslot({
                            ...selectedTimeslot,
                            startTime: e.target.value,
                          })
                        }
                        className="form-input"
                      />
                    </Form.Group>
  
                    <Form.Group controlId="editEndTime" className="mb-3">
                      <Form.Label>End Time</Form.Label>
                      <Form.Control
                        type="time"
                        value={selectedTimeslot.endTime}
                        onChange={(e) =>
                          setSelectedTimeslot({
                            ...selectedTimeslot,
                            endTime: e.target.value,
                          })
                        }
                        className="form-input"
                      />
                    </Form.Group>
                  </Form>
                )}
              </Modal.Body>
              <Modal.Footer>
                <Button
                  variant="secondary"
                  onClick={() => setShowModal(false)}
                  className="cancel-button"
                >
                  Cancel
                </Button>
                <Button
                  variant="primary"
                  onClick={handleSaveChanges}
                  className="save-button"
                >
                  Save Changes
                </Button>
              </Modal.Footer>
            </Modal>
          </Container>
        </div>
      </div>
    </>
  );    
};

export default tutortimeslots;
