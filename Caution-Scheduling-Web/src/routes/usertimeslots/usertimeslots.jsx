import { Link, useSearchParams } from "react-router-dom";
import React, { useEffect, useState } from "react";
import useFetch, { CachePolicies } from "use-http";

function UserTimeslots(){

    return(
        <>
        <div className="underNavBarSpacer"></div>


        <div>
             <FilterUserTimeSlots/>
        </div>
        
       
        
        </>
    )
}

function FilterUserTimeSlots(){
    const [params] = useSearchParams();
    const [hourId, setHourId] = useState("")
    const [tutorId, setTutorId] = useState("")
    const [dayOfWeek, setDayOfWeek] = useState("")

    const searchTermsHourId = params.get("hourId")
    const searchTermsTutorId = params.get("tutorId")
    const searchTermsDayOfWeek  = params.get("dayOfWeek")


    
      

    const { data: UserTimeslots, loading, error } = useFetch(
        `/api/availability?tutorId=`+searchTermsTutorId+`&hourId=`+searchTermsHourId+`&dayOfWeek=`+searchTermsDayOfWeek,
        { method: "get" },
        []
      );


      if (loading) {
        return (
          <center>
            <div className="loader"></div>
          </center>
        );
      }

    return(
        <>
            <table>
                <thead>
                    <tr>
                        <th>Id</th>
                        <th>First Name</th>
                        <th>Last Name</th>
                        <th>Day Of The Week</th>
                        <th>Start Time</th>
                        <th>EndTime</th>
                        <th>Hour Details</th>
                        <th>Tutor Details</th>
                    </tr>
                </thead>
                <tbody>
            {UserTimeslots &&
            Object.keys(UserTimeslots).map((timeSlot, i) => (
              <tr key={i}>
                <td>{UserTimeslots[timeSlot].id}</td>
                <td>{UserTimeslots[timeSlot].firstName}</td>
                <td>{UserTimeslots[timeSlot].lastName}</td>
                <td>{UserTimeslots[timeSlot].dayOfWeek}</td>
                <td>{UserTimeslots[timeSlot].startTime}</td>
                <td>{UserTimeslots[timeSlot].endTime}</td>
                <td>
                  <button/>
                </td>
                <td>
                  <button/>
                </td>
              </tr>
            ))}
                </tbody>
            </table>
        </>
    )
}

export default UserTimeslots

