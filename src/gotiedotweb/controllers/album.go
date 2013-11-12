// album
package controllers

import (
	//	"encoding/json"
	"fmt"
	"github.com/QLeelulu/goku"
	"gotiedotweb"
	"gotiedotweb/models"
	//	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"strconv"
)

var AlbumController = goku.Controller("album").
	Get("index", func(ctx *goku.HttpContext) goku.ActionResulter {
	rawId := ctx.RouteData.Params["id"]
	fmt.Println("rawId =", rawId)
	ctx.ViewData["Id"] = rawId
	bandId := ToObjectId(rawId)
	fmt.Println("bandId = ", bandId)
	bandDoc := models.GetDoc(bandId, gotiedotweb.BAND_COL)
	//	ctx.ViewData["Title"] = bandDoc.Value["Name"].(string)
	//	band := bandDoc.Value.(models.Band)
	band := bandDoc.Value
	ctx.ViewData["Title"] = band["name"]
	return ctx.View(bandDoc)
}).Get("add", func(ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Title"] = "Add Album"
	ctx.ViewData["Id"] = ToObjectId(ctx.RouteData.Params["id"])
	genres := models.GetAll(gotiedotweb.GENRE_COL)
	return ctx.View(genres)
}).Post("verify", func(ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Title"] = "Verifying Album"
	rawId := ToObjectId(ctx.RouteData.Params["id"])
	ctx.ViewData["Id"] = rawId
	name := ctx.Request.FormValue("name")
	yearString := ctx.Request.FormValue("year")
	year, _ := strconv.Atoi(yearString)
	genretype := ctx.Request.FormValue("genretype")
	var genreId uint64
	errorString := "no errors"
	var err error
	switch genretype {
	case "existing":
		if ctx.Request.FormValue("genre_id") == "" {
			errorString = "No genre was selected"
		} else {
			genreId = ToObjectId(ctx.Request.FormValue("genre_id"))
		}
		break
	case "new":
		if ctx.Request.FormValue("genre_name") != "" {
			//			genreId = models.GenerateId()
			genre := map[string]interface{}{"name": ctx.Request.FormValue("genre_name")}

			//			doc := models.MyDoc{Id: genreId, Value: bson.M{"Name": genre.Name}}
			genreId, err = models.AddDoc(genre, gotiedotweb.GENRE_COL)
			if err != nil {
				errorString = fmt.Sprintf("Genre: %s", err.Error())
			}
		} else {
			errorString = "Genre name is required"
		}
		break
	}

	if errorString == "no errors" {

		bandDoc := models.GetDoc(rawId, gotiedotweb.BAND_COL)
		album := models.Album{Name: name, Year: year, GenreId: genreId}
		err := bandDoc.AddAlbum(album)
		if err != nil {
			errorString = fmt.Sprintf("Album: %s", err.Error())
		}
	}
	ctx.ViewData["Message"] = errorString
	return ctx.View(nil)
})
