package interfaces

import (
	"github.com/GuibuAdrian/go-RestFutbet/models"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Player struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	Name		string	`json:"name" bson:"name"`
	Id			int		`json:"id" bson:"id"`
	Number		int		`json:"number" bson:"number"`
	mgm.DefaultModel	`bson:",inline"`
	Position 	string	`json:"position" bson:"position"`
	TeamId		primitive.ObjectID
}

type PlayerDaoI interface {
	Create(player *models.Player)
	ReadByPlayerName(playerName string, teamId primitive.ObjectID) (*models.Player, error)
	Read(idP int) (*models.Player, error)
	ReadByTeam(team models.Team) []Player
	Update(playerId string , player *models.Player)
	Delete(playerId string)
}
