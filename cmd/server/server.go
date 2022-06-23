package main

import (
	config "github.com/nhatvu148/business-day-go/config"
	handlers "github.com/nhatvu148/business-day-go/handlers"
)

func main() {
	config.SetConfig()
	handlers.HandleRequests()
}
