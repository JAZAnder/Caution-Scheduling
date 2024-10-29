import React, { useEffect, useState } from "react";
import Background from '../../background';
import './timeslots.css';
import useFetch from "use-http";
import useMediaQuery from "../../context/useMediaQuery.jsx";
import { Container, Row, Col, Button, Form, Table } from 'react-bootstrap';
import TimeslotDetailsButton from "./details/details"
import NewTimeslotButton from "./create/create"

const timeslots = () => {
  const isMobile = useMediaQuery("(max-width: 900px)");


  const [loading, setLoading] = useState(false);
  const [dayOfWeek, setDayOfWeek] = useState('')
  const [startTime, setStartTime] = useState('')
  const [endTime, setEndTime] = useState('')
  const [debounce, setDebounce] = useState(true);

  useEffect(() => {
    setDebounce(!debounce);
  }, []);

  React.useEffect(() => {
    const getData = setTimeout(() => {
      setDebounce(!debounce);
    }, 100);
    return () => clearTimeout(getData);
  }, [dayOfWeek, startTime, endTime]);



  const resetSearch = async (event) => {
    setDayOfWeek('');
    setStartTime('');
    setEndTime('');
    setDebounce(!debounce);
  };


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
                  <Form.Control
                    id='id' 
                    type="text" 
                    placeholder="Id" 
                    className="filter-input"
                  />
                </Form.Group>
                <Form.Group className="mb-2">
                  <Form.Control 
                    id='startTime'
                    value={startTime}
                    onChange={(e) => setStartTime(e.target.value)}
                    type="text" 
                    placeholder="Start Time" 
                    className="filter-input" 
                  />
                </Form.Group>
                <Form.Group className="mb-2">
                  <Form.Control 
                    id='endTime'
                    value={endTime}
                    onChange={(e) => setEndTime(e.target.value)}
                    type="text" 
                    placeholder="End Time" 
                    className="filter-input"  
                  />
                </Form.Group>
                <Form.Group className="mb-2">
                  <Form.Control 
                    type="text" 
                    placeholder="DOW" 
                    className="filter-input" 
                  />
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
              <NewTimeslotButton/>
            </div>

            <div className="timeslots-filter">
              <label className="timeslots-filter-label">Filter on:</label>
              <input 
                type="text" 
                placeholder="Id" 
                className="timeslots-filter-input" 
              />
              <input 
                id='startTime'
                value={startTime}
                onChange={(e) => setStartTime(e.target.value)}
                type="text" 
                placeholder="Start Time" 
                className="timeslots-filter-input" 
              />
              <input 
                id='endTime'
                value={endTime}
                onChange={(e) => setEndTime(e.target.value)}
                type="text" 
                placeholder="End Time" 
                className="timeslots-filter-input"  
              />
              <select 
                className="timeslots-filter-input" 
                name="dayOfWeek" 
                id="dayOfWeek"
                onChange={(e) => setDayOfWeek(e.target.value)}
              > 
                <option value=""> Day of Week </option>
                <option value="1"> Monday</option>
                <option value="2"> Tuesday</option>
                <option value="3"> Wednesday</option>
                <option value="4"> Thursday</option>
                <option value="5"> Friday</option>
                <option value="6"> Saturday</option>
                <option value="0"> Sunday</option>
                
              </select>
              

              <button
                type="button" 
                disabled={loading} 
                onClick={resetSearch}
                className="timeslots-search-button"
              >
                  {loading ? 'Waiting' : 'Reset Search'}
              </button>
            </div>

            <ListFilteredTimeSlots
              FLdayOfWeek={dayOfWeek}
              FLstartTime={startTime}
              FLendTime={endTime}
              debounce={debounce}
            />


            
          </div>
        </div>
      )}
    </div>
  );
};


function ListFilteredTimeSlots({FLdayOfWeek, FLstartTime, FLendTime, debounce}){
  const { data: timeSlots, loading, error } = useFetch(
    `/api/timeslots`,
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
  return(
    <>
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
                {timeSlots &&
                  Object.keys(timeSlots).map((timeslot, i) =>(
                    <tr key={i}>
                      <td>{timeSlots[timeslot].id}</td>
                      <td>{timeSlots[timeslot].dayOfWeek}</td>
                      <td>{timeSlots[timeslot].startTime}</td>
                      <td>{timeSlots[timeslot].endTime}</td>
                      <td><TimeslotDetailsButton timeslot={timeSlots[timeslot]}/></td>
                    </tr>
                  ))

                }
              </tbody>
            </table>
    </>
  )
}

export default timeslots;
