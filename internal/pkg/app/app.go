package app

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"middleware/internal/app/endpoint"
	"middleware/internal/app/mw"
	"middleware/internal/app/service"
	"os"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
	log  *slog.Logger
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
	address  = ":8088"
)

func New() *App {
	a := &App{}

	a.log = setupLogger("local")

	a.log.Info("starting server...")

	a.s = service.New()

	a.e = endpoint.New(a.s, a.log)

	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck(a.log))

	a.echo.GET("/status", a.e.StatusHandler)

	return a
}

func (a *App) MustRun() {
	err := a.echo.Start(address)
	if err != nil {
		panic(err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
