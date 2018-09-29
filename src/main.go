package main

import (
	dlvr "message/delivery"
	usrcase "message/useCase"

	msgRepository "message/repository"

	"log"

	config "config/configuration"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func init() {

}
func main() {

	config.InitConfiguration()
	e := echo.New()
	repository := msgRepository.NewMessageRepository()

	scase := usrcase.NewMessageUseCase(repository, nil)
	imgscase := usrcase.NewImageUseCase()
	//handler init injecting useCase "scase"
	dlvr.NewMesajeHttpHandler(e, scase, imgscase)
	e.Start(viper.GetString("port"))

	log.Print("started")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

}
