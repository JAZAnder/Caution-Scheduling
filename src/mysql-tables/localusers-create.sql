CREATE TABLE IF NOT EXISTS localusers(
  `userName` varchar(255) PRIMARY KEY,
  `firstName` varchar(255) NOT NULL,
  `lastName` varchar(225) NOT NULL,
  `email` varchar(225) NOT NULL,
  `password` varchar(225) NOT NULL,
  `isAdmin` boolean NOT NULL);