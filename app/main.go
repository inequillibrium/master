package main

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Node struct {
	IP        string        `bson:"IP" json:"IP"`
	Status    int32         `bson:"Status" json:"Status"`
	Instances string        `bson:"Instances" json:"Instances"`
	CPU       string        `bson:"CPU" json:"CPU"`
	RAM       string        `bson:"RAM" json:"RAM"`
}

var s, err = mgo.Dial("localhost")

func main() {
	r := gin.Default()
	if err != nil {
		panic(err)
	}
	defer s.Close()
	s.SetMode(mgo.Monotonic, true)

	r.GET("/ping", ping)
	r.GET("/nodes", listNodes)
	r.Run()
}

func ping(co *gin.Context) {
	co.JSON(200, gin.H{
		"message": "pong",
	})
}

func listNodes(co *gin.Context) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("balancer").C("nodes")
	var nodes []Node
	err := c.Find(bson.M{}).All(&nodes)
	if err != nil {
		panic(err)
	}
	co.JSON(200, gin.H{
		"nodes": nodes,
	})
}