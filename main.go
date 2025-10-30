// Package main Main entry point for the Go Backend Tmpl application.
//
//	@update 2024-10-30 08:50:53
package main

import (
	"github.com/hcd233/go-backend-tmpl/cmd"
)

// @title           Go Backend Tmpl
// @version         1.0
// @description     Go Backend Tmpl

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host           localhost:8080
// @BasePath       /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
