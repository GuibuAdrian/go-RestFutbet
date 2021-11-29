package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	Name string
	ObjId primitive.ObjectID
}

func InitTeam( name string, objId primitive.ObjectID ) Team {
	return Team{
		Name: name,
		ObjId: objId,
	}
}

func (team Team) GetTeamName() string              { return team.Name }
func (team Team) GetTeamObjId() primitive.ObjectID { return team.ObjId }
