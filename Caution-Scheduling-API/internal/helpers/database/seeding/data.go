package seeding

import (
	"database/sql"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"

)

var database *sql.DB

func SeedData(db *sql.DB) {
	database = db
	seedUsers()

}

func seedUsers() {
	var Admin user.LocalUser = user.LocalUser{
			UserName: "Admin",
			FirstName: "System",
			LastName: "Administrator",
			Email: "admin@localhost.com",
			Password: "P@33word123!",
			IsAdmin: true,
		}

	err := Admin.SignUp(database)

	if err != nil {
		logger.Log(2, "database", "Seeding Data", "System", Admin.UserName + " user is Created")
	}else{
		logger.Log(3, "database", "Seeding Data", "System", err.Error())
	}


	
}
