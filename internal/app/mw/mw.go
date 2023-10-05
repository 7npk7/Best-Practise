package mw

import (
	"github.com/labstack/echo"
	"log"
	"strings"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		val := ctx.Request().Header.Get("user-role")

		if strings.EqualFold(val, roleAdmin) {
			log.Println("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
