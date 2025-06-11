package main

import (
	"ewallet-notification/cmd"
	helpers "ewallet-notification/helpers"
)

func main() {
	//load config
	helpers.SetupConfig()

	//load log
	helpers.SetupLogger()

	//load database
	helpers.SetupMySQL()

	//run grpc
	cmd.ServeGRPC()

	//run http
	// cmd.ServeHTTP()

}
