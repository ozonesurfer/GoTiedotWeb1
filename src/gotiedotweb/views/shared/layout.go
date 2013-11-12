// layout
package shared

import (
	"gomgoweb"
	"gomgoweb/models"
	"labix.org/v2/mgo/bson"
)

func GetGenreName2(id bson.ObjectId) string {
	database, session := models.GetDB()
	defer session.Close()
	collection := database.C(gomgoweb.GENRE_COL)
	var doc models.MyDoc
	collection.Find(id).One(&doc)
	var genre models.Genre
	m := doc.Value
	for _, value := range m {
		genre.Name = value.(string)
	}
	//	genre := doc.Value.(Genre)
	return genre.Name
}
