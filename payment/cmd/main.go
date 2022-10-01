package main

import (
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/payment/config"
)

func main() {
	conf := config.InitConfig()
	logger.InitLogger(&conf.Log)
}
