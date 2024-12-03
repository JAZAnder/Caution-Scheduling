import React, { useState } from 'react';
import axios from 'axios';

const TimeslotToggle = ({ timeslot }) => {
    const [isEnabled, setIsEnabled] = useState(timeslot.isEnabled);

    const handleToggle = async () => {
        try {
            const response = await axios.patch(`/api/timeslots/${timeslot.id}`, { 
                id: timeslot.id,
                isEnabled: !isEnabled 
            });
            setIsEnabled(response.data.isEnabled);
        } catch (error) {
            console.error('Error updating timeslot status:', error);
        }
    };

    return (
        <div className="timeslot-toggle">
            <label>
                <input 
                    type="checkbox" 
                    checked={isEnabled} 
                    onChange={handleToggle} 
                />
                Timeslot {timeslot.time}
            </label>
        </div>
    );
};

export default TimeslotToggle;
