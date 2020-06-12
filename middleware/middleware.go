package middleware

import (
	"github.com/labstack/echo/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
