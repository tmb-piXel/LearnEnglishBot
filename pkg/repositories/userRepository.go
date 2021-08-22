package repositories

import (
	m "github.com/tmb-piXel/LearnEnglishBot/pkg/models"
)

var users = make(map[int64]*m.User)

func GetUser(id int64) *m.User {
	return users[id]
}

func SaveUser(u *m.User) {
	id := u.GetChatID()
	users[id] = u
}
