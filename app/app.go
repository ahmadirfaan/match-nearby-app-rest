package app

import "github.com/ahmadirfaan/match-nearby-app-rest/config"

type Application struct {
	Config *config.Config
}

func Init() *Application {
	application := &Application{
		Config: config.Init(),
	}

	return application
}
