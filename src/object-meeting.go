package main

import(
	"database/sql"
	"strconv"
)

type meeting struct{
	Id int `json:"id"`
	UserHourId int `json:"userHourId"`
	LabId int `json:"labId"`
	StudentName string `json:"studentName"`
	StudentEmail string `json:"studentEmail"`
}

func (m *meeting) getMeeting(db *sql.DB) error{
	var tempTutorHourId string
	var tempLabId string
	query := "SELECT userHourId, labId, studentName, studentEmail FROM meetings WHERE id=" + strconv.Itoa(m.Id)
	err := db.QueryRow(query).Scan(&tempUserHourId, &tempLabId, &m.StudentName, &m.StudentEmail)
	m.TutorHourId, err = strconv.Atoi(tempUserHourId)
	m.LabId, err = strconv.Atoi(tempLabId)
	return err
}



func (m *meeting) updateMeeting(db *sql.DB) error{
	query := "Update `meetings` SET `userHourId` = '"+strconv.Itoa(m.UserHourId)+"', `labId` = '"+strconv.Itoa(m.LabId)+"', `studentName` = '"+m.StudentName+"', `studentEmail` = '"+m.StudentEmail+"' WHERE `meetings`.`id` ="+strconv.Itoa(m.Id)+""
	_, err := db.Exec(query)
	return err
}

func (m *meeting) createMeeting(db *sql.DB) error{
	var available bool
	query = "SELECT `Available` FROM `UserHours` WHERE `userHourId` = " + m.userHourId
	err = db.QueryRow(query).Scan(&available)

	if err != nil{
		return err.Err()
	}
	
	if !available{
		return errors.New("This tutor is not available for that given time")
	}

	query := "INSERT INTO `meetings` (`userHourId`, `labId`, `studentName`, `studentEmail`) VALUES ('"+strconv.Itoa(m.UserHourId)+"', '"+strconv.Itoa(m.LabId)+"', '"+m.StudentName+"', '"+m.StudentEmail+"');"
	err := db.QueryRow(query)

	if err != nil{
		return err.Err()
	}

	query := "UPDATE `UserHours` SET `Available` = 'false' WHERE `userHourId` = " + m.userHourId
	err := db.QueryRow(query)

	if err != nil{
		return err.Err()
	}

	return nil
}

func getMeetings(db *sql.DB) ([]meeting, error){
	rows, err := db.Query("SELECT id, userHourId, labId, studentName, studentEmail FROM meetings")

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	meetings := []meeting{}

	for rows.Next(){
		var m meeting
		if err := rows.Scan(&m.Id, &m.UserHourId, &m.LabId, &m.StudentName, &m.StudentEmail); err != nil{
			return nil, err
		}
		meetings = append(meetings, m)
	}
	return meetings, nil
}

func getMyMeetings(db *sql.DB) ([]meeting, error){
	var c sessionCookie
    cookie, err := r.Cookie("key")
    if err != nil {
        if errors.Is(err, http.ErrNoCookie) {
            respondWithError(w, http.StatusUnauthorized, "Cookie not Found")
            return
        } else {
            respondWithError(w, http.StatusInternalServerError, err.Error())
            return
        }
    }

    c.Cookie = cookie.Value

    currentUser, err := c.checkSession(a.DB)
    if err != nil {
        if err == sql.ErrNoRows {
            respondWithError(w, http.StatusUnauthorized, "Session Expired")
            return
        } else {
            respondWithError(w, http.StatusInternalServerError, err.Error())
            return
        }
    } 

	rows, err := db.Query("SELECT id, userHourId, labId, studentName, studentEmail FROM meetings m join userHours u on m.userHourId = u.id where u.userId = " + currentUser)

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	meetings := []meeting{}

	for rows.Next(){
		var m meeting
		if err := rows.Scan(&m.Id, &m.UserHourId, &m.LabId, &m.StudentName, &m.StudentEmail); err != nil{
			return nil, err
		}
		meetings = append(meetings, m)
	}
	return meetings, nil
}

func (m *meeting) deleteMeeting(db *sql.DB) error{
	query := "DELETE FROM `meetings` WHERE `meetings`.`Id`="+strconv.Itoa(m.Id)+""
	_, err := db.Exec(query)
	return err
}