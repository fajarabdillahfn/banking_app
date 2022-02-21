package main

import (
	"github.com/fajarabdillahfn/banking_app/app"
	"github.com/fajarabdillahfn/banking_app/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
