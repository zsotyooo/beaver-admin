package main

import (
	"api/cmd"
	_ "api/docs"
)

// @title Beaver Admin API
// @version 1.0
// @description Beaver Admin API
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3001
// @BasePath /api
func main() {
	cmd.Execute()
}
