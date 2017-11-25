package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v3"
)

var rd *redis.Client

func init() {

	rd = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rd.Ping().Result()
	fmt.Println(pong, err)
}

func main() {
	router := gin.Default()
	router.Static("/css", "css/")
	router.Static("/js", "js/")
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", Index)
	router.GET("/schedule", ScheduleGet)
	router.POST("/schedule", SchedulePost)

	router.Run(":8080")
}

func Index(c *gin.Context) {
	t := rd.ZRange("t_hist", -1, -1)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": t.Val()[0],
	})

	fmt.Print(t.Result())
}

func ScheduleGet(c *gin.Context) {
	c.HTML(http.StatusOK, "schedule.tmpl", gin.H{})

}

func SchedulePost(c *gin.Context) {
	message, _ := c.GetRawData()
	fmt.Print("schedule Post\n")
	log.Println(message)

}
