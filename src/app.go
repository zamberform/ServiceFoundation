package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gomodule/redigo/redis"
)

type Person struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Age  int           `bson:"age"`
}

type Happy struct {
	Money  int
	Moral  bool
	Health bool
}

func main() {
	//mongodb 部分
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

	//mongodb 終わり
	// =============================================
	//redis 部分
	var key = flag.String("key", "testKey", "Set Key")
	var val = flag.String("value", "redis store is string", "Set Value")

	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer c.Close()

	c.Do("SET", *key, *val)
	s, err := redis.String(c.Do("GET", *key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%#vn", s)
	//redis 終わり

	r := gin.Default()
	r.GET("/mongo", func(c *gin.Context) {
		c.JSON(200, p)
	})

	r.GET("/redis", func(c *gin.Context) {
		c.String(200, "%s", s)
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	r.Run(":" + port)

}
