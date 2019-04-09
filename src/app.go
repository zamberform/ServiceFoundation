package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Person struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Age  int           `bson:"age"`
}

func main() {
	session, _ := mgo.Dial("mongodb://root:example@localhost:27017")
	defer session.Close()
	db := session.DB("test")

	ritsu := &Person{
		ID:   bson.NewObjectId(),
		Name: "田井中律",
		Age:  17,
	}
	col := db.C("persons")
	if err := col.Insert(ritsu); err != nil {
		log.Fatalln(err)
	}

	p := new(Person)
	query := db.C("persons").Find(bson.M{})
	query.One(&p)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, p)
	})
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	r.Run(":" + port)

}
