package main

import (
	"gowriter/file"
	"gowriter/server"
	"log"
)

func main() {
	err := file.CheckSiteDir()
	if err != nil {
		log.Println(err.Error())
	} else {
		server.Start()
	}
}
