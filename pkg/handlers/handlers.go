package handlers

import (
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

var globalMap = make(map[int]int)
var globalID = 0

func PostPoints(c *gin.Context) {
	var request Receipt
	var points int
	var totalCost float64

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data", "description": "The receipt is invalid.", "details": err.Error()})
		return
	}

	globalID++
	// id := uuid.New()
	// log.Println("Received JSON:", request)
	points += CountAlphaNumeric(request.Retailer)
	totalCost = GetTotalCost(request.Items)
	totalCost = math.Round(totalCost*100) / 100 // round to 2 decimal places
	numItems := len(request.Items)
	// log.Println("Retailer points: ", points)
	// log.Println("Total cost: ", totalCost)
	// log.Println("Is round dollar: ", math.Floor(totalCost) == math.Ceil(totalCost))
	// log.Println("Mod 0.25: ", math.Mod(totalCost, 0.25) == 0)
	if IsRoundDollar(totalCost) {
		points += 50
	}
	if math.Mod(totalCost, 0.25) == 0 {
		points += 25
	}
	// log.Println("Points added by /2: ", 5*math.Floor(float64(numItems)/2))
	points += 5 * int(math.Floor(float64(numItems)/2))
	points += TrimPoints(request.Items)
	// log.Println("Points added by trimming: ", TrimPoints(request.Items))
	points += AddOddDay(request.PurchaseDate)
	points += AddTime(request.PurchaseTime)
	// log.Println("Total points: ", points)
	globalMap[globalID] = points
	c.JSON(http.StatusOK, gin.H{"id": strconv.Itoa(globalID)})
}

func GetID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID", "description": "The ID is invalid."})
		return
	}

	points, exists := globalMap[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID not found", "description": "No receipt found for that ID."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
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
			// log.Println(s)
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

func AddOddDay(purchaseDateStr string) int {
	points := 0
	purchaseDate, err := time.Parse("2006-01-02", purchaseDateStr)
	if err == nil {
		day := purchaseDate.Day()
		// log.Println("Day: ", day)
		if day%2 != 0 {
			points += 6
		}
	}
	// log.Println("Points added by odd day: ", points)
	return points
}

func AddTime(purchaseTimeStr string) int {
	points := 0
	purchaseTime, err := time.Parse("15:04", purchaseTimeStr)
	if err == nil {
		hour := purchaseTime.Hour()
		// log.Println("Hour: ", hour)
		if hour >= 14 && hour <= 16 {
			points += 10
		}
	}
	// log.Println("Points added by time: ", points)
	return points
}
