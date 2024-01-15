package main

import (
	"clean-project-exp/internal/factory"
	"clean-project-exp/internal/infra/database/memory"
	"clean-project-exp/internal/infra/web/webserver"
)

func main() {
	ws := webserver.NewWebServer("8080")

	db := memory.InitInMemoryDB()
	userHandler := factory.MakeWebUserHandler(db)

	ws.AddHandler("/users", userHandler.Handle)

	go ws.Start()

	select {}
}
