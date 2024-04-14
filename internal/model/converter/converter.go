package converter

import (
	"api-survey-go/internal/entity"
	"api-survey-go/internal/model"
)

func UserToModel(user *entity.User) *model.User {
	return &model.User{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Role:        RoleToModel(&user.Role),
		Client:      ClientToModel(&user.Client),
	}
}

func RoleToModel(role *entity.Role) model.Role {
	return model.Role{
		ID:          role.ID,
		Name:        role.Name,
		DisplayName: role.DisplayName,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}

func ClientToModel(client *entity.Client) model.Client {
	return model.Client{
		ID:         client.ID,
		ClientName: client.ClientName,
		ClientLogo: client.ClientLogo,
		Status:     client.Status,
		CreatedAt:  client.CreatedAt,
		UpdatedAt:  client.UpdatedAt,
	}
}
