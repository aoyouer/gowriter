package main

import (
	"gowriter/config"
	"gowriter/server"
	"log"
)


func main() {
	err := config.CheckConfig()
	if err != nil {
		log.Println(err.Error())
	} else {
		server.Start(config)
	}
}
