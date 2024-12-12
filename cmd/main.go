package main

import (
	"pet_pr/log"
	"pet_pr/tools/configs"
	"pet_pr/tools/handlers"
	"pet_pr/tools/storage"
)

func main() {
	cfg := configs.InitServerConfig()
	log := log.LoggerSetup(cfg.Env)
	storage := storage.InitStorage(cfg.DBConfig, log)
	if storage == nil {
		log.Error("The application is not initialised")
		return
	}
	storage.CreateTables()
	serv := handlers.InitHandlers(storage)
	if serv == nil {
		log.Error("Server not init")
		return
	}
	if err := serv.Router.Run(cfg.Address); err != nil {
		log.Error("Server stoped")
		return
	}
	log.Info("Server started")
}