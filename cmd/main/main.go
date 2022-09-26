package main

import (
	"fmt"
	argsparser "micro_service_phone/internal/args_parser"
	"micro_service_phone/internal/environment"
	"micro_service_phone/internal/server"
)

func main() {

	// initialize command_line arguments, app config, services and repositories
	environment := environment.InitEnvironment()

	// initialize handler with dependencies
	handlers := server.NewHandler(environment)

	(*environment.Countries).CacheRemoteCountries()

	// create server instance
	srv := new(server.Server)
	// run server with InitRoutes
	if err := srv.Run(fmt.Sprintf("%d", *argsparser.ServerPort), handlers.InitRoutes(), environment.Logger); err != nil {
		environment.Logger.Fatalln("Error occured while runnig http server: ", err.Error())
	}

}
