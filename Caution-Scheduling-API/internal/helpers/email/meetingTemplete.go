package email

import (
	"strconv"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/meeting"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
)

func NewMeeting(student user.LocalUser, tutor user.LocalUser, meeting meeting.Meeting) {
	newMeetingStudentEmail(student, meeting)
	newMeetingTutorEmail(tutor, meeting)
}

func newMeetingStudentEmail(student user.LocalUser, meeting meeting.Meeting) {
	if !student.Settings.ReceiveMeetingEmails {
		logger.Log(2, "Email", "New Meeting Student", student.UserName, student.FirstName+" "+student.LastName+" has declined to receive emails.")
		return
	}

	htmlContent := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Meeting Confirmation</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            color: #333;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }
        .container {
            width: 100%;
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            border-radius: 8px;
            overflow: hidden;
            box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
        }
        .header {
            background-color: #4CAF50;
            padding: 20px;
            text-align: center;
            color: #fff;
        }
        .content {
            padding: 20px;
        }
        .button {
            display: inline-block;
            padding: 10px 20px;
            margin-top: 20px;
            background-color: #4CAF50;
            color: #fff;
            text-decoration: none;
            border-radius: 4px;
        }
        .footer {
            padding: 10px;
            text-align: center;
            font-size: 12px;
            color: #777;
            background-color: #f4f4f4;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Meeting Scheduled</h1>
        </div>
        <div class="content">
            <p>Hi ` + student.FullName + `,</p>
            <p>Thank you for scheduling a meeting with Caution Scheduling!</p>
            <p><strong>Meeting Details:</strong></p>
            <ul>
                <li><strong>Date:</strong> ` + strconv.Itoa(meeting.Date) + `</li>
                <li><strong>Time:</strong> ` + strconv.Itoa(meeting.UserHourId) + `</li>
                <li><strong>Location:</strong> ` + strconv.Itoa(meeting.LabId) + `</li>
            </ul>
            <p>If you need to reschedule or have any questions before the meeting, feel free to contact us.</p>
            <a href="[Insert Calendar Link]" class="button">Add to Calendar</a>
        </div>
        <div class="footer">
            <p>Caution Scheduling</p>
        </div>
    </div>
</body>
</html>
`
	plainTextContent := `Hi ` + student.FullName + `, Thank you for scheduling a meeting with Caution Scheduling! Meeting Details: ` + strconv.Itoa(meeting.Date) + ` -- ` + strconv.Itoa(meeting.UserHourId) + ``

	sendEmail(student.Email, student.FullName, "New Meeting Scheduled", htmlContent, plainTextContent)

	return

}

func newMeetingTutorEmail(tutor user.LocalUser, meeting meeting.Meeting) {
	if !tutor.Settings.ReceiveMeetingEmails {
		logger.Log(2, "Email", "New Meeting Student", tutor.UserName, tutor.FirstName+" "+tutor.LastName+" has declined to receive emails.")
		return
	}

	htmlContent := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Meeting Confirmation</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            color: #333;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }
        .container {
            width: 100%;
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            border-radius: 8px;
            overflow: hidden;
            box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
        }
        .header {
            background-color: #4CAF50;
            padding: 20px;
            text-align: center;
            color: #fff;
        }
        .content {
            padding: 20px;
        }
        .button {
            display: inline-block;
            padding: 10px 20px;
            margin-top: 20px;
            background-color: #4CAF50;
            color: #fff;
            text-decoration: none;
            border-radius: 4px;
        }
        .footer {
            padding: 10px;
            text-align: center;
            font-size: 12px;
            color: #777;
            background-color: #f4f4f4;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Meeting Scheduled</h1>
        </div>
        <div class="content">
            <p>Hi ` + tutor.FullName + `,</p>
            <p>Thank you for scheduling a meeting with Caution Scheduling!</p>
            <p><strong>Meeting Details:</strong></p>
            <ul>
                <li><strong>Date:</strong> ` + strconv.Itoa(meeting.Date) + `</li>
                <li><strong>Time:</strong> ` + strconv.Itoa(meeting.UserHourId) + `</li>
                <li><strong>Location:</strong> ` + strconv.Itoa(meeting.LabId) + `</li>
            </ul>
            <p>If you need to reschedule or have any questions before the meeting, feel free to contact us.</p>
            <a href="[Insert Calendar Link]" class="button">Add to Calendar</a>
        </div>
        <div class="footer">
            <p>Caution Scheduling</p>
        </div>
    </div>
</body>
</html>
`
	plainTextContent := `Hi ` + tutor.FullName + `, Thank you for scheduling a meeting with Caution Scheduling! Meeting Details: ` + strconv.Itoa(meeting.Date) + ` -- ` + strconv.Itoa(meeting.UserHourId) + ``

	sendEmail(tutor.Email, tutor.FullName, "New Meeting Scheduled", htmlContent, plainTextContent)
	return
}
