package main

import (
	"fmt"
	"github.com/GuibuAdrian/go-RestFutbet/dao"
	"github.com/GuibuAdrian/go-RestFutbet/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Player struct {
	Name		string	`json:"name" bson:"name"`
	Id			int		`json:"id" bson:"id"`
	Number		int		`json:"number" bson:"number"`
	Position 	string	`json:"position" bson:"position"`
	Team		string	`json:"team" bson:"team"`
}

func getPlayerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	player,_ := dao.PlayerDaoGetInstance().Read(id)
	if player != nil {
		fmt.Println(player)
		playerJSON := Player{
			Name:     player.GetName(),
			Id:       player.GetId(),
			Number:   player.GetNumber(),
			Position: player.GetPosition(),
			Team: player.GetTeam().GetTeamName(),
		}
		c.IndentedJSON(http.StatusOK, playerJSON)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// getAlbums responds with the list of all albums as JSON.
func getPlayersByTeam(c *gin.Context) {
	tName := c.Param("teamName")
	teamDao := dao.TeamDaoGetInstance()

	team, _ := teamDao.ReadByName(tName)

	players := dao.PlayerDaoGetInstance().ReadByTeam(team)

	c.IndentedJSON(http.StatusOK, players)
}

func main() {
	db.InitMongoDB()
	router := gin.Default()
	router.GET("/players/:id", getPlayerByID)
	router.GET("/teamPlayers/:teamName", getPlayersByTeam)

	router.Run("localhost:8080")
}
