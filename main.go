package main

import (
	"log"

	"github.com/keiko233/V2Board-Bot/service"
)

func main() {
	_, err := service.InitDB()
	if err != nil {
		log.Fatalln(err)
	}
	service.Start()
}
