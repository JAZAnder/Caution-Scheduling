package user

import (
	"database/sql"
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)



func (u *LocalUser) ToTutorInformation() (TutorInformation, error){
	
	if (u.checkValidUser()) {
		return TutorInformation{}, errors.New("user missing information")
	}
	
	return TutorInformation{
		UserId: u.UserId,
		FirstName: u.FirstName,
		LastName: u.LastName,
		FullName: u.FullName,
	}, nil
}

func (u *LocalUser) ToStandardUserInformation() (StandardUserInformation, error){
	
	if (u.checkValidUser()) {
		return StandardUserInformation{}, errors.New("user missing information")
	}
	
	return StandardUserInformation{
		UserName: u.UserName,
		FirstName: u.FirstName,
		LastName: u.LastName,
		Email: u.Email,
	}, nil
}

func (u *LocalUser) ToSelfViewInformation() (SelfViewInformation, error){
	
	if (u.checkValidUser()) {
		return SelfViewInformation{}, errors.New("user missing information")
	}
	
	var userRole string;

	if u.Role == 1 {
		userRole = "Student"
	}else if u.Role == 2 {
		userRole = "Tutor"
	}else if u.Role == 3 {
		userRole = "Supervisor"
	}else if u.Role == 4 {
		userRole = "Administrator"
	}else{
		userRole = "Deactivated"
	}
	return SelfViewInformation{
		UserName: u.UserName,
		FirstName: u.FirstName,
		LastName: u.LastName,
		FullName: u.FullName,
		Email: u.Email,
		Role: userRole,
		Settings: u.Settings,
	}, nil
}

func (u *LocalUser) checkValidUser() (bool){
	if (u.UserId == 0 || u.UserName == ""|| u.Email == "" || u.FirstName == "" || u.LastName == "" || u.Role == 0) {
		return false
	}
	return true;
}

func (u *LocalUser) Login(db *sql.DB) error {
	query := "SELECT `password` FROM `localusers` WHERE `userName` = '" + u.UserName + "';"
	result := db.QueryRow(query)

	var storedCreds LocalUser

	err := result.Scan(&storedCreds.Password)
	if err == sql.ErrNoRows {
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(u.Password)); err != nil {
		return err
	}

	query = "SELECT `firstName`, `lastName`, `email`, `isAdmin`, `role`, `fullName`, `googleId` FROM `localusers` WHERE `userName` = '" + u.UserName + "';"
	result = db.QueryRow(query)

	var isAdmin string
	err = result.Scan(&u.FirstName, &u.LastName, &u.Email, &isAdmin, &u.Role, &u.FullName, &u.GoogleId)
	if err != nil {
		return err
	}
	u.IsAdmin, _ = strconv.ParseBool(isAdmin)

	return nil
}

func (u *LocalUser) SignUp(db *sql.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return err
	}

	isAdmin := "0"
	if u.IsAdmin {
		isAdmin = "1"
	}

	query := "INSERT INTO `userSettings` (`userName`, `ReceiveMeetingEmails`) VALUES ('" + u.UserName + "', '" + "1" + "');"
	db.QueryRow(query)

	query = "INSERT INTO `localusers` (`userName`, `firstName`, `lastName`, `email`, `password`, `isAdmin`, `role`, `fullName`, `googleId`) VALUES ('" + u.UserName + "', '" + u.FirstName + "', '" + u.LastName + "', '" + u.Email + "', '" + string(hashedPassword) + "', '" + isAdmin + "', '" + strconv.Itoa(u.Role)  + "', '" + u.FullName + "', '" + u.GoogleId + "');"
	sqlerr := db.QueryRow(query)

	if sqlerr != nil {
		return sqlerr.Err()
	}

	return nil
}

func GetLusers(db *sql.DB, isAdmin bool) ([]LocalUser, error) {
	rows, err := db.Query("SELECT * FROM `localusers`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	lusers := []LocalUser{}

	for rows.Next() {
		var u LocalUser
		var uAdmin string
		if err := rows.Scan(&u.UserName, &u.FirstName, &u.LastName, &u.Email, &u.Password, &uAdmin); err != nil {
			return nil, err
		}
		u.Password = "REDACTED"
		u.IsAdmin, _ = strconv.ParseBool(uAdmin)
		if !isAdmin {
			u.IsAdmin = false
		}
		lusers = append(lusers, u)
	}
	return lusers, nil
}

func (u *LocalUser) ChangePassword(db *sql.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return err
	}
	query := "UPDATE `localusers` SET `password` = '" + string(hashedPassword) + "'WHERE `userName` = '" + u.UserName + "';"
	sqlerr := db.QueryRow(query)

	if sqlerr != nil {
		return sqlerr.Err()
	}

	return nil
}

func (u *LocalUser) GetUser(db *sql.DB) error {
	query := "SELECT `firstName`, `lastName` FROM `localusers` WHERE `userName` = '" + u.UserName + "'"
	err := db.QueryRow(query).Scan(&u.FirstName, &u.LastName)
	u.Email = "REDACTED"
	u.UserName = "REDACTED"
	u.Password = "REDACTED"
	u.IsAdmin = false
	return err
}

func (u *LocalUser) HasStudentRights() (bool, error) {
	if (!u.checkValidUser()){
		return false, errors.New("not Valid User")
	}
	if(u.Role >= 1){
		return true, nil
	}else{
		return false, errors.New("insufficient permissions")
	}
}

func (u *LocalUser) HasTutorRights() (bool, error) {
	if (!u.checkValidUser()){
		return false, errors.New("not Valid User")
	}
	if(u.Role >= 2){
		return true, nil
	}else{
		return false, errors.New("insufficient permissions")
	}
}

func (u *LocalUser) HasSupervisorRights() (bool, error) {
	if (!u.checkValidUser()){
		return false, errors.New("not Valid User")
	}
	if(u.Role >= 3){
		return true, nil
	}else{
		return false, errors.New("insufficient permissions")
	}
}

func (u *LocalUser) HasAdministratorRights() (bool, error) {
	if (!u.checkValidUser()){
		return false, errors.New("not valid user")
	}
	if(u.Role >= 4){
		return true, nil
	}else{
		return false, errors.New("insufficient permissions")
	}
}


