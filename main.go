package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User has all the details of the User
type User struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

// Users an array of User
type Users struct {
	Users []User
}

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Hello there!"})
}

func fetchTopContributers(c *gin.Context) {
	var org = c.Param("org")
	var repo = c.Param("repo")
	var users []User
	response, err := http.Get("https://api.github.com/repos/" + org + "/" + repo + "/contributors?q=contributions&order=desc")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "error invoking git api's"})
		return
	}
	data, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "error while unmarshalling the json"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
	return
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", sayHello)
		v1.GET("/:org/:repo", fetchTopContributers)
	}
	router.Run()
}
