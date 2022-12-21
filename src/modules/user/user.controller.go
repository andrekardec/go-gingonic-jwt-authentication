package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerInterface interface {
	Create(ctx *gin.Context)
}

type Controller struct {
	userService ServiceInterface
}

func NewUserController(userService ServiceInterface) ControllerInterface {
	return &Controller{
		userService,
	}
}

func (ctr *Controller) Create(ctx *gin.Context) {
	user := CreateUserRequestDto{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ctr.userService.create(context.Background(), user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "created",
	})
}
