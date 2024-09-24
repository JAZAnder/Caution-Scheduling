import React from "react";
import "./labschedule.css";
import EnvocLab from "../../assets/EnvocLab.png";
import FayLab from "../../assets/Fay125Lab.png";

const LabSchedule = () => {
  return (
    <div className="lab-schedule-page">
      <div className="lab-schedule-container1">
        <img src={EnvocLab} alt="Envoc Lab" className="lab-schedule-image" />
      </div>
      <div className="lab-schedule-container2">
        <img src={FayLab} alt="Fay 125 Lab" className="lab-schedule-image" />
      </div>
    </div>
  );
};

export default LabSchedule;
