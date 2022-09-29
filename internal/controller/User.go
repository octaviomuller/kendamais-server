package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/kendamais-server/internal/interfaces"
	"github.com/octaviomuller/kendamais-server/internal/model"
)

type UserController struct {
	userService interfaces.UserService
}

func NewUserController(userService interfaces.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (p *UserController) Post(ctx *gin.Context) {
	body := model.CreateUser{}

	bodyErr := ctx.BindJSON(&body)
	if bodyErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid request body",
		})

		return
	}

	email := body.Email
	password := body.Password
	name := body.Name
	cpf := &body.Cpf
	cnpj := &body.Cnpj
	cellphone := body.Cellphone

	birthdayTime, err := time.Parse("2006-01-02T15:04:05.000Z", body.Birthday)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	birthday := &[]time.Time{birthdayTime}[0]

	if *cpf == "" {
		cpf = nil
	}

	if *cnpj == "" {
		cnpj = nil
	}

	err = p.userService.Create(email, password, name, cpf, cnpj, cellphone, birthday)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User registered with success",
	})

	return
}
