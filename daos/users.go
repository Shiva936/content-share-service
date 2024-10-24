package daos

import (
	"content-share/daos/models"
	"errors"
	"sync"

	"github.com/google/uuid"
)

type IUsers interface {
	GetUsers() []models.User
	CreateUser(user models.User)
	UpdateUser(user models.User) error
}

type Users struct {
	sync.Mutex
	Users []models.User
}

func NewUsers() IUsers {
	return &Users{
		Users: make([]models.User, 0),
	}
}

func (u *Users) GetUsers() []models.User {
	return u.Users
}

func (u *Users) CreateUser(user models.User) {
	u.Lock()
	defer u.Unlock()

	user.ID = uuid.NewString()
	u.Users = append(u.Users, user)
}

func (u *Users) UpdateUser(user models.User) error {
	u.Lock()
	defer u.Unlock()

	var hasFound bool
	for i := range u.Users {
		if u.Users[i].ID == user.ID {
			hasFound = true
			u.Users[i] = user
		}
	}
	if !hasFound {
		return errors.New("user not found")
	}
	return nil
}
