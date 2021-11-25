package dao

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"go-RestFutbet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamDao struct {
	teamSlice []models.Team
}

type Team struct {
	Name	string		`json:"name" bson:"name"`
	mgm.DefaultModel	`bson:",inline"`
}

var teamDaoInstance *TeamDao

func TeamDaoGetInstance() *TeamDao {
	if teamDaoInstance == nil {
		teamDaoInstance = &TeamDao{}
	}
	return teamDaoInstance
}

func NewTeam(Name string) *Team {
	return &Team{
		Name: Name,
	}
}
func (teamDao *TeamDao) Create(team models.Team) {
	teamCreate := NewTeam(team.GetTeamName())
	errCreate := mgm.Coll(teamCreate).Create(teamCreate)
	if errCreate != nil {
		fmt.Println(errCreate)
	}
	//teamDao.teamSlice = append(teamDao.teamSlice, team)
}

func (teamDao TeamDao) Read(id primitive.ObjectID) (models.Team, error) {
	team := &Team{}
	coll := mgm.Coll(team)
	err := coll.FindByID(id, team)
	teamM := models.InitTeam(team.Name, team.DefaultModel.ID)
	return teamM, err
	/*team := models.Team{}
	for _, team := range teamDao.teamSlice {
		if team.GetTeamId() == id {
			return team, nil
		}
	}

	return team, errors.New("team not found")*/
}

func (teamDao *TeamDao) ReadByName(teamName string) (models.Team, error) {
	team := &Team{}
	coll := mgm.Coll(team)

	err := coll.First(bson.M{"name":teamName}, team)
	teamM := models.InitTeam(team.Name, team.DefaultModel.ID)
	return teamM, err
}

func (teamDao TeamDao) Update(team models.Team)  {
	//Todo
}

func (teamDao *TeamDao) Delete(team models.Team)  {
	for posV, teamVal := range teamDao.teamSlice {
		if teamVal == team{
			teamDao.teamSlice = append(teamDao.teamSlice[:posV], teamDao.teamSlice[posV+1:]...)
			break
		}
	}
}

func (teamDao TeamDao) GetTeamSlice() []models.Team {
	return teamDao.teamSlice
}

func initializeTeamDao() {
	TeamDaoGetInstance().Create(models.InitTeam("River", primitive.ObjectID{}))
	TeamDaoGetInstance().Create(models.InitTeam("Newells", primitive.ObjectID{}))
	TeamDaoGetInstance().Create(models.InitTeam("Banfield", primitive.ObjectID{}))
	TeamDaoGetInstance().Create(models.InitTeam("Velez", primitive.ObjectID{}))
}