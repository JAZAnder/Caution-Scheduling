CREATE TABLE IF NOT EXISTS meetings(
  `Id` int AUTO_INCREMENT PRIMARY KEY,
  `tutorHourId` int DEFAULT NULL /*FOREIGN KEY REFERENCES userHour(Id)*/,
  `labId` int DEFAULT NULL,
  `studentName` varchar(255) NOT NULL,
  `studentEmail` varchar(255) NOT NULL
);