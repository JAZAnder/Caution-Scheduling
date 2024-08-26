CREATE TABLE IF NOT EXISTS userHours(
  `Id` int AUTO_INCREMENT PRIMARY KEY,
  `hourId` INT NOT NULL,
  `username` VARCHAR(255) NOT NULL,
  `available` tinyint(1) NOT NULL DEFAULT '1'
);