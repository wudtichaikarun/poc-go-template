package server

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/wudtichaikarun/poc-go-template/api/rest/app"
	"github.com/wudtichaikarun/poc-go-template/api/rest/handler"
	"github.com/wudtichaikarun/poc-go-template/config"
	"github.com/wudtichaikarun/poc-go-template/service"
)

type Serverer interface {
	Start()
}

type server struct {
	config *config.EnvConfig
	// store  *database.Store
}

func NewServer(config *config.EnvConfig) Serverer {
	return &server{
		config: config,
		// store:  database.NewStore(config.DBConfig()),
	}
}

func (s *server) Start() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	app := app.NewApp()
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowHeaders:  "*",
		AllowMethods:  "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		ExposeHeaders: "content-disposition",
	}))
	h := handler.New(service.New())

	go func() {
		defer wg.Done()
		<-c
		app.Shutdown()
	}()

	go func() {
		defer wg.Done()
		handler.RegisterRouter(app, h)
		host := fmt.Sprintf("%s:%d", s.config.AppHost, s.config.AppPort)
		fmt.Print("Server is running at " + host)
		app.Listen(host)
	}()

	wg.Wait()
}
