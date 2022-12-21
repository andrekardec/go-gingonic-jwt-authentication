package main

import (
	"fmt"
	"github.com/andrekardec/go-gingonic-jwt-authentication/src/modules/user"
	pgKit "github.com/andrekardec/go-gingonic-jwt-authentication/src/shared/providers/storage/pgsql"
	"github.com/andrekardec/go-gingonic-jwt-authentication/src/shared/router"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

var cfg struct {
	DBUser string `envconfig:"DB_USER" default:"postgres"`
	DBPass string `envconfig:"DB_PASS" default:"Password!123#"`
	DBName string `envconfig:"DB_NAME" default:"postgres"`
	DBHost string `envconfig:"DB_HOST" default:"localhost"`
	DBPort int    `envconfig:"DB_PORT" default:"5432"`
}

func main() {
	ginServer := gin.Default()

	if err := envconfig.Process("LIST", &cfg); err != nil {
		err = errors.Wrap(err, "parse environment variables")
		return
	}

	db := pgKit.NewPgDb(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort),
		User:     cfg.DBUser,
		Password: cfg.DBPass,
		Database: cfg.DBName,
	})

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userController := user.NewUserController(userService)

	r := router.NewRouter(ginServer, userController)
	r.Init()

	err := ginServer.Run(":3000")
	if err != nil {
		err = errors.Wrap(err, "parse environment variables")
		return
	}
}
