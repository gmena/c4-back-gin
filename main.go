package main

import (
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

const WIN = 4
const ROWS = 6
const COLS = 7

var idCounter int32 = 0

type gamestate struct {
	ID         int             `json:"id"`
	NextPlayer int             `json:"nextPlayer"`
	Winner     int             `json:"winner"`
	Board      [ROWS][COLS]int `json:"board"`
}

var ongoingGames = make(map[int]gamestate)

func newId() int {
	return int(atomic.AddInt32(&idCounter, 1))
}

func newGame(c *gin.Context) {
	var id = newId()
	var board [ROWS][COLS]int
	var newGame = gamestate{ID: id, NextPlayer: 1, Winner: 0, Board: board}

	ongoingGames[id] = newGame

	c.JSON(http.StatusOK, newGame)
}

func getGame(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, ongoingGames[id])
}

func main() {
	router := gin.Default()

	router.GET("/connectfour/new", newGame)
	router.GET("/connectfour/:id", getGame)

	router.Run("localhost:8080")
}
