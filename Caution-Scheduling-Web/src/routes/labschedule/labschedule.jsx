import React from "react";
import "./labschedule.css";
import EnvocLab from "../../assets/EnvocLab.png";
import FayLab from "../../assets/Fay125Lab.png";
import Background from "../../background";

const LabSchedule = () => {
  return (
    <div className="lab-schedule-page">
      <Background />
      <div className="lab-schedule-container1">
        <div className="lab-schedule-content">
          <img src={EnvocLab} alt="Envoc Lab" className="lab-schedule-image" />
          <a href="http://meet.google.com/cgc-diaj-gnx" className="labschedulebutton-link">
            <button className="labschedulebutton">Envoc Google Link</button>
          </a>
        </div>
      </div>
      <div className="lab-schedule-container2">
        <div className="lab-schedule-content">
          <img src={FayLab} alt="Fay 125 Lab" className="lab-schedule-image" />
          <a href="http://meet.google.com/fda-vanx-mnd" className="labschedulebutton-link">
            <button className="labschedulebutton">Fayard Google Link</button>
          </a>
        </div>
      </div>
    </div>
  );
};

export default LabSchedule;
