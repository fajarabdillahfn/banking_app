package main

import (
	"github.com/fajarabdillahfn/banking-lib/logger"
	"github.com/fajarabdillahfn/banking_app/app"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
