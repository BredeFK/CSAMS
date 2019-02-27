package config

import (
	"encoding/json"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/email"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/server"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/view"
	"io/ioutil"
	"log"
	"os"
)

// Configuration struct
type Configuration struct {
	Database *db.MySQLInfo  `json:"database"`
	Server   *server.Server `json:"server"`

	View *view.View `json:"view"`

	Session *session.Session `json:"session"`
	Email   *email.SMTPInfo  `json:"email"`

	Template      *view.Template `json:"template"`
	TemplateAdmin *view.Template `json:"template_admin"`
}

// Load a JSON file making a Configuration pointer
func Load(configFile string) (*Configuration, error) {
	// Open file
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("could not open config file: %v\n", err)
		return &Configuration{}, err
	}

	// Close file
	defer file.Close()
	// Read file
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("could not read all from file: %v\n", err)
		return &Configuration{}, err
	}

	// Declare Configuration
	cfg := Configuration{}
	// Unmarshal JSON to Configuration-object
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		log.Fatalf("could not unmarshal json: %v\n", err)
		return &Configuration{}, err
	}

	return &cfg, nil
}

// Initialize the configuration
func Initialize(filePath string) *Configuration {
	var cfg = &Configuration{}
	cfg, err := Load(filePath)
	if err != nil {
		panic(err)
	}

	// Configure Session
	session.Configure(cfg.Session)

	// Configure Database
	db.ConfigureDB(cfg.Database)
	db.OpenDB()

	// Configure View
	view.Configure(cfg.View)
	view.LoadTemplate(cfg.Template.Root, cfg.Template.Children)
	view.LoadAdminTemplate(cfg.TemplateAdmin.Root, cfg.TemplateAdmin.Children)

	return cfg
}