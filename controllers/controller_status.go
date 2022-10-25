package controllers

import (
	"assignment_3/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func GetStatus() {
	content_status := models.WeatherStatus{
		Status: models.Weather{
			Water: rand.Intn(18),
			Wind:  rand.Intn(18),
		},
	}
	content_json, _ := json.MarshalIndent(content_status, "", " ")

	_ = ioutil.WriteFile("status.json", content_json, 0644)
	fmt.Printf("Done")
}

func Status(ctx *gin.Context) {

}
