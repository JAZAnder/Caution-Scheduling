
-- ALTER TABLE `meetings` ADD CONSTRAINT `Tutor/Hour -> tutorHour.Id` FOREIGN KEY (`tutorHourId`) REFERENCES `userHours`(`Id`) ON DELETE RESTRICT;