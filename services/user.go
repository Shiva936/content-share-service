package services

import (
	"content-share/daos"
	"content-share/daos/models"
	"content-share/dtos"
	"context"
	"time"
)

func GetUsers(ctx *context.Context) (*dtos.Users, error) {
	response := &dtos.Users{}
	users := daos.UsersDB.GetUsers()
	for i := range users {
		response.Users = append(response.Users, dtos.User{
			ID:       users[i].ID,
			Name:     users[i].Name,
			Email:    users[i].Email,
			Password: users[i].Password,
		})
	}
	return response, nil
}

func CreateUser(ctx *context.Context, user *dtos.User) (*dtos.User, error) {
	daos.UsersDB.CreateUser(models.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
	})
	return user, nil
}

func UpdateUser(ctx *context.Context, user *dtos.User) (*dtos.User, error) {
	err := daos.UsersDB.UpdateUser(models.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
