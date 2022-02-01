package main

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var checkAuth_ = middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	if subtle.ConstantTimeCompare([]byte(username), []byte(config.Username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(config.Password)) == 1 {
		return true, nil
	}
	return false, nil
})

func checkAuth(next echo.HandlerFunc) echo.HandlerFunc {
	ca := checkAuth_(next)

	return func(c echo.Context) error {
		err := ca(c)
		// This header makes the browser ask the user for a username-password
		// combo, which we don't want
		c.Response().Header().Del(echo.HeaderWWWAuthenticate)
		return err
	}
}
