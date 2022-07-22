package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//var movementDataCollection *mongo.Collection = configs.GetCollection(configs.DB, "mousemovementdata")

type userDataRes struct {
	Status                          string  `json:"status"`
	Date                            int     `json:"data"`
	StandardDeviation               float64 `json:"standardDeviation"`
	AveragMovementDistance          float64 `json:"averagMovementDistance"`
	AverageMovementTime             float64 `json:"averageMovementTime"`
	AverageMovementDistanceOverTime float64 `json:"AverageMovementDistanceOverTime"`
}

func ViewData(c *gin.Context) {
	avantCookie, err := c.Request.Cookie("avant")
	if err != nil {
		c.String(http.StatusBadRequest, "untrusted")

	}

	userFingerprint := Verify(avantCookie.Value, os.Getenv("AVANTPRIVATEAPIKEY"), "authuser").UserIdentification
	fmt.Println(userFingerprint)

	var storedUserData userDataRes
	for i := 0; i < 3; i++ {

		mongoErr := movementDataCollection.FindOne(context.TODO(), bson.M{
			"_id": userFingerprint,
		}).Decode(&storedUserData)
		if mongoErr == mongo.ErrNoDocuments {
			fmt.Println("looping again")
		} else if mongoErr != nil {
			fmt.Println("err")
			fmt.Println(mongoErr)
			c.String(500, "internal error")
			return
		} else {
			storedUserData.Status = "good"
			c.JSON(200, storedUserData)
			return
		}
		time.Sleep(250 * time.Millisecond)
	}
	c.String(400, "no data found")
}
