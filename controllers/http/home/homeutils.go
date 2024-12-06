package httpHome

import (
	gormDB "LOOTERZ_backend/config/database"
	"LOOTERZ_backend/models/modelsDB"
)

func createUser(user *modelsDB.User) error {
	result := gormDB.DB.Select("UserID", "UserName").Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func checkUserInRoom(userId string) (string, error) {
	var user modelsDB.User
	result := gormDB.DB.Select("status").First(&user, "userid = ?", userId)

	if result.Error != nil || result.RowsAffected < 1 {
		return "", result.Error
	}

	statusMap := map[int8]string{
		0: "free",
		1: "lobby",
		2: "play",
	}

	status, ok := statusMap[user.Status]
	if !ok {
		status = "unknown"
	}

	return status, nil
}
