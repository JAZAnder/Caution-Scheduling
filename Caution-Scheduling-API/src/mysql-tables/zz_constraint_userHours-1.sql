ALTER TABLE `userHours` 
ADD CONSTRAINT `username-hour -> localuser.userName` 
FOREIGN KEY (`username`) REFERENCES `localusers`(`userName`) 
ON DELETE RESTRICT; 
