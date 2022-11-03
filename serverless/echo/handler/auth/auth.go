package auth

import (
	conf "yhidetoshi/serverless-aws-go-local/conf"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	id = conf.BASIC_AUTH_ID
	pw = conf.BASIC_AUTH_PASS
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		if username == id && password == pw {
			return true, nil
		}
		return false, nil
	})
}
