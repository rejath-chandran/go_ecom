package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func (app *App)CreateSeller(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
  fmt.Fprintf(w,"hey running")
}
func (app *App)LoginSeller(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
  fmt.Fprintf(w,"hey running")
}