package main

import (
	"github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/contrib/jwt"
)

func (app *App)Routes() *fiber.App{

	r:=fiber.New()

	
	r.Get("/login",app.Login)


	//protected routes
	r.Use(jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte("rejathchandran")},
    }))
   
	r.Get("/details",GetToken,app.Details)


// TODO: category 
// TODO: order 
// TODO: cart 


	return r
}