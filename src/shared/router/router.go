package router

import (
	"github.com/andrekardec/go-gingonic-jwt-authentication/src/modules/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	server         *gin.Engine
	userController user.ControllerInterface
}

func NewRouter(server *gin.Engine, userController user.ControllerInterface) *Router {
	return &Router{
		server,
		userController,
	}
}

func (r *Router) Init() {
	basePath := r.server.Group("/api/v1")

	basePath.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
	})

	userRoutes := basePath.Group("/users")
	{
		userRoutes.POST("/", r.userController.Create)
	}
}
