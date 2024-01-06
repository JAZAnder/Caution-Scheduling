CREATE TABLE IF NOT EXISTS hours(
  `Id` int AUTO_INCREMENT PRIMARY KEY,
  `startTime` varchar(255) NOT NULL,
  `endTime` varchar(225) NOT NULL
  `dayOfWeek` int DEFAULT NULL
);