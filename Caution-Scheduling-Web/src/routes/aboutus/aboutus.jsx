import React from "react";
import "./aboutus.css";
import Placeholder from "../../assets/placeholder.jpg";
import Background from "../../background";
import useMediaQuery from '../../context/useMediaQuery.jsx';

const AboutUs = () => {
  const isMobile = useMediaQuery('(max-width: 900px)'); // Checks if the device is mobile
  const cards = [
    {
      name: "Joshua Cantu",
      description: "Team Leader - Senior - Information Technology",
      descriptionextended: "Things about you! Placeholders to fill",
      imgSrc: Placeholder,
    },
    {
      name: "Chase Leimbach",
      description: "Senior - Information Technology",
      descriptionextended: "Things about you! Placeholders to fill",
      imgSrc: Placeholder,
    },
    {
      name: "Ethan Stoulig",
      description: "Senior - Scientific",
      descriptionextended: "Things about you! Placeholders to fill",
      imgSrc: Placeholder,
    },
    {
      name: "Dmitriy Levytskyi",
      description: "Senior - Information Technology",
      descriptionextended: "Things about you! Placeholders to fill",
      imgSrc: Placeholder,
    },
  ];

  return (
    <>
      <Background />
      <div className="slide-container">
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
