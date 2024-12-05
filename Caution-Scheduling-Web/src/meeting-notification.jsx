import React from 'react';
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert';
import { Card, CardContent } from '@/components/ui/card';
import { Clock, Calendar } from 'lucide-react';

const NotificationCard = ({ 
  type="confirmed",
  date,
  time,
  studentName = "[STUDENT_NAME]",
  tutorName = "[TUTOR_NAME]",
  subject = "[SUBJECT]"
}) => {
  const notificationConfig = {
    confirmed: {
      title: "Meeting Confirmed",
      heading: "Hello there,",
      subheading: "Thank you for scheduling your meeting! We're looking forward to connecting with you.",
      statusColor: "bg-green-100 dark:bg-green-900",
      buttonText: "Add to Calendar"
    },
    cancelled: {
      title: "Meeting Cancelled",
      heading: "Sometimes plans change...",
      subheading: "Your meeting has been cancelled - no worries! Feel free to reschedule whenever you're ready.",
      statusColor: "bg-red-100 dark:bg-red-900",
      buttonText: "Schedule New Meeting"
    },
    rescheduled: {
      title: "Meeting Rescheduled ✨",
      heading: "Time for a change of plans!",
      subheading: "No worries - we've got your meeting rescheduled and ready to go. New time, same great conversation!",
      statusColor: "bg-blue-100 dark:bg-blue-900",
      buttonText: "Update Calendar"
    }
  };

  const config = notificationConfig[type] || notificationConfig.confirmed;

  return (
    <div className="max-w-2xl mx-auto p-6 bg-white dark:bg-gray-800 rounded-lg shadow">
      {/* Status Banner */}
      <div className={`${config.statusColor} px-4 py-2 rounded-t-lg text-center font-medium`}>
        {config.title}
      </div>

      {/* Main Content */}
      <div className="mt-8">
        <h2 className="text-2xl mb-4">{config.heading}</h2>
        <p className="text-gray-600 dark:text-gray-300 mb-8">{config.subheading}</p>

        {/* Meeting Details Section */}
        <div className="bg-gray-50 dark:bg-gray-700 rounded-lg p-6">
          <h3 className="text-lg font-semibold mb-4">Meeting Details:</h3>
          
          <div className="space-y-4">
            <div className="flex items-center gap-3">
              <Calendar className="h-5 w-5 text-gray-500" />
              <span>Date: <span className="font-medium">{date}</span></span>
            </div>
            <div className="flex items-center gap-3">
              <Clock className="h-5 w-5 text-gray-500" />
              <span>Time: <span className="font-medium">{time}</span></span>
            </div>
            {type !== 'cancelled' && (
              <div className="space-y-2 mt-4">
                <p><strong>Student:</strong> {studentName}</p>
                <p><strong>Tutor:</strong> {tutorName}</p>
                <p><strong>Subject:</strong> {subject}</p>
              </div>
            )}
          </div>
        </div>

        {/* Action Button */}
        <div className="mt-8 text-center">
          <button className="bg-blue-500 hover:bg-blue-600 text-white px-6 py-2 rounded-lg font-medium">
            {config.buttonText}
          </button>
        </div>

        {/* Footer */}
        <footer className="mt-8 text-center text-sm text-gray-500 space-y-2">
          <p>Questions? Contact us at support@tutoring.cantusolutions.com</p>
          <p>
            <a href="#" className="text-blue-500 hover:underline">
              Unsubscribe from these notifications
            </a>
          </p>
        </footer>
      </div>
    </div>
  );
};

export default NotificationCard;
