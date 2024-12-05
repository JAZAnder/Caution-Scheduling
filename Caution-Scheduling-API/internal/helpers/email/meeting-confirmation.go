package email

import (
	"strconv"


	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/hour"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/meeting"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/userHour"
)


func MeetingCreated(student user.LocalUser, tutor user.LocalUser, meeting meeting.Meeting) {
	meetingCreated(student, meeting)
	meetingCreated(tutor, meeting)
}

func meetingCreated(student user.LocalUser, meeting meeting.Meeting) {
	// if !student.Settings.ReceiveMeetingEmails {
	// 	logger.Log(2, "Email", "New Meeting Student", student.UserName, student.FirstName+" "+student.LastName+" has declined to receive emails.")
	// 	return
	// }

	var uh = userHour.UserHour{Id: meeting.Id}
	uh.GetUserHour(db)
	
	var h = hour.Hour{Id: uh.HourId}
	h.GetHour(db)

	htmlContent := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Meeting Confirmation</title>
    <style>
        :root {
            --southeastern-green: #2B5234;  /* Pantone 357 C */
            --southeastern-gold: #FFC72C;   /* Pantone 123 C */
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            line-height: 1.6;
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .email-container {
            background-color: white;
            border-radius: 8px;
            padding: 30px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
            border-top: 4px solid var(--southeastern-green);
        }
        .status-banner {
            background-color: var(--southeastern-green);
            color: white;
            padding: 15px 20px;
            border-radius: 6px;
            margin-bottom: 25px;
            font-size: 24px;
            text-align: center;
        }
        .meeting-details {
            background-color: #f8f9fa;
            border-radius: 6px;
            padding: 20px;
            margin: 20px 0;
            border-left: 4px solid var(--southeastern-gold);
        }
        .detail-item {
            margin: 10px 0;
            display: flex;
            align-items: center;
        }
        .detail-item svg {
            margin-right: 10px;
            color: var(--southeastern-green);
        }
        .button {
            background-color: var(--southeastern-green);
            color: white;
            padding: 12px 24px;
            border-radius: 6px;
            text-decoration: none;
            display: inline-block;
            margin: 20px 0;
            text-align: center;
            font-weight: 500;
            transition: all 0.2s;
            border: 2px solid var(--southeastern-green);
        }
        .button:hover {
            background-color: white;
            color: var(--southeastern-green);
        }
        .footer {
            text-align: center;
            color: #6b7280;
            font-size: 14px;
            margin-top: 30px;
            padding-top: 20px;
            border-top: 1px solid #e5e7eb;
        }
        .footer a {
            color: var(--southeastern-green);
            text-decoration: none;
            font-weight: 500;
        }
        .footer a:hover {
            color: var(--southeastern-gold);
        }
        h2, h3 {
            color: var(--southeastern-green);
        }
        .highlight {
            color: var(--southeastern-green);
            font-weight: 500;
        }
    </style>
</head>
<body>
    <div class="email-container">
        <div class="status-banner">
            Meeting Confirmed
        </div>
        
        <h2>Hello there,</h2>
        
        <p>Thank you for scheduling your meeting! We're looking forward to connecting with you.</p>
        
        <div class="meeting-details">
            <h3 style="margin-top: 0;">Meeting Details:</h3>
            <div class="detail-item">
                <svg width="20" height="20" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                <span>Date: <span class="highlight">`+ strconv.Itoa(meeting.Date)+`</span></span>
            </div>
            <div class="detail-item">
                <svg width="20" height="20" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span>Time: <span class="highlight">`+h.StartTime+`</span></span>
            </div>
        </div>

        <p>Need to make changes? No problem! You can reschedule or cancel your meeting if needed.</p>
        
        <a href="#" class="button">Add to Calendar</a>
        
        <div class="footer">
            <p>Questions? Contact us at <a href="mailto:support@tutoring.cantusolutions.com">support@tutoring.cantusolutions.com</a></p>
            <p>To unsubscribe from these notifications, <a href="#">click here</a></p>
        </div>
    </div>
</body>
</html>

`
	plainTextContent := `Hi ` + student.FullName + `, Thank you for scheduling a meeting with Caution Scheduling! Meeting Details: ` + strconv.Itoa(meeting.Date) + ` -- ` + strconv.Itoa(meeting.UserHourId) + ``

	logger.Log(1, "Email", "Meeting", "emailManager", "Email has been send to "+student.FullName+" at "+student.Email+" informing them about a meeting they scheduled")
	sendEmail(student.Email, student.FullName, "New Meeting Scheduled", htmlContent, plainTextContent)

}