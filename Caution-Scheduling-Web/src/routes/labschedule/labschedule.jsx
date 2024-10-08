import React from "react";
import "./labschedule.css";
import EnvocLab from "../../assets/EnvocLab.png";
import FayLab from "../../assets/Fay125Lab.png";
import Background from "../../background";
import { Container, Row, Col, Button } from "react-bootstrap";
import useMediaQuery from "../../context/useMediaQuery.jsx";

const LabSchedule = () => {
  const isMobile = useMediaQuery("(max-width: 900px)");

  return (
    <div className="lab-schedule-page">
      <Background />
      {isMobile ? (
        // React Bootstrap layout for mobile
        <Container className="mobile-lab-schedule">
          <div className="mobile-lab-content">
            <img
              src={EnvocLab}
              alt="Envoc Lab"
              className="img-fluid mobile-lab-image"
            />
            <a
              href="http://meet.google.com/cgc-diaj-gnx"
              className="d-block mt-2"
            >
              <Button
                variant="primary"
                style={{ backgroundColor: "#1a5632", color: "white" }}
              >
                Envoc Google Link
              </Button>
            </a>
          </div>
          <div className="mobile-lab-content">
            <img
              src={FayLab}
              alt="Fay 125 Lab"
              className="img-fluid mobile-lab-image"
            />
            <a
              href="http://meet.google.com/fda-vanx-mnd"
              className="d-block mt-2"
            >
              <Button
                variant="primary"
                style={{ backgroundColor: "#1a5632", color: "white" }}
              >
                Fayard Google Link
              </Button>
            </a>
          </div>
        </Container>
      ) : (
        // Custom CSS layout for desktop
        <>
          <div className="lab-schedule-container1">
            <div className="lab-schedule-content">
              <img
                src={EnvocLab}
                alt="Envoc Lab"
                className="lab-schedule-image"
              />
              <a
                href="http://meet.google.com/cgc-diaj-gnx"
                className="labschedulebutton-link"
              >
                <button className="labschedulebutton">Envoc Google Link</button>
              </a>
            </div>
          </div>
          <div className="lab-schedule-container2">
            <div className="lab-schedule-content">
              <img
                src={FayLab}
                alt="Fay 125 Lab"
                className="lab-schedule-image"
              />
              <a
                href="http://meet.google.com/fda-vanx-mnd"
                className="labschedulebutton-link"
              >
                <button className="labschedulebutton">
                  Fayard Google Link
                </button>
              </a>
            </div>
          </div>
        </>
      )}
    </div>
  );
};

export default LabSchedule;
