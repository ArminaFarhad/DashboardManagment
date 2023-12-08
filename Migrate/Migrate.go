package main

import (
	"fmt"
	"log"

	"github.com/ArminaFarhad/DashboardManagment/Initializers"
	"github.com/ArminaFarhad/DashboardManagment/Models"
)

func init() {
	config, err := Initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	Initializers.ConnectDB(&config)
}

func main() {
	Initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	Initializers.DB.AutoMigrate(&Models.Dashboard{}, &Models.Widget{}, &Models.DashboardWidget{})
	fmt.Println("👍 Migration complete")
}
