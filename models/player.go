package models

type Player struct {
	name		string
	number		int
	id			int
	position	string
	team		Team
}

func InitPlayer(name string, number, id  int, position string, team Team) *Player {
	return &Player{
		name:     name,
		number:   number,
		id:       id,
		position: position,
		team: team,
	}
}

func (player Player) GetName() string { return player.name }
func (player Player) GetNumber() int  { return player.number }
func (player Player) GetId() int { return player.id }
func (player Player) GetPosition() string { return player.position }
func (player Player) GetTeam() Team { return player.team }