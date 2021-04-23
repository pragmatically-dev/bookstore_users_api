package users

import "encoding/json"

type PublicUser struct {
	ID        int64  `json:"user_id" db:"user_id" `
	CreatedAt string `json:"created_at" db:"created_at" `
	Status    string `json:"status" db:"status" `
}

type PrivateUser struct {
	ID        int64  `json:"user_id" db:"user_id" `
	FirstName string `json:"first_name" db:"first_name" `
	LastName  string `json:"last_name" db:"last_name" `
	Email     string `json:"email" db:"email" `
	CreatedAt string `json:"created_at" db:"created_at" `
	Status    string `json:"status" db:"status" `
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	usrJson, _ := json.Marshal(user)
	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(usrJson, &publicUser)
		return publicUser
	}
	var privateUser PrivateUser
	json.Unmarshal(usrJson, &privateUser)
	return privateUser

}
