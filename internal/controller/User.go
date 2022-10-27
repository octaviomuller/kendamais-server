package controller

import (
	"net/http"

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

	err := ctx.BindJSON(&body)
	if err != nil {
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

	if *cpf == "" {
		cpf = nil
	}

	if *cnpj == "" {
		cnpj = nil
	}

	err = p.userService.Create(email, password, name, cellphone, cpf, cnpj)
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

func (p *UserController) Login(ctx *gin.Context) {
	body := &model.LoginUser{}

	err := ctx.BindJSON(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	email := body.Email
	password := body.Password

	user, err := p.userService.Login(email, password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)

	return
}

func (p *UserController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := p.userService.Get(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)

	return
}

func (p *UserController) Patch(ctx *gin.Context) {
	body := model.UpdateUser{}

	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	id := ctx.Param("id")
	email := body.Email
	name := body.Name
	cellphone := body.Cellphone
	cpf := &body.Cpf
	cnpj := &body.Cnpj

	if *cpf == "" {
		cpf = nil
	}
	if *cnpj == "" {
		cnpj = nil
	}

	err = p.userService.Update(id, email, name, cellphone, cpf, cnpj)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated with success",
	})

	return
}
