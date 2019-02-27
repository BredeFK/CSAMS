package main

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/route"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/config"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/server"
)

func main() {
	// Initialize config
	var cfg = config.Initialize("config/config.json")

	defer db.CloseDB()

	// Run Server
	server.Run(route.Load(), route.LoadHTTPS(), cfg.Server)
}
