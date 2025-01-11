package handler

import "github.com/qw1281065768/trash_be/database"

type UserInfo struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
	Money int    `json:"money"`
	Exp   int    `json:"exp"`
}

func GetUserInfo(userID int64) (*UserInfo, error) {
	user, err := database.GetUserInfo(userID)
	if err != nil {
		return nil, err
	}
	info := &UserInfo{
		Name:  user.Name,
		Level: user.Level,
		Money: user.Money,
		Exp:   user.Exp,
	}

	return info, nil
}
