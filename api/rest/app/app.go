package app

import "github.com/wudtichaikarun/poc-go-template/api/rest/router"

type Apper interface {
	Listen(addr string) error
	Shutdown() error
}

type App struct {
	*router.FiberRouter
}

func NewApp() *App {
	return &App{router.NewFiberRouter()}
}

func (a *App) Listen(addr string) error {
	return a.App.Listen(addr)
}

func (a *App) Shutdown() error {
	return a.App.Shutdown()
}
