package main

import (
	"encoding/csv"
	"log"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type Data struct {
	ID       int    `csv:"id"`
	Name     string `csv:"name"`
	Age      int    `csv:"age"`
	Location string `csv:"location"`
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/data", fileExists(), getData)
	router.GET("/data/:id", fileExists(), getOneData)
	router.POST("/data", fileExists(), createData)
	router.PUT("/data/:id", fileExists(), updateData)
	router.DELETE("/data/:id", fileExists(), deleteData)

	router.Run("localhost:5700")
}

// Checking if file exists
func fileExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := os.Stat("user.csv")
		if os.IsNotExist(err) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "CSV file not found"})
			return
		}
		c.Next()
	}
}

// Getting data from CSV file using GET method
func getData(c *gin.Context) {
	file, err := os.Open("user.csv")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result []Data
	for i, row := range records {
		// Skipping the header row as it would be column names
		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(row[0])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		age, err := strconv.Atoi(row[2])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":
