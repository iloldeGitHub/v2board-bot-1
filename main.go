package main

import "github.com/keiko233/V2Board-Bot/service"

func main() {
	service.InitDB()
	service.Start()
}
