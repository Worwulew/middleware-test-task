package mw

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"middleware/internal/pkg/logger/sl"
	"strings"
)

const admin = "admin"

func RoleCheck(log *slog.Logger) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			const fn = "mw.RoleCheck"

			role := ctx.Request().Header.Get("User-Role")

			if strings.EqualFold(role, admin) {
				log.Info("User is admin", slog.String("fn", fn))
			} else {
				log.Info("User is not admin", slog.String("fn", fn))
			}

			err := next(ctx)
			if err != nil {
				log.Error("some error", sl.Err(fmt.Errorf("%s: %w", fn, err)))

				return err
			}

			return nil
		}
	}
}
