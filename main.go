package main

import (
	"github.com/jeffleon/banking-hexarch/app"
	"github.com/jeffleon/banking-hexarch/logger"
)

func main() {
	logger.Info("Starting our aplication...")
	app.Start()
}
