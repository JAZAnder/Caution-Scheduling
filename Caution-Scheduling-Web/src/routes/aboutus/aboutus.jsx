import React from 'react';
import './aboutus.css'; // Ensure this path is correct
import Placeholder from "../../assets/placeholder.jpg"; // Adjust the path if necessary
import EnvocLab from "../../assets/EnvocLab.png"; // Adjust the path if necessary

const AboutUs = () => {
  const cards = [
    {
      name: "Joshua Cantu",
      description: "Team Leader - Senior - Information Technology",
      descriptionextended: "Things about you! Placeholders to fill",
      imgSrc: Placeholder
    },
    {
      name: "Chase Leimbach",
      description: "Senior - Information Technology",
      descriptionextended: "Things about you! Placeholders to fill",
      imgSrc: Placeholder
    },
    {
      name: "Ethan Stoulig",
      description: "Senior - Scientific",
      descriptionextended: "Things about you! Placeholders to fill",
      imgSrc: Placeholder
    },
    {
      name: "Dmitriy Levytskyi",
      description: "Senior - Information Technology",
      descriptionextended: "Things about you! Placeholders to fill",
      imgSrc: Placeholder
    },
  ];

  return (
    <div className="slide-container">
      <div className="slide-content">
        <div className="card-wrapper">
          {cards.map((card, index) => (
            <div className="card" key={index}>
              <div className="image-content">
                <span className="overlay"></span>
                <div className="card-image">
                  <img src={card.imgSrc} alt={card.name} className="card-img" />
                </div>
              </div>
              <div className="card-content">
                <h2 className="name">{card.name}</h2>
                <p className="description">{card.description}</p>
                <p className="descriptionextended">{card.descriptionextended}</p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default AboutUs;
