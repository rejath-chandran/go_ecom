package main

import (
	"github.com/julienschmidt/httprouter"
)

func (app *App)Routes() *httprouter.Router{
	
	r:=httprouter.New()


	r.GET("/v1/register",app.CreateSeller)



	return r
}