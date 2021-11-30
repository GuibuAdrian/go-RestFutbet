package dao

import (
	"errors"
	"fmt"
	"github.com/GuibuAdrian/go-RestFutbet/models"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayerDao struct {
	playerSlice []*models.Player
}

type Player struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	Name		string	`json:"name" bson:"name"`
	Id			int		`json:"id" bson:"id"`
	Number		int		`json:"number" bson:"number"`
	mgm.DefaultModel	`bson:",inline"`
	Position 	string	`json:"position" bson:"position"`
	TeamId		primitive.ObjectID
}

func NewPlayer(name string, id, number int, position string, teamId primitive.ObjectID) *Player {
	return &Player{
		Name:     name,
		Id:       id,
		Number:   number,
		Position: position,
		TeamId:   teamId,
	}
}

var playerDaoInstance *PlayerDao

func PlayerDaoGetInstance() *PlayerDao {
	if playerDaoInstance == nil {
		playerDaoInstance = &PlayerDao{}
	}

	return playerDaoInstance
}

func (playerDao *PlayerDao) Create(player *models.Player)  {
	playerCreate := NewPlayer(player.GetName(), player.GetId(), player.GetNumber(), player.GetPosition(), player.GetTeam().GetTeamObjId())

	// Make sure to pass the model by reference.
	errorCreate := mgm.Coll(playerCreate).Create(playerCreate)
	check(errorCreate)
	//playerDao.playerSlice = append(playerDao.playerSlice, player)
}

func (playerDao PlayerDao) ReadByPlayerName(playerName string, teamId primitive.ObjectID) (*models.Player, error) {
	player := &Player{}
	coll := mgm.Coll(player)

	_ = coll.First(bson.M{"name":playerName, "teamid": teamId}, player)
	team,err := TeamDaoGetInstance().Read(teamId)
	check(err)

	playerM := models.InitPlayer(player.Name, player.Number, player.Number, player.Position, team)
	return playerM, errors.New("player not found")
}

func (playerDao PlayerDao) ReadByTeam(team models.Team) []Player {
	player := &Player{}
	coll := mgm.Coll(player)
	result := []Player{}

	err := coll.SimpleFind(&result, bson.M{"teamid": team.GetTeamObjId()})
	check(err)

	return result
}

func (playerDao PlayerDao) Read(idP int) (*models.Player, error) {
	player := &Player{}
	coll := mgm.Coll(player)
	_ = coll.First(bson.M{"id":idP}, player)
	team,err := TeamDaoGetInstance().Read(player.TeamId)
	check(err)

	playerM := models.InitPlayer(player.Name, player.Number, player.Number, player.Position, team)
	fmt.Println("Player found:", playerM)
	return playerM, errors.New("player not found")
}

func (playerDao *PlayerDao) Update(playerId string , player *models.Player)  {
	playerU := &Player{}
	coll := mgm.Coll(playerU)
	err := coll.FindByID(playerId, playerU)
	check(err)

	playerU.Name = player.GetName()
	playerU.Id = player.GetId()
	playerU.TeamId = player.GetTeam().GetTeamObjId()
	playerU.Number = player.GetNumber()
	playerU.Position = player.GetPosition()

	err = coll.Update(playerU)
	check(err)
}

func (playerDao *PlayerDao) Delete(playerId string)  {
	/*for posV, playerVal := range playerDao.playerSlice {
		if playerVal == &player{
			playerDao.playerSlice = append(playerDao.playerSlice[:posV], playerDao.playerSlice[posV+1:]...)
			break
		}
	}*/
	playerD := &Player{}
	coll := mgm.Coll(playerD)
	err := coll.FindByID(playerId, playerD)
	check(err)
	err = coll.Delete(playerD)
	check(err)
}

func (playerDao PlayerDao) GetPlayerSlice() []*models.Player {
	return playerDao.playerSlice
}

func initializePlayerDao() {
	river, _ := TeamDaoGetInstance().ReadByName("River")
	newells, _ := TeamDaoGetInstance().ReadByName("Newells")
	PlayerDaoGetInstance().Create(models.InitPlayer("Armani", 1, 1111221, "arquero", river ))
	PlayerDaoGetInstance().Create(models.InitPlayer("Angileri", 3, 3322333, "defensor", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("P. Diaz", 17, 1717171, "defensor", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Rojas", 2, 2122222, "defensor", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Vigo", 16, 1616162, "defensor", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Carrascal", 10, 1010101, "medio C", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("E. Perez", 24, 2424242, "medio C", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Zuculini", 5, 5555555, "medio C", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("De la Cruz", 11, 1121111, "medio C", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Alvarez", 9, 9999999, "delantero", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Romero", 19, 1919323, "delantero", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Aguerre", 1, 1111111, "arquero", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Compagnucci", 27, 2727272, "defensor", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Lema", 2, 2222222, "defensor", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Z. F. Mansilla", 19, 1919191, "defensor", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Bittolo", 28, 2828282, "defensor", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("P. Perez", 8, 8888888, "medio C", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("J. Fernandez", 20, 2020202, "medio C", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Comba", 33, 3333333, "medio C", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Castro", 36, 3636363, "medio C", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Sordo", 26, 2626262, "medio C", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Scocco", 32, 3232323, "delantero", newells))
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}