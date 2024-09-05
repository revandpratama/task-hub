package server

import (
	"fmt"

	"github.com/revandpratama/task-hub/config"
	"github.com/revandpratama/task-hub/database"
)

func Run() {
	//load all the configurations
	config.LoadConfig()
	database.LoadDB()

	//Initialize routers
	router := InitRouters()

	//Listen and serve to given ports
	router.Listen(fmt.Sprintf(":%s", config.ENV.PORT))
}
