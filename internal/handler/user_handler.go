package handler

import (
	"api-survey-go/internal/model"
	"api-survey-go/internal/model/converter"
	"api-survey-go/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	Log     *logrus.Logger
	Service *service.UserService
}

func NewUserHandler(service *service.UserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		Log:     logger,
		Service: service,
	}
}

func (s *UserHandler) GetCurrentUser(c *fiber.Ctx) error {
	auth := c.Locals("auth").(*model.Auth)
	user, err := s.Service.GetUserById(auth.ID)
	if err != nil {
		s.Log.Warnf("failed to get user : %v", err)
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(model.Response{
		Status: "success",
		Data:   converter.UserToModel(user),
	})
}

func (s *UserHandler) FindById(c *fiber.Ctx) error {
	user, err := s.Service.GetUserById(c.Params("id"))
	if err != nil {
		s.Log.Warnf("failed to get user : %v", err)
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(model.Response{
		Status: "success",
		Data:   converter.UserToModel(user),
	})
}

func (s *UserHandler) Login(c *fiber.Ctx) error {
	request := new(model.UserLoginRequest)
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := s.Service.Login(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(model.Response{Data: user, Status: "success"})
}

func (s *UserHandler) Create(c *fiber.Ctx) error {
	request := new(model.CreateUserRequest)
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := s.Service.Validate.Struct(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := s.Service.Create(c.UserContext(), request)
	if err != nil {
		return err
	}

	return c.JSON(model.Response{
		Status: "success",
		Data:   user,
	})
}

func (s *UserHandler) Update(c *fiber.Ctx) error {
	request := new(model.UpdateUserRequest)
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := s.Service.Validate.Struct((request)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := s.Service.GetUserById(c.Params("id"))
	if err != nil {
		s.Log.Warnf("failed to get user : %v", err)
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	userUpdated, err := s.Service.Update(c.UserContext(), request, user)
	if err != nil {
		return err
	}

	return c.JSON(model.Response{
		Status: "success",
		Data:   userUpdated,
	})
}
