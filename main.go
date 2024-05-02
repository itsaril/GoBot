// main.go
package main

import (
	"gobot/api"
	"gobot/config"
	"gobot/database"
)

func main() {
	config.Init()
	database.Init()
	api.Init()
}
