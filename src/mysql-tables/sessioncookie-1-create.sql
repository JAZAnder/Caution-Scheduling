CREATE TABLE IF NOT EXISTS `sessionCookie`(
    `Id` INT AUTO_INCREMENT PRIMARY KEY,
    `cookie` VARCHAR(255) NOT NULL UNIQUE KEY,
    `username` VARCHAR(255) NOT NULL
);
