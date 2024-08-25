ALTER TABLE
    `sessionCookie` ADD CONSTRAINT `username -> localuser.userName` FOREIGN KEY(`username`) REFERENCES `localusers`(`userName`) ON DELETE CASCADE;