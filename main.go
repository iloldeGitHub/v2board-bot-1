package main

import (
	"github.com/miyaUU/v2board-bot/service"
)

func main() {
	service.InitDB()
	service.Start()
}
