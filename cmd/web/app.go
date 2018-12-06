package main

import (
	"stock-level-api/cmd/models"
)

// Define an App struct to hold the application-wide dependencies
type App struct {
	Addr     string
	Database *models.Database
	TLSCert  string
	TLSKey   string
}
