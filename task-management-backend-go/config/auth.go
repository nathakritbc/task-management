package config

import (
	jwtware "github.com/gofiber/contrib/jwt"
)

var secretKey = Config("SECRET_KEY")
var ConfigAuth = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{Key: []byte(secretKey)},
})
