package main

import (
	"log"

	"blog/app/application"
	"blog/app/config"
	"blog/app/domain"
	"blog/app/infrastructure"
	"blog/app/infrastructure/persistence/db"
	"blog/app/interfaces"

	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	container.Provide(config.NewConfig)
	container.Provide(db.NewDB)
	container.Provide(db.NewTransaction)
	container.Provide(infrastructure.NewArticleRepository)
	container.Provide(infrastructure.NewSystemRepository)
	container.Provide(domain.NewArticleService)
	container.Provide(domain.NewSystemService)
	container.Provide(application.NewArticleInteractor)
	container.Provide(application.NewSystemInteractor)
	container.Provide(interfaces.NewArticleServer)
	container.Provide(interfaces.NewServer)

	err := container.Invoke(func(server *interfaces.Server) error {
		return server.Run()
	})

	if err != nil {
		log.Fatalln(err)
	}
}
