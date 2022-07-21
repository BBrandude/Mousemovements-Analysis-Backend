package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mouse-mousements-thesis-backend/calculations"
	"mouse-mousements-thesis-backend/configs"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type dataPayload struct {
	MouseMovement []struct {
		X int     `json:"x"`
		Y int     `json:"y"`
		T float64 `json:"t"`
	} `json:"mouseMovement"`
}

var movementDataCollection *mongo.Collection = configs.GetCollection(configs.DB, "mousemovementdata")

func ProccessData(c *gin.Context) {

	var userData dataPayload

	if err := c.BindJSON(&userData); err != nil {
		log.Fatal(err)
		c.String(http.StatusInternalServerError, "invalid data submitted")
		return
	}
	avantCookie, err := c.Request.Cookie("avant")
	if err != nil {
		c.String(http.StatusBadRequest, "untrusted")
	}
	cookieVerifcationRes := verify(avantCookie.Value, os.Getenv("AVANTPRIVATEAPIKEY"))
	if cookieVerifcationRes.Status == "deny" {
		c.String(http.StatusBadRequest, "untrusted")
		return
	} else if cookieVerifcationRes.Status != "allow" {
		c.String(http.StatusInternalServerError, "internal error")
		return
	}

	movementData := userData.MouseMovement

	mouseMoveDistances := make([]float64, 0)
	mouseMoveTimes := make([]float64, 0)

	for i := 0; i < len(movementData)-1; i++ {
		mouseMoveDistances = append(mouseMoveDistances, calculations.CalculateDistance(movementData[i].X, movementData[i+1].X, movementData[i].Y, movementData[i+1].Y))
		mouseMoveTimes = append(mouseMoveTimes, (movementData[i+1].T - movementData[i].T))
	}

	mouseMovementsOverTimes := calculations.CalcDistanceOverTime(mouseMoveDistances, mouseMoveTimes)

	averageMovementDistance := calculations.CalculateAverage(mouseMoveDistances)
	averageMovementTimes := calculations.CalculateAverage(mouseMoveTimes)
	averageMovementDistanceOverTime := calculations.CalculateAverage(mouseMovementsOverTimes)
	standardDeviation := calculations.CalculateStandardDeviation(movementData)

	c.String(http.StatusOK, "thank you for the data")

	movementDataCollectionInsertion, err := movementDataCollection.InsertOne(context.TODO(), bson.M{
		"_id":                             cookieVerifcationRes.UserIdentification,
		"date":                            time.Now().UnixMilli(),
		"movements":                       movementData,
		"standardDeviation":               standardDeviation,
		"averagMovementDistance":          averageMovementDistance,
		"averageMovementTime":             averageMovementTimes,
		"averageMovementDistanceOverTime": averageMovementDistanceOverTime,
	})
	_ = err
	_ = movementDataCollectionInsertion

}

type cookieRes struct {
	Status             string `json:"status"`
	Reason             string `json:"reason"`
	UserIdentification string `json:"userIdentification"`
}

func verify(cookie string, apiKey string) cookieRes {

	url := "http://avantsecure.net/endpointprotection/" + cookie
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("x-api-key", apiKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	var avantResponse cookieRes
	err = json.Unmarshal([]byte(body), &avantResponse) // here!
	if err != nil {
		fmt.Println(err)
	}

	return avantResponse
}
