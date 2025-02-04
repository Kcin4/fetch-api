package handlers

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type PointsRequest struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

var globalMap = make(map[int]int)

func PostPoints(c *gin.Context) {
	var request PointsRequest
	var points int
	var totalCost float64

	// Bind the JSON payload to the struct
	c.ShouldBindJSON(&request)

	// Log the received data
	// log.Println("Received JSON:", request)
	points += CountAlphaNumeric(request.Retailer)
	totalCost = GetTotalCost(request.Items)
	totalCost = math.Round(totalCost*100) / 100
	numItems := len(request.Items)
	log.Println("Retailer points: ", points)
	log.Println("Total cost: ", totalCost)
	log.Println("Is round dollar: ", math.Floor(totalCost) == math.Ceil(totalCost))
	log.Println("Mod 0.25: ", math.Mod(totalCost, 0.25) == 0)
	if IsRoundDollar(totalCost) {
		points += 50
	}
	if math.Mod(totalCost, 0.25) == 0 {
		points += 25
	}
	log.Println("Points added by /2: ", math.Floor(float64(numItems)/2))
	points += int(math.Floor(float64(numItems) / 2))
	// log.Println("Total cost: ", totalCost)
	// Update the global map (example logic, you can modify as needed)
	// Here we just log the received data
	c.JSON(http.StatusOK, gin.H{"message": "got to the post function", "data": points})
}

func GetID(c *gin.Context) {
	log.Println("got to the get function")
	c.JSON(200, "got to the get function")
}

func CountAlphaNumeric(s string) int {
	count := 0
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			count++
		}
	}
	return count
}

func GetTotalCost(items []Item) float64 {
	totalCost := float64(0)
	for _, item := range items {
		price := item.Price
		if s, err := strconv.ParseFloat(price, 64); err == nil {
			// fmt.Println(s)
			totalCost += float64(s)
		}
	}
	return totalCost
}

func IsRoundDollar(totalCost float64) bool {
	return math.Floor(totalCost) == math.Ceil(totalCost)
}

func TrimPoints(items []Item) int {
	points := 0
	for _, item := range items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}
	return points
}
