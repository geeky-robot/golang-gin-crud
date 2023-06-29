package handler

import (
	"log"

	m "github.com/geeky-robot/golang-gin-crud/model"
	s "github.com/geeky-robot/golang-gin-crud/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserService *s.UserService
}

func NewUserHandler(userService *s.UserService) UserHandler {
	return UserHandler{userService}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var user m.User
	err := ctx.Bind(&user)
	if err != nil {
		log.Println("Something went wrong while unmarshalling request", err)
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": h.UserService.CreateUser(user),
	})

}

func (h *UserHandler) CreateUsers(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var users []m.User
	err := ctx.Bind(&users)
	if err != nil {
		log.Println("Something went wrong while unmarshalling request", err)
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": h.UserService.CreateUsers(users),
	})
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var user m.User
	err := ctx.Bind(&user)
	if err != nil {
		log.Println("Something went wrong while unmarshalling request", err)
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": h.UserService.UpdateUser(user),
	})
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	id := ctx.Params.ByName("id")
	user := h.UserService.GetUser(uuid.MustParse(id))
	if user.Id == uuid.Nil {
		ctx.JSON(204, gin.H{
			"data": "No Data",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, gin.H{
		"data": h.UserService.GetUsers(),
	})
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	id := ctx.Params.ByName("id")
	isDeleted := h.UserService.DeleteUser(uuid.MustParse(id))
	if !isDeleted {
		ctx.JSON(204, gin.H{
			"status": 204,
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"data": "Success",
		})
		return
	}
}
