import React from 'react';
import Background from '../../background';
import './timeslots.css';
import useMediaQuery from "../../context/useMediaQuery.jsx";
import { Container, Row, Col, Button, Form, Table } from 'react-bootstrap';

const timeslots = () => {
  const isMobile = useMediaQuery("(max-width: 900px)");

  return (
    <div className="timeslots-page">
      <Background />
      {isMobile ? (
        <Container className="mt-4">
          <Row className="justify-content-end mt-3">
            <Col xs="auto">
              <Button variant="light" className="timeslots-add-button">
                Add New Timeslot
              </Button>
            </Col>
          </Row>

          <Row className="mt-3">
            <Col xs={12}>
              <Form className="filter-form">
                <Form.Group className="mb-2">
                  <Form.Label className="filter-label">Filter on:</Form.Label>
                </Form.Group>
                <Form.Group className="mb-2">
                  <Form.Control type="text" placeholder="Id" className="filter-input" />
                </Form.Group>
                <Form.Group className="mb-2">
                  <Form.Control type="text" placeholder="DOW" className="filter-input" />
                </Form.Group>
                <Form.Group className="mb-2">
                  <Form.Control type="text" placeholder="Start" className="filter-input" />
                </Form.Group>
                <Form.Group className="mb-2">
                  <Form.Control type="text" placeholder="End" className="filter-input" />
                </Form.Group>
                <Button variant="light" className="search-button">Search</Button>
              </Form>
            </Col>
          </Row>

          <Row className="mt-4">
            <Col xs={12}>
              <Table striped bordered hover responsive className="timeslots-table">
                <thead>
                <tr className="timeslots-tablerow">
                    <th>Id</th>
                    <th>Day of Week</th>
                    <th>Start Time</th>
                    <th>End Time</th>
                    <th>Details</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td></td>
                    <td></td>
                    <td></td>
                    <td></td>
                    <td><Button variant="info" className="details-button">Details</Button></td>
                  </tr>
                </tbody>
              </Table>
            </Col>
          </Row>
        </Container>
      ) : (
        <div className="timeslots-outer-container">
          <div className="timeslots-container">
            <div className="timeslots-header">
              <button className="timeslots-add-button">Add New Timeslot</button>
            </div>

            <div className="timeslots-filter">
              <label className="timeslots-filter-label">Filter on:</label>
              <input type="text" placeholder="Id" className="timeslots-filter-input" />
              <input type="text" placeholder="DOW" className="timeslots-filter-input" />
              <input type="text" placeholder="Start" className="timeslots-filter-input" />
              <input type="text" placeholder="End" className="timeslots-filter-input" />
              <button className="timeslots-search-button">Search</button>
            </div>

            <table className="timeslots-table">
              <thead>
              <tr className="timeslots-tablerow">
                  <th>Id</th>
                  <th>Day of Week</th>
                  <th>Start Time</th>
                  <th>End Time</th>
                  <th>Details</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td><button className="timeslots-details-button">Details</button></td>
                </tr>
                <tr>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td><button className="timeslots-details-button">Details</button></td>
                </tr>
                <tr>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td><button className="timeslots-details-button">Details</button></td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      )}
    </div>
  );
};

export default timeslots;
