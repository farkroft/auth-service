package main

import (
	"gitlab.com/farkroft/auth-service/application/controller"
	"gitlab.com/farkroft/auth-service/application/repository"
	"gitlab.com/farkroft/auth-service/application/usecase"
	"gitlab.com/farkroft/auth-service/external/config"
	"gitlab.com/farkroft/auth-service/external/constants"
	"gitlab.com/farkroft/auth-service/external/database"
	"gitlab.com/farkroft/auth-service/external/log"
	"gitlab.com/farkroft/auth-service/external/redis"
	"gitlab.com/farkroft/auth-service/external/server"
)

func main() {
	log.NewLogger()
	v := config.NewConfig(constants.EnvConfigFile)
	db := database.NewDatabase(v)
	err := db.Migrate()
	if err != nil {
		log.Errorf("migrate", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()
	rd := redis.NewRedis(v)
	userRepo := repository.NewUserRepository(db)
	usecase := usecase.NewUsecase(userRepo, v, rd)
	ctl := controller.NewController(usecase)
	server.NewServer(v, ctl)
}
