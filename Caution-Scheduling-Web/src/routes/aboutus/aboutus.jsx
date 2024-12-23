import React from "react";
import "./aboutus.css";
import Placeholder from "../../assets/placeholder.jpg";
import ChaseImage from "../../assets/chase.webp";
import JoshImage from "../../assets/josh.webp";
import DimaImage from "../../assets/dima.webp";
import EthanImage from "../../assets/ethan.webp";
import Background from "../../background";
import useMediaQuery from '../../context/useMediaQuery.jsx';

const AboutUs = () => {
  const isMobile = useMediaQuery('(max-width: 900px)'); // Checks if the device is mobile
  const cards = [
    {
      name: "Joshua Cantu",
      description: "Team Leader - Senior - Information Technology",
      descriptionextended: "GoLang, MySQL, ReactJS, NGinx, Documentation",
      imgSrc: JoshImage,
    },
    {
      name: "Chase Leimbach",
      description: "Senior - Information Technology",
      descriptionextended: "HTML, CSS, ReactJS, Documentation",
      imgSrc: ChaseImage,
    },
    {
      name: "Ethan Stoulig",
      description: "Senior - Scientific",
      descriptionextended: "HTML, CSS, GoLang, Documentation",
      imgSrc: EthanImage,
    },
    {
      name: "Dmitriy Levytskyi",
      description: "Senior - Information Technology",
      descriptionextended: "ReactJS, CSS, HTML, GoLang, Documentation",
      imgSrc: DimaImage,
    },
  ];

  return (
    <>
      <Background />
      <div className="aboutus-page slide-container">
        <div className="slide-content">
          <div className={`card-wrapper ${isMobile ? "mobile" : ""}`}>
            {cards.map((card, index) => (
              <div className="card" key={index}>
                <div className="image-content">
                  <span className="overlay"></span>
                  <div className="card-image">
                    <img
                      src={card.imgSrc}
                      alt={card.name}
                      className="card-img"
                    />
                  </div>
                </div>
                <div className="card-content">
                  <h2 className="name">{card.name}</h2>
                  <p className="description">{card.description}</p>
                  <p className="descriptionextended">
                    {card.descriptionextended}
                  </p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
};

export default AboutUs;
