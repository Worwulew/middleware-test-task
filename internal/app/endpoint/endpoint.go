package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"middleware/internal/pkg/logger/sl"
	"net/http"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s   Service
	log *slog.Logger
}

func New(s Service, log *slog.Logger) *Endpoint {
	return &Endpoint{
		s:   s,
		log: log,
	}
}

func (e *Endpoint) StatusHandler(ctx echo.Context) error {
	const fn = "endpoint.StatusHandler"

	d := e.s.DaysLeft()

	s := fmt.Sprintf("Days left until January 1 2025: %d", d)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		e.log.Error("could not send response", sl.Err(fmt.Errorf("%s: %w", fn, err)))

		return err
	}

	e.log.Info("success", slog.String("fn", fn))

	return nil
}
