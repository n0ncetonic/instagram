package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/n0ncetonic/instagram/profile"
)

var router *gin.Engine

func idof(c *gin.Context) {
	username := c.Param("username")

	data := getSharedData(username)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": data.GetUser().GetID()})
}

func getSharedData(u string) profile.Data {
	res, err := http.Get("https://instagram.com/" + u)
	if err != nil {
		log.Println("failed to connect to Instagram")
		return profile.Data{}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("http status code: %d\nhttp status: %s", res.StatusCode, res.Status)
		return profile.Data{}
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("%v", err)
	}

	rawScript := doc.Find("body script:nth-child(2)").Text()

	rawJSON := strings.TrimPrefix(rawScript, "window._sharedData = ")
	rawJSON = strings.TrimSuffix(rawJSON, ";")

	var profileData = new(profile.Data)

	json.Unmarshal([]byte(rawJSON), &profileData)
	return *profileData

}

func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	igRoutes := router.Group("/instagram")
	{
		igRoutes.GET("/idof/:username", idof)
	}

	router.Run("127.0.0.1:8080")
}
