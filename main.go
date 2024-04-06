package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"poetic-interludes/db"
	"strconv"

	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/walles/env"
)

// type author struct {
// 	ID        string `json:"id"`
// 	FirstName string `json:"firstName" db:"first_name"`
// 	LastName  string `json:"lastName" db:"last_name"`
// 	Email     string `json:"email" db:"email"`
// }

// type authorWithPassword struct {
// 	author
// 	password string
// }

type poem struct {
	ID         string `json:"id" db:"id"`
	AuthorName string `json:"authorName" db:"author_name"`
	Text       string `json:"text" db:"text"`
	Title      string `json:"title" db:"title"`
}

func main() {
	loadEnv()

	db.Connect()
	defer db.Connection.Close()

	router := gin.Default()
	router.GET("/poems", getPoems)
	router.POST("/poems", addPoem)
	router.GET("/poems/:id", getPoemByID)

	router.Run("localhost:" + strconv.Itoa(env.GetOr("PORT", strconv.Atoi, 8080)))
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getPoems(c *gin.Context) {
	var poems []*poem
	err := sqlscan.Select(context.Background(), db.Connection, &poems /* sql */, `
		SELECT
		  p.id,
		  a.first_name || ' ' || a.last_name as author_name,
		  p.title,
		  p.text
		FROM poems p
		LEFT JOIN authors a ON a.id = p.author_id
	`)
	if err != nil {
		// TODO: proper error handling
		fmt.Println("Error fetching poems", err)
	} else {
		c.IndentedJSON(http.StatusOK, poems)
	}
}

// postAlbums adds an album from JSON received in the request body.
func addPoem(c *gin.Context) {
	var newPoem poem

	if err := c.BindJSON(&newPoem); err != nil {
		// TODO: proper error handling
		return
	} else {
		c.IndentedJSON(http.StatusCreated, newPoem)
	}
}

func getPoemByID(c *gin.Context) {
	id := c.Param("id")

	var poems []poem
	err := sqlscan.Select(context.Background(), db.Connection, &poems /* sql */, `
		SELECT
		  p.id,
		  a.first_name || ' ' || a.last_name as author_name,
		  p.title,
		  p.text
		FROM poems p
		LEFT JOIN authors a ON a.id = p.author_id
		WHERE p.id = $1
	`, id)
	if err != nil {
		fmt.Println("err", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	} else if len(poems) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "poem not found"})
	} else {
		c.IndentedJSON(http.StatusOK, poems[0])

	}

}
