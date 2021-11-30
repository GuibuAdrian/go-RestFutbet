package db

import (
	"encoding/json"
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
)


type Response1 struct {
	User		string `json:"user"`
	Pass		string `json:"pass"`
	Url			string `json:"url"`
	ClusterName	string `json:"cluster_name"`
	DBName		string `json:"db_name"`
}

func InitMongoDB() {
	file, _ := ioutil.ReadFile("../file.json")
	data := Response1{}
	_ = json.Unmarshal([]byte(file), &data)

	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, data.DBName, options.Client().ApplyURI("mongodb+srv://"+data.User+":"+data.Pass+"@"+data.ClusterName+".mongodb.net/"+data.Url))
	if err != nil {
		fmt.Println(err)
	}
}

