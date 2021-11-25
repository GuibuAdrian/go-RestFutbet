package main

import (
	"fmt"
	"go-RestFutbet/dao"
	"go-RestFutbet/db"
	"go-RestFutbet/models"
)

func main() {
	db.InitMongoDB()
	teamDao := dao.TeamDaoGetInstance()

	team, _ := teamDao.ReadByName("River")
	fmt.Println(team.GetTeamObjId())

	playerDao := dao.PlayerDaoGetInstance()
	p, _ := playerDao.ReadByPlayerName("Armani", team.GetTeamObjId())
	fmt.Println(p)
	fmt.Println(p.GetName())

	teamPrueba, _ := teamDao.ReadByName("Newells")
	playerPrueba := models.InitPlayer("Prueba2", 000000002, 000000002, "prueba2", teamPrueba)
	//playerDao.Create(playerPrueba)
	playerDao.Update("619ce00c194bc222fcf65c77", playerPrueba)

	playerDao.Delete("619ce00c194bc222fcf65c77")
	fmt.Println("Hello World!")
}
