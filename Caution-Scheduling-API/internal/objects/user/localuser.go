package user

import (
	"database/sql"
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func (u *LocalUser) GetUserByEmail(db *sql.DB) error {
	query := "SELECT `Id`, `UserName`, `firstName`, `lastName`, `email`, `role`, `googleId` FROM `localusers` WHERE `email` = ?;"
	result := db.QueryRow(query, u.Email)

	err := result.Scan(&u.UserId, &u.UserName, &u.FirstName, &u.LastName, &u.Email, &u.Role, &u.GoogleId)
	if err == sql.ErrNoRows {
		return errors.New("user not found")
	}
	return err
}

func (su *SQLLocalUser) toLocalUser() (LocalUser, error) {

	user := LocalUser{
		GoogleId:  su.GoogleId,
		UserName:  su.UserName,
		FirstName: su.FirstName,
		LastName:  su.LastName,
		FullName:  su.FirstName + " " + su.LastName,
		Email:     su.Email,
		Password:  su.Password,
	}
	user.UserId, _ = strconv.Atoi(su.UserId)

	user.Role, _ = strconv.Atoi(su.Role)

	if !user.checkValidUser() {
		return LocalUser{}, errors.New("user missing information")
	}

	return user, nil
}

func (dto *CreateLocalUserDto) ToLocalUser() (LocalUser, error) {

	user := LocalUser{
		UserName:  dto.UserName,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		FullName:  dto.FirstName + " " + dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password,
	}

	user.Role, _ = strconv.Atoi(dto.Role)

	if !user.checkValidUserWithoutId() {
		return LocalUser{}, errors.New("user missing information")
	}

	return user, nil
}

func (u *LocalUser) ToTutorInformation() (TutorInformation, error) {

	if !u.checkValidUser() {
		return TutorInformation{}, errors.New("user missing information")
	}

	return TutorInformation{
		UserId:    u.UserId,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		FullName:  u.FullName,
	}, nil
}

func (u *LocalUser) ToStandardUserInformation() (StandardUserInformation, error) {

	if !u.checkValidUser() {
		return StandardUserInformation{}, errors.New("user missing information")
	}

	return StandardUserInformation{
		UserName:  u.UserName,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}, nil
}

func (u *LocalUser) ToAdminViewUserInformation() (AdminViewUserInformation, error) {

	if !u.checkValidUser() {
		return AdminViewUserInformation{}, errors.New("user missing information 2")
	}

	var userRole string

	if u.Role == 1 {
		userRole = "Student"
	} else if u.Role == 2 {
		userRole = "Tutor"
	} else if u.Role == 3 {
		userRole = "Supervisor"
	} else if u.Role == 4 {
		userRole = "Administrator"
	} else {
		userRole = "Deactivated"
	}

	Id := strconv.Itoa(u.UserId)
	return AdminViewUserInformation{
		UserId:    Id,
		UserName:  u.UserName,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		FullName:  u.FullName,
		Email:     u.Email,
		Role:      userRole,
	}, nil
}
func (u *LocalUser) ToSelfViewInformation() (SelfViewInformation, error) {

	if !u.checkValidUser() {
		return SelfViewInformation{}, errors.New("user missing information")
	}

	var userRole string

	if u.Role == 1 {
		userRole = "Student"
	} else if u.Role == 2 {
		userRole = "Tutor"
	} else if u.Role == 3 {
		userRole = "Supervisor"
	} else if u.Role == 4 {
		userRole = "Administrator"
	} else {
		userRole = "Deactivated"
	}
	return SelfViewInformation{
		UserName:  u.UserName,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		FullName:  u.FullName,
		Email:     u.Email,
		Role:      userRole,
		Settings:  u.Settings,
	}, nil
}

func (u *LocalUser) checkValidUser() bool {
	if u.UserId == 0 || u.UserName == "" || u.Email == "" || u.FirstName == "" || u.LastName == "" || u.Role == 0 {
		return false
	}
	return true
}

func (u *LocalUser) checkValidUserWithoutId() bool {
	if u.UserName == "" || u.Email == "" || u.FirstName == "" || u.LastName == "" || u.Role == 0 {
		return false
	}
	return true
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

	query = "SELECT `Id`, `firstName`, `lastName`, `email`, `isAdmin`, `role`, `fullName`, `googleId` FROM `localusers` WHERE `userName` = '" + u.UserName + "';"
	result = db.QueryRow(query)

	var isAdmin string
	err = result.Scan(&u.UserId, &u.FirstName, &u.LastName, &u.Email, &isAdmin, &u.Role, &u.FullName, &u.GoogleId)
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

	query = "INSERT INTO `localusers` (`userName`, `firstName`, `lastName`, `email`, `password`, `isAdmin`, `role`, `fullName`, `googleId`) VALUES ('" + u.UserName + "', '" + u.FirstName + "', '" + u.LastName + "', '" + u.Email + "', '" + string(hashedPassword) + "', '" + isAdmin + "', '" + strconv.Itoa(u.Role) + "', '" + u.FullName + "', '" + u.GoogleId + "');"
	sqlerr := db.QueryRow(query)

	if sqlerr != nil {
		return sqlerr.Err()
	}

	return nil
}

func GetLusers(db *sql.DB) ([]TutorInformation, error) {
	rows, err := db.Query("SELECT `Id`, `UserName`, `firstName`, `lastName`, `email`, `role` FROM localusers;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	userToReturn := []TutorInformation{}

	for rows.Next() {
		var su SQLLocalUser
		if err := rows.Scan(&su.UserId, &su.UserName, &su.FirstName, &su.LastName, &su.Email, &su.Role); err != nil {
			return nil, err
		}

		user, err := su.toLocalUser()
		if err != nil {
			return nil, err
		}

		viewableUser, _ := user.ToTutorInformation()

		userToReturn = append(userToReturn, viewableUser)
	}
	return userToReturn, nil
}

func GetUsersByFilter(db *sql.DB, filter AdminViewUserInformation) ([]AdminViewUserInformation, error) {
	query := "SELECT `Id`, `UserName`, `firstName`, `lastName`, `email`, `role` FROM localusers WHERE `Id` lIKE'%" + filter.UserId + "%' AND `Username` like '%" + filter.UserName + "%' AND `firstName` Like '%" + filter.FirstName + "%' AND `lastName` Like '%" + filter.LastName + "%' AND `email` like '%" + filter.Email + "%' AND `role` Like '%" + filter.Role + "%';"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	userToReturn := []AdminViewUserInformation{}

	for rows.Next() {
		var su SQLLocalUser
		if err := rows.Scan(&su.UserId, &su.UserName, &su.FirstName, &su.LastName, &su.Email, &su.Role); err != nil {
			return nil, err
		}

		user, err := su.toLocalUser()
		if err != nil {
			return nil, err
		}

		viewableUser, _ := user.ToAdminViewUserInformation()

		userToReturn = append(userToReturn, viewableUser)
	}
	return userToReturn, nil
}

func GetTutors(db *sql.DB) ([]TutorInformation, error) {
	query := "SELECT `Id`, `UserName`, `firstName`, `lastName`, `email`, `role` FROM localusers WHERE `role` = 2;"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	userToReturn := []TutorInformation{}

	for rows.Next() {
		var su SQLLocalUser
		if err := rows.Scan(&su.UserId, &su.UserName, &su.FirstName, &su.LastName, &su.Email, &su.Role); err != nil {
			return nil, err
		}

		user, err := su.toLocalUser()
		if err != nil {
			return nil, err
		}

		viewableUser, _ := user.ToTutorInformation()

		userToReturn = append(userToReturn, viewableUser)
	}
	return userToReturn, nil
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
	u.Password = "REDACTED"
	u.IsAdmin = false
	return err
}



func (u *LocalUser) HasStudentRights() (bool, error) {
	if !u.checkValidUser() {
		return false, errors.New("not Valid User")
	}
	if u.Role >= 1 {
		return true, nil
	} else {
		return false, errors.New("insufficient permissions")
	}
}

func (u *LocalUser) HasTutorRights() (bool, error) {
	if !u.checkValidUser() {
		return false, errors.New("not Valid User")
	}
	if u.Role >= 2 {
		return true, nil
	} else {
		return false, errors.New("insufficient permissions")
	}
}

func (u *LocalUser) HasSupervisorRights() (bool, error) {
	if !u.checkValidUser() {
		return false, errors.New("not Valid User")
	}
	if u.Role >= 3 {
		return true, nil
	} else {
		return false, errors.New("insufficient permissions")
	}
}

func (u *LocalUser) HasAdministratorRights() (bool, error) {
	if !u.checkValidUser() {
		return false, errors.New("not valid user")
	}
	if u.Role >= 4 {
		return true, nil
	} else {
		return false, errors.New("insufficient permissions")
	}
}
