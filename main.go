// main
package main

import (
	"github.com/QLeelulu/goku"
	"gotiedotweb"
	"gotiedotweb/controllers"
	"gotiedotweb/models"
	"log"
)

func main() {
	//	fmt.Println("Hello World!")
	database := models.GetDB()
	database.Create(gotiedotweb.BAND_COL)
	database.Create(gotiedotweb.LOCATION_COL)
	database.Create(gotiedotweb.GENRE_COL)
	database.Close()
	rt := &goku.RouteTable{Routes: gotiedotweb.Routes}
	middlewares := []goku.Middlewarer{}
	s := goku.CreateServer(rt, middlewares, gotiedotweb.Config)
	goku.Logger().Logln("Server start on", s.Addr)
	log.Fatal(s.ListenAndServe())
}

var home = controllers.HomeController
var band = controllers.BandController
var album = controllers.AlbumController
